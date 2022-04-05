package main

import (
	"awaken-media-service/pkg/aws-s3"
	"awaken-media-service/pkg/env"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "github.com/aws/aws-sdk-go/service/s3/s3manager"

func main() {
	env.LoadEnv()
	awsAccessKeyID := env.GetEnv("AWS_ACCESS_KEY_ID")
	fmt.Println("My access key ID is ", awsAccessKeyID)

	sess := aws_s3.GetAWSSession()
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("sess", sess)
		c.Next()
	})

	router.POST("/upload", UploadImage)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Failed to upload file",
			"uploader": "up",
		})
	})

	_ = router.Run(":4000")
}

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
