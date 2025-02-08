package main

import (
	"bot/botdb"
	"context"
	"encoding/base64"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
)

func main() {
	mongoClient, _ := botdb.GetMongoClient()

	bot, botInitErr := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if botInitErr != nil {
		log.Fatalf("Ошибка инициализации бота: %s", botInitErr.Error())
	}

	log.Printf("Бот инициализирован с именем %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			log.Print("Обновление произошло без сообщения")
			continue
		}

		if len(update.Message.NewChatMembers) > 0 {
			authorRefTgId := strconv.FormatInt(update.Message.From.ID, 64)
			newMemberTgId := strconv.FormatInt(update.Message.NewChatMembers[0].ID, 64)

			mongoClient.Database(os.Getenv("MONGO_APP_NAME")).
				Collection("employees").
				InsertOne(context.Background(),
					botdb.Employee{
						TgId: newMemberTgId,
						Ref:  authorRefTgId,
					},
				)

			newMember := update.Message.NewChatMembers[0]

			welcomeMessage := fmt.Sprintf("Добро пожаловать, %s! 🙌", newMember.UserName)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcomeMessage)

			_, sendErr := bot.Send(msg)
			if sendErr != nil {
				log.Fatalf("Ошибка отправления приветственного сообщения: %s", sendErr.Error())
			}
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "invite":
				authorLink := update.Message.From.ID
				authorLinkB64 := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(authorLink, 10)))
				inviteLink := os.Getenv("LINK_TEMPLATE") + authorLinkB64

				msgBody := fmt.Sprintf("Ссылка на вступление: %s", inviteLink)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgBody)

				_, err := bot.Send(msg)
				if err != nil {
					log.Fatalf("Ошибка выполнения команды /invite: %s", err.Error())
				}

				log.Print("Ссылка приглашение успешно сгенерирована")
			default:
				log.Printf("Неизвестная команда: %s", update.Message.Command())
			}
		}

	}
}
