package question_router

import (
	"que-sera-de-mi/src/controllers/question"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	router.GET("", question.GetQuestionHandler)
}
