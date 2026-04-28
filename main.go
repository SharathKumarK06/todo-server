package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var todos []Todo
var Id int = 1

type Todo struct {
	ID          int       `json:"id"`
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	Done        *bool     `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func stringPtr(s string) *string { return &s }

func boolPtr(b bool) *bool { return &b }

func getTodoIdx(id int) (int, error) {
	if id <= 0 {
		return -1, fmt.Errorf("Bad id %d!", id)
	}

	for i := 1; i < len(todos); i++ {
		if id == todos[i].ID {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Todo with id %d not found!", id)
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	if *todo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title required!",
		})
		return
	}

	todo.ID = Id
	Id++
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	todos = append(todos, todo)

	c.JSON(http.StatusCreated, todo)
}

func updateTodo(c *gin.Context) {
	var todo Todo
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	todoIdx, err := getTodoIdx(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Todo with id %d not found!", id),
		})
		return
	}
	oldTodo := &todos[todoIdx]

	// Update only if the field is provided (not nil)
	if todo.Title != nil {
		oldTodo.Title = todo.Title
	}
	if todo.Description != nil {
		oldTodo.Description = todo.Description
	}
	if todo.Done != nil {
		oldTodo.Done = todo.Done
	}
	oldTodo.UpdatedAt = time.Now()

	c.JSON(http.StatusAccepted, oldTodo)
}

func main() {
	todo := Todo{
		ID:          Id,
		Title:       stringPtr("First todo"),
		Description: stringPtr("This is my first test todo"),
		Done:        boolPtr(false),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	todos = append(todos, todo)
	Id++

	router := gin.Default()
	fmt.Println("Server started...")

	// Router
	router.GET("/todos", getTodos)
	router.POST("/todos", createTodo)
	router.PATCH("/todos/update/:id", updateTodo)

	router.Run("localhost:8080")
}
