package sieve

func Sieve(limit int) (output []int) {
    mark := map[int]bool{} // true = not prime
    for i:=2; i<=limit; i++ {
        if mark[i] == true {
            continue
        } else {
    		output = append(output, i)
        }
    	for multiplier:=i; ; multiplier++ {
            tmp := i * multiplier
            mark[tmp] = true
            if tmp > limit {
                break
            }
        }
    }
	return
}