package baseapp

import (
	storetypes "github.com/verzth/cosmos-sdk/store/types"

	sdk "github.com/verzth/cosmos-sdk/types"
)

type state struct {
	ms  storetypes.CacheMultiStore
	ctx sdk.Context
}

// CacheMultiStore calls and returns a CacheMultiStore on the state's underling
// CacheMultiStore.
func (st *state) CacheMultiStore() storetypes.CacheMultiStore {
	return st.ms.CacheMultiStore()
}

// Context returns the Context of the state.
func (st *state) Context() sdk.Context {
	return st.ctx
}
