// go generate
package main

import (
	"log"

	"github.com/alecthomas/kong"

	grpcservergenerator "github/linhbkhn95/grpc-scaffolding/codegen/grpcserver"
)

// CLI define list command to generate file.
var CLI struct {
	GrpcServer grpcservergenerator.ServerData `name:"grpc_server" help:"Generate code server grpc" cmd:"GrpcServer"`
}

var processorRegistry = map[string]func() error{
	"grpc_server": func() error {
		return grpcservergenerator.Execute(CLI.GrpcServer)
	},
}

func main() {
	ctx := kong.Parse(&CLI)
	err := processorRegistry[ctx.Command()]()
	if err != nil {
		log.Fatalf("error when executed err=%s", err.Error())
	}
}
