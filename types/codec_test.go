package types

import (
	"testing"

	"github.com/verzth/cosmos-sdk/collections/colltest"
)

func TestIntValue(t *testing.T) {
	colltest.TestValueCodec(t, IntValue, NewInt(10005994859))
}
