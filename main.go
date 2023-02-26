package main

import (
	"log"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/bson/primitive"



	"github.com/gofiber/fiber/v2"
)

type MongoInstance struct {
	Client *mongo.Client
	Db	   *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017" + dbName

type Employee struct {
	ID     string
	Name   string
	Salary float64
	Age    float64
}

func Connect() error {
	client , err := mongo.NewClient(options.Client.ApplyURL(mongoURI))
	ctx,cancel := context.WithTimeout(context.Background(),30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db: db,
	}
	return nil


}

func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {

	})
	app.Post("/employee")
	app.Put("/employee/:id")
	app.Delete("/employee/:id")

}
