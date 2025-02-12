package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/internal/api/routers"
	"server/pkg/mongoose"
	"time"
)

func main() {
	//botApi, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	//if err != nil {
	//	log.Fatalf("Ошибка инициализации бота: %s", err.Error())
	//}
	//myBot := &bot.Bot{
	//	Bot: botApi,
	//}

	mongoClient, _ := mongoose.GetMongoClient()
	appHandlers := routers.AppRouters(mongoClient.Database(os.Getenv("MONGODB")))

	server := &http.Server{
		WriteTimeout: time.Second * 120,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Addr:         fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
		Handler:      appHandlers,
	}

	//myBot.Run()

	log.Fatalf("Ошибка запуска сервера: %v", server.ListenAndServe().Error())

	log.Printf("Сервер запущен на порту: %s", os.Getenv("SERVER_PORT"))
}
