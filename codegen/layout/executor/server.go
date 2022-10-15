package executor

import (
	"fmt"
	"text/template"

	"github.com/linhbkhn95/grpc-scaffolding/codegen/grpcserver"
	"github.com/linhbkhn95/grpc-scaffolding/codegen/layout/tmpl"
)

const (
	serverTmplName = "server"

	serverHandlerTmplName = "handler"
)

type serverExecutor struct {
	outputPath string
	t          *template.Template
	tmplData   serverTmplData
}

type serverTmplData struct {
	ProjectName      string
	EnablePrometheus bool
	Services         []grpcserver.Service
}

func NewServerExecutor(projectName string, enablePrometheus bool, services []grpcserver.Service, outputPath string) (*serverExecutor, error) {
	t := template.New(serverTmplName)

	t, err := t.Parse(tmpl.ServerStr)
	if err != nil {
		return nil, err
	}
	return &serverExecutor{
		t:          t,
		outputPath: outputPath,
		tmplData: serverTmplData{
			ProjectName:      projectName,
			EnablePrometheus: enablePrometheus,
			Services:         services,
		},
	}, nil
}

func (c serverExecutor) Execute() error {
	return executeTemplate(c.t, c.tmplData, c.outputPath, true, false)
}

type serverHandlerExecutor struct {
	handlerDir string
	services   []grpcserver.Service
	t          *template.Template
}

func NewServerHandlerExecutor(handlerDir string, services []grpcserver.Service) (*serverHandlerExecutor, error) {
	t := template.New(serverHandlerTmplName)

	t, err := t.Parse(tmpl.ServerHandlerStr)
	if err != nil {
		return nil, err
	}
	return &serverHandlerExecutor{
		t:          t,
		services:   services,
		handlerDir: handlerDir,
	}, nil
}

func (c serverHandlerExecutor) Execute() error {
	for _, service := range c.services {
		outputPath := fmt.Sprintf("%s/%s.go", c.handlerDir, extractServiceName(ToSnakeCase(service.ServiceName)))
		if err := executeTemplate(c.t, service, outputPath, true, false); err != nil {
			return err
		}
	}
	return nil
}

type grpcServerExecutor struct {
	tmplData grpcserver.ServerData
}

func NewGRPCServerExecutor(tmplData grpcserver.ServerData) *grpcServerExecutor {
	return &grpcServerExecutor{
		tmplData,
	}
}

func (c grpcServerExecutor) Execute() error {
	return grpcserver.Execute(c.tmplData)
}
