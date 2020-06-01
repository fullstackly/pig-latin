package translator

import "testing"

func TestWordTranslate(t *testing.T) {
	t.Run("translating words that begin with consonant sounds", func(t *testing.T) {
		for _, tt := range consTests {
			got := wordTranslate(tt.initial)
			assertTranslation(t, tt.initial, tt.translated, got)
		}
	})

	t.Run("translating words begin with consonant clusters", func(t *testing.T) {
		for _, tt := range consClustersTests {
			got := wordTranslate(tt.initial)
			assertTranslation(t, tt.initial, tt.translated, got)
		}
	})

	t.Run("translating words that begin with vowel sounds", func(t *testing.T) {
		for _, tt := range vowelTests {
			got := wordTranslate(tt.initial)
			assertTranslation(t, tt.initial, tt.translated, got)
		}
	})

	t.Run("translating empty string", func(t *testing.T) {
		got := wordTranslate("")
		assertTranslation(t, "", "", got)
	})
}

func assertTranslation(t *testing.T, initial, want, got string) {
	t.Helper()

	if got != want {
		t.Errorf("expect %q translate to %q instead got %q", initial, want, got)
	}
}

var consTests = []struct {
	initial    string
	translated string
}{
	{initial: "pig", translated: "igpay"},
	{initial: "latin", translated: "atinlay"},
	{initial: "banana", translated: "ananabay"},
	{initial: "will", translated: "illway"},
}

var consClustersTests = []struct {
	initial    string
	translated string
}{
	{initial: "smile", translated: "ilesmay"},
	{initial: "string", translated: "ingstray"},
	{initial: "stupid", translated: "upidstay"},
	{initial: "glove", translated: "oveglay"},
}

var vowelTests = []struct {
	initial    string
	translated string
}{
	{initial: "eat", translated: "eatyay"},
	{initial: "omelet", translated: "omeletyay"},
	{initial: "are", translated: "areyay"},
	{initial: "egg", translated: "eggyay"},
}
