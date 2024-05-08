package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/udodinho/daily-standup-app/http/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err.Error())
	}
	s := server.NewServer()
	s.Start()
}