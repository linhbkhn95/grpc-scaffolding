package executor

import (
	"fmt"
	"text/template"

	"github/linhbkhn95/grpc-scaffolding/codegen/layout/tmpl"
)

const (
	dockerTmplName = "docker"
)

type dockerExecutor struct {
	outputPath string
	t          *template.Template
}

func NewDockerExecutor(outputPath string) (*dockerExecutor, error) {
	t := template.New(dockerTmplName)

	t, err := t.Parse(tmpl.DockerfileStr)
	if err != nil {
		return nil, err
	}
	return &dockerExecutor{
		t:          t,
		outputPath: outputPath,
	}, nil
}

func (c dockerExecutor) Execute() error {
	err := executeTemplate(c.t, nil, c.outputPath, true, true)
	if err != nil {
		return fmt.Errorf("error when generated docker file cause by %v", err)
	}
	return nil
}
