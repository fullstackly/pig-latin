package translator

import "testing"

func TestTranslateText(t *testing.T) {
	for _, tt := range puncTests {
		got := TranslateText(tt.initial)
		want := tt.translated
		initial := tt.initial

		if got != want {
			t.Errorf("expect %q translate to %q instead got %q", initial, want, got)
		}
	}
}

var puncTests = []struct {
	initial    string
	translated string
}{
	{initial: "Eat, Pray, Love", translated: "Eatyay, ayPray, oveLay"},
	{initial: "Hello everybody: Dolly, Molly & Polly!",
		translated: "elloHay everybodyyay: ollyDay, ollyMay & ollyPay!"},
	{initial: "Awesome programming languages: - Golang, - Ruby, - TypeScript, - Swift!",
		translated: "Awesomeyay ogrammingpray anguageslay: - olangGay, - ubyRay, - ypeScriptTay, - iftSway!"},
}
