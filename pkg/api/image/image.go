package image

import (
	"bytes"
	"github.com/Behemoth11/awaken-email-service/pkg/api/custom_error"
	"github.com/Behemoth11/awaken-email-service/pkg/api/image/manipulations"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
)

// NewService : Creates a image service
func NewService() Service {
	return Service{}
}

type Service struct{}

func (service Service) RegisterRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/", handlePost)
}

// handlePost : handles post requests
func handlePost(context *gin.Context) {
	file, _, error := context.Request.FormFile("upload")

	if error != nil {
		context.Error(custom_error.BadRequestError("Request did not contain any file"))
		return
	}
	srcImage, header, _ := image.Decode(file)
	finalImage := manipulations.Crop(srcImage)

	print(header)

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, finalImage, nil)

	context.Data(200, "image/jpeg", buf.Bytes())
}
