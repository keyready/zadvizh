package mongoose

import (
	"context"
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
