package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token      = "your-token-here"
	challenges = NewChallenges()
	statistics = NewStatistics()
)

func ConnectToDiscord() {

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore messages from bot himself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// !challengeBot - bot will always accept challenges
	if m.Content == "!challengeBot" {
		handleChallengeBot(s, m)
		return
	}

	// !challenge <username>
	if strings.HasPrefix(m.Content, "!challenge") && len(strings.Split(m.Content, " ")) == 2 {
		handleChallenge(s, m)
		return
	}

	// !accepct <username>
	if strings.HasPrefix(m.Content, "!accept") && len(strings.Split(m.Content, " ")) == 2 {
		handleAcceptChallenge(s, m)
		return
	}

	// !challenges
	if m.Content == "!challenges" {
		handleChallenges(s, m)
		return
	}

	// !leaderboard
	if m.Content == "!leaderboard" {
		handleLeaderboard(s, m)
		return
	}

	// !help
	if m.Content == "!help" {
		handleHelp(s, m)
		return
	}

	// help as default
	handleHelp(s, m)
}

func handleChallengeBot(s *discordgo.Session, m *discordgo.MessageCreate) {
	f1 := NewFighter("The Battle Bot")
	f2 := NewFighter(m.Author.Username)

	result := Fight(f1, f2)

	statistics.AddStatistic(result.Winner.Username, result.Loser.Username)

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s has beaten %s\n", result.Winner.Username, result.Loser.Username))
}

func handleChallenge(s *discordgo.Session, m *discordgo.MessageCreate) {

	splits := strings.Split(m.Content, " ")
	if len(splits) != 2 {
		s.ChannelMessageSend(m.ChannelID, "wrong usage of !challenge. Use !challenge <username> to challenge the user with <username>")
		return
	}

	user1 := m.Author.Username
	user2 := splits[1]

	exists := challenges.ChallengeExists(user1, user2)
	if exists {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("a challenge between you and %s already exists\nYou can accept the challenge using !accept %s", user2, user2))
		return
	}

	err := challenges.AddChallengers(user1, user2)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("an error occured: %s", err.Error()))
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("you have challenged %s to an epic fight. He has to accecpt now", user2))
}

func handleAcceptChallenge(s *discordgo.Session, m *discordgo.MessageCreate) {
	splits := strings.Split(m.Content, " ")
	if len(splits) != 2 {
		s.ChannelMessageSend(m.ChannelID, "wrong usage of !accept. Use !accept <username> to accept a challenge from the user with <username>")
		return
	}

	user1 := m.Author.Username
	user2 := splits[1]

	exists := challenges.ChallengeExists(user1, user2)
	if !exists {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("no challenge between you and %s exists\nYou can challenge start a challenge with !challenge %s", user2, user2))
		return
	}

	f1 := NewFighter(user1)
	f2 := NewFighter(user2)

	result := Fight(f1, f2)

	statistics.AddStatistic(result.Winner.Username, result.Loser.Username)

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s has beaten %s\n", result.Winner.Username, result.Loser.Username))

	err := challenges.RemoveChallengers(user1, user2)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("an error occured: %s", err.Error()))
		return
	}

}

func handleChallenges(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author.Username

	challengers := challenges.OpenChallengers(user)

	if len(challengers) == 0 {
		s.ChannelMessageSend(m.ChannelID, "you have no open challenges")
		return
	}

	msg := "Your open challenges are:\n"

	for _, challenger := range challengers {
		msg = fmt.Sprintf("%s\t%s\n", msg, challenger)
	}

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleLeaderboard(s *discordgo.Session, m *discordgo.MessageCreate) {
	stats := statistics.GetStatistics()

	msg := "Leaderboard:\n"

	for idx, stat := range stats {
		msg = fmt.Sprintf("%s%d. Place - %s\n", msg, idx+1, stat.Username)
	}

	s.ChannelMessageSend(m.ChannelID, msg)
}

func handleHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "commands:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "!help - shows all available commands")
	msg = fmt.Sprintf("%s\t%s\n", msg, "!challengeBot - bot will always accept challenges")
	msg = fmt.Sprintf("%s\t%s\n", msg, "!challenges - shows all your open challenges")
	msg = fmt.Sprintf("%s\t%s\n", msg, "!challenge <username> - challenge another user")
	msg = fmt.Sprintf("%s\t%s\n", msg, "!accepct <username> - accept a challenge from another user")
	msg = fmt.Sprintf("%s\t%s\n", msg, "!leaderboard - shows the leaderboard")

	s.ChannelMessageSend(m.ChannelID, msg)
}
