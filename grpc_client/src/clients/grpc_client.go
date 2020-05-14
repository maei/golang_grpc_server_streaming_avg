package clients

import (
	"github.com/maei/shared_utils_go/logger"
	"google.golang.org/grpc"
)

var GRPCClient grpcClientInterface = &grpcClient{}

type grpcClientInterface interface {
	SetClient() (*grpc.ClientConn, error)
}

type grpcClient struct{}

const (
	ServerHost = "localhost:50051"
)

func (*grpcClient) SetClient() (*grpc.ClientConn, error) {
	logger.Info("starting gRPC Client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		logger.Error("could not connect", err)
		return nil, err
	}

	return conn, nil
}
