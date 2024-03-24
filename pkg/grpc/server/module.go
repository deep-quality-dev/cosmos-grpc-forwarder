package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/deep-quality-dev/cosmos-grpc-forwarder/pkg/configs"
	"github.com/deep-quality-dev/cosmos-grpc-forwarder/pkg/jsonconv"
	"github.com/deep-quality-dev/cosmos-grpc-forwarder/pkg/log"
)

// InitialiazeNewGRPCServer initializes the gRPC server module.
func InitialiazeNewGRPCServer(
	ctx context.Context,
	conf *configs.Config,
	logger log.Logger,
	jsonConverter *jsonconv.JSONConverter,
) *Server {
	serverAddress := fmt.Sprintf("%s:%d", conf.ServerHost, conf.ServerPort)

	lis, err := NewListener(serverAddress)
	if err != nil {
		logger.Panic("error: cannot create server listener: ", log.Error(err))
	}

	return NewGRPCServer(
		conf.ServerName,
		serverAddress,
		lis,
		logger,
		[]grpc.UnaryServerInterceptor{
			NewLoggingInterceptor(logger, jsonConverter),
		},
	)
}
