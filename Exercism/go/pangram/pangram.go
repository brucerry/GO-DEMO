package pangram

import (
    "strings"
    "unicode"
)

func IsPangram(input string) bool {
    if input == "" {
        return false
    }
	input = strings.ToLower(input)
    m := map[rune]int{}
    for _,ch := range input {
        if unicode.IsLetter(ch) {
            m[ch]++
        }
    }
	for i:=0; i<26; i++ {
        if m[rune('a' + i)] == 0 {
            return false
        }
    }
	return true
}