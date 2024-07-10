package main

import (
	"math/rand/v2"
	"time"

	"github.com/bwmarrin/discordgo"
)

func DoKick(minutes int, session *discordgo.Session, guild *discordgo.Guild, voiceChannelId string) {
	time.Sleep(time.Duration(minutes) * time.Minute)
	//time.Sleep(time.Duration(minutes) * time.Second)

	vc, err := session.ChannelVoiceJoin(guild.ID, voiceChannelId, false, true)
	if err != nil {
		return
	}
	vc.Speaking(true)

	switch kickEvent := rand.IntN(3); kickEvent {
	case 0:
		kickEveryoneNow(session, guild, voiceChannelId, vc)
	case 1:
		kickOneAtTime(session, guild, voiceChannelId, vc)
	case 2:
		kickOneAtTimeZed(session, guild, voiceChannelId, vc)
	}

	vc.Speaking(false)
	time.Sleep(250 * time.Millisecond)
	vc.Disconnect()

}

func kickOneAtTime(session *discordgo.Session, guild *discordgo.Guild, voiceChannelId string, vc *discordgo.VoiceConnection) {

	for {
		guild2, err := session.State.Guild(guild.ID)
		if err != nil {
			return
		}

		voiceStateQuantity := len(guild2.VoiceStates)
		if voiceStateQuantity >= 1 {
			break
		}

		kickIndex := rand.IntN(voiceStateQuantity)
		kickPerson := guild2.VoiceStates[kickIndex]
		if kickPerson.ChannelID == voiceChannelId && kickPerson.UserID == session.State.User.ID {
			continue
		}

		Gnomed(vc)
		time.Sleep(200 * time.Millisecond)
		session.GuildMemberMove(guild.ID, kickPerson.UserID, nil)
		time.Sleep(600 * time.Millisecond)
	}
}

func kickOneAtTimeZed(session *discordgo.Session, guild *discordgo.Guild, voiceChannelId string, vc *discordgo.VoiceConnection) {
	voiceIndex := 0

	for {
		guild2, err := session.State.Guild(guild.ID)
		if err != nil {
			return
		}

		voiceStateQuantity := len(guild2.VoiceStates)
		if voiceStateQuantity >= 1 {
			break
		}

		kickIndex := rand.IntN(voiceStateQuantity)
		kickPerson := guild2.VoiceStates[kickIndex]
		if kickPerson.ChannelID == voiceChannelId && kickPerson.UserID == session.State.User.ID {
			continue
		}

		Zed(vc, voiceIndex)
		voiceIndex = (voiceIndex + 1) % len(ZedBuffer)
		time.Sleep(200 * time.Millisecond)
		session.GuildMemberMove(guild.ID, kickPerson.UserID, nil)
		time.Sleep(600 * time.Millisecond)
	}
}

func kickEveryoneNow(session *discordgo.Session, guild *discordgo.Guild, voiceChannelId string, vc *discordgo.VoiceConnection) {
	guild, err := session.State.Guild(guild.ID)
	if err != nil {
		return
	}

	Yasou(vc)
	for _, vs := range guild.VoiceStates {

		if vs.ChannelID == voiceChannelId && vs.UserID == session.State.User.ID {
			continue
		}
		session.GuildMemberMove(guild.ID, vs.UserID, nil)
	}
}

func DoTalk(minutes int, session *discordgo.Session, voiceChannelId, guildId string) {
	// wait 0.93 percent of minutes
	time.Sleep(time.Duration(minutes) * time.Second * time.Duration(56))

	vc, err := session.ChannelVoiceJoin(guildId, voiceChannelId, false, true)
	if err != nil {
		return
	}

	vc.Speaking(true)
	Tuturu(vc)
	vc.Speaking(false)

	time.Sleep(250 * time.Millisecond)
	vc.Disconnect()
}
