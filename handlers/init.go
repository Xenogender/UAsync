
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
		}
		if m.Content == config.BotPrefix+"pong" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "ping")
		}
	}
}

func (h *HandlersProps) Img() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content != "" && strings.HasPrefix(m.Content, config.BotPrefix+"picture") {
			data := huf.ParamSeparator(m.Content)

			go huf.PicGenerator(data, responseAi, errC)

			select {
			case res := <-responseAi:
				var ogSize string

				for _, p := range res.Photos {
					ogSize = p.Src.Original
				}
				_, _ = s.ChannelMessageSend(m.ChannelID, "Aqui esta o que você pediu! \n"+ogSize)

			case err := <-errC:
				if err != nil {
					_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())