package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var todos []Todo

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func createTodos(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}
	todo.ID = todos[len(todos)-1].ID + 1
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	todos = append(todos, todo)

	c.JSON(http.StatusAccepted, todo)
}

func main() {
	todo := Todo{
		ID:          1,
		Title:       "First todo",
		Description: "This is my first test todo",
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	todos = append(todos, todo)

	router := gin.Default()
	fmt.Println("Server started...")

	// Router
	router.GET("/todos", getTodos)
	router.POST("/todos/create", createTodos)

	router.Run("localhost:8080")
}
