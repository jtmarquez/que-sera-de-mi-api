package question

import (
	"net/http"
	"que-sera-de-mi/src/models/question"

	"github.com/gin-gonic/gin"
)

func GetQuestionHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, question.QuestionData)
}
