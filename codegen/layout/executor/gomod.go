package executor

import (
	"fmt"
	"text/template"

	"github.com/linhbkhn95/grpc-scaffolding/codegen/layout/tmpl"
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
	ModuleName string
}

func NewGomodExecutor(moduleName string,
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
			ModuleName: moduleName,
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
