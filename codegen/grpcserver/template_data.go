package grpcserver

// ServerData define flag of gen grpc server.
type ServerData struct {
	Services        []Service
	TemplateSource  string
	ServiceProtos   []string `name:"service_protos" help:"List service from project grpc" flag:"ServiceProtos" default:"health/v1/service.proto"`
	EnableGateway   bool     `name:"enable_gateway" help:"Option to enable gateway" flag:"EnableGateway" default:"false"`
	OutputPath      string   `name:"output_path" help:"Goal path then generate" flag:"OutputPath" default:"z_server_grpc.go"`
	EnableMetric    bool     `name:"enable_metric" help:"Option to enable metric prometheus" flag:"EnableMetric" default:"true"`
	EnableHTTP      bool     `name:"enable_http" help:"Option expose port http" flag:"EnableHTTP" default:"true"`
	ServiceProtoDir string   `name:"service_proto_dir" help:"Folder contain list service" flag:"ServiceProtoDir" default:"service-proto/proto"`
}
