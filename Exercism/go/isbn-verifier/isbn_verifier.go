package isbn

import (
    "strconv"
    "unicode"
)

func IsValidISBN(isbn string) bool {
    if isbn == "" {
        return false
    }
	isbn_2 := ""
	for _,ch := range isbn {
        if ch != '-' {
            isbn_2 += string(ch)
        }
    }
	if len(isbn_2) != 10 {
        return false
    }
	sum := 0
    multiplier := 10
	for i,ch := range isbn_2 {
        if i < len(isbn_2) - 1 {
            if !unicode.IsDigit(ch) {
            	return false
            } else {
            	d,_ := strconv.Atoi(string(ch))
            	sum += (d * multiplier)
                multiplier--
            }
        } else {
        	if ch == 'X' {
            	sum += (10 * multiplier)
                multiplier--
            } else if unicode.IsDigit(ch) {
            	d,_ := strconv.Atoi(string(ch))
            	sum += (d * multiplier)
                multiplier--
            } else {
            	return false
            }
        }
    }
	return sum % 11 == 0
}