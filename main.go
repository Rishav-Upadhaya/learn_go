package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"learn_golang/todo"
	"learn_golang/post"
	// "learn_golang/todo/todo_controller"
)

func main() {
	g := gin.Default()
	// g.Handle("/user/", func)
	g.GET("/", HandleInitialRoute)
	g.GET("/todos", todo.HandleGetAllTodos)
	g.POST("/todo", todo.HandleAddTodo)
	g.DELETE("/todo/:id", todo.HandleDeleteTodo)
	g.Run(":8000")
}

func HandleInitialRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Sucessfully Connected to golang API",
	})
}