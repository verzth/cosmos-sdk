package types

import (
	"github.com/verzth/cosmos-sdk/x/evidence/exported"

	"github.com/verzth/cosmos-sdk/codec"
	"github.com/verzth/cosmos-sdk/codec/legacy"
	"github.com/verzth/cosmos-sdk/codec/types"
	cryptocodec "github.com/verzth/cosmos-sdk/crypto/codec"
	sdk "github.com/verzth/cosmos-sdk/types"
	"github.com/verzth/cosmos-sdk/types/msgservice"
	authzcodec "github.com/verzth/cosmos-sdk/x/authz/codec"
	govcodec "github.com/verzth/cosmos-sdk/x/gov/codec"
	groupcodec "github.com/verzth/cosmos-sdk/x/group/codec"
)

// RegisterLegacyAminoCodec registers all the necessary types and interfaces for the
// evidence module.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*exported.Evidence)(nil), nil)
	legacy.RegisterAminoMsg(cdc, &MsgSubmitEvidence{}, "cosmos-sdk/MsgSubmitEvidence")
	cdc.RegisterConcrete(&Equivocation{}, "cosmos-sdk/Equivocation", nil)
}

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSubmitEvidence{})
	registry.RegisterInterface(
		"cosmos.evidence.v1beta1.Evidence",
		(*exported.Evidence)(nil),
		&Equivocation{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz  and gov Amino codec so that this can later be
	// used to properly serialize MsgGrant, MsgExec and MsgSubmitProposal instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
	RegisterLegacyAminoCodec(govcodec.Amino)
	RegisterLegacyAminoCodec(groupcodec.Amino)
}
