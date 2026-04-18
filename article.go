package main

import (

	"strings"
)

func fixArticle(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word)-1; i++ {
		vow := strings.ContainsAny("aeiouhAEIOUH", string(word[i+1][:1]))
		if word[i] == "a" && vow {
			word[i] = "an"
		}else if word[i] == "A" && vow {
			word[i] = "An"
		}else if word[i] == "an" && !vow {
			word[i] = "a"
		}else if word[i] == "An" && !vow {
			word[i] = "A"
		}

	}

	return strings.Join(word, " ")
}

// func main() {
// 	k := "A orage is good, but sometimes try a apple and A egg and an mango with An grape, spice it with A honey"
// 	fmt.Println(fixArticle(k))
// }
