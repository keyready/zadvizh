package usecases

import (
	"server/internal/domain/repositories"
	"server/internal/domain/types/models"
	"server/internal/domain/types/request"
)

type TeacherUsecase interface {
	GetAllTeachers() (httpCode int, usecaseError error, teachers []models.Teacher)
	WriteComment(newComment request.WriteNewComment) (httpCode int, usecaseErr error)
	LikeDislike(likeDislike request.LikeDislike) (httpCode int, usecaseErr error)
}

type TeacherUsecaseImpl struct {
	teacherRepo repositories.TeacherRepository
}

func NewTeacherRepostoryImpl(teacherRepo repositories.TeacherRepository) *TeacherUsecaseImpl {
	return &TeacherUsecaseImpl{teacherRepo: teacherRepo}
}

func (tUsecase *TeacherUsecaseImpl) LikeDislike(likeDislike request.LikeDislike) (httpCode int, usecaseErr error) {
	httpCode, usecaseErr = tUsecase.teacherRepo.LikeDislike(likeDislike)
	if usecaseErr != nil {
		return httpCode, usecaseErr
	}
	return httpCode, nil
}

func (tUsecase *TeacherUsecaseImpl) WriteComment(newComment request.WriteNewComment) (httpCode int, usecaseErr error) {
	httpCode, usecaseErr = tUsecase.teacherRepo.WriteComment(newComment)
	if usecaseErr != nil {
		return httpCode, usecaseErr
	}
	return httpCode, nil
}

func (tUsecase *TeacherUsecaseImpl) GetAllTeachers() (httpCode int, usecaseError error, teachers []models.Teacher) {
	httpCode, usecaseError, teachers = tUsecase.teacherRepo.GetAllTeachers()
	if usecaseError != nil {
		return httpCode, usecaseError, nil
	}
	return httpCode, nil, teachers
}
