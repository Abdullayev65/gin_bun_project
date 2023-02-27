package myTest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func Test(c *gin.Context) {
	file, header,
		err := c.Request.FormFile("upload")

	file_, er1 := c.FormFile("file")
	//fmt.Println(file_.Filename)
	a := c.Request.Form
	_, _, _ = a, er1, file_
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create("tmp/" + filename + ".png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
}
