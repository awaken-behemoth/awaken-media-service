package aws_s3

import (
	"awaken-media-service/pkg/env"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(c *gin.Context) {
	sess := c.MustGet("sess").(*session.Session)
	uploader := s3manager.NewUploader(sess)
	MyBucket := env.GetEnv("BUCKET_NAME")
	file, header, err := c.Request.FormFile("photo")
	filename := header.Filename
	AwsRegion := env.GetEnv("AWS_REGION")
	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(MyBucket),
		Key:    aws.String(filename),
		Body:   file,
	})

	fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    err,
			"uploader": up,
		})
		return
	}
	filepath := "https://" + MyBucket + "." + "s3-" + AwsRegion + ".amazonaws.com/" + filename
	c.JSON(http.StatusOK, gin.H{
		"filepath": filepath,
	})
}
