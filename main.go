package main

import (
	"awaken-media-service/pkg/env"
	"awaken-media-service/pkg/router_control"
	"github.com/gin-gonic/gin"
)

func main() {
	env.LoadEnv()

	router := gin.Default()

	router_control.Init(router)

	_ = router.Run(":4000")
}
