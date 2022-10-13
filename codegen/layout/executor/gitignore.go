package executor

import (
	"fmt"
	"text/template"

	"github/linhbkhn95/grpc-scaffolding/codegen/layout/tmpl"
)

const (
	gitIgnoreTmplName = "gitIgnore"
)

type gitIgnoreExecutor struct {
	outputPath string
	t          *template.Template
}

func NewGitIgnoreExecutor(outputPath string) (*gitIgnoreExecutor, error) {
	t := template.New(gitIgnoreTmplName)

	t, err := t.Parse(tmpl.GitignoreStr)
	if err != nil {
		return nil, err
	}
	return &gitIgnoreExecutor{
		t:          t,
		outputPath: outputPath,
	}, nil
}

func (c gitIgnoreExecutor) Execute() error {
	err := executeTemplate(c.t, nil, c.outputPath, true, true)
	if err != nil {
		return fmt.Errorf("error when generated gitignore file cause by %v", err)
	}
	return nil
}
