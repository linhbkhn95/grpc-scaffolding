package executor

import (
	"fmt"
	"text/template"

	"github.com/linhbkhn95/grpc-scaffolding/codegen/layout/tmpl"
)

const (
	cfgTmplName = "config"
)

type configExecutor struct {
	outputPath string
	t          *template.Template
	tmplData   configTmplData
}

type configTmplData struct {
	ProjectName string
}

func NewConfigExecutor(projectName, outputPath string) (*configExecutor, error) {

	t := template.New(cfgTmplName)
	t, err := t.Parse(tmpl.ConfigStr)
	if err != nil {
		return nil, err
	}
	return &configExecutor{
		t:          t,
		outputPath: outputPath,
		tmplData: configTmplData{
			ProjectName: projectName,
		},
	}, nil
}

func (c configExecutor) Execute() error {
	err := executeTemplate(c.t, c.tmplData, c.outputPath, true, false)
	if err != nil {
		return fmt.Errorf("error when generated config file cause by %v", err)
	}
	return nil
}
