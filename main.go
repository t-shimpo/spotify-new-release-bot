package main

import (
	"context"
	"fmt"
	"log"
	"os"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2/clientcredentials"
	// "github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// bot, err := linebot.New(
	// 	os.Getenv("LINE_CHANNEL_SECRET"),
	// 	os.Getenv("LINE_CHANNEL_TOKEN"),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// message := linebot.NewTextMessage("Hello World")

	// if _, err := bot.BroadcastMessage(message).Do(); err != nil {
	// 	log.Fatal(err)
	// }

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

	res, err := client.NewReleases(ctx, spotify.Country("JP"), spotify.Limit(10))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
