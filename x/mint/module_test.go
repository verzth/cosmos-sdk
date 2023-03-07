package mint_test

import (
	"testing"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/require"

	simtestutil "github.com/verzth/cosmos-sdk/testutil/sims"
	authkeeper "github.com/verzth/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/verzth/cosmos-sdk/x/auth/types"
	"github.com/verzth/cosmos-sdk/x/mint/testutil"
	"github.com/verzth/cosmos-sdk/x/mint/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	var accountKeeper authkeeper.AccountKeeper

	app, err := simtestutil.SetupAtGenesis(testutil.AppConfig, &accountKeeper)
	require.NoError(t, err)

	ctx := app.BaseApp.NewContext(false, cmtproto.Header{})
	acc := accountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	require.NotNil(t, acc)
}
