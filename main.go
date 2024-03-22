package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/iuan95/apigo/db"
	"github.com/iuan95/apigo/route"
)

func init(){
	err:=godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	  }
	
}

func main(){
	err:=db.Connection(context.Background())
	if err != nil{
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer db.DB.Close()
	app:= fiber.New()
	app.Use(func(c *fiber.Ctx) error {
	   c.Set("Access-Control-Allow-Origin", "*")
	   c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	   c.Set("Access-Control-Allow-Headers", "*")
	   if c.Method() == "OPTIONS" {
		   return c.SendStatus(fiber.StatusOK)
	   }
	   return c.Next()
	})
	
	route.InitRoute(app)
	app.Listen(":3001")
}