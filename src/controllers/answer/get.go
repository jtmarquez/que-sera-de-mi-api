package answer

import (
	"net/http"
	"que-sera-de-mi/src/models/answer"

	"github.com/gin-gonic/gin"
)

func GetAnswerHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, answer.AnswerData)
}
