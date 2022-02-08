package cryptosquare

import (
	"unicode"
	"math"
)

func Encode(text string) (out string) {
    tmp := []rune{}
    for _,ch := range text {
        if unicode.IsLetter(ch) || unicode.IsNumber(ch) {
            tmp = append(tmp, unicode.ToLower(ch))
        }
    }
	len := len(tmp)
	sqrt := math.Sqrt(float64(len))
	var r, c int
	if sqrt - float64(int(sqrt)) > 0 {
		r, c = int(sqrt), int(sqrt) + 1
	} else {
		r, c = int(sqrt), int(sqrt)
	}
	if r * c < len {
		r++
	}
	tmp_2 := make([][]rune, r)
	for i := range tmp_2 {
		tmp_2[i] = make([]rune, c)
	}
	for i:=0; i<r; i++ {
		for j:=0; j<c; j++ {
			if j + i * c >= len {
				tmp_2[i][j] = ' '
			} else {
				tmp_2[i][j] = tmp[j + i * c]
			}
		}
	}
	for i:=0; i<c; i++ {
		for j:=0; j<r; j++ {
			out += string(tmp_2[j][i])
		}
		if i < c - 1 {
			out += " "
		}
	}
	return out
}