package algo

// Insertion sorting algorithm
func Insertion(list []int) []int {
	for idx := range list {
		if idx > 0 {
			if list[idx] < list[idx-1] {
				list[idx], list[idx-1] = list[idx-1], list[idx]
				// reverse iteration in slice of array, breaks if find place
				slice := list[:idx]
				len := len(slice)
				for idx := range slice {
					if slice[len-idx-1] < slice[len-idx-2] {
						slice[len-idx-1], slice[len-idx-2] = slice[len-idx-2], slice[len-idx-1]
					} else {
						break
					}
				}
			}

		}
	}
	return list
}
