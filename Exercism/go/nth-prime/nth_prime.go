package prime

import "math"

func isPrime(n int) bool {
    for i:=2; i<=int(math.Sqrt(float64(n))); i++ {
        if n % i == 0 {
            return false
        }
    }
	return true
}

func Nth(n int) (int, bool) {
	if n <= 0 {
        return 0, false
    }
	counter := 0
    for i:=2; ; i++ {
        if yes := isPrime(i); yes {
            counter++
            if counter == n {
                return i, true
            }
        }
    }
}
