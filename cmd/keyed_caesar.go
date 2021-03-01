package main

import (
	"flag"
	"fmt"
	"github.com/kjkondratuk/runic-cipher/static"
	"os"
	"strconv"
	"strings"
)

func main() {
	decode := flag.Bool("d", false, "-d")
	encode := flag.Bool("e", false, "-e")
	key := flag.String("k", "", "-k <string>")
	text := flag.String("t", "", "-t <string>")
	flag.Parse()

	if *encode {
		newString := ""
		words := strings.Fields(*text)
		fmt.Printf("Length: %d\n", len(words))
		for i, word := range words {
			newWord := ""
			keyVal := *key
			shift, err := strconv.ParseInt(keyVal[i:i+1], 10, 64)
			fmt.Printf("Shift: %d\n", shift)
			for _, letter := range word {
				if err != nil {
					fmt.Printf("%s is not a valid key shift value", keyVal[i:i+1])
				}
				currIndex := indexOf(static.EnglishAlphabet, strings.ToUpper(string(letter)))
				newIndex := int64(currIndex) + shift
				if newIndex > int64(len(static.EnglishAlphabet) - 1) {
					newIndex = newIndex - int64(len(static.EnglishAlphabet) - 1)
				}
				newWord += string(static.EnglishAlphabet[newIndex].Rune)
			}
			newString += newWord + " "
		}

		fmt.Printf("New string is: [%s]", newString)
	} else if *decode {
		os.Exit(-1)
	}
}

func indexOf(list []static.RunicRune, letter string) int {
	for i, item := range list{
		if letter == string(item.Rune) {
			return i
		}
	}
	return -1
}