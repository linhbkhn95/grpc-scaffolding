package executor

import (
	"fmt"
	"text/template"

	"github/linhbkhn95/grpc-scaffolding/codegen/layout/tmpl"
)

const (
	gomodTmplName = "gomod"
)

type gomodExecutor struct {
	outputPath string
	t          *template.Template
	tmplData   gomodTmplData
}

type gomodTmplData struct {
	ProjectName string
}

func NewGomodExecutor(projectName string,
	outputPath string) (*gomodExecutor, error) {
	t := template.New(gomodTmplName)

	t, err := t.Parse(tmpl.GomodStr)
	if err != nil {
		return nil, err
	}
	return &gomodExecutor{
		t:          t,
		outputPath: outputPath,
		tmplData: gomodTmplData{
			ProjectName: projectName,
		},
	}, nil
}

func (c gomodExecutor) Execute() error {
	err := executeTemplate(c.t, c.tmplData, c.outputPath, true, true)
	if err != nil {
		return fmt.Errorf("error when generated gomod file cause by %v", err)
	}
	return nil
}
