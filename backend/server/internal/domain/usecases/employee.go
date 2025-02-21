package usecases

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"server/internal/domain/repositories"
	"server/internal/domain/types/enum/field"
	"server/internal/domain/types/models"
	"server/internal/domain/types/request"
	"server/internal/domain/types/response"
)

type EmployeeUsecase interface {
	AuthEmployee(authEmployee request.AuthEmployee) (httpCode int, usecaseError error, inviteLink string)
	GetAllEmployers() (httpCode int, usecaseError error, three response.Node)
	GetAccessToken(tgId string) (httpCode int, usecaseError error, token response.Token)
	VerifyLink(tgId string) (httpCode int, usecaseError error)
}

type EmployeeUsecaseImpl struct {
	employeeRepo repositories.EmployeeRepository
}

func NewEmployeeUsecase(employeeRepo repositories.EmployeeRepository) *EmployeeUsecaseImpl {
	return &EmployeeUsecaseImpl{employeeRepo: employeeRepo}
}

func (eUsecase *EmployeeUsecaseImpl) VerifyLink(tgId string) (httpCode int, usecaseError error) {
	check := eUsecase.employeeRepo.VerifyLink(tgId)
	if !check {
		return http.StatusUnauthorized, fmt.Errorf("Нет доступа")
	}
	return http.StatusOK, nil
}

func (eUsecase *EmployeeUsecaseImpl) GetAccessToken(tgId string) (httpCode int, usecaseError error, tokenData response.Token) {
	check, userId := eUsecase.employeeRepo.GetAccessToken(tgId)
	if !check {
		return http.StatusUnauthorized, fmt.Errorf("Данный пользователь не авторизован"), tokenData
	}

	claims := jwt.MapClaims{
		"tgId": tgId,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := t.SignedString([]byte("IimULHg9FRS0XleGnPZo"))

	tokenData = response.Token{
		AccessToken: tokenString,
		ID:          userId,
	}

	return http.StatusOK, nil, tokenData
}

func (eUsecase *EmployeeUsecaseImpl) GetAllEmployers() (httpCode int, usecaseError error, three response.Node) {
	httpCode, usecaseErr, employers := eUsecase.employeeRepo.GetAllEmployees()
	if usecaseErr != nil {
		return httpCode, usecaseErr, three
	}

	var DN models.Employee
	var AS models.Employee
	var RO models.Employee
	var OA models.Employee
	for _, e := range employers {
		switch e.Lastname {
		case "Бирюков":
			DN = e
		case "Дудкин":
			AS = e
		case "Крюков":
			RO = e
		case "Мишуков":
			OA = e
		}
	}

	var devNodes []response.Node

	var secNodes []response.Node

	var devopsNodes []response.Node

	var sciNodes []response.Node

	devTeams := eUsecase.employeeRepo.GetAllTeamNames(field.DEV)
	secTeams := eUsecase.employeeRepo.GetAllTeamNames(field.SEC)

	for _, t := range devTeams {
		devNode := response.Node{
			ID: uuid.NewString(),
			Data: response.Data{
				Label: t,
			},
		}
		for _, e := range employers {
			if e.TeamName == t {
				dvn := response.Node{
					ID: uuid.NewString(),
					Data: response.Data{
						Label:    e.Lastname + " " + e.Firstname,
						Employee: e,
					},
				}
				devNode.Children = append(devNode.Children, dvn)
			}
		}
		devNodes = append(devNodes, devNode)
	}

	for _, s := range secTeams {
		secNode := response.Node{
			ID: uuid.NewString(),
			Data: response.Data{
				Label: s,
			},
		}
		for _, e := range employers {
			if e.TeamName == s {
				svn := response.Node{
					ID: uuid.NewString(),
					Data: response.Data{
						Label:    e.Lastname + " " + e.Firstname,
						Employee: e,
					},
				}
				secNode.Children = append(secNode.Children, svn)
			}
		}
		secNodes = append(secNodes, secNode)
	}

	for _, e := range employers {
		switch e.Field {
		case field.DEVOPS:
			devopsNode := response.Node{
				ID: uuid.NewString(),
				Data: response.Data{
					Label:    e.Lastname + " " + e.Firstname,
					Employee: e,
				},
				Children: []response.Node{},
			}
			devopsNodes = append(devopsNodes, devopsNode)
		case field.SCIENCE:
			sciNode := response.Node{
				ID: uuid.NewString(),
				Data: response.Data{
					Label:    e.Lastname + " " + e.Firstname,
					Employee: e,
				},
				Children: []response.Node{},
			}
			sciNodes = append(sciNodes, sciNode)
		}
	}

	three = response.Node{}

	three.ID = uuid.NewString()
	three.Data = response.Data{
		Label:     DN.Lastname + " " + DN.Firstname,
		DataLabel: field.ORG,
	}
	three.Children = []response.Node{
		response.Node{
			ID: uuid.NewString(),
			Data: response.Data{
				Label:     AS.Lastname + " " + AS.Firstname,
				DataLabel: field.ORG,
			},
			Children: []response.Node{
				response.Node{
					ID: uuid.NewString(),
					Data: response.Data{
						Label:     "Разработка",
						DataLabel: field.DEV,
					},
					Children: devNodes,
				},
				response.Node{
					ID: uuid.NewString(),
					Data: response.Data{
						Label:     "Безопасность",
						DataLabel: field.SEC,
					},
					Children: secNodes,
				},
				response.Node{
					ID: uuid.NewString(),
					Data: response.Data{
						Label:     "DevOps",
						DataLabel: field.DEVOPS,
					},
					Children: []response.Node{
						response.Node{
							ID: uuid.NewString(),
							Data: response.Data{
								Label: RO.Lastname + " " + RO.Firstname,
							},
							Children: devopsNodes,
						},
					},
				},
				response.Node{
					ID: uuid.NewString(),
					Data: response.Data{
						Label:     "Научная деятельность",
						DataLabel: field.SCIENCE,
					},
					Children: []response.Node{
						response.Node{
							ID: uuid.NewString(),
							Data: response.Data{
								Label: OA.Lastname + " " + OA.Firstname,
							},
							Children: sciNodes,
						},
					},
				},
			},
		},
	}

	return httpCode, nil, three
}

func (eUsecae *EmployeeUsecaseImpl) AuthEmployee(authEmployee request.AuthEmployee) (httpCode int, usecaseError error, inviteLink string) {
	httpCode, usecaseError, inviteLink = eUsecae.employeeRepo.AuthEmployee(authEmployee)
	if usecaseError != nil {
		return httpCode, usecaseError, inviteLink
	}
	return httpCode, nil, inviteLink
}
