package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
)

func main() {
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
				inviteLink := os.Getenv("LINK_TEMPLATE") + strconv.FormatInt(authorLink, 10)

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
