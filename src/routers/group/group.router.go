package group_router

import (
	group_controllers "que-sera-de-mi/src/controllers/group"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	router.GET("", group_controllers.GetGroupHandler)
	router.POST("", group_controllers.CreateGroupController)
}
