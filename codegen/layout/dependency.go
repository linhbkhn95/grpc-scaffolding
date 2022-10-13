package layout

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github/linhbkhn95/grpc-scaffolding/codegen"
)

func (g generator) installDependence() error {
	err := os.Chdir(g.projectName)
	if err != nil {
		log.Printf("error when cd to %s folder \n", g.projectName)
		return err
	}
	log.Println("installing dependencies ...")

	var processes []func() error
	processes = append(processes, g.setPrivateRepo())
	processes = append(processes, g.installViper(), g.installKitchen(), g.installGRPCServiceInternal())
	processes = append(processes, g.installGRPCEcosystem()...)
	processes = append(processes, g.runGomodTidy())
	for _, process := range processes {
		if err := process(); err != nil {
			return err
		}
	}

	log.Println("install successfully ...")
	return nil
}

func (g generator) runGomodTidy() func() error {
	return func() error {
		log.Println("run go mod tidy ...")

		cmd := exec.Command("go", "mod", "tidy")
		stdout, err := cmd.CombinedOutput()
		log.Println(string(stdout))
		if err != nil {
			return fmt.Errorf("error when install dependency with err=%s", err.Error())
		}
		return nil
	}
}

func (g generator) installViper() func() error {
	return func() error {
		log.Println("installing viper...")
		cmd := exec.Command("go", "get", "github.com/spf13/viper")
		_, err := cmd.CombinedOutput()
		return err
	}
}

func (g generator) setPrivateRepo() func() error {
	return func() error {
		log.Println("setting private repo ...")

		err, _, stderr := codegen.Shellout("export GOPRIVATE=github.com/KyberNetwork/*")
		if err != nil {
			log.Printf("error: %s ,%v\n", stderr, err)
		}
		err, _, stderr = codegen.Shellout("export GIT_TERMINAL_PROMPT=1")
		if err != nil {
			log.Printf("error: %s ,%v\n", stderr, err)
		}

		return err

	}
}

func (g generator) installKitchen() func() error {
	return func() error {
		log.Println("installing github.com/KyberNetwork/kitchen ...")
		err, _, stderr := codegen.Shellout("GOSUMDB=off go get github.com/KyberNetwork/kitchen ")
		if err != nil {
			log.Printf("error: %s ,%v\n", stderr, err)
		}
		return err
	}
}

func (g generator) installGRPCServiceInternal() func() error {
	return func() error {
		log.Println("installing github.com/KyberNetwork/grpc-service ...")
		err, _, stderr := codegen.Shellout("GOSUMDB=off go get github.com/KyberNetwork/grpc-service ")
		if err != nil {
			log.Printf("error: %s %v\n", stderr, err)
		}
		return err
	}
}

func (g generator) installGRPCEcosystem() []func() error {
	var ecosystems []func() error
	log.Println("installing grpc ecosystem ...")
	ecosystems = append(ecosystems, func() error {
		log.Println("installing google.golang.org/grpc ...")
		cmd := exec.Command("go", "get", "google.golang.org/grpc")
		_, err := cmd.CombinedOutput()
		return err
	})
	ecosystems = append(ecosystems, func() error {
		log.Println("installing github.com/grpc-ecosystem/go-grpc-middleware ...")
		cmd := exec.Command("go", "get", "github.com/grpc-ecosystem/go-grpc-middleware")
		_, err := cmd.CombinedOutput()
		return err
	})

	ecosystems = append(ecosystems, func() error {
		log.Println("installing github.com/grpc-ecosystem/grpc-gateway/v2/runtime ...")
		cmd := exec.Command("go", "get", "github.com/grpc-ecosystem/grpc-gateway/v2/runtime")
		_, err := cmd.CombinedOutput()
		return err
	})

	ecosystems = append(ecosystems, func() error {
		log.Println("installing github.com/prometheus/client_golang/prometheus/promhttp ...")
		cmd := exec.Command("go", "get", "github.com/prometheus/client_golang/prometheus/promhttp")
		_, err := cmd.CombinedOutput()
		return err
	})

	ecosystems = append(ecosystems, func() error {
		log.Println("installing google.golang.org/grpc ...")
		cmd := exec.Command("go", "get", "google.golang.org/grpc")
		_, err := cmd.CombinedOutput()
		return err
	})

	ecosystems = append(ecosystems, func() error {
		log.Println("installing google.golang.org/grpc/credentials/insecure ...")
		cmd := exec.Command("go", "get", "google.golang.org/grpc/credentials/insecure")
		_, err := cmd.CombinedOutput()
		return err
	})

	ecosystems = append(ecosystems, func() error {
		log.Println("installing google.golang.org/protobuf/encoding/protojson ...")
		cmd := exec.Command("go", "get", "google.golang.org/protobuf/encoding/protojson")
		_, err := cmd.CombinedOutput()
		return err
	})

	ecosystems = append(ecosystems, func() error {
		log.Println("installing github.com/grpc-ecosystem/go-grpc-prometheus ...")
		cmd := exec.Command("go", "get", "github.com/grpc-ecosystem/go-grpc-prometheus")
		_, err := cmd.CombinedOutput()
		return err
	})
	return ecosystems
}
