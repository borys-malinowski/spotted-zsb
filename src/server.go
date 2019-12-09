package main

import (
	"database/sql"
	"fmt"
	"github.com/bdwilliams/go-jsonify/jsonify"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/microcosm-cc/bluemonday"
)

type Post struct {
	title string `json:"title"`
	content string `json:"content"`
	id uint `json:"id"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	db, err := sql.Open("mysql", "root:@/spottedzsb")
	checkError(err)
	defer db.Close()
	router := gin.Default()
	go router.Use(database(db))
	go router.Use(static.Serve("/", static.LocalFile("public", false)))
	go router.POST("/api/add-post", addPost)
	go router.GET("/api/get-posts", getPosts)
	router.Run()
}

func database(db *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("DB", db)
		context.Next()
	}
}

func addPost(context *gin.Context) {
	sanitizer := bluemonday.UGCPolicy()
	topic := sanitizer.Sanitize(context.PostForm("topic"))
	text := sanitizer.Sanitize(context.PostForm("text"))
	if topic != "" && text != "" {
		db := context.MustGet("DB").(*sql.DB)
		result, err := db.Query("INSERT INTO posts (title, content) VALUES ('" + topic + "', '" + text + "')")
		checkError(err)
		defer result.Close()
		context.Redirect(301, "/")
	}
}

func getPosts(context *gin.Context) {
	db := context.MustGet("DB").(*sql.DB)
	result, err := db.Query("SELECT * FROM posts")
	checkError(err)
	defer result.Close()
	resultJSON := jsonify.Jsonify(result)
	context.JSON(200, resultJSON)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}