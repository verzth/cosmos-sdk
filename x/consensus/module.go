package consensus

import (
	"context"
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	modulev1 "github.com/verzth/cosmos-sdk/api/cosmos/consensus/module/v1"
	"github.com/verzth/cosmos-sdk/core/appmodule"
	"github.com/verzth/cosmos-sdk/depinject"

	store "github.com/verzth/cosmos-sdk/store/types"

	"github.com/verzth/cosmos-sdk/baseapp"
	"github.com/verzth/cosmos-sdk/client"
	"github.com/verzth/cosmos-sdk/codec"
	codectypes "github.com/verzth/cosmos-sdk/codec/types"
	"github.com/verzth/cosmos-sdk/runtime"
	sdk "github.com/verzth/cosmos-sdk/types"
	"github.com/verzth/cosmos-sdk/types/module"
	authtypes "github.com/verzth/cosmos-sdk/x/auth/types"
	"github.com/verzth/cosmos-sdk/x/consensus/keeper"
	"github.com/verzth/cosmos-sdk/x/consensus/types"
	govtypes "github.com/verzth/cosmos-sdk/x/gov/types"
)

// ConsensusVersion defines the current x/consensus module consensus version.
const ConsensusVersion = 1

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the consensus_param module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the consensus_param module's name.
func (AppModuleBasic) Name() string { return types.ModuleName }

// RegisterLegacyAminoCodec registers the consensus_param module's types on the LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the consensus_param
// module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	// nil is returned since default genesis of consensus params is handled by tendermint
	return nil
}

// ValidateGenesis performs genesis state validation
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	return nil
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// GetTxCmd returns the root tx command
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

// GetQueryCmd returns no root query command
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

// RegisterInterfaces registers interfaces and implementations of the bank module.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// AppModule implements an application module
type AppModule struct {
	AppModuleBasic

	keeper keeper.Keeper
}

var _ appmodule.AppModule = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), keeper.NewQuerier(am.keeper))
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
	}
}

// Name returns the consensus_param module's name.
func (AppModule) Name() string { return types.ModuleName }

// InitGenesis is handled by for init genesis of consensus_param
func (am AppModule) InitGenesis(sdk.Context, codec.JSONCodec, json.RawMessage) []abci.ValidatorUpdate {
	// nil is returned since initgenesis of consensus params is handled by tendermint
	return nil
}

// ExportGenesis is handled by CometBFT export of genesis
func (am AppModule) ExportGenesis(sdk.Context, codec.JSONCodec) json.RawMessage {
	// nil is returned since ExportGenesis of consensus params is handled by CometBFT and baseapp
	return nil
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

// RegisterInvariants does nothing, there are no invariants to enforce
func (am AppModule) RegisterInvariants(sdk.InvariantRegistry) {}

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

//nolint:revive
type ConsensusInputs struct {
	depinject.In

	Config *modulev1.Module
	Cdc    codec.Codec
	Key    *store.KVStoreKey
}

//nolint:revive
type ConsensusOutputs struct {
	depinject.Out

	Keeper        keeper.Keeper
	Module        appmodule.AppModule
	BaseAppOption runtime.BaseAppOption
}

func ProvideModule(in ConsensusInputs) ConsensusOutputs {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	k := keeper.NewKeeper(in.Cdc, in.Key, authority.String())
	m := NewAppModule(in.Cdc, k)
	baseappOpt := func(app *baseapp.BaseApp) {
		app.SetParamStore(&k)
	}

	return ConsensusOutputs{
		Keeper:        k,
		Module:        m,
		BaseAppOption: baseappOpt,
	}
}
