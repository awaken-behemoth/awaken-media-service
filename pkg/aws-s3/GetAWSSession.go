package aws_s3

import (
	"awaken-media-service/pkg/env"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetAWSSession() *session.Session {
	AwsAccessKeyId := env.GetEnv("AWS_ACCESS_KEY_ID")
	AwsSecretAccessKey := env.GetEnv("AWS_SECRET_ACCESS_KEY")
	AwsRegion := env.GetEnv("AWS_REGION")

	session, error := session.NewSession(
		&aws.Config{
			Region: aws.String(AwsRegion),
			Credentials: credentials.NewStaticCredentials(
				AwsAccessKeyId,
				AwsSecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})

	if error != nil {
		panic(session)
	}
	return session
}
