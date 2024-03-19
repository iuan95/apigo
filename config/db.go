package config

import (
	"fmt"
	"os"
)

const (
    port     = 5432

)

func GetDbConfig() string {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), port, os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	return psqlInfo
}