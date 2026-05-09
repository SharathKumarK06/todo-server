package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SharathKumarK06/todo-server/models"
	"github.com/SharathKumarK06/todo-server/repository"
	"github.com/gin-gonic/gin"
)

var repo = repository.NewInMemoryRepo()

func getTodos(c *gin.Context) {
	todos, err := repo.List()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, todos)
}

func getOneTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo, err := repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, todo)
}

func createTodo(c *gin.Context) {
	var req models.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo := models.Todo{}
	if *req.Title != "" {
		todo.Title = *req.Title
	}
	if req.Description != nil {
		todo.Description = *req.Description
	}
	todo, err := repo.Create(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, todo)
}

func updateTodo(c *gin.Context) {
	var req models.UpdateTodoRequest
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

	todo, err := repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Update only if the field is provided (which is not `nil``)
	if req.Title != nil && *req.Title != "" && *req.Title != todo.Title {
		todo.Title = *req.Title
	}
	if req.Description != nil && *req.Description != todo.Description {
		todo.Description = *req.Description
	}
	if req.Done != nil && *req.Done != todo.Done {
		todo.Done = *req.Done
	}
	todo, err = repo.Update(id, todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, todo)
}

func deleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := repo.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func main() {

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
