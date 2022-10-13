package tmpl

import _ "embed"

//go:embed .gitignore.tmpl
var GitignoreStr string

//go:embed go.mod.tmpl
var GomodStr string

//go:embed Makefile.tmpl
var MakefileStr string

//go:embed README.md.tmpl
var ReadmeStr string

//go:embed Dockerfile.tmpl
var DockerfileStr string

//go:embed .envrc.example.tmpl
var EnvStr string

//go:embed config/config.go.tmpl
var ConfigStr string

//go:embed server/server.go.tmpl
var ServerStr string

//go:embed server/grpc/server.go.tmpl
var GRPCServerStr string

//go:embed server/handler/service.go.tmpl
var ServerHandlerStr string

//go:embed cmd/main.go.tmpl
var CMDStr string
