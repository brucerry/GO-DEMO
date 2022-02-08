package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(Frequency)
	phrase = strings.ToLower(phrase)
	word := ""
	for i,ch := range phrase {
		if unicode.IsSpace(ch) {
			freq[word]++
			word = ""
		} else if i == len(phrase) - 1 && !unicode.IsPunct(ch) {
			word += string(ch)
			freq[word]++
		} else if unicode.IsPunct(ch) {
			if ch == '\'' {
				if i == 0 || i == len(phrase) - 1 { // where ' at the beginning / end
					freq[word]++
					word = ""
				} else if i < len(phrase) - 1 && unicode.IsLetter(rune(phrase[i-1])) && unicode.IsLetter(rune(phrase[i+1])) {
					word += string(ch)
				}
			} else {
				freq[word]++
				word = ""
			}
		} else if ch != '$' && ch != '^' {
			word += string(ch)
		}
	}
	delete(freq, "")
	return freq
}