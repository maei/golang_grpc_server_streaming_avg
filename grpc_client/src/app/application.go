package app

import (
	"github.com/labstack/echo/v4"
)

var (
	router = echo.New()
)

func StartApplication() {
	mapUrls()
	router.Logger.Fatal(router.Start(":8002"))
}
