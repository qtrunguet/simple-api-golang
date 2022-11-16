package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func main() {

	router := gin.Default()

	port := "9090"

	router.GET("/register", func(c *gin.Context) {
		// dang ky user
		c.String(http.StatusOK, "dang ky user")
	})

	router.Run(":" + port)
}
