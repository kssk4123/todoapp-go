package main

import (
        "todo2/config"
        "todo2/controllers"
        "net/http"

        "github.com/gin-contrib/sessions"
        "github.com/gin-contrib/sessions/cookie"
        "github.com/gin-gonic/gin"
)

func main() {
        config.InitDB()

        r := gin.Default()
        r.Static("/css", "./static/css")
        r.Static("/js", "./static/js")
        r.Static("/icon", "./static/icon")

        store := cookie.NewStore([]byte("secret"))
        store.Options(sessions.Options{
                Path:     "/",
                MaxAge:   3600, // セッションの有効期限
                HttpOnly: true, // JavaScriptからクッキーをアクセス不可にする
                Secure:   true, // HTTPS経由でのみ送信
        })
        r.Use(sessions.Sessions("session", store))

        r.LoadHTMLGlob("templates/*")

        r.GET("/", func(c *gin.Context) {
                c.Redirect(http.StatusFound, "/home")
        })
        r.GET("/register", controllers.ShowRegisterPage)
        r.GET("/login", controllers.ShowLoginPage)
        r.GET("/home", controllers.ShowHomePage)
        r.GET("/logout", controllers.LogoutUser)
        r.POST("/api/register", controllers.RegisterUser)
        r.POST("/api/login", controllers.LoginUser)
        r.GET("/api/todos", controllers.GetTodos)
        r.POST("/api/todos", controllers.AddTodo)
        r.DELETE("/api/todos/:id", controllers.DeleteTodo)
        r.PUT("/api/todos/:id", controllers.UpdateTodoCompletion)
        r.DELETE("/api/delete-user", controllers.DeleteUser)

        r.Run(":8080")
}
