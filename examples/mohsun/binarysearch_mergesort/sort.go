package main

func selectionsort(slice []int) []int {
	var temp = 0
	for index := 0; index < len(slice); index++ {
		for j := index + 1; j < len(slice); j++ {
			// fmt.Println(j)
			if slice[index] > slice[j] {
				// fmt.Println("before reindex", slice)
				temp = slice[index]
				slice[index] = slice[j]
				slice[j] = temp
				// fmt.Println("after reindex", slice)
			}
		}
	}
	return slice
}

func mergeSort[T string | int](items []T) []T {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge[T string | int](a []T, b []T) []T {
	final := []T{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			// fmt.Println("a[i]:", a[i], "b[i]: ", b[i])
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
