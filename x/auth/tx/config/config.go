package tx

import (
	"fmt"

	txconfigv1 "github.com/verzth/cosmos-sdk/api/cosmos/tx/config/v1"
	"github.com/verzth/cosmos-sdk/baseapp"
	"github.com/verzth/cosmos-sdk/client"
	"github.com/verzth/cosmos-sdk/codec"
	"github.com/verzth/cosmos-sdk/core/appmodule"
	"github.com/verzth/cosmos-sdk/depinject"
	"github.com/verzth/cosmos-sdk/runtime"
	sdk "github.com/verzth/cosmos-sdk/types"
	"github.com/verzth/cosmos-sdk/x/auth/ante"
	"github.com/verzth/cosmos-sdk/x/auth/posthandler"
	"github.com/verzth/cosmos-sdk/x/auth/signing"
	"github.com/verzth/cosmos-sdk/x/auth/tx"
	authtypes "github.com/verzth/cosmos-sdk/x/auth/types"
)

func init() {
	appmodule.Register(&txconfigv1.Config{},
		appmodule.Provide(ProvideModule),
	)
}

//nolint:revive
type TxInputs struct {
	depinject.In

	Config              *txconfigv1.Config
	ProtoCodecMarshaler codec.ProtoCodecMarshaler

	AccountKeeper ante.AccountKeeper `optional:"true"`
	// BankKeeper is the expected bank keeper to be passed to AnteHandlers
	BankKeeper authtypes.BankKeeper `optional:"true"`
	// TxBankKeeper is the expected bank keeper to be passed to Textual
	TxBankKeeper   BankKeeper
	FeeGrantKeeper ante.FeegrantKeeper `optional:"true"`

	CustomSignModeHandlers func() []signing.SignModeHandler `optional:"true"`
}

//nolint:revive
type TxOutputs struct {
	depinject.Out

	TxConfig      client.TxConfig
	BaseAppOption runtime.BaseAppOption
}

func ProvideModule(in TxInputs) TxOutputs {
	textual := NewTextualWithBankKeeper(in.TxBankKeeper)
	var txConfig client.TxConfig
	if in.CustomSignModeHandlers == nil {
		txConfig = tx.NewTxConfigWithTextual(in.ProtoCodecMarshaler, tx.DefaultSignModes, textual)
	} else {
		txConfig = tx.NewTxConfigWithTextual(in.ProtoCodecMarshaler, tx.DefaultSignModes, textual, in.CustomSignModeHandlers()...)
	}

	baseAppOption := func(app *baseapp.BaseApp) {
		// AnteHandlers
		if !in.Config.SkipAnteHandler {
			anteHandler, err := newAnteHandler(txConfig, in)
			if err != nil {
				panic(err)
			}
			app.SetAnteHandler(anteHandler)
		}

		// PostHandlers
		if !in.Config.SkipPostHandler {
			// In v0.46, the SDK introduces _postHandlers_. PostHandlers are like
			// antehandlers, but are run _after_ the `runMsgs` execution. They are also
			// defined as a chain, and have the same signature as antehandlers.
			//
			// In baseapp, postHandlers are run in the same store branch as `runMsgs`,
			// meaning that both `runMsgs` and `postHandler` state will be committed if
			// both are successful, and both will be reverted if any of the two fails.
			//
			// The SDK exposes a default empty postHandlers chain.
			//
			// Please note that changing any of the anteHandler or postHandler chain is
			// likely to be a state-machine breaking change, which needs a coordinated
			// upgrade.
			postHandler, err := posthandler.NewPostHandler(
				posthandler.HandlerOptions{},
			)
			if err != nil {
				panic(err)
			}
			app.SetPostHandler(postHandler)
		}

		// TxDecoder/TxEncoder
		app.SetTxDecoder(txConfig.TxDecoder())
		app.SetTxEncoder(txConfig.TxEncoder())
	}

	return TxOutputs{TxConfig: txConfig, BaseAppOption: baseAppOption}
}

func newAnteHandler(txConfig client.TxConfig, in TxInputs) (sdk.AnteHandler, error) {
	if in.BankKeeper == nil {
		return nil, fmt.Errorf("both AccountKeeper and BankKeeper are required")
	}

	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   in.AccountKeeper,
			BankKeeper:      in.BankKeeper,
			SignModeHandler: txConfig.SignModeHandler(),
			FeegrantKeeper:  in.FeeGrantKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create ante handler: %w", err)
	}

	return anteHandler, nil
}
