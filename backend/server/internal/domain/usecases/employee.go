package usecases

import (
	"server/internal/domain/repositories"
	"server/internal/domain/types/request"
)

type EmployeeUsecase interface {
	AuthEmployee(authEmployee request.AuthEmployee) (httpCode int, usecaseError error)
	//GetAllEmployers() (httpCode int, usecaseError error, candidates []models.Employee)
}

type EmployeeUsecaseImpl struct {
	employeeRepo repositories.EmployeeRepository
}

func NewEmployeeUsecase(employeeRepo repositories.EmployeeRepository) *EmployeeUsecaseImpl {
	return &EmployeeUsecaseImpl{employeeRepo: employeeRepo}
}

func (eUsecae *EmployeeUsecaseImpl) AuthEmployee(authEmployee request.AuthEmployee) (httpCode int, usecaseError error) {
	httpCode, usecaseError = eUsecae.employeeRepo.AuthEmployee(authEmployee)
	if usecaseError != nil {
		return httpCode, usecaseError
	}
	return httpCode, nil
}
