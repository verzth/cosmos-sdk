package distribution_test

import (
	"testing"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"gotest.tools/v3/assert"

	simtestutil "github.com/verzth/cosmos-sdk/testutil/sims"
	authkeeper "github.com/verzth/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/verzth/cosmos-sdk/x/auth/types"
	"github.com/verzth/cosmos-sdk/x/distribution/testutil"
	"github.com/verzth/cosmos-sdk/x/distribution/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	var accountKeeper authkeeper.AccountKeeper

	app, err := simtestutil.SetupAtGenesis(testutil.AppConfig, &accountKeeper)
	assert.NilError(t, err)

	ctx := app.BaseApp.NewContext(false, cmtproto.Header{})
	acc := accountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	assert.Assert(t, acc != nil)
}
