package binarysearch

import "fmt"

func search(item int, sortedList []int) string {
	var mid int
	low := 0
	high := len(sortedList) - 1

	for low <= high {
		mid = (high + low)/2

		guess := sortedList[mid] 

		if guess == item {
			return fmt.Sprintf("%d found at position %d",item,mid)
		}

		if guess > item {
			high = mid -1
		}else {
			low = mid + 1
		}

	}


	return fmt.Sprintf("Item %d does not exist", item)
}
