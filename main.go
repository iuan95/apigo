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
	route.InitRoute(app)
	app.Listen(":3001")
}