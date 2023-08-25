package response

import (
	"fmt"
	"net/http"
	dotenv "que-sera-de-mi/src/config/env_vars"

	"github.com/gin-gonic/gin"
)

type StatusCode = int

const (
	OK           StatusCode = http.StatusOK
	BAD_REQUEST  StatusCode = http.StatusBadRequest
	CREATED      StatusCode = http.StatusCreated
	FORBIDDEN    StatusCode = http.StatusForbidden
	SERVER_ERROR StatusCode = http.StatusInternalServerError
)

func Response(c *gin.Context, status StatusCode, body any) {
	stage := dotenv.GetEnvironmentVariable("STAGE")

	fmt.Println("stage", stage)
	if stage == "production" {
		fmt.Println("prod")
		c.JSON(status, body)
	}

	fmt.Println("dev", status, body)
	c.IndentedJSON(status, body)
}
