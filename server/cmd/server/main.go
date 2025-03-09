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
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}

	r.Run(":8080")
}
