package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	new           = false
	tries         int
	word          string
	wordCharSlice []string
	wordNumSlice  []string
)

func newGame(new bool) {

	if new == false {
		var letter string
		var counter int

		fmt.Println("NewWordChoice")
		word = "batman"

		wordCharSlice := make([]string, int(len(word)))
		wordNumSlice := make([]string, int(len(word)))
		//print user prompt
		fmt.Print("Enter letter guess: ")

		//get user input
		fmt.Scanf("%s", &letter)
		wordCharSlice = strings.Split(word, "")

		// for i := 0; wordCharSlice > i; i++ {
		// 	buffer.WriteString("_ ")
		// }
		/*
			iterate over word array and user array, comparing to see what,
			if any answer, that the user got right
		*/

		for _, value := range wordCharSlice {
			if letter == value {
				wordNumSlice[i] = letter
			}
			i++
		}

		//isTrue tracks if user entered incorrect answer
		isTrue := false

		it := 0
		for _, value := range wordCharSlice {
			_ = value
			if wordCharSlice[it] == wordNumSlice[it] {
				fmt.Print(wordCharSlice[it], "")
				counter++
				isTrue = true
			} else {
				fmt.Print(" _")
			}
			it++

		}

		tries = 0
		choice(wordChoice)
	} else {
		fmt.Println("Not a New Game!")
	}

}
func choice(wordChoice string) {

	if tries <= 4 {
		fmt.Print("Please choose a letter: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Contains(wordChoice, text) {
			fmt.Println("yes")
			fmt.Println(text)
		} else {
			fmt.Println("no")
			fmt.Println(text)

		}
	}
}

func main() {
	if new == false {
		newGame(false)
	}
	// fmt.Print("Enter choice: ")
	// text, _ := reader.ReadString('\n')
	// text = strings.Replace(text, "\n", "", -1)
	// choice(text)

}

// func f(x int) {

// 	fmt.Println(x)
// }

// func main() {

// 	x := 5
// 	f(x)
// }
