package main

import (
	"strconv"
	"strings"
)

func HandleMessage(m messageEnter) {
	message := m.message
	session := m.session

	if message.Author.ID == session.State.User.ID {
		return
	}

	args := strings.Split(message.Content, " ")
	if len(args) < 2 || args[0] != "meido" {
		return
	}

	if len(args) >= 3 && args[1] == "kick" {
		minutes, err := strconv.Atoi(args[2])
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "For kick command say meido kick [number of minutes until kick]")
			session.ChannelMessageSend(message.ChannelID, "eg: meido kick 4")
			return
		}

		//find which chat channel command was requested
		chat, err := session.State.Channel(message.ChannelID)
		if err != nil {
			return
		}

		// Find the guild for that channel.
		guild, err := session.State.Guild(chat.GuildID)
		if err != nil {
			return
		}

		// Look for the message sender in that guild's current voice states and kick all user in same voice chat
		for _, vs := range guild.VoiceStates {
			if vs.UserID == message.Author.ID {
				session.MessageReactionAdd(message.ChannelID, message.ID, "🧹")
				session.ChannelMessageSend(message.ChannelID, "よろしく お願いします♥主人様")
				go DoKick(minutes, session, guild, vs.ChannelID)
				go DoTalk(minutes, session, vs.ChannelID, guild.ID)
			}
		}
	}
}
