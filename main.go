package main

import (
	"fmt"
	"strings"
)

var (
	words         string
	wordCharSlice []string
	wordState     []string
	blankWord     []string
	letter        string
	guesses       int
	inputChoice   string
	gameState     bool
)

func newGame(gameState bool) {
	if gameState == true {

		fmt.Println("Random word choice in process..")
		// words = []string{"batman", "holla", "giraffe", "superman"}
		words = "batman"
		wordCharSlice = strings.Split(words, "")
		for i := 0; len(wordCharSlice) > i; i++ {

			blankWord = append(blankWord, "_")
			// buffer.WriteString("_ ")
		}
		fmt.Println(blankWord)
		fmt.Println("Please choose a letter")
		guesses = 1
		fmt.Scanf("%s\n", &letter)
		gameState = false
		guess(blankWord, wordCharSlice, letter, guesses)
	}

	letter = ""
	guess(blankWord, wordCharSlice, letter, guesses)
}

func guess(blankWord []string, wordCharSlice []string, letter string, guesses int) {
	fmt.Println("b4 guess")

	fmt.Println(guesses)
	for guesses < 10 {
		fmt.Println("a guess")
		if letter == "" {
			fmt.Println("Please choose a lettera")
			fmt.Scanf("%s\n", &letter)

		}
		var i = 0
		for _, l := range wordCharSlice {
			if letter == l {
				blankWord[i] = l

			}
			i++

		}
		fmt.Println(blankWord)

		gameState = false
		newGame(gameState)
	}
}

func main() {

	//to start game
	fmt.Println("Type 'start' to begin.")

	//input of user
	fmt.Scanf("%s\n", &inputChoice)

	//check if true
	if inputChoice == "start" {
		gameState := true
		newGame(gameState)
	}

}
