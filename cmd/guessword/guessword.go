package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	fmt.Print("\nHangman v1.0\n")

	possibleWords := loadWordList("word_list.txt")
	word := []rune(possibleWords[rand.Intn(len(possibleWords))])
	numLives := 5
	matchedLetters := []rune(strings.Repeat("_", len(word)))
	foundMatch := false

	for numLives != 0 {
		displayMatchedLetters(matchedLetters)

		letter := getLetter()

		foundMatch = false
		for x, char := range word {
			if letter == char {
				matchedLetters[x] = letter
				foundMatch = true
			}
		}

		if !foundMatch {
			numLives -= 1
			fmt.Printf("Letter '%s' is in not word. %d lives remaining!\n", string(letter), numLives)
		}

		if numLives == 0 {
			fmt.Println("You lose. Try again.")
			os.Exit(0)
		}

		if string(word) == string(matchedLetters) {
			fmt.Printf("'%s' is correct. You win!\n", string(matchedLetters))
			os.Exit(0)
		}
	}

}

func getLetter() rune {
	numLettersEntered := 0
	letter := []rune("")

	for numLettersEntered != 1 {
		fmt.Print("Enter a letter: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}

		numLettersEntered = len(scanner.Text())

		switch {
		case numLettersEntered == 1:
			letter = []rune(scanner.Text())
		case numLettersEntered < 1:
			fmt.Println("No letter entered! Please enter a single letter.")
		case numLettersEntered > 1:
			fmt.Println("Too many letters! Please enter a single letter.")
		}
	}

	return letter[0]
}

func displayMatchedLetters(matchedLetters []rune) {
	fmt.Print("\nMatched Letters: ")
	for _, letter := range matchedLetters {
		fmt.Printf(" %s ", string(letter))
	}
	fmt.Println()
}

func loadWordList(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var wordList []string
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}

	return wordList
}
