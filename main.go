package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {
	timeNow := time.Now()
	mid := timeNow.Nanosecond()
	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		displayStr := fmt.Sprintf("Hello from machine %d", mid)
		return ctx.String(http.StatusOK, displayStr)
	})
	e.Start(":3000")
}
