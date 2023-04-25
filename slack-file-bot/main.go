package main

import (
	"fmt"
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
	fileArr := []string{"sample.png"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return

		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)

	}

}
