package translator

import (
	"bytes"
	"strings"
)

// word interpretation as : punct[0] + base_word + punct[1]
type token struct {
	word    string
	symbols [2]string
}

func (t *token) string() string {
	return t.symbols[0] + t.word + t.symbols[1]
}

// TranslateText translates english text according to pig-latin rules
func TranslateText(text string) string {
	tokens := tockenizeText(text)
	var pigWords []string

	for _, token := range tokens {
		if token.word != "" {
			token.word = wordTranslate(token.word)
		}
		pigWords = append(pigWords, token.string())
	}

	return strings.Join(pigWords, " ")
}

func tockenizeText(text string) (tokens []token) {
	for _, chunk := range strings.Fields(text) {
		tokens = append(tokens, tockenizeWord(chunk))
	}
	return
}

func tockenizeWord(chunk string) (token token) {
	var word bytes.Buffer

	for _, let := range chunk {
		if strings.Contains(alphabet, string(let)) {
			word.WriteRune(let)
		}
	}

	symbols := strings.Split(chunk, word.String())

	token.word = word.String()
	copy(token.symbols[:], symbols)
	return
}
