package main

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/verzth/cosmos-sdk/orm/internal/codegen"
)

func main() {
	protogen.Options{}.Run(codegen.QueryProtoPluginRunner)
}
