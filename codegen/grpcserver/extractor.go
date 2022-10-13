package grpcserver

import (
	"strings"

	"github.com/emicklei/proto"
)

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

type protoExtractor struct {
}

func NewProtoExtractor() ProtoExtractor {
	return &protoExtractor{}
}

func (p protoExtractor) ExtractService(protoFilePath string) (Service, error) {
	definition, err := parseProto(protoFilePath)
	if err != nil {
		return Service{}, err
	}
	service := Service{
		ServiceAliasName: extractServiceAliasName(protoFilePath),
	}

	proto.Walk(definition,
		proto.WithService(func(s *proto.Service) {
			service.ServiceName = s.Name
		}),
		proto.WithRPC(func(r *proto.RPC) {
			methodInfo := MethodInfo{
				Name:        r.Name,
				RequestType: r.RequestType,
				ReturnType:  r.ReturnsType,
			}
			if r.Comment != nil {
				lines := []string{}
				for _, line := range r.Comment.Lines {
					line = strings.Trim(line, " ")
					if isEligibleCommentLine(line) {
						lines = append(lines, line)
					}
				}
				methodInfo.CommentLines = lines
			}
			service.Methods = append(service.Methods, methodInfo)
		}),
		proto.WithOption(func(s *proto.Option) {
			if len(s.Constant.Source) > 0 {
				service.PackageName = s.Constant.Source
			}
		}))
	service.ShortServiceName = extractShortServiceName(service.ServiceName)
	return service, nil
}

func isEligibleCommentLine(line string) bool {
	return line != "" && line != "\n" && len(line) > 0
}
