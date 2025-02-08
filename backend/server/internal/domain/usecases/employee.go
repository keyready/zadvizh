package usecases

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"server/internal/domain/repositories"
	"server/internal/domain/types/enum/field"
	"server/internal/domain/types/models"
	"server/internal/domain/types/request"
	"server/internal/domain/types/response"
	"server/pkg/utils"
	"strings"
)

type EmployeeUsecase interface {
	AuthEmployee(authEmployee request.AuthEmployee) (httpCode int, usecaseError error)
	VerifyLink(ref string) (httpCode int, usecaseError error)
	GetAllEmployers() (httpCode int, usecaseError error, three response.Three)
}

type EmployeeUsecaseImpl struct {
	employeeRepo repositories.EmployeeRepository
}

func NewEmployeeUsecase(employeeRepo repositories.EmployeeRepository) *EmployeeUsecaseImpl {
	return &EmployeeUsecaseImpl{employeeRepo: employeeRepo}
}

func (eUsecase *EmployeeUsecaseImpl) VerifyLink(ref string) (httpCode int, usecaseError error) {
	authorLinkId := strings.Split(ref, "_")[len(strings.Split(ref, "_"))-1]
	authorLinkIdByte, _ := base64.StdEncoding.DecodeString(authorLinkId)
	authorLinkId = string(authorLinkIdByte)

	verifyLink := eUsecase.employeeRepo.VerifyLink(authorLinkId)
	if !verifyLink {
		return http.StatusUnauthorized, fmt.Errorf("Невалидная ссылка приглашение")
	}

	return http.StatusOK, nil
}

func (eUsecase *EmployeeUsecaseImpl) GetAllEmployers() (httpCode int, usecaseError error, three response.Three) {
	httpCode, usecaseError, employees := eUsecase.employeeRepo.GetAllEmployees()
	if usecaseError != nil {
		return httpCode, usecaseError, three
	}

	three = response.Three{}

	for _, employee := range employees {
		switch employee.Field {
		case field.ORG:
			three.Org = append(three.Org, employee)

		case field.DEV:
			dev := make(map[string][]models.Employee)

			teamNames := eUsecase.employeeRepo.GetAllTeamNames(field.DEV)
			if teamNames == nil {
				usecaseError = fmt.Errorf("Ошибка получения всех имен команд")
				return http.StatusInternalServerError, usecaseError, three
			}

			teamNames = utils.RemoveDuplicates(teamNames)

			for _, teamName := range teamNames {
				if employee.TeamName == teamName {
					dev[teamName] = append(dev[teamName], employee)
				}
			}

			three.Dev = dev

		case field.SEC:
			sec := make(map[string][]models.Employee)

			teamNames := eUsecase.employeeRepo.GetAllTeamNames(field.SEC)
			if teamNames == nil {
				usecaseError = fmt.Errorf("Ошибка получения всех имен команд")
				return http.StatusInternalServerError, usecaseError, three
			}

			teamNames = utils.RemoveDuplicates(teamNames)

			for _, teamName := range teamNames {
				if employee.TeamName == teamName {
					sec[teamName] = append(sec[teamName], employee)
				}
			}

			three.Sec = sec

		case field.DEVOPS:
			three.Devops = append(three.Devops, employee)

		case field.SCIENCE:
			science := make(map[string][]models.Employee)

			scidirs := eUsecase.employeeRepo.GetAllScidirs(field.SCIENCE)
			if scidirs == nil {
				usecaseError = fmt.Errorf("Ошибка получения всех научруков")
				return http.StatusInternalServerError, usecaseError, three
			}

			scidirs = utils.RemoveDuplicates(scidirs)

			for _, scidir := range scidirs {
				if scidir == employee.Scidir {
					science[scidir] = append(science[scidir], employee)
				}
			}

			three.Science = science
		}
	}

	return httpCode, nil, three
}

func (eUsecae *EmployeeUsecaseImpl) AuthEmployee(authEmployee request.AuthEmployee) (httpCode int, usecaseError error) {
	httpCode, usecaseError = eUsecae.employeeRepo.AuthEmployee(authEmployee)
	if usecaseError != nil {
		return httpCode, usecaseError
	}
	return httpCode, nil
}
