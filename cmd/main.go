package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/bavatech/envloader"

	"github.com/Barboxxa/DojoLebre/internal/environment"
	"github.com/Barboxxa/DojoLebre/internal/interface/rest"
	"github.com/Barboxxa/DojoLebre/internal/service"
)

func main() {

	if err := envloader.Load(&environment.Env); err != nil {
		panic(fmt.Sprintf("error on loading variables: %s", err.Error()))
	}

	env := environment.Env

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(env.Region),
		Credentials: credentials.NewStaticCredentials(env.AccessKeyID, env.SecretAccessKey, ""),
	})

	if err != nil {
		panic(fmt.Sprintf("error on create session: %s", err.Error()))
	}

	rekoClient := rekognition.New(sess)

	textractClient := textract.New(sess)

	uploadService := service.NewUploadService(rekoClient, textractClient)

	controller := rest.Controllers{
		UploadController: rest.NewUploadController(uploadService),
	}

	controller.NewServer()
}
