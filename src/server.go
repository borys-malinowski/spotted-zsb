package main

import (
	"database/sql"
	"fmt"
	"github.com/bdwilliams/go-jsonify/jsonify"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("public", false)))
	router.POST("/api/add-post", addPost)
	router.GET("/api/get-posts", getPosts)
	router.Run()
}

func addPost(context *gin.Context) {
	topic := context.PostForm("topic")
	text := context.PostForm("text")
	if topic != "" && text != "" {
		db, err := sql.Open("mysql", "root:@/spottedzsb")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer db.Close()
		result, err := db.Query("INSERT INTO posts (title, content) VALUES ('" + topic + "', '" + text + "')")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer result.Close()
		context.File("public/index.html")
		context.Status(200)
	}
}

func getPosts(context *gin.Context) {
	db, err := sql.Open("mysql", "root:@/spottedzsb")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	result, err := db.Query("SELECT * FROM posts")
	resultJSON := jsonify.Jsonify(result)
	context.JSON(200, resultJSON)
}
