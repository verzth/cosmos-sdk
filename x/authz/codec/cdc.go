package codec

import (
	"github.com/verzth/cosmos-sdk/codec"
	cryptocodec "github.com/verzth/cosmos-sdk/crypto/codec"
	sdk "github.com/verzth/cosmos-sdk/types"
)

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	cryptocodec.RegisterCrypto(Amino)
	codec.RegisterEvidences(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}
