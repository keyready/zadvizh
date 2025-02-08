package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/api/controllers"
)

func NewEmployeeRouters(eContr *controllers.EmployeeControllers, router *gin.Engine) {
	employeeRouters := router.Group("/api/v1/employers")

	employeeRouters.GET("", eContr.GetAllEmployers)

	router.POST("/api/v1/auth", eContr.AuthEmployee)
	router.GET("/", eContr.VerifyLink)
}
