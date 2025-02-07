package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
	"net/http"
	"server/internal/domain/types/models"
	"server/internal/domain/types/request"
)

var (
	ctx context.Context
)

type EmployeeRepository interface {
	AuthEmployee(authEmployee request.AuthEmployee) (httpCode int, repoError error)
	GetAllEmployees() (httpCode int, repoError error, employees []models.Employee)
}

type EmployeeRepositoryImpl struct {
	mongoDB *mongo.Database
}

func NewEmployeeRepositoryImpl(mongoDB *mongo.Database) *EmployeeRepositoryImpl {
	return &EmployeeRepositoryImpl{mongoDB: mongoDB}
}

func (eRepo *EmployeeRepositoryImpl) AuthEmployee(authEmployee request.AuthEmployee) (httpCode int, repoError error) {
	result, mongoErr := eRepo.mongoDB.Collection("employees").
		InsertOne(ctx, authEmployee)
	if mongoErr != nil {
		repoError = fmt.Errorf("Ошибка добавления нового участника: %w", mongoErr.Error())
		return http.StatusInternalServerError, repoError
	}

	log.Printf("Новый участник: %v", result.InsertedID)

	return http.StatusOK, nil
}

func (eRepo *EmployeeRepositoryImpl) GetAllEmployees() (httpCode int, repoErr error, employees []models.Employee) {
	//cur, mongoErr := eRepo.mongoDB.Collection("employees").
	//	Find(ctx, bson.D{})
	return http.StatusOK, nil, nil
}
