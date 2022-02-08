package allyourbase

import (
    "fmt"
    "math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) (output []int, e error) {
    if inputBase < 2 {
        return []int(nil), fmt.Errorf("input base must be >= 2")
    } else if outputBase < 2 {
        return []int(nil), fmt.Errorf("output base must be >= 2")
    }
	decimal := 0
	for i,j:=len(inputDigits)-1,0; i>=0; i,j=i-1,j+1 {
        if !(0 <= inputDigits[i] && inputDigits[i] < inputBase) {
        	return []int(nil), fmt.Errorf("all digits must satisfy 0 <= d < input base")
        }
    	decimal += (inputDigits[i] * int(math.Pow(float64(inputBase), float64(j))))
    }
	for decimal != 0 {
		d := decimal % outputBase
		decimal /= outputBase
		output = append(output, 0)
		copy(output[1:], output[:len(output)-1])
		output[0] = d
	}
	if len(output) == 0 {
        output = append(output, 0)
    }
	return output, nil
}