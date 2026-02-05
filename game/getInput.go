package game

import (
	"bufio"
	"fmt"
	"strings"
)

func GetInput(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Println("Scan() returned false (EOF or stdin closed)")
		return ""
	}
	return strings.TrimSpace(scanner.Text())
}
