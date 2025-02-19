package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"server/internal/domain/types/dto"
	"server/internal/domain/types/models"
	"server/internal/domain/types/request"
	"server/pkg/utils"
	"time"
)

type TeacherRepository interface {
	GetAllTeachers() (httpCode int, repoErr error, teachers []models.Teacher)
	WriteComment(newComment request.WriteNewComment) (httpCode int, repoErr error)
	LikeDislike(likeDislike request.LikeDislike) (httpCode int, repoErr error)
}

type TeacherRepositoryImpl struct {
	mongoDB *mongo.Database
}

func NewTeacherRepositoryImpl(mongoDB *mongo.Database) *TeacherRepositoryImpl {
	return &TeacherRepositoryImpl{mongoDB: mongoDB}
}

func (tRepo *TeacherRepositoryImpl) LikeDislike(likeDislike request.LikeDislike) (httpCode int, repoErr error) {
	teacherId, _ := bson.ObjectIDFromHex(likeDislike.TeacherID)

	var author models.Employee
	_ = tRepo.mongoDB.Collection("employees").
		FindOne(context.Background(), bson.D{{"tgid", likeDislike.AuthorID}}).Decode(&author)

	var teacher models.Teacher
	_ = tRepo.mongoDB.Collection("teachers").
		FindOne(context.Background(), bson.D{{"_id", teacherId}}).Decode(&teacher)

	switch likeDislike.Action {
	case "like":
		if !utils.Contains(teacher.Likes.Authors, author.ID) {
			teacher.Likes.Value += 1
			teacher.Likes.Authors = append(teacher.Likes.Authors, author.ID)
		} else {
			teacher.Likes.Value -= 1
			teacher.Likes.Authors = utils.RemoveElement(author.ID, teacher.Likes.Authors)
		}
	case "dislike":
		if !utils.Contains(teacher.Dislikes.Authors, author.ID) {
			teacher.Dislikes.Value += 1
			teacher.Dislikes.Authors = append(teacher.Dislikes.Authors, author.ID)
		} else {
			teacher.Dislikes.Value -= 1
			teacher.Dislikes.Authors = utils.RemoveElement(author.ID, teacher.Dislikes.Authors)
		}
	}

	_, mongoErr := tRepo.mongoDB.Collection("teachers").ReplaceOne(ctx, bson.M{"_id": teacherId}, teacher)
	if mongoErr != nil {
		return http.StatusInternalServerError, fmt.Errorf("Ошибка лайка/дизлайка: %s", mongoErr.Error())
	}

	return http.StatusOK, nil
}

func (tRepo *TeacherRepositoryImpl) WriteComment(newComment request.WriteNewComment) (httpCode int, repoErr error) {
	var author models.Employee
	_ = tRepo.mongoDB.Collection("employees").FindOne(ctx, bson.M{"tgid": newComment.AuthorID}).Decode(&author)
	teacherId, _ := bson.ObjectIDFromHex(newComment.TeacherID)

	filter := bson.M{"_id": teacherId}
	update := bson.M{
		"$push": bson.M{
			"comments": dto.Comment{
				Author:    author.ID,
				Content:   newComment.Content,
				CreatedAt: time.Now(),
			},
		},
	}

	_, mongoErr := tRepo.mongoDB.Collection("teachers").
		UpdateOne(ctx, filter, update)
	if mongoErr != nil {
		repoErr = fmt.Errorf("Ошибка добавления нового комментария: %s", mongoErr.Error())
		return http.StatusInternalServerError, repoErr
	}

	return http.StatusOK, nil
}

func (tRepo *TeacherRepositoryImpl) GetAllTeachers() (httpCode int, repoErr error, teachers []models.Teacher) {
	cur, mongoErr := tRepo.mongoDB.Collection("teachers").
		Find(ctx, bson.D{})
	defer cur.Close(ctx)

	if mongoErr != nil {
		repoErr = fmt.Errorf("Ошибка извлечения всех преподавателей: %s", mongoErr.Error())
		return http.StatusInternalServerError, repoErr, nil
	}

	for cur.Next(ctx) {
		var t models.Teacher
		decodeErr := cur.Decode(&t)
		if decodeErr != nil {
			repoErr = fmt.Errorf("Ошибка анмаршалинга препода: %s", decodeErr.Error())
			return http.StatusInternalServerError, repoErr, nil
		}
		teachers = append(teachers, t)
	}

	return http.StatusOK, nil, teachers
}
