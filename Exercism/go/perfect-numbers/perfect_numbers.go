package perfect

import "errors"

type Classification int

var ErrOnlyPositive error = errors.New("")

const (
    ClassificationDeficient Classification = 1
	ClassificationPerfect Classification = 2
	ClassificationAbundant Classification = 3
)

func Classify(number int64) (Classification, error) {
	if number <= 0 {
        return -1, ErrOnlyPositive
    }
	var aliquot_sum int64 = 0
	visited := map[int64]bool{}
	for i:=int64(1); i<=number; i++ {
        if visited[i] {
            break
        }
        visited[i] = true
    	if number != i && i != number/i && number % i == 0 {
            visited[number/i] = true
            aliquot_sum += i + number/i
        }
    }
	aliquot_sum -= number
    switch {
        case aliquot_sum == number:
    		return ClassificationPerfect, nil
        case aliquot_sum > number:
    		return ClassificationAbundant, nil
        case aliquot_sum < number:
    		return ClassificationDeficient, nil
        default:
    		return -1, nil
    }
}