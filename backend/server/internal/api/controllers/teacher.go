package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/domain/types/request"
	"server/internal/domain/usecases"
)

type TeacherController struct {
	teacherUsecase usecases.TeacherUsecase
}

func NewTeacherControllers(teacherUsecase usecases.TeacherUsecase) *TeacherController {
	return &TeacherController{teacherUsecase: teacherUsecase}
}

func (tContr *TeacherController) LikeDislike(ctx *gin.Context) {
	var likeDislike request.LikeDislike
	authorId, _ := ctx.Get("tgId")

	likeDislike.AuthorID = authorId.(string)

	bindErr := ctx.ShouldBindJSON(&likeDislike)
	if bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
		return
	}

	httpCode, contrErr := tContr.teacherUsecase.LikeDislike(likeDislike)
	if contrErr != nil {
		ctx.AbortWithStatusJSON(httpCode, gin.H{"error": contrErr.Error()})
		return
	}

	ctx.JSON(httpCode, gin.H{})
}

func (tContr *TeacherController) WriteComment(ctx *gin.Context) {
	var newComment request.WriteNewComment

	authorID, _ := ctx.Get("tgId")
	newComment.AuthorID = authorID.(string)

	bindErr := ctx.ShouldBindJSON(&newComment)
	if bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
		return
	}

	httpCode, contrErr := tContr.teacherUsecase.WriteComment(newComment)
	if contrErr != nil {
		ctx.AbortWithStatusJSON(httpCode, gin.H{"error": contrErr.Error()})
		return
	}

	ctx.JSON(httpCode, gin.H{})
}

func (tController *TeacherController) GetAllTeachers(ctx *gin.Context) {
	httpCode, contrErr, teachers := tController.teacherUsecase.GetAllTeachers()
	if contrErr != nil {
		ctx.AbortWithStatusJSON(httpCode, gin.H{"error": contrErr.Error()})
		return
	}
	ctx.JSON(httpCode, teachers)
}
