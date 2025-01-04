package controllers

import (
        "todo2/config"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AddTodo(c *gin.Context) {
        session := sessions.Default(c)
        userID := session.Get("user_id")
        if userID == nil {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
                return
        }

	type TodoRequest struct {
		// UserID int    `json:"user_id" binding:"required"`
		Title  string `json:"title" binding:"required"`
	}

	var req TodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := config.DB.Exec("INSERT INTO todos (user_id, title) VALUES (?, ?)", userID, req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo created successfully"})
}

func GetTodos(c *gin.Context) {
        session := sessions.Default(c)
        userID := session.Get("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	rows, err := config.DB.Query("SELECT id, title, is_completed FROM todos WHERE user_id = ?", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
                return
	}
	defer rows.Close()

	var todos []gin.H
	for rows.Next() {
		var id int
		var title string
		var isCompleted bool
		if err := rows.Scan(&id, &title, &isCompleted); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse todos"})
		}
		todos = append(todos, gin.H{"id": id, "title": title, "is_completed": isCompleted})
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

// DELETE /api/todos/:id
func DeleteTodo(c *gin.Context) {
        todoID := c.Param("id")

        // Todoを削除するSQLクエリ
    _, err := config.DB.Exec("DELETE FROM todos WHERE id = ?", todoID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

// PUT /api/todos/:id
func UpdateTodoCompletion(c *gin.Context) {
    todoID := c.Param("id")
    var request struct {
        IsCompleted bool `json:"is_completed"`
    }
    
    // リクエストの内容をバインド
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // 完了状態を更新するSQLクエリ
    _, err := config.DB.Exec("UPDATE todos SET is_completed = ? WHERE id = ?", request.IsCompleted, todoID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}
