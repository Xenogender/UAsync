
package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Leoff00/go-diego-bot/config"
	"github.com/Leoff00/go-diego-bot/handlers"
	"github.com/bwmarrin/discordgo"
)

var (
	BotID string
	goBot *discordgo.Session
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	handler := handlers.HandlersProps{}

	goBot.AddHandler(handler.MessagePingPong())
	goBot.AddHandler(handler.HelpJava())
	goBot.AddHandler(handler.Greeting())
	goBot.AddHandler(handler.Img())