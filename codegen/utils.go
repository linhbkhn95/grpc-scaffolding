package codegen

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const (
	defaultPerm = 0644
	ShellToUse  = "bash"
)

// FormatSourceCode to format code.
func FormatSourceCode(source []byte) ([]byte, error) {
	return format.Source(source)
}

// WriteToFile to write byte data to filePath.
func WriteToFile(filePath string, buf []byte) error {
	if err := ioutil.WriteFile(filePath, buf, defaultPerm); err != nil {
		return fmt.Errorf("error when write file to %s: %w", filePath, err)
	}
	return nil
}

// CheckGoVet to find errors not caught by the compiler.
func CheckGoVet(filePath string) error {
	cmd := exec.Command("go", "vet", filePath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		os.Remove(filePath)
		return fmt.Errorf("error when run go vet to check file %s: %s", filePath, string(out))
	}
	return nil
}

func IsDirExisted(dirPath string) bool {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func InstallServiceProto() error {
	err, _, _ := Shellout("export GOPRIVATE=github.com/KyberNetwork/*")
	if err != nil && !os.IsExist(err) {
		log.Printf("error: %v\n", err)
		return err
	}
	log.Println("clone repo git@github.com:linhbkhn95/rpc-proto.git ...")
	err, _, stderr := Shellout("git clone git@github.com:linhbkhn95/rpc-proto.git")
	if err != nil && !os.IsExist(err) {
		log.Printf("error: %s %v\n", stderr, err)
		return err
	}
	return nil
}

func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
