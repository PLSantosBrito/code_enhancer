package main

import (
	"github.com/gin-gonic/gin"
	"github.com/PLSantosBrito/server/internal/auth/handler"
)

func main() {
	r := gin.Default()

	authHandler := handler.NewAuthHandler()

	auth := r.Group("/auth")
	{
		auth.GET("/github/callback", authHandler.LoginCallback)
	}

	r.Run(":8080")
}
