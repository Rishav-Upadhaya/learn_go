package post

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"

)

func HandleReadPost(ctx *gin.Context) {
	allPosts := GetDummyPosts()
	ctx.JSON(http.StatusOK, allPosts)
}

func HandleAddPost(ctx *gin.Context) {
	var post Post
	err := ctx.ShouldBindJSON(&post)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Invalid input",
		})
		return
	}
	// store post
	post.ID = len(GetDummyPosts()) + 1
	log.Println("Post added:", post)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post added successfully",
		"data": post,
	})
}

func HandleUpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Println("id: ", id)
	var post Post
	err := ctx.ShouldBindJSON(&post)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Invalid input",
		})
		return
	}
	// logic to update post
	log.Println("Post updated:", post)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
		"data": post,
	})
}

func HandleDeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Println("id: ", id)
	// logic to delete post
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
		"post_id": id,
	})
}