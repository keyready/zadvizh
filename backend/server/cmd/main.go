package main

import (
	"log"
	"net/http"
	"os"
	"server/internal/api/routers"
	"server/pkg/mongoose"
	"time"
)

func main() {
	mongoClient, _ := mongoose.GetMongoClient()
	appHandlers := routers.AppRouters(mongoClient.Database(os.Getenv("MONGODB")))

	server := &http.Server{
		WriteTimeout: time.Second * 120,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Addr:         ":5000",
		Handler:      appHandlers,
	}

	log.Println("Сервер запущен на порту 5000")

	log.Fatalf("Ошибка запуска сервера: %v", server.ListenAndServe().Error())
}
