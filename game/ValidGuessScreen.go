package game

import (
	"fmt"
	c "koodWordle/constants"
	"strings"
)

func ValidGuessScreen(game *Game) bool {

	if game.secret_word == game.guess_letter {
		return true
	}

	Feedback(game)
	return false

}

func Feedback(game *Game) {
	feedback := ""
	remaining := "Remaining letters: "

	for i, ch := range game.guess_letter {

		letter := strings.ToUpper(string(ch))

		if strings.ContainsRune(game.secret_word, ch) {

			if game.secret_word[i] == byte(ch) {
				feedback += c.Green + letter + c.Reset
			} else {
				feedback += c.Yellow + letter + c.Reset
			}

		} else {
			feedback += c.White + letter + c.Reset
			game.seen[ch-'a'+'A'] = false
		}
	}

	for r := 'A'; r <= 'Z'; r++ {
		if game.seen[r] {
			remaining += string(r) + " "
		}
	}

	fmt.Println("Feedback:", feedback)
	fmt.Println(remaining)
	fmt.Printf("Attempts remaining:  %d\n", game.attempts)
}
