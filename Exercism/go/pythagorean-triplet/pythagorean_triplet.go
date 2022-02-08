package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
    list := []Triplet{}
    for i:=min; i<=max; i++ {
        for j:=i+1; j<=max; j++ {
            for k:=j+1; k<=max; k++ {
                if i*i + j*j == k*k {
                    list = append(list, Triplet{i, j, k})
                }
            }
        }
    }
	return list
}

func Sum(p int) []Triplet {
    list := []Triplet{}
    for _,triplet := range Range(1, p) {
		if triplet[0] + triplet[1] + triplet[2] == p {
			list = append(list, triplet)
		}
    }
	return list
}