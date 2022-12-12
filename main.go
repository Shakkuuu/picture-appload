package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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
	r.GET("/findimage", findimage)

	r.Run()
}

func index(c *gin.Context) {
	aaa := "aaa"
	c.HTML(200, "index.html", gin.H{"aaa": aaa})
}

func upload(c *gin.Context) {
	image, err := c.FormFile("image")
	if err != nil {
		c.HTML(200, "error.html", gin.H{"err": err})
	}
	log.Println(image.Filename)

	if err := c.SaveUploadedFile(image, "data/"+image.Filename); err != nil {
		c.HTML(200, "error.html", gin.H{"err": err})
	}
	filename := image.Filename

	ex := filepath.Ext(filename)

	i := c.PostForm("id")
	if i == "" {
		id := filename
		c.HTML(200, "uploaded_check.html", gin.H{"filename": filename, "id": id})
		return
	}

	id := i + ex

	if err := os.Rename("data/"+image.Filename, "data/"+id); err != nil {
		c.HTML(200, "error.html", gin.H{"err": err})
	}

	c.HTML(200, "uploaded_check.html", gin.H{"filename": filename, "id": id})
}

func imagelist(c *gin.Context) {
	image, err := os.ReadDir("data")
	if err != nil {
		c.HTML(200, "error.html", gin.H{"err": err})
	}
	c.HTML(200, "imageview.html", gin.H{"uploadedimage": image})
}

func deletecheck(c *gin.Context) {
	d := c.Param("id")

	c.HTML(200, "delete.html", gin.H{"id": d})
}

func imagedelete(c *gin.Context) {
	d := c.Param("id")

	if err := os.Remove("data/" + d); err != nil {
		c.HTML(200, "error.html", gin.H{"err": err})
	}
	aaa := d + "は削除されました。"
	c.HTML(200, "index.html", gin.H{"aaa": aaa})
}

func findimage(c *gin.Context) {
	f := c.Query("find")

	c.HTML(200, "imageview.html", gin.H{"find": f})
}

//ファイル複数アップロード、画像表示の見栄え、UI、
