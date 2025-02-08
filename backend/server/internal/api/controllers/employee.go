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

func (eCont *EmployeeControllers) VerifyLink(ctx *gin.Context) {
	ref := ctx.Query("ref")

	httpCode, contrErr := eCont.employeeUsecase.VerifyLink(ref)
	if contrErr != nil {
		ctx.AbortWithStatusJSON(httpCode, gin.H{"error": fmt.Sprintf("Ошибка валидации ссылки: %s", contrErr.Error())})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (eCont *EmployeeControllers) GetAllEmployers(ctx *gin.Context) {
	httpCode, contrErr, three := eCont.employeeUsecase.GetAllEmployers()
	if contrErr != nil {
		ctx.AbortWithStatusJSON(httpCode, gin.H{"error": contrErr.Error()})
		return
	}

	ctx.JSON(httpCode, gin.H{"three": three})
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
