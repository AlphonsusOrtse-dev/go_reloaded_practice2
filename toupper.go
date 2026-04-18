package main

import (
	"strconv"
	"strings"
)

func toUpper(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		switch word[i] {
		case "(up)":
			word[i-1] = strings.ToUpper(word[i-1])
			word[i] = ""

			for _, r := range word {
				if r == "" {
					word = append(word, r)
				}
			}
		}

		if word[i] == "(up," {
			word[i+1] = strings.Trim(word[i+1], ")")
			value, err := strconv.Atoi(word[i+1])
			if err == nil {
				for j := 1; j <= value && i-j >= 0; j++ {
					word[i-j] = strings.ToUpper(word[i-j])
				}
			}

			word = append(word[:i], word[i+2:]...)

		}
	}
	return strings.Join(word, " ")

}
