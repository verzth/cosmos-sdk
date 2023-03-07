package gov_test

import (
	"testing"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"gotest.tools/v3/assert"

	"github.com/verzth/cosmos-sdk/testutil/configurator"
	simtestutil "github.com/verzth/cosmos-sdk/testutil/sims"
	authkeeper "github.com/verzth/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/verzth/cosmos-sdk/x/auth/types"
	_ "github.com/verzth/cosmos-sdk/x/distribution"
	"github.com/verzth/cosmos-sdk/x/gov/types"
	_ "github.com/verzth/cosmos-sdk/x/mint"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	var accountKeeper authkeeper.AccountKeeper
	app, err := simtestutil.SetupAtGenesis(
		configurator.NewAppConfig(
			configurator.ParamsModule(),
			configurator.AuthModule(),
			configurator.StakingModule(),
			configurator.BankModule(),
			configurator.GovModule(),
			configurator.DistributionModule(),
			configurator.ConsensusModule(),
		),
		&accountKeeper,
	)
	assert.NilError(t, err)

	ctx := app.BaseApp.NewContext(false, cmtproto.Header{})
	acc := accountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	assert.Assert(t, acc != nil)
}
