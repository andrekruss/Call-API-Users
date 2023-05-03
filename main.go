package main

import (
	"call-ms-users/controllers"
	"call-ms-users/database"
	"call-ms-users/helpers"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
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

func routesSetup(app *fiber.App) {
	app.Get("/", controllers.GetRoot)
	app.Post("/create-user", controllers.PostUser)
}

func main() {

	// sets env variables through viper package
	readConfigFile()

	// validator
	helpers.Validate = validator.New()

	// connecting to mongoDB
	database.DBConn = database.CreateMongoClient(viper.GetString("database.connectionstring"))
	database.ConnectMongo(database.DBConn)

	// fiber app
	app := fiber.New()
	routesSetup(app)
	app.Listen(fmt.Sprintf(":%v", viper.GetString("app.port")))
}
