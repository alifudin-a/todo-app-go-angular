package main

import (
	"path"
	"path/filepath"

	"github.com/alifudin-a/todo-app-go-angular/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("/ui/dist/ui/index.html")
		} else {
			c.File("/ui/dist/ui" + path.Join(dir, file))
		}
	})

	r.GET("/todo", handlers.GetTodoListHandler)
	r.POST("/todo", handlers.AddTodoHandler)
	r.DELETE("/todo/:id", handlers.DeleteTodoHandler)
	r.PUT("/todo", handlers.CompleteTodoHandler)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
