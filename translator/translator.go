package translator

import (
	"bytes"
	"strings"
)

const (
	consonants  = "B, C, D, F, G, H, J, K, L, M, N, P, Q, R, S, T, V, W, X, Z"
	vowels      = "A, E, I, O, U, Y"
	consEnging  = "ay"
	vowelEnding = "yay"
)

var (
	allConsonants = consonants + strings.ToLower(consonants)
	allVowels     = vowels + strings.ToLower(vowels)
)

// Translate translates english text according to pig-latin rules
func Translate(text string) string {
	var buf bytes.Buffer

	if strings.Contains(allConsonants, string(text[0])) {
		prefix, base := splitPrefix(text)
		buf.WriteString(base)
		buf.WriteString(prefix)
		buf.WriteString(consEnging)
	} else if strings.Contains(allVowels, string(text[0])) {
		buf.WriteString(text)
		buf.WriteString(vowelEnding)
	}

	return buf.String()
}

func splitPrefix(word string) (prefix, base string) {
	var subBuf bytes.Buffer

	for i, let := range word {
		if strings.Contains(allVowels, string(let)) {
			prefix = subBuf.String()
			base = word[i:]
			return
		}
		subBuf.WriteRune(let)
	}

	base = word
	return
}
