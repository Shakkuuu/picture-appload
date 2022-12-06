package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("start")

	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")

	r.GET("/", index)
	r.POST("/upload", upload)

	r.Run()
}

func index(c *gin.Context) {
	aaa := "aaa"
	c.HTML(200, "index.html", gin.H{"aaa": aaa})
}

func upload(c *gin.Context) {
	image, _ := c.FormFile("image")
	log.Println(image.Filename)

	c.SaveUploadedFile(image, "data/"+image.Filename)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded", image.Filename))
}
