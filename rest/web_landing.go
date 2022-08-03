package rest

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func pingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": os.Getenv("VERSION"),
	})
}
