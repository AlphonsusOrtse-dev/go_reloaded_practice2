package main

import (

	"strconv"
	"strings"
)

func hexDec(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if word[i] == "(hex)" {
			value, err := strconv.ParseInt(word[i-1], 16, 64)
			
			if err == nil {
				word[i-1] = strconv.FormatInt(value, 10)
			}
			word = append(word[:i], word[i+1:]...)
		}
	}
	return strings.Join(word, " ")

}
// func main() {
// 	fmt.Println(hexDec("there were 1E (hex) files added"))
// }
