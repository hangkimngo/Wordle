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

	// if index < 0 || index >= len(playgame.WordList) {
	// 	fmt.Printf("Index must be between 0 and %d\n", len(playgame.WordList)-1)
	// 	return
	// }

	scanner := bufio.NewScanner(os.Stdin)
	// model.StatFile()
	game.StartScreen(scanner, index, &playgame)
	//Print stats
	fmt.Print("Do you want to see your stats? (yes/no)")
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
