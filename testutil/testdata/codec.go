package testdata

import (
	amino "github.com/tendermint/go-amino"

	"github.com/verzth/cosmos-sdk/codec/types"
	sdk "github.com/verzth/cosmos-sdk/types"
	"github.com/verzth/cosmos-sdk/types/msgservice"
	tx "github.com/verzth/cosmos-sdk/types/tx"
)

func NewTestInterfaceRegistry() types.InterfaceRegistry {
	registry := types.NewInterfaceRegistry()
	RegisterInterfaces(registry)
	return registry
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &TestMsg{})

	registry.RegisterInterface("Animal", (*Animal)(nil))
	registry.RegisterImplementations(
		(*Animal)(nil),
		&Dog{},
		&Cat{},
	)
	registry.RegisterImplementations(
		(*HasAnimalI)(nil),
		&HasAnimal{},
	)
	registry.RegisterImplementations(
		(*HasHasAnimalI)(nil),
		&HasHasAnimal{},
	)
	registry.RegisterImplementations(
		(*tx.ExtensionOptionI)(nil),
		&Cat{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func NewTestAmino() *amino.Codec {
	cdc := amino.NewCodec()
	cdc.RegisterInterface((*Animal)(nil), nil)
	cdc.RegisterConcrete(&Dog{}, "testpb/Dog", nil)
	cdc.RegisterConcrete(&Cat{}, "testpb/Cat", nil)

	cdc.RegisterInterface((*HasAnimalI)(nil), nil)
	cdc.RegisterConcrete(&HasAnimal{}, "testpb/HasAnimal", nil)

	cdc.RegisterInterface((*HasHasAnimalI)(nil), nil)
	cdc.RegisterConcrete(&HasHasAnimal{}, "testpb/HasHasAnimal", nil)

	return cdc
}
