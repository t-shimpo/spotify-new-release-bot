package main

import (
	"context"
	"log"
	"os"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	res, err := client.NewReleases(ctx, spotify.Country("US"), spotify.Limit(10))
	if err != nil {
		log.Fatal(err)
	}

	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	firstMessage := linebot.NewTextMessage("新作リリース情報をお届けします！")
	if _, err := bot.BroadcastMessage(firstMessage).Do(); err != nil {
		log.Fatal(err)
	}

	var messages []linebot.SendingMessage
	// NOTE: 1回のリクエストで送信できるメッセージオブジェクトは5つ
	for _, v := range res.Albums {
		messages = append(messages, linebot.NewTextMessage(v.ExternalURLs["spotify"]))
		if len(messages) == 5 {
			if _, err := bot.BroadcastMessage(messages...).Do(); err != nil {
				log.Fatal(err)
			}
			messages = nil
		}
	}
}
