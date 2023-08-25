package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/textract"

	"github.com/Barboxxa/DojoLebre/internal/interface/rest"
	"github.com/Barboxxa/DojoLebre/internal/service"
)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("", "", ""),
	})

	if err != nil {
		panic(fmt.Sprintf("Erro ao criar sessão: %s", err.Error()))
	}

	rekoClient := rekognition.New(sess)

	textractClient := textract.New(sess)

	uploadService := service.NewUploadService(rekoClient, textractClient)

	controller := rest.Controllers{
		UploadController: rest.NewUploadController(uploadService),
	}

	controller.NewServer()
}
