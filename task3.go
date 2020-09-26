package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func collectWords(line string) []string {
	var word string
	var words []string
	for _, c := range line {
		if c != ' ' {
			word += string(c)
		} else {
			words = append(words, word)
			word = ""
		}
	}
	return words
}
func sortWords(words []string) []string {
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if words[i] < words[j] {
				tmpWord := words[i]
				words[i] = words[j]
				words[j] = tmpWord
			}
		}
	}
	return words
}

func writeToOutput(sortedWords []string) string {
	var output string
	stringSpace := " "
	for _, word := range sortedWords {
		for _, c := range word {
			output = output + string(c)
		}
		output = output + stringSpace
	}
	output += "\n"
	return output
}

func createOutputFile(output string) {
	file, _ := os.Open("data.txt")
	file, err := os.Create("res.txt") // write output to the file
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range []byte(output) {
		_, err2 := file.Write([]byte{c})
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}

func main() {
	var buffer = make([]byte, 1)
	file, _ := os.Open("data.txt")
	var err error
	var line = ""
	// Read loop
	for ; err == nil; _, err = file.Read(buffer) {
		// I think we need to wait here for character to be read
		time.Sleep(time.Millisecond) // appending each byte to a line
		if buffer[0] != '\n' {
			line = line + string(buffer[0])
		} else {
			fmt.Println(line)
			// Collect words for each line
			words := collectWords(line)
			// Sort the words
			sortedWords := sortWords(words)
			// Write words to the output buffer
			output := writeToOutput(sortedWords)
			// Create output file
			createOutputFile(output)
		}
	}
}
