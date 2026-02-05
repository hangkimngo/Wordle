package game

import (
	"fmt"
	"slices"
)

func ValidWord(word_list []string, game *Game) bool {
	if !slices.Contains(word_list, game.guess_letter) {
		fmt.Println("Word not in list. Please enter a valid word.")
		return false
	} else {
		return true
	}
}
