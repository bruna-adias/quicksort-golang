package pkg

// Sort sorts a slice of integers using the quicksort algorithm
func Sort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := partition(arr)
	Sort(arr[:pivotIndex])
	Sort(arr[pivotIndex+1:])
}

func partition(arr []int) int {
	pivot := arr[len(arr)-1]
	i := 0

	for j := 0; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return i
}
