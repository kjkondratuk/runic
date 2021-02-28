package main

import (
	"fmt"
	"github.com/kjkondratuk/runic-cipher/static"
	"os"
	"strings"
)

var (
	defaultExcludes = []string {
		"LONG-BRANCH",
		"SHORT-TWIG",
		"DOTTED",
		"PUNCTUATION",
		"OPEN",
		"ICELANDIC",
		"FRANKS-CASKET",
		"TOLKEIN",
		"GOLDEN-NUMBER",
	}
)

func main() {
	text := os.Args[1]

	// invert the map first so we can go from characters to runes
	reverse := invertMap(static.RunicAlphabet, defaultExcludes...)

	for k, v := range reverse {
		//if k == "S" {
			fmt.Println(k, string(v))
		//}
	}

	var glyph rune
	for _, r := range text {
		if r == ' ' {
			glyph = r
		} else if r == '.' {
			glyph = '\u16eb'
		} else {
			glyph = reverse[strings.ToUpper(string(r))][0]
		}
		fmt.Print(string(glyph))
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsAny(s []string, e ...string) bool {
	for _, i := range e {
		if contains(s, i) {
			return true
		}
	}
	return false
}

func invertMap(in map[rune][]string, exclusions ...string) map[string][]rune {
	newMap := make(map[string][]rune)
	for k, value := range in {
		if !containsAny(value, exclusions...) {
			for _, v := range value {
				if len(newMap[v]) > 0 {
					newMap[v] = append(newMap[v], k)
				} else {
					newArr := make([]rune, 0)
					newMap[v] = append(newArr, k)
				}
			}
		}
	}
	return newMap
}