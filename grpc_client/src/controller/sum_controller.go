package controller

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/labstack/echo/v4"
	"github.com/maei/golang_grpc_server_streaming_avg/grpc_client/src/domain/dto/primepb"
	"github.com/maei/golang_grpc_server_streaming_avg/grpc_client/src/services"
	"github.com/maei/shared_utils_go/logger"
	"io/ioutil"
	"net/http"
)

var SumController sumControllerInterface = &sumController{}

type sumControllerInterface interface {
	GetPrime(c echo.Context) error
}

type sumController struct{}

func (*sumController) GetPrime(c echo.Context) error {
	defer c.Request().Body.Close()
	bodyJSON := c.Request().Body

	bs, err := ioutil.ReadAll(bodyJSON)
	if err != nil {
		logger.Error("decoding byte failed", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "decoding byte failed",
		})
	}

	var prime = &primepb.PrimeNumber{}
	protoErr := jsonpb.UnmarshalString(string(bs), prime)
	if protoErr != nil {
		logger.Error("error unmarshal string to proto", protoErr)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "error unmarshal string to proto",
		})
	}

	res := services.PrimeService.GetPrimeNumber(prime)

	return c.JSON(http.StatusOK, res)
}
