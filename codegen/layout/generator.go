package layout

//go:generate go run ../../main.go layout --project_name=example

import "github/linhbkhn95/grpc-scaffolding/codegen/grpcserver"

type TemplateData struct {
	ProjectName      string `name:"project_name" help:"Project's Name" flag:"project-name" default:"example"`
	EnablePrometheus bool   `name:"enable_prometheus" help:"Project should inject prometheus to collect metric" flag:"enable-prometheus" default:"true"`
	grpcserver.ServerData
}

type Generate interface {
	Generate() error
}

type Executor interface {
	Execute() error
}
