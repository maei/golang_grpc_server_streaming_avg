package main

import (
	"fmt"
	"github.com/maei/golang_grpc_server_streaming_avg/grpc_client/src/services"
	"github.com/maei/shared_utils_go/logger"
)

func main() {
	logger.Info(fmt.Sprintf("Starting GRPC-Client listening on port: 50051"))
	services.PrimeService.GetPrimeNumber()
}
