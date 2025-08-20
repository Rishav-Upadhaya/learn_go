package main

import (
	"net/http"

	"learn_golang/database/myquery"
	"learn_golang/handlers"
	"learn_golang/models"
	"learn_golang/post"
	"learn_golang/todo"

	"github.com/gin-gonic/gin"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	//auth route
	g.POST("/auth/register", handlers.RegisterUser)
	g.GET("/auth/users", handlers.HandleGetAllUsers)

	dsn := "host=localhost user=rishav password=rishav dbname=gopractice port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Error conecting to database")
		return
	}
	log.Println(db)
	log.Println("Connected to database :5433")

	// migrate the User model
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println("Error migrating User model:", err)
		return
	}

	//setting db on myquery
	myquery.SetDefault(db)

	g.Run(":8000")
}

func HandleInitialRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Sucessfully Connected to golang API",
	})
}
