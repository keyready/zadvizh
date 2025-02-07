package routers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"server/internal/api/controllers"
	v1 "server/internal/api/routers/v1"
	"server/internal/domain/repositories"
	"server/internal/domain/usecases"
)

func AppRouters(mongoDB *mongo.Database) *gin.Engine {
	router := gin.Default()

	eRepo := repositories.NewEmployeeRepositoryImpl(mongoDB)
	eUsecase := usecases.NewEmployeeUsecase(eRepo)
	eContr := controllers.NewEmployeeControllers(eUsecase)
	v1.NewEmployeeRouters(eContr, router)

	return router
}
