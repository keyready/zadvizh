package repositories

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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
	GetAllTeamNames(field string) (teamNames []string)
	GetAllScidirs(field string) (scidirs []string)
	VerifyLink(authorLinkTgId string) (verifyLink bool)
}

type EmployeeRepositoryImpl struct {
	mongoDB *mongo.Database
}

func NewEmployeeRepositoryImpl(mongoDB *mongo.Database) *EmployeeRepositoryImpl {
	return &EmployeeRepositoryImpl{mongoDB: mongoDB}
}

func (eRepo *EmployeeRepositoryImpl) VerifyLink(authorLinkTgId string) (verifyLink bool) {
	mongoErr := eRepo.mongoDB.Collection("employees").
		FindOne(ctx, bson.M{"TgId": authorLinkTgId}).Err()

	if errors.Is(mongoErr, mongo.ErrNoDocuments) {
		return false
	}

	return true
}

func (eRepo *EmployeeRepositoryImpl) GetAllScidirs(field string) (scidirs []string) {
	prj := bson.M{
		"scidir": 1,
	}

	cur, mongoErr := eRepo.mongoDB.Collection("employees").
		Find(ctx, bson.D{{"field", field}}, options.Find().SetProjection(prj))
	defer cur.Close(ctx)

	if mongoErr != nil {
		return nil
	}

	for cur.Next(ctx) {
		var scidir bson.M
		if decodeErr := cur.Decode(&scidir); decodeErr != nil {
			return nil
		}

		if sd, exist := scidir["scidir"]; exist {
			scidirs = append(scidirs, sd.(string))
		}
	}

	return scidirs
}

func (eRepo *EmployeeRepositoryImpl) GetAllTeamNames(field string) (teamNames []string) {
	projection := bson.M{
		"teamName": 1,
	}

	cur, mongoErr := eRepo.mongoDB.Collection("employees").
		Find(ctx, bson.D{{"field", field}}, options.Find().SetProjection(projection))
	defer cur.Close(ctx)

	if mongoErr != nil {
		return nil
	}

	for cur.Next(ctx) {
		var teamName bson.M
		if decodeErr := cur.Decode(&teamName); decodeErr != nil {
			return nil
		}

		if tN, exist := teamName["teamName"]; exist {
			teamNames = append(teamNames, tN.(string))
		}
	}

	return teamNames
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
	cur, mongoErr := eRepo.mongoDB.Collection("employees").
		Find(ctx, bson.D{})
	defer cur.Close(ctx)

	if mongoErr != nil {
		repoErr = fmt.Errorf("Ошибка извлечения всех участнико: %w", mongoErr.Error())
		return http.StatusInternalServerError, repoErr, nil
	}

	for cur.Next(ctx) {
		var employee models.Employee
		if decodeErr := cur.Decode(&employee); decodeErr != nil {
			repoErr = fmt.Errorf("Ошибка анмаршалинга одного участника: %w", decodeErr.Error())
			return http.StatusInternalServerError, repoErr, nil
		}
		employees = append(employees, employee)
	}

	return http.StatusOK, nil, employees
}
