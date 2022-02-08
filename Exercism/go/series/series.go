package series

func All(n int, s string) []string {
	if n > len(s) {
        return nil
    }
	buf := []string{}
    for i:=0; i<len(s)-n+1; i++ {
        buf = append(buf, s[i:n+i])
    }
	return buf
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}
