package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("go run . sample.txt result.txt")
		os.Exit(1)
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	read, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)

	}
	data := string(read)
	result := process(data)


	err = os.WriteFile(outputFile, []byte(result), 0664)
	if err != nil {
		fmt.Println("error wtrinting to file")
		os.Exit(1)
	}
}


func process(text string) string {
	word := strings.Split(text, "\n")

	for i := range word {
		word[i] = fixArticle(word[i])
		word[i] = hexBin(word[i])
		word[i] = hexDec(word[i])
		word[i] = capitalize(word[i])
		word[i] = fixPunct(word[i])
		word[i] = fixQuote(word[i])
		word[i] = toLower(word[i])
		word[i] = toUpper(word[i])
	}

	return strings.Join(word, "\n")
}