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

	fmt.Printf("Converting [%s]...\n", text)

	for _, r := range text {
		for _, c := range static.RunicAlphabet {
			if contains(c.Nemonics, strings.ToUpper(string(r))) {
				fmt.Printf("%s[%s]", string(c.Rune), string(r))
			}
		}
	}

	fmt.Println("\n...Finished!")
	os.Exit(0)
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