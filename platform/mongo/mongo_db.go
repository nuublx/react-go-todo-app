package mongodb

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client                           *mongo.Client
	TodosCollection, UsersCollection *mongo.Collection
)

func OpenMongoConnection() {
	dbURI := os.Getenv("MONGO_DB_URI")
	clientOptions := options.Client().ApplyURI(dbURI)

	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	clientOptions.SetMaxConnecting(uint64(maxConn))

	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_TIME"))
	clientOptions.SetMaxConnIdleTime(time.Duration(maxIdleConn))

	maxPoolSizeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_POOL_SIZE"))
	clientOptions.SetMaxPoolSize(uint64(maxPoolSizeConn))

	Client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = Client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	TodosCollection = Client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("TODOS_COLLECTION_NAME"))
	UsersCollection = Client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("USERS_COLLECTION_NAME"))
}
