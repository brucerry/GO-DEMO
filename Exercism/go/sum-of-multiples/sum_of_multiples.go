package summultiples

func SumMultiples(limit int, divisors ...int) (sum int) {
    counter := map[int]int{}
    for _,d := range divisors {
        if d == 0 {
            break
        }
        for i:=1; d*i<limit; i++ {
            counter[d*i]++
        }
    }
	for k := range counter {
		sum += k
    }
	return
} 