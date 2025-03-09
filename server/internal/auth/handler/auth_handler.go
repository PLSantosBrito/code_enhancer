package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type AuthHandler struct {}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *gin.Context) {
	//TODO: Implement login logic
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token": "mock-jwt-token",
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	//TODO: Implement register logic
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
	})
} 