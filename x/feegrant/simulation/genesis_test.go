package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	sdkmath "github.com/verzth/cosmos-sdk/math"
	moduletypes "github.com/verzth/cosmos-sdk/types/module"
	moduletestutil "github.com/verzth/cosmos-sdk/types/module/testutil"
	simtypes "github.com/verzth/cosmos-sdk/types/simulation"
	"github.com/verzth/cosmos-sdk/x/feegrant"
	"github.com/verzth/cosmos-sdk/x/feegrant/module"
	"github.com/verzth/cosmos-sdk/x/feegrant/simulation"
)

func TestRandomizedGenState(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(module.AppModuleBasic{})
	s := rand.NewSource(1)
	r := rand.New(s)

	accounts := simtypes.RandomAccounts(r, 3)

	simState := moduletypes.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          encCfg.Codec,
		Rand:         r,
		NumBonded:    3,
		Accounts:     accounts,
		InitialStake: sdkmath.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)
	var feegrantGenesis feegrant.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[feegrant.ModuleName], &feegrantGenesis)

	require.Len(t, feegrantGenesis.Allowances, len(accounts)-1)
}
