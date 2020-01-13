
package handlers

import (
	"fmt"
	"strings"

	"github.com/Leoff00/go-diego-bot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	str        string
	arr        []string
	huf        *HandlerUtilFunctions
	responseAi = make(chan *AiResponse)
	errC       = make(chan error)
)

func (h *HandlersProps) MessagePingPong() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == config.BotPrefix+"ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")