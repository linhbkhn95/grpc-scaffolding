package grpcserver

type ProtoExtractor interface {
	ExtractService(protoFilePath string) (Service, error)
}

// Service ...
type Service struct {
	ServiceAliasName string
	PackageName      string
	ServiceName      string
	ShortServiceName string
	Methods          []MethodInfo
}

type MethodInfo struct {
	Name         string
	RequestType  string
	ReturnType   string
	CommentLines []string
}
