package main

import (
	"bufio"
	"fmt"
	"koodWordle/game"
	"koodWordle/store"
	"os"
	"strconv"
)

func main() {
	var playgame game.Game
	playgame.WordList = game.GetFile()

	if len(os.Args) != 2 {
		fmt.Println("Please provide a number as command line argument")
		return
	}

	index_word := os.Args[1]
	index, err := strconv.Atoi(index_word)

	if err != nil {
		fmt.Println("Invalid command-line argument. Please launch with a valid number.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	username := game.User(scanner)
	if username == "" {
		return
	}

	_ = username

	if index < 0 || index >= len(playgame.WordList) {
		fmt.Println("Invalid word number.\nPress Enter to exit...")
		return
	}

	attemptsUsed, outcome, secret := game.StartScreen(scanner, index, &playgame)

	_ = store.AppendGameStat("stats.csv", store.GameStat{
		Username: username,
		Secret:   secret, // or playgame.secret_word if you expose it
		Attempts: attemptsUsed,
		Outcome:  outcome,
	})

	fmt.Println("Do you want to see your stats? (yes/no):")
	show, ok := game.GetInput(scanner)
	if !ok {
		return
	}
	if show == "yes" {
		us, _ := store.ReadUserStats("stats.csv", username)
		fmt.Printf("Stats for %s:\n", username)
		fmt.Printf("Games played: %d\n", us.Played)
		fmt.Printf("Games won: %d\n", us.Won)
		fmt.Printf("Average attempts per game: %.2f\n", us.AvgAttempts)
		fmt.Println("Press Enter to exit...")
		_, _ = game.GetInput(scanner)
	} else {
		fmt.Println("Press Enter to exit...")
		_, _ = game.GetInput(scanner)
	}

}
