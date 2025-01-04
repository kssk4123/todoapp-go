package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// ShowHomePage displays the user's TODO list.
func ShowHomePage(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")

	if username == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"username": username,
	})
}

