package routers

import (
	answer_router "que-sera-de-mi/src/routers/answer"
	group_router "que-sera-de-mi/src/routers/group"
	question_router "que-sera-de-mi/src/routers/question"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	answerRouter := router.Group("/answers")
	answer_router.Setup(answerRouter)

	groupRouter := router.Group("/groups")
	group_router.Setup(groupRouter)

	questionRouter := router.Group("/questions")
	question_router.Setup(questionRouter)
}
