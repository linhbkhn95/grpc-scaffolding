// Code generated by go generate; DO NOT EDIT.
// This file was generated by @generated
// source: codegen/grpcserver/tmpl/server.go.tmpl

package server

import (
	"context"
	"fmt"
	"net"

	http "net/http"

	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"

	examplev1 "github.com/linhbkhn95/rpc-service/go/example/v1"
	healthv1 "github.com/linhbkhn95/rpc-service/go/health/v1"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

type (
	// Server is server struct which contains both grpc and http server.
	Server struct {
		gRPC                 *grpc.Server
		mux                  *runtime.ServeMux
		cfg                  Config
		isRunningDevelopment bool
	}

	// Config hold http/grpc server config
	Config struct {
		GRPC ServerListen `yaml:"grpc" mapstructure:"grpc"`
		HTTP ServerListen `yaml:"http" mapstructure:"http"`
	}

	// ServerListen config for host/port socket listener
	// nolint:revive
	ServerListen struct {
		Host string `yaml:"host" mapstructure:"host"`
		Port int    `yaml:"port" mapstructure:"port"`
	}
)

// DefaultConfig return a default server config
func DefaultConfig() Config {
	return NewConfig(10443, 10080)
}

// NewConfig return a optional config with grpc port and http port.
func NewConfig(grpcPort, httpPort int) Config {
	return Config{
		GRPC: ServerListen{
			Host: "0.0.0.0",
			Port: grpcPort,
		},

		HTTP: ServerListen{
			Host: "0.0.0.0",
			Port: httpPort,
		},
	}
}

// String return socket listen DSN
func (l ServerListen) String() string {
	return fmt.Sprintf("%s:%d", l.Host, l.Port)
}

func NewServer(cfg Config, isRunningDevelopment bool, opt ...grpc.ServerOption) *Server {
	return &Server{
		gRPC: grpc.NewServer(opt...),
		mux: runtime.NewServeMux(
			runtime.WithMarshalerOption(runtime.MIMEWildcard,
				&runtime.JSONPb{
					MarshalOptions: protojson.MarshalOptions{
						UseProtoNames:   false,
						UseEnumNumbers:  false,
						EmitUnpopulated: true,
					},
					UnmarshalOptions: protojson.UnmarshalOptions{
						DiscardUnknown: true,
					},
				})),
		cfg:                  cfg,
		isRunningDevelopment: isRunningDevelopment,
	}
}
func (s *Server) Register(grpcServer ...interface{}) error {
	for _, srv := range grpcServer {
		switch _srv := srv.(type) {

		case healthv1.HealthServiceServer:
			healthv1.RegisterHealthServiceServer(s.gRPC, _srv)
			if err := healthv1.RegisterHealthServiceHandlerFromEndpoint(context.Background(), s.mux, s.cfg.GRPC.String(), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
				return err
			}

		case examplev1.ExampleServiceServer:
			examplev1.RegisterExampleServiceServer(s.gRPC, _srv)
			if err := examplev1.RegisterExampleServiceHandlerFromEndpoint(context.Background(), s.mux, s.cfg.GRPC.String(), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
				return err
			}

		default:
			return fmt.Errorf("Unknown GRPC Service to register %#v", srv)
		}
	}
	return nil
}

// Serve server listen for HTTP and GRPC
func (s *Server) Serve() error {
	stop := make(chan os.Signal, 1)
	errch := make(chan error)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	httpMux := http.NewServeMux()

	httpMux.Handle("/metrics", promhttp.Handler())

	httpMux.Handle("/", s.mux)
	httpServer := http.Server{
		Addr:    s.cfg.HTTP.String(),
		Handler: httpMux,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			errch <- err
		}
	}()

	go func() {
		listener, err := net.Listen("tcp", s.cfg.GRPC.String())
		if err != nil {
			errch <- err
			return
		}
		if err := s.gRPC.Serve(listener); err != nil {
			errch <- err
		}
	}()
	for {
		select {
		case <-stop:
			ctx, cancelFn := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancelFn()
			s.gRPC.GracefulStop()

			if err := httpServer.Shutdown(ctx); err != nil {
				fmt.Println("failed to stop server: %w", err)
			}

			if !s.isRunningDevelopment {
				fmt.Println("Shutting down. Wait for 15 seconds")
				time.Sleep(15 * time.Second)
			}
			return nil
		case err := <-errch:
			return err
		}
	}
}