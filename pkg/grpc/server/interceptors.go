package server

import (
	"context"

	"google.golang.org/grpc"

	"github.com/deep-quality-dev/cosmos-grpc-forwarder/pkg/log"
)

func NewLoggingInterceptor(logger log.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		return nil, nil
	}
}
