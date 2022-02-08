package anagram

import (
    "strings"
    "sort"
)

type Str []rune

func (s Str) Len() int {
    return len(s)
}

func (s Str) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s Str) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func Detect(subject string, candidates []string) []string {
    output := []string{}
    subject = strings.ToLower(subject)
	subject_rune := []rune(subject)
    sort.Sort(Str(subject_rune))
    for _, candidate := range candidates {
        tester := strings.ToLower(candidate)
        if subject == tester {
            return []string{}
        }
		tester_rune := []rune(tester)
    	sort.Sort(Str(tester_rune))
    	if string(subject_rune) == string(tester_rune) {
            output = append(output, candidate)
        }
    }
	return output
}