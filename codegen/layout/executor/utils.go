package executor

import (
	"bytes"
	"log"
	"regexp"
	"strings"
	"text/template"

	"github/linhbkhn95/grpc-scaffolding/codegen"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func executeTemplate(t *template.Template, tmplData interface{}, outputPath string, skipGoVet bool, skipFormat bool) error {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, tmplData)
	if err != nil {
		return err
	}
	var buf = tmplBytes.Bytes()
	if !skipFormat {
		// format code before write to file.
		buf, err = codegen.FormatSourceCode(tmplBytes.Bytes())
		if err != nil {
			return err
		}
	}

	// write code generate to file.
	err = codegen.WriteToFile(outputPath, buf)
	if err != nil {
		return err
	}
	if !skipGoVet {
		// use go vet to can find errors not caught by the compilers.
		err = codegen.CheckGoVet(outputPath)
		if err != nil {
			log.Println(string(buf))
			return err
		}
	}
	log.Printf("File was generate at %s\n", outputPath)
	return nil
}

// GetFilenameWithExt ... example : path: /server.go.tmpl -> return server.go.tmpl
func GetFilenameWithExt(path string) string {
	results := strings.Split(path, "/")
	len := len(results)
	return results[len-1]
}

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)

}

// extractServiceName will extract service name from service name with snake case like health_service => health
func extractServiceName(serviceName string) string {
	terms := strings.Split(serviceName, "_")
	return strings.Join(terms[:len(terms)-1], "_")
}
