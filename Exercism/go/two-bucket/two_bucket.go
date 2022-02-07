package twobucket

import "errors"

type Bucket struct {
    name string
    capacity int
    current int
}

func (b *Bucket) isEmpty() bool {
    return b.current == 0
}

func (b *Bucket) isFull() bool {
    return b.current == b.capacity
}

func (b *Bucket) fillUp() {
    b.current = b.capacity
}

func (b *Bucket) emptyOut() {
    b.current = 0
}

func (b *Bucket) pourInto(c *Bucket) {
    for b.current != 0 && c.current != c.capacity {
        b.current--
        c.current++
    }
}

func Solve(sizeBucketOne,
	sizeBucketTwo,
	goalAmount int,
	startBucket string) (goalBucket string, numSteps, otherBucketLevel int, e error) {
    if sizeBucketOne < 1 || sizeBucketTwo < 1 || goalAmount < 1 || (startBucket != "one" && startBucket != "two") {
        return "", 0, 0, errors.New("invalid parameters")
    }
	bucketOne := &Bucket{"one", sizeBucketOne, 0}
    bucketTwo := &Bucket{"two", sizeBucketTwo, 0}
    var first, second *Bucket
    if startBucket == "one" {
        first, second = bucketOne, bucketTwo
    } else {
    	first, second = bucketTwo, bucketOne
    }
    for first.current != goalAmount && second.current != goalAmount {
        switch {
            case first.isEmpty():
                first.fillUp()
            case second.isFull():
        		second.emptyOut()
            case first.capacity == goalAmount:
        		first.fillUp()
            case second.capacity == goalAmount:
        		second.fillUp()
            default:
        		first.pourInto(second)
        }
    	numSteps++
    }
	if first.current == goalAmount {
        goalBucket = first.name
        otherBucketLevel = second.current
    } else {
        goalBucket = second.name
        otherBucketLevel = first.current
    }
	return
}