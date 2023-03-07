package keeper

import (
	sdk "github.com/verzth/cosmos-sdk/types"
	"github.com/verzth/cosmos-sdk/x/gov/types"
)

func (keeper Keeper) GetConstitution(ctx sdk.Context) (constitution string) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(types.KeyConstitution)

	return string(bz)
}

func (keeper Keeper) SetConstitution(ctx sdk.Context, constitution string) {
	store := ctx.KVStore(keeper.storeKey)
	store.Set(types.KeyConstitution, []byte(constitution))
}
