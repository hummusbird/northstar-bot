package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("bot.env")
	if err != nil {
		log.Fatal("[FATAL] Error loading environment variables. Make sure you are using the file \"bot.env\"")
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("[FATAL] No Discord Token provided in env as TOKEN")
	}

	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Open()
	log.Println("[INFO] NorthstarServers bot is running")
	defer bot.Close()

	Setup(bot)

	log.Println("[INFO] Bot ready")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("[INFO] Shutting down bot gracefully")
}
