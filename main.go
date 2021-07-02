package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	"go.linka.cloud/protoc-gen-go-fields/fields"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		fields.Module(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}


