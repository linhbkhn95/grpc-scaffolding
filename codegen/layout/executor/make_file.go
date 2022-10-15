package executor

import (
	"fmt"
	"text/template"

	"github.com/linhbkhn95/grpc-scaffolding/codegen/layout/tmpl"
)

const (
	makeFileTmplName = "makeFile"
)

type makeFileExecutor struct {
	outputPath string
	t          *template.Template
}

func NewMakeFileExecutor(outputPath string) (*makeFileExecutor, error) {
	t := template.New(makeFileTmplName)

	t, err := t.Parse(tmpl.MakefileStr)
	if err != nil {
		return nil, err
	}
	return &makeFileExecutor{
		t:          t,
		outputPath: outputPath,
	}, nil
}

func (c makeFileExecutor) Execute() error {
	err := executeTemplate(c.t, nil, c.outputPath, true, true)
	if err != nil {
		return fmt.Errorf("error when generated make file cause by %v", err)
	}
	return nil

}
