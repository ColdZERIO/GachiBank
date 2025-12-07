package database

import (
	"context"
	"fmt"
	"gachibank/Backend/models"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client    *mongo.Client
	Database  *mongo.Database
	UsersAuth *mongo.Collection
}

func (d Database) GetUserAuth(login, password string) (*models.UserAuthDB, error) {
	var User models.UserAuthDB

	filter := bson.M{
		"login":    login,
		"password": password,
	}

	err := d.UsersAuth.FindOne(context.TODO(), filter).Decode(&User)
	if err == mongo.ErrNoDocuments {
		// Создаем нового пользователя если его нет в БД (этот кусок пропадет когда появится регистрация)
		User = models.UserAuthDB{
			Login:    login,
			Password: password,
		}
		_, err = d.UsersAuth.InsertOne(context.TODO(), User)
		return &User, err
	}
	return &User, err
}

func (d Database) SaveUserAuth(u *models.UserAuthDB) error {
	_, err := d.UsersAuth.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}

	return nil
}

func connectToMongoDB() (*mongo.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("env is not found")
	}

	mongouri := os.Getenv("MONGODB_URI")
	if mongouri == "" {
		mongouri = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(mongouri)

	clientOptions.SetMaxPoolSize(50)
	clientOptions.SetMinPoolSize(10)
	clientOptions.SetMaxConnIdleTime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к MongoDB: %w", err)
	}

	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()

	if err := client.Ping(pingCtx, nil); err != nil {
		return nil, fmt.Errorf("не удалось проверить подключение к MongoDB: %w", err)
	}

	log.Println("Successful connection to the database")
	return client, nil
}

func DatabaseInit() (*Database, error) {
	client, err := connectToMongoDB()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dbName := os.Getenv("MONGODB_NAME")
	if dbName == "" {
		dbName = "Gachibank"
	}

	db := client.Database(dbName)
	return &Database{
		Client:    client,
		Database:  db,
		UsersAuth: db.Collection("AuthCollections"),
	}, nil
}
