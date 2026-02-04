package game

import (
	"bufio"
	"fmt"
	"strings"
)

func GetInput(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Println("Scan() returned false (EOF or stdin closed)")
		return "-1"
	}
	input := scanner.Text()
	input = strings.TrimSpace(input)
	return input
}
