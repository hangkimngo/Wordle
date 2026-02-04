package game

import (
	"bufio"
	"fmt"
)

type Game struct {
	secret_word  string
	guess_letter string
	seen         map[rune]bool
	attempts     int
	WordList     []string
}

func StartScreen(scanner *bufio.Scanner, index_word int, playgame *Game) {

	//StartScreen
	for {
		fmt.Println("Enter your username:")
		username := GetInput(scanner)

		if username == "-1" {
			fmt.Println("\nGoodbye!")
			return
		}
		if username != "" {
			break
		}
		fmt.Println("Please provide username")
	}

	playgame.attempts = 6

	playgame.secret_word = playgame.WordList[index_word]
	// seen:= map[rune]bool{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H' ,'I', 'J', 'K', 'L', 'M' ,'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	playgame.seen = make(map[rune]bool, 26)
	for i := 0; i < 26; i++ {
		playgame.seen['A'+rune(i)] = true
	}

	// var characters []rune
	// for _,ch:= range word{
	// 	characters= append(characters,ch)
	// }

	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")
	for {
		fmt.Println("Enter your guess:")
		playgame.guess_letter = GetInput(scanner)
		if LenCheck(playgame) {
			if LowercaseCheck(playgame) {
				if ValidWord(playgame.WordList, playgame) {
					if ValidGuessScreen(playgame) {
						fmt.Println("Congratulations!You've guessed the word correctly.")
						fmt.Print("Do you want to see your stats? (yes/no)")
						return
					}
					if playgame.attempts == 0 {
						fmt.Println("Game over. the correct word was:", playgame.secret_word)
						fmt.Print("Do you want to see your stats? (yes/no)")
						return
					}
				}
			}
		}
	}

}
