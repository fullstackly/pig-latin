package translator

import (
	"bytes"
	"strings"
)

// word interpretation as : punct[0] + base_word + punct[1]
type token struct {
	word            string
	preSym, postSym string
}

func (t *token) String() string {
	return t.preSym + t.word + t.postSym
}

// TranslateText translates english text according to pig-latin rules
func TranslateText(text string) string {
	var tokens []token

	for _, chunk := range strings.Fields(text) {
		tokens = append(tokens, tockenize(chunk))
	}

	for i, token := range tokens {
		if len(token.word) != 0 {
			tokens[i].word = wordTranslate(token.word)
		}
	}

	var pigWords []string

	for _, token := range tokens {
		pigWords = append(pigWords, token.String())
	}

	return strings.Join(pigWords, " ")
}

func tockenize(input string) (output token) {
	var buf bytes.Buffer

	for _, let := range input {
		if strings.Contains(alphabet, string(let)) {
			buf.WriteRune(let)
		}
	}

	output.word = buf.String()
	switch puncts := strings.Split(input, output.word); len(puncts) {
	case 2:
		output.preSym = puncts[0]
		output.postSym = puncts[1]
	case 1:
		output.preSym = puncts[0]
	default:
		return
	}

	return
}
