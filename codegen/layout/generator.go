package layout

//go:generate go run ../../main.go layout --project_name=example

import (
	"fmt"
	"log"
	"os"

	"github.com/linhbkhn95/grpc-scaffolding/codegen"
	"github.com/linhbkhn95/grpc-scaffolding/codegen/grpcserver"
	"github.com/linhbkhn95/grpc-scaffolding/codegen/layout/executor"
)

type TemplateData struct {
	ProjectName      string `name:"project_name" help:"Project's Name" flag:"project-name" default:"example"`
	ModuleName       string `name:"module_name" help:"Project's Name" flag:"module-name" default:"github.com/linhbkhn95/example"`
	EnablePrometheus bool   `name:"enable_prometheus" help:"Project should inject prometheus to collect metric" flag:"enable-prometheus" default:"true"`
	grpcserver.ServerData
}

type Generate interface {
	Generate() error
}

type Executor interface {
	Execute() error
}

type generator struct {
	executors                 []Executor
	projectName               string
	shouldInstallServiceProto bool
}

// TODO: REFACTOR CODE: separate to many small function.
func NewGenerator(tmplData TemplateData) (*generator, error) {
	var shouldInstallServiceProto bool
	if !codegen.IsDirExisted(tmplData.RPCProtoDir) {
		shouldInstallServiceProto = true
		log.Println("rpc proto dir is not existed!")
		if err := codegen.InstallServiceProto(); err != nil {
			return nil, err
		}
	}
	executors := []Executor{}
	err := ensureDir(tmplData.ProjectName)
	if err != nil {
		return nil, err
	}
	err = os.Chmod(tmplData.ProjectName, os.ModePerm)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	gomodOutputPath := fmt.Sprintf("%s/%s", tmplData.ProjectName, "go.mod")
	gomodExecutor, err := executor.NewGomodExecutor(tmplData.ModuleName, gomodOutputPath)
	if err != nil {
		return nil, err
	}

	configDir := fmt.Sprintf("%s/config", tmplData.ProjectName)
	err = ensureDir(configDir)

	if err != nil {
		return nil, err
	}

	configOutputPath := fmt.Sprintf("%s/%s", configDir, "config.go")
	configExecutor, err := executor.NewConfigExecutor(tmplData.ProjectName, configOutputPath)
	if err != nil {
		return nil, err
	}

	serverDir := fmt.Sprintf("%s/server", tmplData.ProjectName)
	err = ensureDir(serverDir)

	if err != nil {
		return nil, err
	}

	grpcServerTmplData, err := grpcserver.NewServerData(grpcserver.ServerData{
		RPCProtos:     tmplData.RPCProtos,
		EnableGateway: tmplData.EnableGateway,
		EnableMetric:  tmplData.EnableMetric,
		EnableHTTP:    tmplData.EnableHTTP,
		RPCProtoDir:   tmplData.RPCProtoDir,
	})
	if err != nil {
		return nil, fmt.Errorf("error when create server data cause by %s", err.Error())
	}

	serverOutputPath := fmt.Sprintf("%s/%s", serverDir, "server.go")
	serverExecutor, err := executor.NewServerExecutor(tmplData.ProjectName, tmplData.EnablePrometheus, grpcServerTmplData.Services, serverOutputPath)
	if err != nil {
		return nil, fmt.Errorf("error when generate server grpc instance cause by %s", err.Error())
	}

	grpcServerDir := fmt.Sprintf("%s/grpc", serverDir)
	err = ensureDir(grpcServerDir)

	if err != nil {
		return nil, err
	}

	grpcServerOutputPath := fmt.Sprintf("%s/%s", grpcServerDir, "server.go")
	grpcServerTmplData.OutputPath = grpcServerOutputPath
	grpcServerExecutor := executor.NewGRPCServerExecutor(grpcServerTmplData)

	serverHandlerDir := fmt.Sprintf("%s/handler", serverDir)
	err = ensureDir(serverHandlerDir)
	if err != nil {
		return nil, err
	}
	serverHandlerExecutor, err := executor.NewServerHandlerExecutor(serverHandlerDir, grpcServerTmplData.Services)
	if err != nil {
		return nil, err
	}

	cmdDir := fmt.Sprintf("%s/cmd", tmplData.ProjectName)
	err = ensureDir(cmdDir)

	if err != nil {
		return nil, err
	}
	cmdOutputPath := fmt.Sprintf("%s/%s", cmdDir, "main.go")
	cmdExecutor, err := executor.NewCmdExecutor(tmplData.ProjectName, cmdOutputPath)
	if err != nil {
		return nil, err
	}

	readmeFileOutputPath := fmt.Sprintf("%s/%s", tmplData.ProjectName, "README.md")
	readmeExecutor, err := executor.NewReadmeExecutor(tmplData.ProjectName, readmeFileOutputPath)
	if err != nil {
		return nil, err
	}
	envFileOutputPath := fmt.Sprintf("%s/%s", tmplData.ProjectName, ".envrc.example")
	envExecutor, err := executor.NewENVExecutor(envFileOutputPath)
	if err != nil {
		return nil, err
	}

	gitIgnoreFileOutputPath := fmt.Sprintf("%s/%s", tmplData.ProjectName, ".gitignore")
	gitIgnoreExecutor, err := executor.NewGitIgnoreExecutor(gitIgnoreFileOutputPath)
	if err != nil {
		return nil, err
	}

	makeFileOutputPath := fmt.Sprintf("%s/%s", tmplData.ProjectName, "Makefile")
	makeFileExecutor, err := executor.NewMakeFileExecutor(makeFileOutputPath)

	if err != nil {
		return nil, err
	}
	dockerFileOutputPath := fmt.Sprintf("%s/%s", tmplData.ProjectName, "Dockerfile")
	dockerExecutor, err := executor.NewDockerExecutor(dockerFileOutputPath)
	if err != nil {
		return nil, err
	}
	if err = createEmptyFolder(tmplData.ProjectName); err != nil {
		return nil, err
	}
	executors = append(
		executors,
		configExecutor,
		serverExecutor,
		grpcServerExecutor,
		serverHandlerExecutor,
		cmdExecutor,
		readmeExecutor,
		envExecutor,
		gitIgnoreExecutor,
		gomodExecutor,
		makeFileExecutor,
		dockerExecutor,
		gitIgnoreExecutor)
	return &generator{
		executors:                 executors,
		projectName:               tmplData.ProjectName,
		shouldInstallServiceProto: shouldInstallServiceProto,
	}, nil
}

func (g generator) Generate() error {
	if g.shouldInstallServiceProto {
		defer tearDown()
	}
	for _, executor := range g.executors {
		if err := executor.Execute(); err != nil {
			return err
		}
	}

	log.Println("Generated completed!")

	return g.installDependence()
}

func createEmptyFolder(projectName string) error {
	pkgDir := fmt.Sprintf("%s/pkg", projectName)
	err := ensureDir(pkgDir)
	if err != nil {
		return err
	}

	internalDir := fmt.Sprintf("%s/internal", projectName)
	err = ensureDir(internalDir)
	if err != nil {
		return err
	}
	return nil
}

func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

func tearDown() {
	log.Println("remove rpc-proto folder")
	err, _, _ := codegen.Shellout("cd .. &&rm -rf rpc-proto")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}
