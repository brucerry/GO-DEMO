package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	var output strings.Builder
	for len(input) > 0 {
		ch := input[0]
		initLen := len(input)
		input = strings.TrimLeft(input, string(ch))
		if n := initLen - len(input); n > 1 {
			output.WriteString(strconv.Itoa(n))
		}
		output.WriteByte(ch)
	}
	return output.String()
}

func RunLengthDecode(input string) string {
	var output strings.Builder
	for len(input) > 0 {
		i := strings.IndexFunc(input, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
		n := 1
		if i != 0 {
			n, _ = strconv.Atoi(input[:i])
		}
		output.WriteString(strings.Repeat(string(input[i]), n))
		input = input[i+1:]
	}
	return output.String()
}