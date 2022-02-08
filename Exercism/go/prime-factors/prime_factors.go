package prime

func Factors(n int64) []int64 {
	if n < 2 {
        return []int64{}
    }
	buf := []int64{}
    var i int64
    for i=2; n!=1; i++ {
        if i == n {
            buf = append(buf, i)
            break
        }
        if n % i == 0 {
            buf = append(buf, i)
            n /= i
            i = 1
            if n == 1 {
                break
            }
        }
    }
	return buf
}
