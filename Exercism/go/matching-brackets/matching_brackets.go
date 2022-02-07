package brackets

func Bracket(input string) bool {
	if len(input) == 0 {
        return true
    }
	pairs := map[rune]rune{}
    pairs['('] = ')'
    pairs['['] = ']'
    pairs['{'] = '}'

    marker := []rune{}

    for _,v := range input {
        if _,ok := pairs[v]; ok {
            marker = append(marker, v)
        } else if v == ')' || v == ']' || v == '}' {
            if len(marker) > 0 && v == pairs[marker[len(marker)-1]] {
                marker = marker[:len(marker)-1]
            } else {
                return false
            }
        }
    }
    return len(marker) == 0
}
