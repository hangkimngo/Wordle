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
	var feedback strings.Builder
	remaining := "Remaining letters: "
	for i, ch := range game.guess_letter {
		if strings.ContainsRune(game.secret_word, ch) {
			if i < len(game.secret_word) && byte(game.secret_word[i]) == byte(ch) {
				feedback.WriteString(c.Green + strings.ToUpper(string(ch)) + c.Reset)
			} else {
				feedback.WriteString(c.Yellow + strings.ToUpper(string(ch)) + c.Reset)
			}

		} else {
			feedback.WriteString(strings.ToUpper(string(ch)))
			game.seen[ch-'a'+'A'] = false
		}
	}
	for r := 'A'; r <= 'Z'; r++ {
		if game.seen[r] == true {
			remaining += string(r) + " "
		}
	}

	fmt.Println("Feedback:", feedback.String())
	fmt.Println(remaining)
	fmt.Println("Attempts remaining: ", game.attempts)

}
