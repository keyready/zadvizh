package main

import (
	"log"
	"net/http"
	"server/internal/api/routers"
	"server/pkg/mongoose"
)

func main() {
	mongoClient, _ := mongoose.GetMongoClient()
	appHandlers := routers.AppRouters(mongoClient.Database("zadvizh"))

	server := &http.Server{
		Addr:    ":5000",
		Handler: appHandlers,
	}

	log.Println("Сервер запущен на порту 5000")

	log.Fatalf("Ошибка запуска сервера: %v", server.ListenAndServe().Error())
}
