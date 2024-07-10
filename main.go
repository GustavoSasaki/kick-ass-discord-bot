package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	session, err := StartDiscordSession()
	if err != nil {
		log.Fatal(err)
		return
	}

	LoadSounds()

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

func loopHandleMessages(c chan messageEnter) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	for {
		select {
		case message := <-c:
			HandleMessage(message)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
