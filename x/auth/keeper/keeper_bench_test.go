package keeper_test

import (
	"testing"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/require"
	simtestutil "github.com/verzth/cosmos-sdk/testutil/sims"
	sdk "github.com/verzth/cosmos-sdk/types"
	"github.com/verzth/cosmos-sdk/x/auth/keeper"
	"github.com/verzth/cosmos-sdk/x/auth/testutil"
)

func BenchmarkAccountMapperGetAccountFound(b *testing.B) {
	b.ReportAllocs()
	var accountKeeper keeper.AccountKeeper
	app, err := simtestutil.Setup(testutil.AppConfig, &accountKeeper)
	require.NoError(b, err)

	ctx := app.BaseApp.NewContext(false, cmtproto.Header{})

	// assumes b.N < 2**24
	for i := 0; i < b.N; i++ {
		arr := []byte{byte((i & 0xFF0000) >> 16), byte((i & 0xFF00) >> 8), byte(i & 0xFF)}
		addr := sdk.AccAddress(arr)
		acc := accountKeeper.NewAccountWithAddress(ctx, addr)
		accountKeeper.SetAccount(ctx, acc)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := []byte{byte((i & 0xFF0000) >> 16), byte((i & 0xFF00) >> 8), byte(i & 0xFF)}
		accountKeeper.GetAccount(ctx, sdk.AccAddress(arr))
	}
}

func BenchmarkAccountMapperSetAccount(b *testing.B) {
	b.ReportAllocs()
	var accountKeeper keeper.AccountKeeper
	app, err := simtestutil.Setup(testutil.AppConfig, &accountKeeper)
	require.NoError(b, err)

	ctx := app.BaseApp.NewContext(false, cmtproto.Header{})

	b.ResetTimer()

	// assumes b.N < 2**24
	for i := 0; i < b.N; i++ {
		arr := []byte{byte((i & 0xFF0000) >> 16), byte((i & 0xFF00) >> 8), byte(i & 0xFF)}
		addr := sdk.AccAddress(arr)
		acc := accountKeeper.NewAccountWithAddress(ctx, addr)
		accountKeeper.SetAccount(ctx, acc)
	}
}
