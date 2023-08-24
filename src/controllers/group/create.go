package group_controllers

import (
	"fmt"
	"net/http"
	"que-sera-de-mi/src/models/group"

	"github.com/gin-gonic/gin"
)

func CreateGroupHandler(c *gin.Context) {
	err := group.CreateGroupHandler(group.GroupData[0])

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, group.GroupData[0])
	}

	c.IndentedJSON(http.StatusCreated, group.GroupData[0])
}
