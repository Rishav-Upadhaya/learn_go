package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"learn_golang/todo"
	"learn_golang/post"

	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"log"
	// "learn_golang/todo/todo_controller"
)

func main() {
	g := gin.Default()
	// g.Handle("/user/", func)
	g.GET("/", HandleInitialRoute)
	g.GET("/todos", todo.HandleGetAllTodos)
	g.POST("/todo", todo.HandleAddTodo)
	g.DELETE("/todo/:id", todo.HandleDeleteTodo)

	g.GET("/posts", post.HandleReadPost)
	g.POST("/post", post.HandleAddPost)	
	g.DELETE("/post/delete/:id", post.HandleDeletePost)
	g.PUT("/post/update/:id", post.HandleUpdatePost)
	

	dsn := "host=localhost user=rishav password=rishav dbname=gopractice port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Println("Error conecting to database")
		return
	}
	log.Println(db)
	log.Println("Connected to database :5433")
	
	g.Run(":8000")
}

func HandleInitialRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Sucessfully Connected to golang API",
	})
}