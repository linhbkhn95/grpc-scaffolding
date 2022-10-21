package grpcserver

//go:generate go run ../../main.go grpc_server --rpc_protos=health/v1/service.proto

import (
	"bytes"
	"log"
	"text/template"

	"github.com/linhbkhn95/grpc-scaffolding/codegen"

	"github.com/linhbkhn95/grpc-scaffolding/codegen/grpcserver/tmpl"
)

const (
	defaultServerTemplatePath = "codegen/grpcserver/tmpl/server.go.tmpl"
	defaultTemplateName       = "Template server"
)

// Execute to generate code grpc server with flag.
func Execute(serverData ServerData) error {
	if !codegen.IsDirExisted(serverData.RPCProtoDir) {
		defer tearDown()
		if err := codegen.InstallServiceProto(); err != nil {
			return err
		}
	}
	t := template.New(defaultTemplateName)

	t, err := t.Parse(tmpl.ServerStr)
	if err != nil {
		return err
	}
	templateData := serverData
	if len(serverData.Services) == 0 {
		templateData, err = NewServerData(serverData)
		if err != nil {
			return err
		}
	}

	templateData.TemplateSource = defaultServerTemplatePath

	var tmplBytes bytes.Buffer

	err = t.Execute(&tmplBytes, templateData)
	if err != nil {
		return err
	}

	// format code before write to file.
	buf, err := FormatSourceCode(tmplBytes.Bytes())
	if err != nil {
		log.Println(string(buf))
		return err
	}

	// write code generate to file.
	err = WriteToFile(serverData.OutputPath, buf)
	if err != nil {
		log.Println(string(buf))
		return err
	}
	log.Printf("File was generate at %s\n", serverData.OutputPath)
	return nil
}

func tearDown() {
	log.Println("remove service-proto folder")
	err, _, _ := codegen.Shellout("rm -rf service-proto")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}
