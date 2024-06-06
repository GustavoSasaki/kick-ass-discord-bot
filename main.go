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

	messagesChannel := make(chan messageEnter)

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		messagesChannel <- messageEnter{s, m}
	})

	session.Open()
	defer session.Close()

	loopHandleMessages(messagesChannel)

}

type messageEnter struct {
	session *discordgo.Session
	message *discordgo.MessageCreate
}

func handleMessage(m messageEnter) {
	message := m.message
	session := m.session

	if message.Author.ID == session.State.User.ID {
		return
	}

	if message.Content == "l" || message.Content == "L" {
		session.ChannelMessageSend(message.ChannelID, "⠂⠈⠈⢹⠉⠀⠁⠀⠁⠀")
		session.ChannelMessageSend(message.ChannelID, "⠀⠐⠀⢸⠀⠀⠁⡀⠁")
		session.ChannelMessageSend(message.ChannelID, "⠐⠀⢈⣸⣀⣈⣠⡇⠀")

		return
	}
}

func loopHandleMessages(c chan messageEnter) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	for {
		select {
		case message := <-c:
			handleMessage(message)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
