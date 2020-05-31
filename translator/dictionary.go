package translator

import "strings"

const (
	consonants = "B C D F G H J K L M N P Q R S T V W X Z"
	vowels     = "A E I O U Y"

	consEnding  = "ay"
	vowelEnding = "yay"
)

var (
	allConsonants = consonants + strings.ToLower(consonants)
	allVowels     = vowels + strings.ToLower(vowels)

	alphabet = allConsonants + allVowels
)
