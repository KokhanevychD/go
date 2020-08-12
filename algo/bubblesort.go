package algo

// Bubblesort fun, sorting int arrays
func Bubblesort(list []int) []int {
	length := len(list)
	sorted := true
	for idx := range list {
		if idx != length-1 {
			if list[idx] > list[idx+1] {
				list[idx], list[idx+1] = list[idx+1], list[idx]
				sorted = false
			}
		}
	}
	if !sorted {
		list = Bubblesort(list)
	}
	return list
}
