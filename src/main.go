package main

import (
	"que-sera-de-mi/src/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routers.Setup(router)

	router.Run(":3000")
}
