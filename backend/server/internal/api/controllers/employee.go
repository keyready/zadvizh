package controllers

import (
	"fmt"
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

func (eCont *EmployeeControllers) VerifyInviteLink(ctx *gin.Context) {
	inviteLink := ctx.Query("ref")
	if inviteLink == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Ошибка распознавания ссылки"})
		return
	}

	httpCode, contrErr := eCont.employeeUsecase.VerifyInviteLink(inviteLink)
	if contrErr != nil {
		ctx.AbortWithStatusJSON(httpCode, gin.H{"error": contrErr.Error()})
		return
	}

	ctx.JSON(httpCode, gin.H{})
}

func (eCont *EmployeeControllers) GetAllEmployers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (eCont *EmployeeControllers) AuthEmployee(ctx *gin.Context) {
	var authEmployee request.AuthEmployee

	bindErr := ctx.ShouldBindJSON(&authEmployee)
	if bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("Ошибка получения данных с клиента: %s", bindErr.Error())})
		return
	}

	httpCode, contrErr := eCont.employeeUsecase.AuthEmployee(authEmployee)
	if contrErr != nil {
		ctx.AbortWithStatusJSON(httpCode, gin.H{"error": fmt.Errorf("Ошибка сервера: %s", contrErr.Error())})
		return
	}

	ctx.JSON(httpCode, gin.H{})
}
