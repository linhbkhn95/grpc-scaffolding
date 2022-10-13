package grpcserver

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	protoreader "github.com/emicklei/proto"
)

const (
	defaultPerm = 0644
)

// parseProto read proto file, then parse it.
func parseProto(protoPath string) (*protoreader.Proto, error) {
	reader, err := os.Open(protoPath)
	defer func() {
		if err := reader.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		return nil, err
	}
	parser := protoreader.NewParser(reader)
	return parser.Parse()
}

// getServiceFullPaths return full path where contains proto file.
// example: dir: share/grpc/proto, servicePaths: ['health/v1/heath.proto'] => [share/grpc/proto/health/v1/heath.proto]
func getServiceFullPaths(dir string, servicePaths []string) []string {
	for i, service := range servicePaths {
		servicePaths[i] = dir + "/" + service
	}
	return servicePaths
}

// extractServiceAliasName return service name + version like  healthV1
// example: filPath: heath/v1/heath.proto => result : [...,"health", "v1", ...] => return healthV1
func extractServiceAliasName(servicePath string) string {
	results := strings.Split(servicePath, "/")
	len := len(results)
	return results[len-3] + strings.ToLower(results[len-2])
}

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

// extractShortServiceName return short service name.
// example: serviceName: HealthService => return Health
func extractShortServiceName(serviceName string) string {
	results := strings.Split(serviceName, "Service")
	return results[0]
}
