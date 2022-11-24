package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	COUNT_PER_LINE int = 5
)

func main() {

	// Bitten Sie um Eingabe
	fmt.Printf("Gib Hier den Text ein:   ")

	// Eingabe speichern
	in := bufio.NewReader(os.Stdin)

	text, _ := in.ReadString('\n')

	text = strings.ReplaceAll(text, "\n", "")

	words := strings.Split(text, " ")

	var wordsWithLines [100][5]string

	var wordCount int
	var lineCount int
	for _, word := range words {
		wordsWithLines[lineCount][wordCount] = word

		wordCount++

		if wordCount == COUNT_PER_LINE {
			lineCount++
			wordCount = 0
		}
	}

	fmt.Println()
	fmt.Println()

	for k := 0; k < 14; k++ {
		fmt.Printf(" _")
	}

	fmt.Println()

	fmt.Printf("/ \t\t\t     \\")

	fmt.Println()

	for i := 0; i <= lineCount; i++ {
		line := " "
		for j := 0; j < COUNT_PER_LINE; j++ {
			if j == COUNT_PER_LINE-1 {
				line += wordsWithLines[i][j] + " "
			} else {
				line += wordsWithLines[i][j] + ""
			}
		}
		fmt.Printf(line + "\n")
	}

	fmt.Printf("\\\t\t\t     /\n")

	for k := 0; k < 14; k++ {
		fmt.Printf(" -")
	}

	printTux()
}

func printTux() {
	tux := `   			
	                             
	                    \              
			     \   ^__^
			      \  (oo)\_______
				 (__)\       )\/\
				     ||----w |
				     ||     ||`

	fmt.Println(tux)
}
