package server

import (
	"fmt"
	"github.com/maei/golang_grpc_server_streaming_avg/grpc_server/src/domain/dto/primepb"
	"github.com/maei/shared_utils_go/logger"
	"google.golang.org/grpc"
	"net"
)

type server struct{}

var (
	s = grpc.NewServer()
)

func (*server) GetPrimeNumbers(req *primepb.PrimeNumberRequest, stream primepb.PrimeNumberService_GetPrimeNumbersServer) error {
	number := req.GetPrimeNumber().GetPrimeNumber()
	divisor := int32(2)
	logger.Info("streaming-started")

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&primepb.PrimeNumberResponse{Result: divisor})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v\n", divisor)
		}
	}

	logger.Info("streaming-ended")
	return nil
}

func StartGRPCServer() {
	logger.Info(fmt.Sprintf("Starting GRPC-Server on port: 50051"))

	//lis, err := net.Listen("tcp", os.Getenv("SERVER_PORT"))
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		logger.Error("error while listening gRPC Server", err)
	}

	primepb.RegisterPrimeNumberServiceServer(s, &server{})

	errServer := s.Serve(lis)
	if errServer != nil {
		logger.Error("error while serve gRPC Server", errServer)
	}
}
