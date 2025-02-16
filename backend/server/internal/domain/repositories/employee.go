package repositories

import (
	"context"
	"encoding/base64"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"net/http"
	"server/internal/domain/types/models"
	"server/internal/domain/types/request"
	"slices"
	"strings"
)

var (
	ctx context.Context
)

type EmployeeRepository interface {
	AuthEmployee(authEmployee request.AuthEmployee) (httpCode int, repoError error, inviteLink string)
	GetAllEmployees() (httpCode int, repoError error, employees []models.Employee)
	GetAllTeamNames(field string) (teamNames []string)
	GetAccessToken(tgId string) (check bool)
	VerifyLink(tgId string) (check bool)
}

type EmployeeRepositoryImpl struct {
	mongoDB *mongo.Database
}

func NewEmployeeRepositoryImpl(mongoDB *mongo.Database) *EmployeeRepositoryImpl {
	return &EmployeeRepositoryImpl{mongoDB: mongoDB}
}

func (eRepo *EmployeeRepositoryImpl) VerifyLink(link string) (check bool) {
	decodedRef, decodeErr := base64.StdEncoding.DecodeString(link)
	if decodeErr != nil {
		return false
	}

	tgId := strings.Split(string(decodedRef), ":")[1]
	var res bson.M
	mongoErr := eRepo.mongoDB.Collection("employees").FindOne(ctx, bson.M{"tgid": tgId}).Decode(&res)
	if mongoErr == mongo.ErrNoDocuments {
		return false
	}
	return true
}

func (eRepo *EmployeeRepositoryImpl) GetAccessToken(tgId string) (check bool) {
	var result bson.M
	mongoErr := eRepo.mongoDB.Collection("employees").FindOne(ctx, bson.M{"token": tgId}).Decode(&result)
	if mongoErr == mongo.ErrNoDocuments {
		return false
	}
	return true
}

func (eRepo *EmployeeRepositoryImpl) GetAllTeamNames(field string) (teamNames []string) {
	projection := bson.M{
		"teamname": 1,
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

		if tN, exist := teamName["teamname"]; exist {
			if !slices.Contains(teamNames, tN.(string)) {
				teamNames = append(teamNames, tN.(string))
			}
		}
	}

	return teamNames
}

func (eRepo *EmployeeRepositoryImpl) AuthEmployee(authEmployee request.AuthEmployee) (httpCode int, repoError error, inviteLink string) {
	refTgIdByte, _ := base64.StdEncoding.DecodeString(authEmployee.Ref)
	refTgId := strings.Split(string(refTgIdByte), ":")[1]
	authEmployee.Ref = refTgId

	if authEmployee.Firstname == "Валентин" {
		_, _ = eRepo.mongoDB.Collection("employees").UpdateOne(ctx,
			bson.M{"firstname": "Валентин"},
			authEmployee)
	}
	if authEmployee.Firstname == "Родион" {
		_, _ = eRepo.mongoDB.Collection("employees").UpdateOne(ctx,
			bson.M{"firstname": "Родион"},
			authEmployee)
	}

	_, mongoErr := eRepo.mongoDB.Collection("employees").
		InsertOne(ctx, authEmployee)
	if mongoErr != nil {
		repoError = fmt.Errorf("Ошибка добавления нового участника: %s", mongoErr.Error())
		return http.StatusInternalServerError, repoError, inviteLink
	}

	var a models.Employee
	_ = eRepo.mongoDB.Collection("employees").FindOne(ctx, bson.M{"tgid": authEmployee.Ref}).Decode(&a)

	inviteLink = a.TgInviteLink
	_, _ = eRepo.mongoDB.Collection("employees").UpdateOne(ctx, bson.M{"tgid": authEmployee.Ref}, bson.M{"tgInviteLink": ""})

	return http.StatusOK, nil, inviteLink
}

func (eRepo *EmployeeRepositoryImpl) GetAllEmployees() (httpCode int, repoErr error, employers []models.Employee) {
	cur, mongoErr := eRepo.mongoDB.Collection("employees").
		Find(context.TODO(), bson.D{})
	defer cur.Close(ctx)

	if mongoErr != nil {
		repoErr = fmt.Errorf("Ошибка извлечения всех участнико: %s", mongoErr.Error())
		return http.StatusInternalServerError, repoErr, nil
	}

	for cur.Next(ctx) {
		var e models.Employee
		if decodeErr := cur.Decode(&e); decodeErr != nil {
			repoErr = fmt.Errorf("Ошибка анмаршалинга одного участника: %s", decodeErr.Error())
			return http.StatusInternalServerError, repoErr, nil
		}
		employers = append(employers, e)
	}

	return http.StatusOK, nil, employers
}
