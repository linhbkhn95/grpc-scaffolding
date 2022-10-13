package executor

import (
	"fmt"
	"text/template"

	"github/linhbkhn95/grpc-scaffolding/codegen/layout/tmpl"
)

const (
	cmdTmplName = "cmd"
)

type cmdExecutor struct {
	outputPath string
	t          *template.Template
	tmplData   cmdTmplData
}

type cmdTmplData struct {
	ProjectName string
}

func NewCmdExecutor(projectName, outputPath string) (*cmdExecutor, error) {
	t := template.New(cmdTmplName)

	t, err := t.Parse(tmpl.CMDStr)
	if err != nil {
		return nil, err
	}
	return &cmdExecutor{
		t:          t,
		outputPath: outputPath,
		tmplData: cmdTmplData{
			ProjectName: projectName,
		},
	}, nil
}

func (c cmdExecutor) Execute() error {
	err := executeTemplate(c.t, c.tmplData, c.outputPath, true, false)
	if err != nil {
		return fmt.Errorf("error when generated cmd/main.go file cause by %v", err)
	}
	return nil
}
