package main

import (
	"fmt"
	"net/http"
	"slices"
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

type CreateTodoRequest struct {
	Title       *string `json:"title" binding:"required,max=100"`
	Description *string `json:"description"`
}

type UpdateTodoRequest struct {
	Title       *string `json:"title" binding:"required,min=1,max=100"`
	Description *string `json:"description,omitempty"`
	Done        *bool   `json:"done,omitempty"`
}

func stringPtr(s string) *string { return &s }

func boolPtr(b bool) *bool { return &b }

func getTodoIdx(id int) (int, error) {
	if id <= 0 {
		return -1, fmt.Errorf("Bad id %d!", id)
	}

	for i := 0; i < len(todos); i++ {
		if id == todos[i].ID {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Todo with id %d not found!", id)
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func getOneTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	todoIdx, err := getTodoIdx(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, todos[todoIdx])
}

func createTodo(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo := Todo{
		ID:          Id,
		Title:       req.Title,
		Description: req.Description,
		Done:        boolPtr(false),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	Id++
	todos = append(todos, todo)

	c.JSON(http.StatusCreated, todo)
}

// TODO:
//  1. Only update `updated_at` field if any field updated.
//  2. Only fields that can be updated should respond with `202 accepted` others should respond with error.
//  3. `title` field should not be able changed to empty
//  4. Take `string` filed value as text even if it's number
func updateTodo(c *gin.Context) {
	var req UpdateTodoRequest
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todoIdx, err := getTodoIdx(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	oldTodo := &todos[todoIdx]
	updated := false

	// Update only if the field is provided (which is not `nil``)
	if req.Title != nil {
		oldTodo.Title = req.Title
		updated = true
	}
	if req.Description != nil {
		oldTodo.Description = req.Description
		updated = true
	}
	if req.Done != nil {
		oldTodo.Done = req.Done
		updated = true
	}
	if updated {
		oldTodo.UpdatedAt = time.Now()
	}

	c.JSON(http.StatusAccepted, oldTodo)
}

// TODO:
//  1. Deleting id 0, -1 return no error message (err code 404)
func deleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	todoIdx, err := getTodoIdx(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Errorf("Todo with id %d not found!", id),
		})
		return
	}

	todos = slices.Delete(todos, todoIdx, todoIdx+1)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Todo with id %d is successfully deleted.", id),
	})
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
	router.GET("/todos/:id", getOneTodo)
	router.POST("/todos", createTodo)
	router.PATCH("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodo)

	router.Run("localhost:8080")
}
