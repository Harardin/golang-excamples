package main

import (
	"strings"
	"unicode"
)

// removeChars removes unwanted chars from a string
func removeChars(r string, chars string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(chars, r) < 0 {
			return r
		}
		return -1
	}
	return strings.Map(filter, r)
}

// removeUNI removes unicode from text `r` is for source string `s` is for replacement string
func removeUNI(source string) string {
	return strings.Map(func(r rune) rune {
		if r > unicode.MaxASCII {
			return -1
		}
		return r
	}, source)
}
