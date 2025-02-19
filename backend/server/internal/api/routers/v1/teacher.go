package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/api/controllers"
	"server/internal/api/middleware"
)

func NewTeacherRouters(tContr *controllers.TeacherController, router *gin.Engine) {
	router.GET("/api/v1/teachers", middleware.TokenMiddleware(), tContr.GetAllTeachers)
	router.POST("/api/v1/teachers/addComment", middleware.TokenMiddleware(), tContr.WriteComment)
	router.POST("/api/v1/teachers/like", middleware.TokenMiddleware(), tContr.LikeDislike)
}
