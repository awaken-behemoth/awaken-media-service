package router_control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/h2non/bimg"
	"io/ioutil"
	"net/http"
	"os"
)

func Init(router *gin.Engine) {
	router.POST("/upload", Optimize)
}

func Optimize(context *gin.Context) {
	file, header, error := context.Request.FormFile("photo")

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":    error,
			"uploader": header,
		})
	}
	//
	buffer, err := ioutil.ReadAll(file)
	//
	newImage, err := bimg.NewImage(buffer).Resize(320, 240)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	context.Data(200, "img/png", newImage)

}
