package keeper

import (
	"context"
	"testing"

	storetypes "github.com/verzth/cosmos-sdk/store/types"
	sdk "github.com/verzth/cosmos-sdk/types"
)

type baseFixture struct {
	t   *testing.T
	err error
	ctx context.Context

	// k        Keeper //TODO uncomment this after implementing
	addrs    []sdk.AccAddress
	storeKey *storetypes.KVStoreKey
	sdkCtx   sdk.Context
}

func initFixture(t *testing.T) *baseFixture {
	s := &baseFixture{t: t}

	return s
}
