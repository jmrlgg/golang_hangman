package main

import (
	"bytes"
	"fmt"
	"strings"
)

var (
	tries int

	words         string
	wordCharSlice []string
	wordNumSlice  []string
	exists        bool
)

func newGame() {

	fmt.Println("Random word choice in process..")
	// words = []string{"batman", "holla", "giraffe", "superman"}
	words = "batman"

	// wordCharSlice := make([]string, int(len(words)))
	// wordNumSlice := make([]string, int(len(words)))
	wordCharSlice = strings.Split(words, "")
	var buffer bytes.Buffer

	for i := 0; len(wordCharSlice) > i; i++ {
		buffer.WriteString("_ ")
	}

	//get user input
	var letter string
	fmt.Scanf("%s\n", &letter)
	if letter == "" {
		fmt.Println("error")
		return
	}

	inArray(letter, wordCharSlice)

}

func inArray(val string, array []string) (exists bool, index int) {
	exists = false
	index = -1

	for i, v := range array {
		if val == v {
			index = i
			exists = true
			fmt.Println("yes")
			return
		}
	}
	fmt.Println("no")
	return
}

// func choice(wordChoice string) {

// 	if tries <= 4 {
// 		fmt.Print("Please choose a letter: ")
// 		reader := bufio.NewReader(os.Stdin)
// 		text, _ := reader.ReadString('\n')
// 		text = strings.Replace(text, "\n", "", -1)

// 		if strings.Contains(wordChoice, text) {
// 			fmt.Println("yes")
// 			fmt.Println(text)
// 		} else {
// 			fmt.Println("no")
// 			fmt.Println(text)

// 		}
// 	}
// }

func main() {

	//to start game
	fmt.Println("Type 'start' to begin.")

	//input of user
	var inputChoice string
	fmt.Scanf("%s\n", &inputChoice)

	//check if true
	if inputChoice == "start" {
		newGame()
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
