package codec

import (
	"testing"

	"github.com/cosmos/gogoproto/types"
	codectypes "github.com/verzth/cosmos-sdk/codec/types"
	"github.com/verzth/cosmos-sdk/collections/colltest"
)

func TestCollectionsCorrectness(t *testing.T) {
	cdc := NewProtoCodec(codectypes.NewInterfaceRegistry())
	t.Run("CollValue", func(t *testing.T) {
		colltest.TestValueCodec(t, CollValue[types.UInt64Value](cdc), types.UInt64Value{
			Value: 500,
		})
	})
}
