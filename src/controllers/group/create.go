package group_controllers

import (
	"fmt"
	group_actions "que-sera-de-mi/src/models/group/actions"
	response "que-sera-de-mi/src/utils"

	"github.com/gin-gonic/gin"
)

func CreateGroupController(c *gin.Context) {
	var groupInput group_actions.CreateGroupInput

	if err := c.BindJSON(&groupInput); err != nil {
		response.Response(c, response.BAD_REQUEST, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(groupInput)
	res, err := group_actions.CreateGroupHandler(groupInput)

	if err != nil {
		fmt.Println(err)
		response.Response(c, response.SERVER_ERROR, gin.H{"error": err.Error()})
		return
	}

	response.Response(c, response.CREATED, res)
}
