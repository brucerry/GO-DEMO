package romannumerals

import "fmt"

func ToRomanNumeral(arabic int) (roman string, err error) {
    if arabic <= 0 || arabic > 3000 {
        err = fmt.Errorf("")
        return
    }
	for arabic != 0 {
		switch {
		case arabic >= 1000:
			for i:=0; i<arabic/1000; i++ {
				roman += "M"
			}
			arabic %= 1000
		case arabic >= 100:
			if arabic >= 500 {
				if arabic >= 900 {
					roman += "CM"
				} else {
					roman += "D"
					for i:=5; i<arabic/100; i++ {
						roman += "C"
					}
				}
			} else {
				if arabic >= 400 {
					roman += "CD"
				} else {
					for i:=0; i<arabic/100; i++ {
						roman += "C"
					}
				}
			}
			arabic %= 100
		case arabic >= 10:
			if arabic >= 50 {
				if arabic >= 90 {
					roman += "XC"
				} else {
					roman += "L"
					for i:=5; i<arabic/10; i++ {
						roman += "X"
					}
				}
			} else {
				if arabic >= 40 {
					roman += "XL"
				} else {
					for i:=0; i<arabic/10; i++ {
						roman += "X"
					}
				}
			}
			arabic %= 10
		default:
			if arabic >= 5 {
				if arabic >= 9 {
					roman += "IX"
				} else {
					roman += "V"
					for i:=5; i<arabic; i++ {
						roman += "I"
					}
				}
			} else {
				if arabic == 4 {
					roman += "IV"
				} else {
					for i:=0; i<arabic; i++ {
						roman += "I"
					}
				}
			}
			arabic = 0
		}
	}
	return
}