package forwarder

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/deep-quality-dev/cosmos-grpc-forwarder/pkg/grpc/client"
	"github.com/deep-quality-dev/cosmos-grpc-forwarder/pkg/log"
)

func NewCosmosSDKGRPCConn(
	ctx context.Context, logger log.Logger, serverAddr string) (*grpc.ClientConn, error) {
	return client.NewGRPCConn(
		ctx,
		serverAddr,
		[]grpc.UnaryClientInterceptor{
			client.NewLoggingInterceptor(logger),
		},
		// The Cosmos SDK doesn't support any transport security mechanism.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
		// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
}
