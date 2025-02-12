package main

import (
	mongoosepackage "bot/mongoose"
	"context"
	"encoding/base64"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
	"strconv"
	"time"
)

type Employee struct {
	ID         bson.ObjectID `bson:"_id" json:"id"`
	Firstname  string        `bson:"firstname" json:"firstname"`
	Lastname   string        `bson:"lastname" json:"lastname"`
	Department string        `bson:"department" json:"department"`
	Field      string        `bson:"field" json:"field"`
	Position   string        `bson:"position" json:"position"`
	TeamName   string        `bson:"teamname" json:"teamName"`
	TeamRole   string        `bson:"teamrole" json:"teamRole"`
	Scidir     string        `bson:"scidir" json:"scidir"`

	TgId string `bson:"tgid" json:"tgId"`
	Ref  string `bson:"ref" json:"ref"`
}

type Bot struct {
	Bot *tgbotapi.BotAPI
}

func (b *Bot) Run() {
	mongoClient, _ := mongoosepackage.GetMongoClient()

	employeesCollection := mongoClient.Database("zadvizh").Collection("employees")
	var lastKnownId bson.ObjectID

	checkNewEmployees := func() {
		for {
			var lastEmployee Employee

			findOptions := options.FindOne().
				SetSort(bson.D{{"_id", -1}})
			err := employeesCollection.FindOne(context.Background(), bson.D{}, findOptions).Decode(&lastEmployee)
			if err != nil {
				fmt.Printf("Ошибка поиска последней записи: %s", err.Error())
			}

			if lastKnownId != lastEmployee.ID {
				msg := tgbotapi.NewMessage(-1002438510106, "")
				msg.Text = fmt.Sprintf(`В нашем коллективе пополнение!
Встречаем %s %s (%s)!
Род деятельности: %s на позиции %s`, lastEmployee.Firstname,
					lastEmployee.Lastname,
					lastEmployee.Department,
					lastEmployee.Field,
					lastEmployee.Position)

				_, sendErr := b.Bot.Send(msg)
				if sendErr != nil {
					log.Fatalf("Ошибка отправки приветствия: %s", sendErr.Error())
				}
				lastKnownId = lastEmployee.ID
			}
			time.Sleep(time.Second * 5)
		}
	}

	go checkNewEmployees()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.Bot.GetUpdatesChan(u)

	for update := range updates {
		chatMember, err := b.Bot.GetChatMember(tgbotapi.GetChatMemberConfig{
			tgbotapi.ChatConfigWithUser{
				ChatID: update.Message.Chat.ID,
				UserID: update.Message.From.ID,
			},
		})
		if err != nil {
			log.Fatalf("Ошибка получения информации о статусе участника чата: %s", err.Error())
		}

		if update.Message != nil && update.Message.IsCommand() {
			if update.Message.Command() == "invite" {
				if chatMember.Status == "administrator" || chatMember.Status == "creator" {
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

					_, sendErr := b.Bot.Send(msg)
					if sendErr != nil {
						log.Fatalf("Ошибка выполнения команды /invite: %s", sendErr.Error())
					}
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
					msg.Text = "У вас нет прав на выполнение данной команды"

					_, sendErr := b.Bot.Send(msg)
					if sendErr != nil {
						log.Fatalf("Ошибка отправки сообщения: %s", sendErr.Error())
					}
				}
			}
		}

		if update.CallbackQuery != nil {
			callback := update.CallbackQuery

			if callback.Message == nil {
				continue
			}

			if callback.Data == "delete_message" && (chatMember.Status == "administrator" || chatMember.Status == "creator") {
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

func main() {
	token := "7950443390:AAFaHXb6prhYCBFXpiLEXdKAav9dz4GrGrk"
	bot, initErr := tgbotapi.NewBotAPI(token)
	if initErr != nil {
		log.Fatalf("Ошибка инициализации бота: %s", initErr.Error())
	}

	myBot := &Bot{
		Bot: bot,
	}

	myBot.Run()
}
