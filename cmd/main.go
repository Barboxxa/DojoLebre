package main

import (
	"github.com/aws/aws-sdk-go-v2/service/rekognition"

	"github.com/Barboxxa/DojoLebre/internal/interface/rest"
	"github.com/Barboxxa/DojoLebre/internal/service"
)

func main() {

	rekoClient := rekognition.New(rekognition.Options{
		Region: "", // TODO: passar credentiais das envs;
	})

	uploadService := service.NewUploadService(rekoClient)

	controller := rest.Controllers{
		UploadController: rest.NewUploadController(uploadService),
	}

	controller.NewServer()
}
