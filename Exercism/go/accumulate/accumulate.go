package accumulate

func Accumulate(s []string, f func(string) string) []string {
    out := make([]string, len(s))
    for i,str := range s {
        out[i] = f(str)
    }
	return out
}