package lsproduct

import (
    "fmt"
    "strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
    if span > len(digits) || span < 0 {
        return -1, fmt.Errorf("")
    }
	if span == 0 {
        return 1, nil
    }
	var max_product int64
	for i:=0; i<len(digits)-span+1; i++ {
        var product int64
        product = 1
        for j:=i; j<span+i; j++ {
            if d,err := strconv.Atoi(string(digits[j])); err != nil {
                return -1, err
            } else {
            	product *= int64(d)
            }
        }
    	if product > max_product {
            max_product = product
        }
    }
	return max_product, nil
}