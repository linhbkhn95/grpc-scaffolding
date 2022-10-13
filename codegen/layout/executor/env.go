package executor

import (
	"fmt"
	"text/template"

	"github/linhbkhn95/grpc-scaffolding/codegen/layout/tmpl"
)

const (
	envTmplName = "env"
)

type envExecutor struct {
	outputPath string
	t          *template.Template
}

func NewENVExecutor(outputPath string) (*envExecutor, error) {
	t := template.New(envTmplName)

	t, err := t.Parse(tmpl.EnvStr)
	if err != nil {
		return nil, err
	}
	return &envExecutor{
		t:          t,
		outputPath: outputPath,
	}, nil
}

func (c envExecutor) Execute() error {
	err := executeTemplate(c.t, nil, c.outputPath, true, true)
	if err != nil {
		return fmt.Errorf("error when generated env file cause by %v", err)
	}
	return nil
}
