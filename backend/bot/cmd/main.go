package main

import (
	"encoding/base64"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
)

type Bot struct {
	Bot *tgbotapi.BotAPI
}

func (b *Bot) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.Bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil && update.Message.IsCommand() {
			if update.Message.Command() == "invite" {
				authorLink := update.Message.From.ID
				authorLinkB64 := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(authorLink, 10)))

				inviteLink := os.Getenv("LINK_TEMPLATE") + authorLinkB64

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				msg.ParseMode = tgbotapi.ModeHTML

				msg.Text = fmt.Sprintf("Сгенерировал <a href=\"%s\">ссылку</a> на вступление. Удалить можно кнопкой ниже.", inviteLink)

				inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Удалить сообщение", "delete_message")),
				)
				msg.ReplyMarkup = inlineKeyboard

				_, err := b.Bot.Send(msg)
				if err != nil {
					log.Fatalf("Ошибка выполнения команды /invite: %s", err.Error())
				}
			}
		}

		if update.CallbackQuery != nil {
			callback := update.CallbackQuery

			if callback.Message == nil {
				continue
			}

			if callback.Data == "delete_message" {
				editMsg := tgbotapi.NewEditMessageText(callback.Message.Chat.ID,
					callback.Message.MessageID, "Сообщение удалено")
				editMsg.ParseMode = tgbotapi.ModeHTML
				editMsg.ReplyMarkup = nil

				if _, err := b.Bot.Send(editMsg); err != nil {
					log.Printf("Ошибка удаления сообщения: %v", err.Error())
				}

				answerCallback := tgbotapi.NewCallback(callback.ID, "")
				if _, err := b.Bot.Request(answerCallback); err != nil {
					log.Printf("Ошибка отправки коллбека удаления: %v", err.Error())
				}
			}
		}
	}
}

func (b *Bot) Send(msg string) {
	chatID := int64(-1002438510106)

	message := tgbotapi.NewMessage(chatID, msg)

	_, sendErr := b.Bot.Send(message)
	if sendErr != nil {
		log.Printf("Ошибка отправки: %s", sendErr.Error())
	}
}

func main() {
	bot, initErr := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if initErr != nil {
		log.Fatal("Ошибка инициализации бота: %s", initErr.Error())
	}

	myBot := &Bot{
		Bot: bot,
	}

	myBot.Run()
}
