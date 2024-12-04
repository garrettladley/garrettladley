package predicate

import _ "embed"

func IntoPrompt(p Predicate) string {
	switch p {
	case Even:
		return EvenPrompt
	case Odd:
		return OddPrompt
	default:
		return ""
	}
}

//go:embed prompts/even.txt
var EvenPrompt string

//go:embed prompts/odd.txt
var OddPrompt string
