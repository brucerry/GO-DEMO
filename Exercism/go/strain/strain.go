package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(f func(int) bool) Ints {
    var result Ints
    for _,v := range ints {
        if f(v) {
            result = append(result, v)
        }
    }
	return result
}

func (ints Ints) Discard(f func(int) bool) Ints {
    var result Ints
    for _,v := range ints {
        if !f(v) {
            result = append(result, v)
        }
    }
	return result
}

func (lists Lists) Keep(f func([]int) bool) Lists {
    var result Lists
    for _,v := range lists {
        if f(v) {
            result = append(result, v)
        }
    }
	return result
}

func (str Strings) Keep(f func(string) bool) Strings {
    var result Strings
    for _,v := range str {
        if f(v) {
            result = append(result, v)
        }
    }
	return result
}