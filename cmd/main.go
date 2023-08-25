package main

import (
	"github.com/Barboxxa/DojoLebre/internal/interface/rest"
	"github.com/Barboxxa/DojoLebre/internal/service"
)

func main() {

	uploadService := service.NewUploadService()

	controller := rest.Controllers{
		UploadController: rest.NewUploadController(uploadService),
	}

	controller.NewServer()
}
