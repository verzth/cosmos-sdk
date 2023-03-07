package keeper_test

import (
	"testing"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/suite"

	storetypes "github.com/verzth/cosmos-sdk/store/types"

	"github.com/verzth/cosmos-sdk/baseapp"
	"github.com/verzth/cosmos-sdk/testutil"
	sdk "github.com/verzth/cosmos-sdk/types"
	moduletestutil "github.com/verzth/cosmos-sdk/types/module/testutil"
	authtypes "github.com/verzth/cosmos-sdk/x/auth/types"
	consensusparamkeeper "github.com/verzth/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/verzth/cosmos-sdk/x/consensus/types"
	govtypes "github.com/verzth/cosmos-sdk/x/gov/types"
)

type KeeperTestSuite struct {
	suite.Suite
	ctx                   sdk.Context
	consensusParamsKeeper *consensusparamkeeper.Keeper

	queryClient consensusparamtypes.QueryClient
	msgServer   consensusparamtypes.MsgServer
}

func (s *KeeperTestSuite) SetupTest() {
	key := storetypes.NewKVStoreKey(consensusparamtypes.StoreKey)
	testCtx := testutil.DefaultContextWithDB(s.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(cmtproto.Header{})
	encCfg := moduletestutil.MakeTestEncodingConfig()

	keeper := consensusparamkeeper.NewKeeper(encCfg.Codec, key, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	s.ctx = ctx
	s.consensusParamsKeeper = &keeper

	consensusparamtypes.RegisterInterfaces(encCfg.InterfaceRegistry)
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	consensusparamtypes.RegisterQueryServer(queryHelper, consensusparamkeeper.NewQuerier(keeper))
	s.queryClient = consensusparamtypes.NewQueryClient(queryHelper)
	s.msgServer = consensusparamkeeper.NewMsgServerImpl(keeper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
