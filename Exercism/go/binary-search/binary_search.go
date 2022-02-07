package binarysearch

func SearchInts(list []int, key int) int {
	if len(list) == 0 {
        return -1
    }
	head := 0
    tail := len(list) - 1
    for head <= tail {
        mid := (head + tail) / 2
        if key == list[mid] {
            return mid
        } else if key < list[mid] {
        	tail = mid - 1
        } else {
        	head = mid + 1
        }
    }
	return -1
}
