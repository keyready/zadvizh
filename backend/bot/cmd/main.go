package main

import (
	mongoosepackage "bot/mongoose"
	"context"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
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
				continue
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
		if update.Message != nil && update.Message.IsCommand() {
			chatMember, err := b.Bot.GetChatMember(tgbotapi.GetChatMemberConfig{
				tgbotapi.ChatConfigWithUser{
					ChatID: update.Message.Chat.ID,
					UserID: update.Message.From.ID,
				},
			})
			if err != nil {
				log.Fatalf("Ошибка получения информации о статусе участника чата: %s", err.Error())
			}
			if update.Message.Command() == "invite" {
				if chatMember.Status == "administrator" || chatMember.Status == "creator" {
					linkConfig := tgbotapi.CreateChatInviteLinkConfig{
						ChatConfig: tgbotapi.ChatConfig{
							ChatID: update.Message.Chat.ID,
						},
						ExpireDate:         int(time.Now().Add(10 * time.Minute).Unix()),
						CreatesJoinRequest: true,
						Name:               "Ссылка-приглашение",
					}

					var responseApi map[string]interface{}
					resp, getErr := b.Bot.Request(linkConfig)
					if getErr != nil {
						log.Fatalf("Ошибка генерации ссылки-приглашения: %s", getErr.Error())
					}

					decodeErr := json.Unmarshal(resp.Result, &responseApi)
					if decodeErr != nil {
						log.Fatalf("Ошибка получения ссылки-приглашения: %s", getErr.Error())
					}

					inviteLink := responseApi["invite_link"].(string)

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
		} else if update.CallbackQuery != nil {
			chatMember, err := b.Bot.GetChatMember(tgbotapi.GetChatMemberConfig{
				tgbotapi.ChatConfigWithUser{
					ChatID: update.CallbackQuery.Message.Chat.ID,
					UserID: update.CallbackQuery.From.ID,
				},
			})
			if err != nil {
				log.Fatalf("Ошибка получения инфы о юзере: %s", err.Error())
			}
			if update.CallbackQuery.Data == "delete_message" {
				if chatMember.Status == "administrator" || chatMember.Status == "creator" {
					messageID := update.CallbackQuery.Message.MessageID
					chatID := update.CallbackQuery.Message.Chat.ID
					inviteLink := update.CallbackQuery.Message.Text

					revokeInviteLinkCfg := tgbotapi.RevokeChatInviteLinkConfig{
						ChatConfig: tgbotapi.ChatConfig{
							ChatID: chatID,
						},
						InviteLink: inviteLink,
					}

					_, apiErr := b.Bot.Request(revokeInviteLinkCfg)
					if apiErr != nil {
						log.Fatalf("Ошибка отзыва ссылки: %s", apiErr.Error())
					}

					editMsg := tgbotapi.NewEditMessageText(chatID, messageID, "Сообщение удалено")
					editMsg.ParseMode = tgbotapi.ModeHTML

					nilInlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
					editMsg.ReplyMarkup = &nilInlineKeyboard

					_, sendErr := b.Bot.Send(editMsg)
					if sendErr != nil {
						log.Fatalf("Ошибка при изменении текста: %s", sendErr.Error())
					}

					callBackConfig := tgbotapi.CallbackConfig{
						CallbackQueryID: update.CallbackQuery.ID,
						Text:            "Сообщение удалено",
						ShowAlert:       false,
					}

					_, apiErr = b.Bot.Request(callBackConfig)
					if apiErr != nil {
						log.Fatalf("Ошибка ответа на callback: %s", apiErr.Error())
					}
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
