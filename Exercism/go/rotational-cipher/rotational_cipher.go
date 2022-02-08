package rotationalcipher

func RotationalCipher(text string, key int) (out string) {
    for _,ch := range text {
        if ch >= 'a' && ch <= 'z' {
            out += string(rune(int('a') + (int(ch) + key - int('a')) % 26))
        } else if ch >= 'A' && ch <= 'Z' {
            out += string(rune(int('A') + (int(ch) + key - int('A')) % 26))
        } else {
        	out += string(ch)
        }
    }
	return
}