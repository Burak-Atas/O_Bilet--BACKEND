package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello")

		q := c.Request.URL.Query()
		c.JSON(http.StatusOK, q)
	})

	router.Run()
}
