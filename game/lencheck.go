package game

import "fmt"

func LenCheck(game *Game) bool {
	if len(game.guess_letter) == 0 {

		return false
	} else if len(game.guess_letter) != 5 {
		fmt.Println("Your guess must be exactly 5 letters long.")
		return false
	}
	return true
}
