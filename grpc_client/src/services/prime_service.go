package services

import (
	"context"
	"fmt"
	"github.com/maei/golang_grpc_server_streaming_avg/grpc_client/src/clients"
	"github.com/maei/golang_grpc_server_streaming_avg/grpc_client/src/domain/dto/primepb"
	"github.com/maei/shared_utils_go/logger"
	"io"
)

var PrimeService primeServiceInterface = &primeService{}

type primeServiceInterface interface {
	GetPrimeNumber()
}

type primeService struct{}

func (*primeService) GetPrimeNumber() {
	conn, clientErr := clients.GRPCClient.SetClient()
	if clientErr != nil {
		logger.Error("error creating grpc client", clientErr)
	}
	client := primepb.NewPrimeNumberServiceClient(conn)

	req := &primepb.PrimeNumberRequest{
		PrimeNumber: &primepb.PrimeNumber{
			PrimeNumber: 1000,
		},
	}

	resStream, errStream := client.GetPrimeNumbers(context.Background(), req)
	if errStream != nil {
		logger.Error("error fetching data from grpc_stream", errStream)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error("error wil fetching messages from stream", err)
		}
		fmt.Printf("Response from stream %v\n", msg)
	}
}
