package main

import (
	"regexp"
	"strings"
)

func fixPunct(s string) string {

	re := regexp.MustCompile(`\s+([,.;:?!])`)
	s = re.ReplaceAllString(s, "$1")

	reg := regexp.MustCompile(`([,.;:?!])([a-zA-Z0-9])`)
	s = reg.ReplaceAllString(s, "$1 $2")
	return strings.Join(strings.Fields(s), " ")

}
