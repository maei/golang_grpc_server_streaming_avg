package app

import "github.com/maei/golang_grpc_server_streaming_avg/grpc_client/src/controller"

func publicRoutes() {
	router.POST("/prime", controller.SumController.GetPrime)
}

func mapUrls() {
	publicRoutes()
}
