package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	githubTokenURL   = "https://github.com/login/oauth/access_token"
	githubGetUserURL = "https://api.github.com/user"
)

const (
	clientID     = "Ov23lim2s3OwlqW3ydmn"
	clientSecret = "0ae2914a5ef8aeea460ea2b3eaf06fc24f400518"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) LoginCallback(c *gin.Context) {
	code, exists := c.GetQuery("code")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No code",
		})
		return
	}

	body := map[string]string{
		"code":          code,
		"client_id":     clientID,
		"client_secret": clientSecret,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't create check request",
		})
		return
	}

	req, err := http.NewRequest("POST", githubTokenURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't create request",
		})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't verify code",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error getting the checker response: " + strconv.Itoa(resp.StatusCode),
		})
		return
	}

	bodyResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error getting the checker response: %v", err),
		})
		return
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	if err := json.Unmarshal(bodyResponse, &tokenResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error getting the checker response %v", err),
		})
		return
	}

	accessToken := tokenResponse.AccessToken
	fmt.Println("==>")
	fmt.Printf("Token Response: %+v\n", tokenResponse)

	// TODO: STORE ACCESS TOKENS AT DB
	req, err = http.NewRequest("GET", githubGetUserURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get user data",
		})
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get user data",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get user data",
		})
		return
	}

	userBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error reading user data",
		})
		return
	}

	var userResponse struct {
		Name      string `json:"name"`
		URL       string `json:"url"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	}

	if err := json.Unmarshal(userBody, &userResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Couldn't get user data",
		})
		return
	}

	fmt.Printf("User data: %+v\n", userResponse)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenResponse.AccessToken,
		"user":    userResponse,
	})
}
