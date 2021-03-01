package main

import (
	"flag"
	"fmt"
	"github.com/kjkondratuk/runic-cipher/static"
	"os"
	"strings"
)

var (
	defaultDialect = ""
)

func main() {
	text := flag.String("t", "<text>", "")
	dialect := flag.String("d", "ELDER-FUTHARK", "ELDER-FUTHARK, MEDIEVAL")
	flag.Parse()

	fmt.Printf("Converting [%s]...\n", *text)

	var convertedString string
	for _, r := range *text {
		var translatedRune string
		for _, c := range static.RunicAlphabet {
			if contains(c.Nemonics, strings.ToUpper(string(r))) && containsAny(c.Tags, *dialect) {
				translatedRune = string(c.Rune)
			}
		}
		if translatedRune != "" {
			//fmt.Printf("%s[%s]", translatedRune, string(r))
			convertedString += translatedRune
		} else {
			convertedString += "#" + translatedRune
		}
	}

	if strings.Contains(convertedString, "#") {
		fmt.Printf("ERROR: Could not convert [%s] to runic error at position [%d]", convertedString, strings.Index(convertedString, "#"))
		os.Exit(-1)
	}

	fmt.Printf("\n%s\n", convertedString)

	fmt.Println("...Finished!")
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