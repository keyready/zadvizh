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
	var update bson.M

	teacherId, _ := bson.ObjectIDFromHex(likeDislike.TeacherID)
	authorId, _ := bson.ObjectIDFromHex(likeDislike.AuthorID)

	var teacher models.Teacher
	_ = tRepo.mongoDB.Collection("teachers").
		FindOne(context.TODO(), bson.M{"_id": teacherId}).Decode(&teacher)

	switch likeDislike.Action {
	case "like":
		if !utils.Contains(teacher.Likes.Authors, authorId) {
			update = bson.M{
				"$inc":  bson.M{"likes.value": 1},
				"$push": bson.M{"likes.authors": authorId},
			}
		} else {
			update = bson.M{
				"$inc":  bson.M{"likes.value": -1},
				"$pull": bson.M{"likes.authors": authorId},
			}
		}
	case "dislike":
		if !utils.Contains(teacher.Likes.Authors, authorId) {
			update = bson.M{
				"$inc":  bson.M{"dislikes.value": 1},
				"$push": bson.M{"dislikes.authors": authorId},
			}
		} else {
			update = bson.M{
				"$inc":  bson.M{"dislikes.value": -1},
				"$pull": bson.M{"dislikes.authors": authorId},
			}
		}
	}

	_, mongoErr := tRepo.mongoDB.Collection("teachers").
		UpdateOne(ctx, bson.M{"_id": teacherId}, update)
	if mongoErr != nil {
		repoErr = fmt.Errorf("Ошибка лайка: %s", mongoErr.Error())
		return http.StatusInternalServerError, repoErr
	}

	return http.StatusOK, nil
}

func (tRepo *TeacherRepositoryImpl) DislikeTeacher() {}

func (tRepo *TeacherRepositoryImpl) WriteComment(newComment request.WriteNewComment) (httpCode int, repoErr error) {
	authorId, _ := bson.ObjectIDFromHex(newComment.AuthorID)
	teacherId, _ := bson.ObjectIDFromHex(newComment.TeacherID)

	filter := bson.M{"_id": teacherId}
	update := bson.M{
		"$push": bson.M{
			"comments": dto.Comment{
				Author:    authorId,
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
