package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/api/controllers"
)

func NewTeacherRouters(tContr *controllers.TeacherController, router *gin.Engine) {
	router.GET("/api/v1/teachers", tContr.GetAllTeachers)
	router.POST("/api/v1/teachers/addComment", tContr.WriteComment)
	router.POST("/api/v1/teachers/like", tContr.LikeDislike)
}
