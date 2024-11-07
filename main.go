package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Create a new Discord session using the provided bot token.

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env : %v", err)
	}
	// get token from env
	token := os.Getenv("DISCORD_TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Just like the ping pong example, we only care about receiving message
	// events in this example.
	//dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
//
// It is called whenever a message is created but only when it's sent through a
// server as we did not request IntentsDirectMessages.

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	log.Printf("Message reçu : %v de la part de %v", m.Content, m.Author.Username)

	switch m.Content {
	case "ping":
		s.ChannelMessageSend(m.ChannelID, "pong")
	case "help":
		_, err := s.ChannelMessageSend(m.ChannelID, "il n'y a pas d'aide encore")
		if err != nil {
			log.Printf("Impossible d'envoyer l'aide : %v", err)
		}
	default:
		log.Printf("message inconnu : %v", m.Content)
	}
	p := playerLoad(m.Author.Username)
	switch p.State {
	case "needPseudo":
		s.ChannelMessageSend(m.ChannelID, "Renseigne ton pseudo")
		p.State = "waitPseudo"
	case "waitPseudo":
		p.Username = m.Content
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Est tu certain de vouloir %v comme pseudo ?", m.Content))
		p.State = "confirmPseudo"
	case "confirmPseudo":
		switch m.Content {
		case "y", "o", "Y", "O":
			p.State = "active"
		default:
			p.State = "waitPseudo"
			s.ChannelMessageSend(m.ChannelID, "Du coup resaisie ton pseudo")
			// p.msg("coucou")
		}
		p.save()
	case "active":
		// Maintenant j'implémente mon jeux

	}
}
