package api

import (
	"github.com/Behemoth11/awaken-email-service/pkg/api/custom_error"
	"github.com/Behemoth11/awaken-email-service/pkg/api/image"
	"github.com/gin-gonic/gin"
)

func New() Api {
	api := Api{
		router: gin.Default(),
	}

	api.Use(custom_error.NewHandler())
	api.Mount(image.NewService(), "/image")

	return api
}
