package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Print("Server Start")
	token := os.Getenv("DISCORD_TOKEN")
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
		return
	}

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "l" || m.Content == "L" {
			s.ChannelMessageSend(m.ChannelID, ".__ ")
			s.ChannelMessageSend(m.ChannelID, "|  |")
			s.ChannelMessageSend(m.ChannelID, "|  | ")
			s.ChannelMessageSend(m.ChannelID, "|  |__")
			s.ChannelMessageSend(m.ChannelID, "|____/")
			return
		}
	})

	session.Open()
	defer session.Close()

	endServer := make(chan os.Signal, 1)
	signal.Notify(endServer, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-endServer
}
