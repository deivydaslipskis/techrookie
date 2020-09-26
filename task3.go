package main

import (
	"os"
	"time"
)

func main() {
	var buffer = make([]byte, 1)
	file, _ := os.Open("data.txt")
	var err error
	var line = ""
	var output = "" // Read loop
	for ; err == nil; _, err = file.Read(buffer) {
		// I think we need to wait here for character to be read
		time.Sleep(time.Millisecond) // appending each byte to a line
		if buffer[0] != '\n' {
			line = line + string(buffer[0])
			// Line is done, let's sort it
		} else {
			// Collect words for each line
			var word = ""
			var words []string
			for _, c := range line {
				if c != ' ' {
					word = word + string(c)
				} else {
					words = append(words, word)
					word = ""
				}
			} // Sort the words
			for i := 0; i < len(words); i++ {
				for j := 0; j < len(words); j++ {
					if words[i] < words[j] {
						tmpWord := words[i]
						words[i] = words[j]
						words[j] = tmpWord
					}
				}
			}
			line = "" // write words to the output buffer
			for _, word := range words {
				for _, c := range word {
					output = output + string(c)
				}
				output = output + " "
			}
			output = output + "\n"
		}
	} // Create output file
	file, _ = os.Create("res.txt") // write output to the file
	for _, c := range []byte(output) {
		file.Write([]byte{c})
	}
}
