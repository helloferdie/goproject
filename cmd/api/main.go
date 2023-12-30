package main

import (
	"fmt"
	"spun/pkg/libecho/libmiddleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	godotenv.Load()
}

func main() {
	fmt.Println("spun v3")

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(libmiddleware.Logger())

	e.Logger.Fatal(e.Start(":1323"))
}
