//go:build gofuzz || go1.18

package tests

import (
	"testing"

	"github.com/verzth/cosmos-sdk/types"
)

func FuzzTypesParseDecCoin(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		types.ParseDecCoin(string(data))
	})
}
