// go generate
package main

import (
	"log"

	"github.com/alecthomas/kong"

	grpcservergenerator "github.com/linhbkhn95/grpc-scaffolding/codegen/grpcserver"
	layourgenerator "github.com/linhbkhn95/grpc-scaffolding/codegen/layout"
)

// CLI define list command to generate file.
var CLI struct {
	GrpcServer    grpcservergenerator.ServerData `name:"grpc_server" help:"Generate code server grpc" cmd:"GrpcServer"`
	LayoutProject layourgenerator.TemplateData   `name:"layout" help:"Generate layout project" cmd:"Layout"`
}

var layoutGenerator layourgenerator.Generate

var processorRegistry = map[string]func() error{
	"grpc_server": func() error {
		return grpcservergenerator.Execute(CLI.GrpcServer)
	},
	"layout": func() error {
		if layoutGenerator == nil {
			l, err := layourgenerator.NewGenerator(CLI.LayoutProject)
			if err != nil {
				log.Fatalf("init layoutGenerator error=%v", err)
			}
			layoutGenerator = l

		}
		return layoutGenerator.Generate()
	},
}

func main() {
	ctx := kong.Parse(&CLI)
	err := processorRegistry[ctx.Command()]()
	if err != nil {
		log.Fatalf("error when executed err=%s", err.Error())
	}
}
