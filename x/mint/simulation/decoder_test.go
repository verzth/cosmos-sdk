package simulation_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/verzth/cosmos-sdk/math"

	"github.com/verzth/cosmos-sdk/types/kv"
	moduletestutil "github.com/verzth/cosmos-sdk/types/module/testutil"
	"github.com/verzth/cosmos-sdk/x/mint"
	"github.com/verzth/cosmos-sdk/x/mint/simulation"
	"github.com/verzth/cosmos-sdk/x/mint/types"
)

func TestDecodeStore(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(mint.AppModuleBasic{})

	dec := simulation.NewDecodeStore(encCfg.Codec)

	minter := types.NewMinter(math.LegacyOneDec(), math.LegacyNewDec(15))

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: types.MinterKey, Value: encCfg.Codec.MustMarshal(&minter)},
			{Key: []byte{0x99}, Value: []byte{0x99}},
		},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Minter", fmt.Sprintf("%v\n%v", minter, minter)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
