package main

import (
	"fmt"
	"log"
	_ "net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("start")

	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")
	r.Static("data", "./data")

	r.GET("/", index)
	r.POST("/upload", upload)
	r.GET("imagelist", imagelist)

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
	filename := image.Filename
	c.HTML(200, "uploaded_check.html", gin.H{"filename": filename})
	// c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded", image.Filename))
}

func imagelist(c *gin.Context) {
	image, _ := os.ReadDir("data")
	c.HTML(200, "imageview.html", gin.H{"uploadedimage": image})
}
