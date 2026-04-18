package main

import (
	"strconv"
	"strings"
)

func capitalize(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		switch word[i] {
		case "(cap)":
			word[i-1] = strings.Title(word[i-1])
			word = append(word[:i], word[i+1:]...)
		}
		if word[i] == "(cap," {
			word[i+1] = strings.Trim(word[i+1], ")")
			value, err := strconv.Atoi(word[i+1])
			if err == nil {
			}
			for j := 1; j <= value && i-j >= 0; j++ {
				word[i-j] = strings.Title(word[i-j])
			}
			word = append(word[:i], word[i+2:]...)
		}

	}
	return strings.Join(word, " ")

}


