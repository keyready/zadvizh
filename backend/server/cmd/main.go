package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/internal/api/routers"
	"server/pkg/mongoose"
)

func main() {
	mongoClient, _ := mongoose.GetMongoClient()
	appHandlers := routers.AppRouters(mongoClient.Database(os.Getenv("MONGO_APP_NAME")))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
		Handler: appHandlers,
	}

	log.Fatalf("Ошибка запуска сервера: %v", server.ListenAndServe().Error())

	log.Printf("Сервер запущен на порту: %s", os.Getenv("SERVER_PORT"))
}
