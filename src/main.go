package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/iqraraminullah/gobot-discord/src/config"
)

var conf config.Configurations = config.InitConfig("config")

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Ignore message created by this bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, conf.Discord.Prefix) {

		//Log the message
		fmt.Printf("[ChannelID]:%s|[User]:%s [Message]:%s (%s)\n",
			m.ChannelID, m.Author.Username, m.Content, time.Now().Format(time.Stamp))

	}

}

func main() {

	//Create new bot session
	var dg, err = discordgo.New()
	if err != nil {
		fmt.Println("error creating Discord Session", err)
		return
	}

	//Set token
	dg.Token = conf.Discord.Token
	//Add message handler
	dg.AddHandler(onMessageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}
