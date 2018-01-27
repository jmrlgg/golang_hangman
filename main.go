package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const maxGuesses = 10

type Hangman struct {
	WordState, blankWord, wordCharSlice []string
	maxGuesses, numOfTries              int
	word, gameStatus, letter            string
	verbose                             bool
}

func getWord() string {
	word := "batman"
	return word
}

func (h *Hangman) isMatch(letter string) bool {
	if strings.Contains(h.word, h.letter) {
		if h.verbose {
			fmt.Println("True")
		}
		return true
	} else {
		if h.verbose {
			fmt.Println("false")
		}
		return false
	}
}

func (h *Hangman) getGuess() string {
	if h.numOfTries < 10 {
		fmt.Println(h.word)
		for i := 0; len(h.wordCharSlice) > i; i++ {
			h.blankWord = append(h.blankWord, "_ ")
			// buffer.WriteString("_ ")
		}
		fmt.Println(h.blankWord)
		fmt.Println("Please choose a letter")
		reader := bufio.NewReader(os.Stdin)
		letter, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Sucks")
		}
		h.numOfTries++
		return letter
	}
	fmt.Println("You have reached the limit of guesses!")
	return h.letter

}

func (h *Hangman) updateWord(letter string) {

	if h.letter == "" {

	} else {
		i := 0
		for _, l := range h.wordCharSlice {
			if h.letter == l {
				h.blankWord[i] = l
				fmt.Println(h.blankWord)
			}
			i++
		}
		return
	}
}

func main() {

	verbose := flag.Bool("v", false, "verbose mode for debugging purposes")
	flag.Parse()

	game := Hangman{
		word:       getWord(),
		maxGuesses: maxGuesses,
		verbose:    *verbose,
	}
	fmt.Println("Ready to start game. y/n")

	letter := game.getGuess()

	if game.isMatch(letter) {
		fmt.Println("Rd")
		game.updateWord(letter)
		fmt.Println(letter)
	}

}
