package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	sdkmath "github.com/verzth/cosmos-sdk/math"
	"github.com/verzth/cosmos-sdk/x/evidence/simulation"
	"github.com/verzth/cosmos-sdk/x/evidence/types"

	"github.com/verzth/cosmos-sdk/codec"
	codectypes "github.com/verzth/cosmos-sdk/codec/types"
	"github.com/verzth/cosmos-sdk/types/module"
	simtypes "github.com/verzth/cosmos-sdk/types/simulation"
)

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abonormal scenarios are not tested here.
func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: sdkmath.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	var evidenceGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &evidenceGenesis)

	require.Len(t, evidenceGenesis.Evidence, 0)
}
