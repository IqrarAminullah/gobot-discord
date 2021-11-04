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
	message_utils "github.com/iqraraminullah/gobot-discord/src/utils"
)

var conf config.Configurations = config.InitConfig("config")

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Ignore message created by this bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, conf.Discord.Prefix) {

		command, param := message_utils.ParseMessage(m.Content, conf.Discord.Prefix)

		//Log the message
		fmt.Printf("[ChannelID]:%s|[User]:%s [Command]:%s [Param]:%q (%s)\n",
			m.ChannelID, m.Author.Username, command, param, time.Now().Format(time.Stamp))

		switch command {
		case "test":
			sendMessage(s, m.ChannelID, "HAI HAI HALO")
		}
	}
}

func sendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)

	if err != nil {
		fmt.Println("Error sending message: ", err)
	}
}

func main() {

	//Create new bot session
	var dg, err = discordgo.New("Bot " + conf.Discord.Token)
	if err != nil {
		fmt.Println("error creating Discord Session", err)
		return
	}

	//Add message handler
	dg.AddHandler(onMessageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	sendMessage(dg, "244700376127766528", "test")

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}
