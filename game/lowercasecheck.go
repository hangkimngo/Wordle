package game

import "fmt"

func LowercaseCheck(game *Game) bool {
	for _, c := range game.guess_letter {
		if c < 'a' || c > 'z' {
			fmt.Println("Your guess must only contain lowercase letters.")
			return false
		}
	}
	return true
}
