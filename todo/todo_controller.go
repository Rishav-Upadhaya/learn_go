package todo

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleGetAllTodos(ctx *gin.Context) {
	allTodos := GetDummyTodos()
	ctx.JSON(http.StatusOK, allTodos)
}

func HandleAddTodo(ctx *gin.Context) {
	var todo Todo
	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Invalid input",
		})
		return
	}
	// store todo
	todo.ID = len(GetDummyTodos()) + 1
	log.Println("Todo added:", todo)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Todo added successfully",
		"data" : todo,
	})
}

func HandleDeleteTodo(ctx *gin.Context){
	id := ctx.Param("id")
	log.Println("id: ", id)
	// logic to delete
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "Todo deleted successfully",
		"todo_id" : id,
	})
}