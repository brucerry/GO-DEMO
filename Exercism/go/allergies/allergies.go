package allergies

import "math"

var arr = [8]string{
    "eggs",
    "peanuts",
    "shellfish",
    "strawberries",
    "tomatoes",
    "chocolate",
    "pollen",
    "cats",
}

func Allergies(allergies uint) (out []string) {
	allergies %= 256
	for _,s := range arr {
		if AllergicTo(allergies, s) {
			out = append(out, s)
		}
	}
	return
}

func AllergicTo(allergies uint, allergen string) bool {
    for i,s := range arr {
        if (uint(math.Pow(2, float64(i))) & allergies) > 0 && s == allergen {
            return true
        }
    }
	return false
}
