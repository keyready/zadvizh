package main

import (
	mongoosepackage "bot/mongoose"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
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

	TgId       string `bson:"tgid" json:"tgId"`
	Ref        string `bson:"ref" json:"ref"`
	InviteLink string `bson:"invitelink" json:"inviteLink"`
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
						ExpireDate:         int(time.Now().Add(3 * time.Minute).Unix()),
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
						log.Fatalf("Ошибка получения ссылки-приглашения: %s", decodeErr.Error())
					}

					authorId := strconv.Itoa(int(update.Message.From.ID))
					bodyRef := "author:" + authorId
					ref := base64.StdEncoding.EncodeToString([]byte(bodyRef))
					inviteLink := "https://zadvizh.tech/?ref=" + ref

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
					msg.ParseMode = tgbotapi.ModeHTML

					msg.Text = fmt.Sprintf("Сгенерировал <a href=\"%s\">ссылку</a> на регистрацию. Удалить можно кнопкой ниже.", inviteLink)

					deleteButton := tgbotapi.NewInlineKeyboardButtonData("Отозвать ссылку", "delete_message")
					row := []tgbotapi.InlineKeyboardButton{deleteButton}
					inlineKeyboard := [][]tgbotapi.InlineKeyboardButton{row}
					msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard...)

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

					editMsg := tgbotapi.NewEditMessageText(chatID, messageID, "Сообщение удалено")
					editMsg.ParseMode = tgbotapi.ModeHTML

					editMsg.ReplyMarkup = nil

					_, sendErr := b.Bot.Send(editMsg)
					if sendErr != nil {
						log.Fatalf("Ошибка при изменении текста: %s", sendErr.Error())
					}

					callBackConfig := tgbotapi.CallbackConfig{
						CallbackQueryID: update.CallbackQuery.ID,
						Text:            "Сообщение удалено",
						ShowAlert:       false,
					}

					_, apiErr := b.Bot.Request(callBackConfig)
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
