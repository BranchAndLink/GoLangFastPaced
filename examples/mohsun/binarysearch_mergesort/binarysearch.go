package main

import "fmt"

func binarysearch(keyword string, list []string) (result int) {
	result = -1
	if len(list) < 1 {

		return
	}
	mid := len(list) / 2
	if list[mid] == keyword {
		result = mid
	} else if keyword > list[mid] {

		result = binarysearch(keyword, list[mid:]) + mid

	} else if keyword < list[mid] {
		result = binarysearch(keyword, list[:mid])

	} else {
		fmt.Println("Keyword is not in the list")
	}
	return result
}

func binarysearch2(keyword string, list []string, low int, high int) (result int) {
	result = -1
	if len(list) < 1 {

		return
	}
	mid := low + (high-low)/2
	if list[mid] == keyword {
		result = mid
	} else if keyword > list[mid] {
		result = binarysearch2(keyword, list, low, mid-1)

	} else if keyword < list[mid] {
		result = binarysearch2(keyword, list, mid+1, low)

	} else {
		fmt.Println("Keyword is not in the list")
	}
	return result
}

//iterativ de yazarsan binarysearch3

func iterBinary(words []string, keyword string, firstIndex bool) int {

	low := 0
	high := len(words) - 1

	for low <= high {
		mid := low + (high-low)/2
		if words[mid] == keyword {
			//
			if firstIndex {
				//lets not return if its not exact the first index
				if mid == 0 || keyword > words[mid-1] {
					return mid
				}

			} else {
				return mid
			}

		}

		if keyword > words[mid] {
			low = mid + 1

		} else {
			//keyword < words[mid]
			high = mid - 1

		}

	}
	return -1
}

func findWordsPrefix(words []string, keyword string, maxCount int) (int, int) {

	first_index := iterBinary(words, keyword, true)
	end_index := first_index + 1
	lenx := len(keyword)
	for i := end_index; i < len(words) && i < first_index+maxCount; i++ {

		if len(words[i]) > lenx && keyword == words[i][:lenx] {
			end_index = i + 1
		} else {
			break
		}
	}
	if first_index < 0 && first_index < end_index {
		first_index = 0
	}
	return first_index, end_index
}
