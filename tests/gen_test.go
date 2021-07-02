package tests

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/stretchr/testify/require"

	"go.linka.cloud/protoc-gen-go-fields/fields"
)

func TestDebugGen(t *testing.T) {
	require := require.New(t)
	f, err := os.Open("code_generator_request.pb.bin")
	require.NoError(err)
	defer f.Close()
	out := &bytes.Buffer{}
	pgs.Init(
		pgs.ProtocInput(f),
		pgs.ProtocOutput(out),
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		fields.Module(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
	fmt.Println(out.String())
}
