package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/iuan95/apigo/db"
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

}