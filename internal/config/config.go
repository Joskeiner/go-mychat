package config

import (
	"log"
	"os"

	"github.com/joskeiner/go-myChat/pkg/env"
)

func LoandingDeps() (string, string) {
	err := env.LoadEnv()
	if err != nil {
		log.Fatal("Error loanding .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	adr := ":" + port
	basePath, err := os.Getwd()
	// log.Println(basePath)
	if err != nil {
		log.Fatal(err)
	}
	return adr, basePath
}
