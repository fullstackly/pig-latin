package translator

import (
	"bytes"
	"strings"
)

// wordTranslate translates single word according to pig-latin rules
func wordTranslate(word string) string {
	var buf bytes.Buffer

	if strings.Contains(allConsonants, string(word[0])) {
		prefix, base := splitPrefix(word)
		buf.WriteString(base + prefix + consEnding)
	} else if strings.Contains(allVowels, string(word[0])) {
		buf.WriteString(word + vowelEnding)
	}

	return buf.String()
}

// splitPrefix splits a word for prefix & base according to big-latin rules
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
