package game

import (
	"bufio"
	"fmt"
	"strings"
)

type Game struct {
	secret_word  string
	guess_letter string
	seen         map[rune]bool
	attempts     int
	WordList     []string
}

func User(scanner *bufio.Scanner) string {
	fmt.Print("Enter your username: ")
	if !scanner.Scan() {
		return ""
	}
	return strings.TrimSpace(scanner.Text())
}

func StartScreen(scanner *bufio.Scanner, index_word int, playgame *Game) (int, string, string) {
	attemptsUsed := 0
	playgame.attempts = 6
	playgame.secret_word = playgame.WordList[index_word]

	playgame.seen = make(map[rune]bool, 26)
	for i := 0; i < 26; i++ {
		playgame.seen['A'+rune(i)] = true
	}

	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")

	for {
		fmt.Print("Enter your guess:")

		guess, ok := GetInput(scanner)
		if !ok { // EOF
			return attemptsUsed, "loss", playgame.secret_word
		}
		playgame.guess_letter = guess

		if LenCheck(playgame) && LowercaseCheck(playgame) && ValidWord(playgame.WordList, playgame) {
			attemptsUsed++

			if playgame.secret_word == playgame.guess_letter {
				fmt.Println("Congratulations!You've guessed the word correctly.")
				return attemptsUsed, "win", playgame.secret_word
			}
			playgame.attempts--
			Feedback(playgame)

			if playgame.attempts == 0 {
				fmt.Println("Game over. the correct word was:", playgame.secret_word)
				return attemptsUsed, "loss", playgame.secret_word
			}
		}
	}
}

// if LenCheck(playgame) {
// 	if LowercaseCheck(playgame) {
// 		if ValidWord(playgame.WordList, playgame) {
// 			if ValidGuessScreen(playgame) {
// 				fmt.Println("Congratulations!You've guessed the word correctly.")
// 				return
// 			}
// 			if playgame.attempts == 0 {
// 				fmt.Println("Game over. the correct word was:", playgame.secret_word)
// 				return
// 			}
// 		}
// 	}
// }
