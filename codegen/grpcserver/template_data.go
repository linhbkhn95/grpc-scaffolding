package grpcserver

// ServerData define flag of gen grpc server.
type ServerData struct {
	Services       []Service
	TemplateSource string
	RPCProtos      []string `name:"rpc_protos" help:"List services from rpc rpc or local folder" flag:"RPCProtos" default:"health/v1/service.proto"`
	EnableGateway  bool     `name:"enable_gateway" help:"Option to enable gateway" flag:"EnableGateway" default:"false"`
	OutputPath     string   `name:"output_path" help:"Goal path then generate" flag:"OutputPath" default:"z_server_grpc.go"`
	EnableMetric   bool     `name:"enable_metric" help:"Option to enable metric prometheus" flag:"EnableMetric" default:"true"`
	EnableHTTP     bool     `name:"enable_http" help:"Option expose port http" flag:"EnableHTTP" default:"true"`
	RPCProtoDir    string   `name:"rpc_proto_dir" help:"Folder contain list service" flag:"RPCProtoDir" default:"rpc-proto/proto"`
}

// NewServerData to return data use for template
func NewServerData(serverData ServerData) (ServerData, error) {

	servicePaths := getServiceFullPaths(serverData.RPCProtoDir, serverData.RPCProtos)

	extractor := NewProtoExtractor()
	for _, path := range servicePaths {
		service, err := extractor.ExtractService(path)
		if err != nil {
			return ServerData{}, err
		}
		serverData.Services = append(serverData.Services, service)
	}
	return serverData, nil
}
