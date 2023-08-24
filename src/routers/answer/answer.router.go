package answer_router

import (
	"que-sera-de-mi/src/controllers/answer"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	router.GET("", answer.GetAnswerHandler)
}
