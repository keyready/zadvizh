package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/api/controllers"
	"server/internal/api/middleware"
)

func NewEmployeeRouters(eContr *controllers.EmployeeControllers, router *gin.Engine) {
	employeeRouters := router.Group("/api/v1/employers")

	employeeRouters.GET("", middleware.TokenMiddleware(), eContr.GetAllEmployers)

	router.POST("/api/v1/auth", eContr.AuthEmployee)
	router.GET("/api/v1/check_ref", eContr.VerifyLink)
	router.GET("/api/v1/get_access", eContr.GetAccessToken)

}
