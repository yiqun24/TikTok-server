package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yiqun24/TikTok-server/service"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
