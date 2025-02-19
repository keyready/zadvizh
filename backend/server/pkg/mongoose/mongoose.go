package mongoose

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"io/ioutil"
	"log"
	"server/internal/domain/types/dto"
	"server/internal/domain/types/models"
	"sync"
)

var (
	clientInstance *mongo.Client
	clientErr      error
	mongoOnce      sync.Once
	ctx            context.Context
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		mongoUri := "mongodb://admin:admin@mongodb:27017/zadvizh?ssl=false&authSource=admin"
		clientOptions := options.Client().
			ApplyURI(mongoUri)

		clientInstance, clientErr = mongo.Connect(clientOptions)
		if clientErr != nil {
			log.Fatalf("Ошибка подключения к MongoDB: %v", clientErr.Error())
		}

		if pingErr := clientInstance.Ping(ctx, nil); pingErr != nil {
			log.Fatalf("Ошибка пинга MongoDB: %v", pingErr.Error())
		}

		adminsData, err := ioutil.ReadFile("/app/data/admins.json")
		if err != nil {
			log.Fatalf("Ошибка чтения файла с данными админов: %v", err.Error())
		}

		var admins []models.Employee
		dErr := json.Unmarshal(adminsData, &admins)
		if dErr != nil {
			log.Fatalf("Ошибка анмаршалинга json-файла: %v", dErr.Error())
		}

		for _, admin := range admins {
			admin.ID = bson.NewObjectID()

			var res bson.M
			mongoErr := clientInstance.Database("zadvizh").
				Collection("employees").FindOne(ctx, bson.D{
				{"firstname", admin.Firstname},
			}).Decode(&res)
			if mongoErr != mongo.ErrNoDocuments {
				continue
			}

			_, insertErr := clientInstance.Database("zadvizh").
				Collection("employees").
				InsertOne(ctx, admin)
			if insertErr != nil {
				log.Fatalf("Ошибка добавления в БД нового препода: %v", insertErr.Error())
			}
		}

		var teachersData, readErr = ioutil.ReadFile("/app/data/teachers.json")
		if readErr != nil {
			log.Fatalf("Ошибка чтения файла с данными преподавателей: %v", readErr.Error())
		}

		var teachers []models.Teacher
		decodeErr := json.Unmarshal(teachersData, &teachers)
		if decodeErr != nil {
			log.Fatalf("Ошибка анмаршалинга json-файла: %v", decodeErr.Error())
		}

		for _, t := range teachers {
			t.ID = bson.NewObjectID()
			t.Comments = []dto.Comment{}
			t.Likes = dto.Like{
				Authors: []bson.ObjectID{},
			}
			t.Dislikes = dto.Dislike{
				Authors: []bson.ObjectID{},
			}

			var res bson.M
			mongoErr := clientInstance.Database("zadvizh").
				Collection("teachers").FindOne(ctx, bson.D{
				{"firstname", t.Firstname},
				{"middlename", t.Middlename},
				{"lastname", t.Lastname},
			}).Decode(&res)
			if mongoErr != mongo.ErrNoDocuments {
				continue
			}

			_, insertErr := clientInstance.Database("zadvizh").
				Collection("teachers").
				InsertOne(ctx, t)
			if insertErr != nil {
				log.Fatalf("Ошибка добавления в БД нового препода: %v", insertErr.Error())
			}
		}

		log.Println("Подключение к MongoDB успешно")
	})
	return clientInstance, clientErr
}
