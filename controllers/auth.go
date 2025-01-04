package controllers

import (
	"net/http"
	"todo2/config"
	"todo2/models"
        "log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// ShowRegisterPage renders the registration form.
func ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

// RegisterUser handles user registration.
// /api/register
func RegisterUser(c *gin.Context) {
	type RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
                log.Println("invalid input")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID, err := models.CreateUser(config.DB, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
        }
        log.Println("userID")
        log.Println(userID)
// 登録後にセッションを設定
	session := sessions.Default(c)
	session.Set("username", req.Username)
	session.Set("user_id", userID)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

// ShowRegisterPage renders the registration form.
func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}


// LoginUser handles user login.
// /api/login
func LoginUser(c *gin.Context) {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	success, err := models.AuthenticateUser(config.DB, req.Username, req.Password)
	if err != nil || !success {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

        log.Println(req.Username, req.Password)
	// ユーザー認証（例: データベースから照合）
	var userID int
	err = config.DB.QueryRow("SELECT id FROM users WHERE username = ?", req.Username).Scan(&userID)
        log.Println("userID", userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	session := sessions.Default(c)
	session.Set("username", req.Username)
	session.Set("user_id", userID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// LogoutUser logs out the user.
// /api/logout
func LogoutUser(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}
