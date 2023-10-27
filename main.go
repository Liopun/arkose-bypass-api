package main

import (
	"net/http"
	"os"

	"github.com/flyingpot/funcaptcha"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Token string `json:"message,omitempty"`
	Error string `json:"error,omitempty"`
}

func main() {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "debug" // default to development mode
	}
	gin.SetMode(ginMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// CORS middleware configuration
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	r.Use(cors.New(config))

	r.GET("/token", func(c *gin.Context) {
		token, err := funcaptcha.GetOpenAITokenV2()
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Error: err.Error(),
			})
		}

		c.JSON(http.StatusOK, Response{
			Token: token,
		})
	})

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	r.Run(":" + port)
}
