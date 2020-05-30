package translator

import (
	"bytes"
	"strings"
)

const (
	consonants = "B, C, D, F, G, H, J, K, L, M, N, P, Q, R, S, T, V, W, X, Z" +
		"b, c, d, f, g, h, j, k, l, m, n, p, q, r, s, t, v, w, x, z"
	vowels = "A, E, I, O, U, Y" +
		" a, e, i, o, u, y"
	consEnging  = "ay"
	vowelEnding = "yay"
)

// Translate translates english text according to pig-latin rules
func Translate(text string) string {
	var buf bytes.Buffer

	if strings.Contains(consonants, string(text[0])) {
		prefix, base := splitPrefix(text)
		buf.WriteString(base)
		buf.WriteString(prefix)
		buf.WriteString(consEnging)
	} else if strings.Contains(vowels, string(text[0])) {
		buf.WriteString(text)
		buf.WriteString(vowelEnding)
	}

	return buf.String()
}

func splitPrefix(word string) (prefix, base string) {
	var subBuf bytes.Buffer

	for i, let := range word {
		if strings.Contains(vowels, string(let)) {
			prefix = subBuf.String()
			base = word[i:]
			return
		}
		subBuf.WriteRune(let)
	}

	base = word
	return
}
