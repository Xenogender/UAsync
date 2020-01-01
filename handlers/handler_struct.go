
package handlers

import "github.com/bwmarrin/discordgo"

type HandlersProps struct {
	msgPingPongHanlder func(s *discordgo.Session, m *discordgo.MessageCreate)
	imgGenerator       func(s *discordgo.Session, m *discordgo.MessageCreate)
	helpJavaHandler    func(s *discordgo.Session, m *discordgo.MessageCreate)
	msgGreeting        func(s *discordgo.Session, m *discordgo.MessageCreate)
	msgHelpCmd         func(s *discordgo.Session, m *discordgo.MessageCreate)
}

type SrcProps struct {
	Original  string `json:"original"`
	Large2X   string `json:"large2x"`
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}

type PhotoProps struct {
	ID              int       `json:"id"`
	Width           int       `json:"width"`
	Height          int       `json:"height"`
	URL             string    `json:"url"`
	Photographer    string    `json:"photographer"`
	PhotographerURL string    `json:"photographer_url"`