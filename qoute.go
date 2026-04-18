package main

import (
	"strings"
)

func fixQuote(s string) string {
	word := strings.Split(s, "'")
	for i := 1; i < len(word); i += 2 {
		word[i] = strings.TrimSpace(word[i])
	}
	return strings.Join(word, "'")

}

// func main() {
// 	fmt.Println(fixQuote("read '     books    ' often ' understood?  '"))

// }
