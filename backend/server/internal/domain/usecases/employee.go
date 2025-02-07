package usecases

import (
	"fmt"
	"net/http"
	"server/internal/domain/repositories"
	"server/internal/domain/types/request"
	"strings"
)

type EmployeeUsecase interface {
	AuthEmployee(authEmployee request.AuthEmployee) (httpCode int, usecaseError error)
	//GetAllEmployers() (httpCode int, usecaseError error, candidates []models.Employee)
	VerifyInviteLink(inviteLink string) (httpCode int, usecaseError error)
}

type EmployeeUsecaseImpl struct {
	employeeRepo repositories.EmployeeRepository
}

func NewEmployeeUsecase(employeeRepo repositories.EmployeeRepository) *EmployeeUsecaseImpl {
	return &EmployeeUsecaseImpl{employeeRepo: employeeRepo}
}

func (eUsecase *EmployeeUsecaseImpl) VerifyInviteLink(inviteLink string) (httpCode int, usecaseError error) {
	authorLinkHash := strings.Split(inviteLink, "_")[len(strings.Split(inviteLink, "_"))-1]

	if authorLinkHash != "" {
		fmt.Println(authorLinkHash)
		return http.StatusOK, nil
	}

	usecaseError = fmt.Errorf("Невалидная инвайт-ссылка")
	return http.StatusBadRequest, usecaseError
}

func (eUsecae *EmployeeUsecaseImpl) AuthEmployee(authEmployee request.AuthEmployee) (httpCode int, usecaseError error) {
	httpCode, usecaseError = eUsecae.employeeRepo.AuthEmployee(authEmployee)
	if usecaseError != nil {
		return httpCode, usecaseError
	}
	return httpCode, nil
}
