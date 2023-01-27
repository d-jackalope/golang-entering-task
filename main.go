package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/handlers"
)

func main() {

	r := gin.Default()

	r.GET("/when/:year", checkPingMiddleware, handlers.GetYear)

	r.Run(":8000")
}

func checkPingMiddleware(c *gin.Context) {

	_, err := strconv.ParseUint(c.Params.ByName("year"), 10, 32)
	if err != nil {
		return
	}

	if _, ok := c.Request.Header["X-Ping"]; ok {
		c.Header("X-Pong", "pong")
	}

}
