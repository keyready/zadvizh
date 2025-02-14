package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/domain/types/request"
	"server/internal/domain/usecases"
)

type EmployeeControllers struct {
	employeeUsecase usecases.EmployeeUsecase
}

func NewEmployeeControllers(employeeUsecase usecases.EmployeeUsecase) *EmployeeControllers {
	return &EmployeeControllers{employeeUsecase: employeeUsecase}
}

func (eCont *EmployeeControllers) GetAccessToken(ctx *gin.Context) {
	tgId := ctx.Query("tgId")

	httpCode, contrErr, token := eCont.employeeUsecase.GetAccessToken(tgId)
	if contrErr != nil {
		ctx.AbortWithStatusJSON(httpCode, contrErr.Error())
		return
	}

	ctx.JSON(httpCode, gin.H{"accessToken": token})
}

func (eCont *EmployeeControllers) GetAllEmployers(ctx *gin.Context) {
	httpCode, contrErr, three := eCont.employeeUsecase.GetAllEmployers()
	if contrErr != nil {
		ctx.AbortWithStatusJSON(httpCode, gin.H{"error": contrErr.Error()})
		return
	}

	ctx.JSON(httpCode, three)
}

func (eCont *EmployeeControllers) AuthEmployee(ctx *gin.Context) {
	var authEmployee request.AuthEmployee

	bindErr := ctx.ShouldBindJSON(&authEmployee)
	if bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
		return
	}

	httpCode, contrErr := eCont.employeeUsecase.AuthEmployee(authEmployee)
	if contrErr != nil {
		ctx.AbortWithStatusJSON(httpCode, gin.H{"error": contrErr.Error()})
		return
	}

	ctx.JSON(httpCode, gin.H{})
}
