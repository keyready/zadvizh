package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
	"net/http"
	"server/internal/domain/types/models"
	"server/internal/domain/types/request"
	"server/internal/domain/types/response"
	"server/pkg/utils"
)

type TeacherRepository interface {
	GetAllTeachers() (httpCode int, repoErr error, teachers []response.Teacher)
	WriteComment(newComment request.WriteNewComment) (httpCode int, repoErr error)
	LikeDislike(likeDislike request.LikeDislike) (httpCode int, repoErr error)
	LikeDislikeComment(likeDislike request.LikeDislikeComment) (httpCode int, repoErr error)
}

type TeacherRepositoryImpl struct {
	mongoDB *mongo.Database
}

func NewTeacherRepositoryImpl(mongoDB *mongo.Database) *TeacherRepositoryImpl {
	return &TeacherRepositoryImpl{mongoDB: mongoDB}
}

func (tRepo *TeacherRepositoryImpl) LikeDislikeComment(likeDislike request.LikeDislikeComment) (httpCode int, repoErr error) {
	commentID, _ := bson.ObjectIDFromHex(likeDislike.CommentID)
	authorID, _ := bson.ObjectIDFromHex(likeDislike.AuthorID)

	var c models.Comment
	_ = tRepo.mongoDB.Collection("comments").FindOne(context.TODO(), bson.M{"_id": commentID}).Decode(&c)

	switch likeDislike.Action {
	case "like":
		if !(utils.Contains(c.Dislikes.Authors, authorID)) {
			if !utils.Contains(c.Likes.Authors, authorID) {
				c.Likes.Value += 1
				c.Likes.Authors = append(c.Likes.Authors, authorID)
			} else {
				c.Likes.Value -= 1
				c.Likes.Authors = utils.RemoveElement(authorID, c.Likes.Authors)
			}
		}
	case "dislike":
		if !utils.Contains(c.Likes.Authors, authorID) {
			if !utils.Contains(c.Dislikes.Authors, authorID) {
				c.Dislikes.Value += 1
				c.Dislikes.Authors = append(c.Dislikes.Authors, authorID)
			} else {
				c.Dislikes.Value -= 1
				c.Dislikes.Authors = utils.RemoveElement(authorID, c.Dislikes.Authors)
			}
		}
	}

	_, mongoErr := tRepo.mongoDB.Collection("comments").ReplaceOne(context.TODO(), bson.M{"_id": commentID}, c)
	if mongoErr != nil {
		repoErr = fmt.Errorf("Ошибка лайка комментария: %s", mongoErr.Error())
		return http.StatusInternalServerError, repoErr
	}

	return http.StatusOK, nil
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
		if !(utils.Contains(teacher.Dislikes.Authors, author.ID)) {
			if !utils.Contains(teacher.Likes.Authors, author.ID) {
				teacher.Likes.Value += 1
				teacher.Likes.Authors = append(teacher.Likes.Authors, author.ID)
			} else {
				teacher.Likes.Value -= 1
				teacher.Likes.Authors = utils.RemoveElement(author.ID, teacher.Likes.Authors)
			}
		}
	case "dislike":
		if !utils.Contains(teacher.Likes.Authors, author.ID) {
			if !utils.Contains(teacher.Dislikes.Authors, author.ID) {
				teacher.Dislikes.Value += 1
				teacher.Dislikes.Authors = append(teacher.Dislikes.Authors, author.ID)
			} else {
				teacher.Dislikes.Value -= 1
				teacher.Dislikes.Authors = utils.RemoveElement(author.ID, teacher.Dislikes.Authors)
			}
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

	newCom, _ := tRepo.mongoDB.Collection("comments").
		InsertOne(context.Background(),
			models.Comment{
				Content: newComment.Content,
				Author:  author.ID,
			})

	filter := bson.M{"_id": teacherId}
	update := bson.M{
		"$push": bson.M{"comments": newCom.InsertedID.(bson.ObjectID)},
	}

	_, mongoErr := tRepo.mongoDB.Collection("teachers").
		UpdateOne(ctx, filter, update)
	if mongoErr != nil {
		repoErr = fmt.Errorf("Ошибка добавления нового комментария: %s", mongoErr.Error())
		return http.StatusInternalServerError, repoErr
	}

	return http.StatusOK, nil
}

func (tRepo *TeacherRepositoryImpl) GetAllTeachers() (httpCode int, repoErr error, teachers []response.Teacher) {
	tCollection := tRepo.mongoDB.Collection("teachers")

	pipeline := mongo.Pipeline{
		bson.D{
			{
				"$lookup", bson.D{
					{"from", "comments"},
					{"localField", "comments"},
					{"foreignField", "_id"},
					{"as", "comments"},
				},
			},
		},
	}

	cur, mongoErr := tCollection.Aggregate(context.Background(), pipeline)
	defer cur.Close(context.Background())

	if mongoErr != nil {
		return http.StatusInternalServerError, mongoErr, nil
	}

	for cur.Next(ctx) {
		var t response.Teacher
		e := cur.Decode(&t)
		if e != nil {
			log.Fatalf(e.Error())
		}
		teachers = append(teachers, t)
	}

	return http.StatusOK, nil, teachers
}
