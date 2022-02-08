package etl

import "strings"

func Transform(input map[int][]string) (output map[string]int) {
    output = map[string]int{}
    for score,collection := range input {
        for _,ch := range collection {
            output[strings.ToLower(ch)] = score
        }
    }
	return
}