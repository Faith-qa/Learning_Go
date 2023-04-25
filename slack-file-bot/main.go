package main

import (
	"log"
	"os"

	"github.com/slack-go/slack"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error loading env file", err)

	}
	log.Println("env successfully loaded")
}

func main() {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{""}

}
