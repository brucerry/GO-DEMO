package armstrong

import "math"

func IsNumber(n int) bool {
	tmp := n
    count := 0
    for tmp != 0 {
        tmp /= 10
        count++
    }
	tmp = n
    sum := 0
    for tmp != 0 {
        sum += int(math.Pow(float64(tmp%10), float64(count)))
        tmp /= 10
    }
	return sum == n
}
