package group_controllers

import (
	"net/http"
	"que-sera-de-mi/src/models/group"

	"github.com/gin-gonic/gin"
)

func GetGroupHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, group.GroupData)
}
