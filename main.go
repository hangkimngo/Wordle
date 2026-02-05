package main

import (
	"bufio"
	"fmt"
	"koodWordle/game"
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

	game.StartScreen(scanner, index, &playgame)

	// model.StatFile()

	//Print stats
	fmt.Print("Do you want to see your stats? (yes/no):")
	ShowStats := game.GetInput(scanner)
	if ShowStats == "yes" {
		fmt.Println("stats for")
	} else {
		fmt.Println("Press Enter to exit...")
	}
	for {
		IsEnter := game.GetInput(scanner)
		if IsEnter == "-1" {
			return
		}
		if IsEnter == "" {
			return
		}
	}
}
