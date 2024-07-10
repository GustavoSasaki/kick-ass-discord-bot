package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func StartDiscordSession() (*discordgo.Session, error) {
	godotenv.Load()
	token := os.Getenv("DISCORD_TOKEN")
	return discordgo.New("Bot " + token)
}
