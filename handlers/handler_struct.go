
package handlers

import "github.com/bwmarrin/discordgo"

type HandlersProps struct {
	msgPingPongHanlder func(s *discordgo.Session, m *discordgo.MessageCreate)