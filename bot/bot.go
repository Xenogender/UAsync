
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