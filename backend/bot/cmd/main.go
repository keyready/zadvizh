package main

import (
	"bot/utils"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
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

		if update.ChatMember != nil && update.ChatMember.NewChatMember.Status == "member" {

			handleNewMember(bot, update.ChatMember)
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "invite":
				authorLink := update.Message.From.ID
				inviteLink := utils.GenerateInviteLink(authorLink)

				msgTemplate := fmt.Sprintf("Твоя ссылка на вступление: %s", inviteLink)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgTemplate)

				_, err := bot.Send(msg)
				if err != nil {
					log.Fatalf("Ошибка выполнения команды /invite: %s", err.Error())
				}
			}
		}
	}
}

func handleNewMember(bot *tgbotapi.BotAPI, chatMember *tgbotapi.ChatMemberUpdated) {
	// Извлекаем информацию о новом пользователе
	user := chatMember.NewChatMember.User

	// Формируем приветственное сообщение
	var welcomeMessage string
	if user.UserName != "" {
		welcomeMessage = fmt.Sprintf("Приветствуем вас, @%s!", user.UserName)
	} else {
		welcomeMessage = fmt.Sprintf("Приветствуем вас, %s %s!", user.FirstName, user.LastName)
		if user.LastName == "" {
			welcomeMessage = fmt.Sprintf("Приветствуем вас, %s!", user.FirstName)
		}
	}

	// Отправляем приветственное сообщение в канал
	msg := tgbotapi.NewMessage(chatMember.Chat.ID, welcomeMessage)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка отправки сообщения: %s", err.Error())
	}
}
