package main

import "github.com/bwmarrin/discordgo"

func Tuturu(vc *discordgo.VoiceConnection) {
	for _, buff := range TuturuBuffer {
		vc.OpusSend <- buff
	}
}

func Gnomed(vc *discordgo.VoiceConnection) {
	for _, buff := range GnomeBuffer {
		vc.OpusSend <- buff
	}
}

func Yasou(vc *discordgo.VoiceConnection) {
	for _, buff := range YasouUltBuffer {
		vc.OpusSend <- buff
	}
}

func Zed(vc *discordgo.VoiceConnection, index int) {
	for _, buff := range ZedBuffer[index] {
		vc.OpusSend <- buff
	}
}
