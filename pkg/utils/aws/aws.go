package aws

import (
	"github.com/Behemoth11/awaken-email-service/pkg/utils/env"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// GetAWSSession: get a aws s3 session than can be used to read and write to s3 storage
func GetAWSSession() *session.Session {
	
	AwsRegion := env.GetEnv("AWS_REGION")
	AwsAccessKeyId := env.GetEnv("AWS_ACCESS_KEY_ID")
	AwsSecretAccessKey := env.GetEnv("AWS_SECRET_ACCESS_KEY")

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