package client

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/deep-quality-dev/cosmos-grpc-forwarder/pkg/jsonconv"
	"github.com/deep-quality-dev/cosmos-grpc-forwarder/pkg/log"
)

// NewLoggingInterceptor is a gRPC client interceptor for logging requests, responses and errors.
func NewLoggingInterceptor(logger log.Logger, jsonConverter *jsonconv.JSONConverter) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req any,
		reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		start := time.Now()
		errResp := invoker(ctx, method, req, reply, cc, opts...)
		duration := time.Since(start)

		reqJSON, err := jsonConverter.Marshal(req)
		if err != nil {
			logger.Error("error: request decoding: ", log.Error(errors.WithStack(err)))
		}

		respJSON, err := jsonConverter.Marshal(reply)
		if err != nil {
			logger.Error("error: response decoding: ", log.Error(errors.WithStack(err)))
		}

		logger.Print("outgoing gRPC request",
			log.String("method", method),
			log.String("request", string(reqJSON)),
			log.String("response", string(respJSON)),
			log.Error(errResp),
			log.Float64("duration", duration.Seconds()),
		)

		return errResp
	}
}
