package main

import (
	"call-ms-users/controllers"
	"call-ms-users/database"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func readConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("%s", err))
	}
}

func createMongoClient(connectionString string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func connectMongo(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func routesSetup(app *fiber.App) {
	app.Get("/", controllers.GetRoot)
	app.Post("/create-user", controllers.PostUser)
}

func main() {

	// sets env variables through viper package
	readConfigFile()

	// connecting to mongoDB
	database.DBConn = createMongoClient(viper.GetString("database.connectionstring"))
	connectMongo(database.DBConn)

	// fiber app
	app := fiber.New()
	routesSetup(app)
	app.Listen(fmt.Sprintf(":%v", viper.GetString("app.port")))
}
