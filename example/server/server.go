package server

import (
	"runtime/debug"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcvalidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"

	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/linhbkhn95/golang-british/logger"

	"github.com/linhbkhn95/example/config"
	"github.com/linhbkhn95/golang-british/appmode"
	"github.com/linhbkhn95/golang-british/grpc/middleware/grpcerror"

	"github.com/linhbkhn95/example/server/grpc"
	"github.com/linhbkhn95/example/server/handler"
)

func Serve() {

	// init connection
	// inject dependency

	loggerInstance, err := logger.InitLogger(config.Log, logger.LoggerBackendZap)
	if err != nil {
		logger.Fatalf("Error when initing logger cause by %v", err)
	}
	zapLogger, err := logger.GetDesugaredZapLoggerDelegate(loggerInstance)
	if err != nil {
		logger.Fatalf("Error when getting zapLogger cause by %v", err)
	}

	grpczap.ReplaceGrpcLoggerV2(zapLogger)

	grpcprometheus.EnableHandlingTimeHistogram()

	internalServerErr := status.Error(codes.Internal, "Something went wrong in our side.")

	recoveryOpt := grpcrecovery.WithRecoveryHandler(func(err interface{}) error {
		logger.WithFields(logger.Fields{"error": err, "stacktrace": string(debug.Stack())}).Error("unexpected error...")
		return internalServerErr
	})
	unaryOpts := []grpc.UnaryServerInterceptor{

		grpcprometheus.UnaryServerInterceptor,

		grpcctxtags.UnaryServerInterceptor(grpcctxtags.WithFieldExtractor(grpcctxtags.CodeGenRequestFieldExtractor)),
		grpcvalidator.UnaryServerInterceptor(),
		grpcrecovery.UnaryServerInterceptor(recoveryOpt),
		grpczap.UnaryServerInterceptor(zapLogger),
		grpcerror.UnaryServerInterceptor(config.Mode == appmode.Development, internalServerErr),
	}

	streamOpts := []grpc.StreamServerInterceptor{

		grpcprometheus.StreamServerInterceptor,

		grpcctxtags.StreamServerInterceptor(grpcctxtags.WithFieldExtractor(grpcctxtags.CodeGenRequestFieldExtractor)),
		grpcvalidator.StreamServerInterceptor(),
		grpcrecovery.StreamServerInterceptor(recoveryOpt),
		grpczap.StreamServerInterceptor(zapLogger),
	}

	s := server.NewServer(config.Server, config.Mode == appmode.Development,
		grpcmiddleware.WithUnaryServerChain(unaryOpts...),
		grpcmiddleware.WithStreamServerChain(streamOpts...),
	)
	healthv1Server := handler.NewHealthServer()
	examplev1Server := handler.NewExampleServer()

	// Register your services here.

	if err := s.Register(
		healthv1Server,
		examplev1Server,
	); err != nil {
		logger.Fatalf("Error register servers %v", err)
	}

	logger.WithFields(logger.Fields{"grpc_addr": config.Server.GRPC.Host}).
		WithFields(logger.Fields{"grpc_port": config.Server.GRPC.Port}).
		WithFields(logger.Fields{"http_addr": config.Server.HTTP.Host}).
		WithFields(logger.Fields{"http_port": config.Server.HTTP.Port}).
		Info("Starting server...")
	if err := s.Serve(); err != nil {
		logger.Fatalf("Error start server %v", err)
	}

}
