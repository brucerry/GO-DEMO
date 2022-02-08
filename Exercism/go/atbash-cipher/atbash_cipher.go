package atbash

import (
    "strings"
)

func Atbash(text string) (out string) {
    group := ""
    for _,ch := range text {
        ch = rune(strings.ToLower(string(ch))[0])
    	if ch >= 'a' && ch <= 'z' {
            s := string(rune(int('z') + int('a') - int(ch)))
            group += s
            out += s
        } else if ch >= '0' && ch <= '9' {
        	group += string(ch)
            out += string(ch)
        }
    	if len(group) == 5 {
            out += " "
            group = ""
        }
    }
	return strings.TrimSpace(out)
}