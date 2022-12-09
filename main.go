package main

import (
	"fmt"
	"log"
	_ "net/http"
	"os"
	_ "strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("start")

	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")
	r.Static("data", "./data")

	r.GET("/", index)
	r.POST("/upload", upload)
	r.GET("/imagelist", imagelist)
	r.GET("/deletecheck/:id", deletecheck)
	r.POST("/imagedelete/:id", imagedelete)

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

func deletecheck(c *gin.Context) {
	d := c.Param("id")
	// id, err := strconv.Atoi(d)
	// if err != nil {
	// 	panic("ERROR")
	// }

	c.HTML(200, "delete.html", gin.H{"id": d})
}

func imagedelete(c *gin.Context) {
	d := c.Param("id")
	// id, err := strconv.Atoi(d)
	// if err != nil {
	// 	panic("ERROR")
	// }

	os.Remove("data/" + d)
	aaa := d + "は削除されました。"
	c.HTML(200, "index.html", gin.H{"aaa": aaa})
}
