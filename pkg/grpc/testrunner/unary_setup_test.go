package testrunner_test

import (
	"context"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"

	"github.com/deep-quality-dev/cosmos-grpc-forwarder/pkg/configs"
	"github.com/deep-quality-dev/cosmos-grpc-forwarder/pkg/grpc/testrunner"
	"github.com/deep-quality-dev/cosmos-grpc-forwarder/pkg/jsonconv"
	"github.com/deep-quality-dev/cosmos-grpc-forwarder/pkg/log"
)

func TestClientConnection(t *testing.T) {
	ctx := context.Background()

	_ = godotenv.Load("../../../.env.test.dist")

	conf := configs.InitializeConfig()

	logger := log.InitializeLogger(conf.LogLevel, conf.LogFormat)

	jsonConverter := jsonconv.NewJSONConverter()

	config := testrunner.NewDefaultTestConfig(logger, conf, jsonConverter)

	conn, closer, err := testrunner.NewUnaryTestSetup(ctx, config)

	if err != nil {
		t.Error(err)
	}

	if conn == nil {
		t.Error(errors.New("connection cannot be nil"))
	}

	closer()
}
