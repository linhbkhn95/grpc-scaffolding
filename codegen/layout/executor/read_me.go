package executor

import (
	"fmt"
	"text/template"

	"github.com/linhbkhn95/grpc-scaffolding/codegen/layout/tmpl"
)

const (
	readmeTmplName = "readme"
)

type readmeExecutor struct {
	outputPath string
	t          *template.Template
	tmplData   readmeTmplData
}

type readmeTmplData struct {
	ProjectName string
}

func NewReadmeExecutor(projectName string,
	outputPath string) (*readmeExecutor, error) {
	t := template.New(readmeTmplName)

	t, err := t.Parse(tmpl.ReadmeStr)
	if err != nil {
		return nil, err
	}
	return &readmeExecutor{
		t:          t,
		outputPath: outputPath,
		tmplData: readmeTmplData{
			ProjectName: projectName,
		},
	}, nil
}

func (c readmeExecutor) Execute() error {
	err := executeTemplate(c.t, c.tmplData, c.outputPath, true, true)
	if err != nil {
		return fmt.Errorf("error when generated readme file cause by %v", err)
	}
	return nil
}
