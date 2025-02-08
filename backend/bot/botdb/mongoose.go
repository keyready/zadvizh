package botdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
	"sync"
	"time"
)

var (
	clientInstance *mongo.Client
	clientErr      error
	mongoOnce      sync.Once
)

type Employee struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Firstname  string             `bson:"firstname" json:"firstname"`
	Lastname   string             `bson:"lastname" json:"lastname"`
	Department string             `bson:"department" json:"department"`
	Field      string             `bson:"field" json:"field"`
	TeamName   string             `bson:"teamName" json:"teamName"`
	TeamRole   string             `bson:"teamRole" json:"teamRole"`
	Scidir     string             `bson:"scidir" json:"scidir"`

	TgId string `bson:"tgid" json:"tgid"`
	Ref  string `bson:"ref" json:"ref"`
}

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientOptions := options.Client().
			ApplyURI(os.Getenv("MONGO_URI")).
			SetAppName(os.Getenv("MONGO_APP_NAME"))

		clientInstance, clientErr = mongo.Connect(clientOptions)
		if clientErr != nil {
			log.Fatalf("Ошибка подключения к MongoDB: %v", clientErr.Error())
		}

		if pingErr := clientInstance.Ping(ctx, nil); pingErr != nil {
			log.Fatalf("Ошибка пинга MongoDB: %v", pingErr.Error())
		}

		log.Println("Подключение к MongoDB успешно")
	})

	return clientInstance, clientErr
}
