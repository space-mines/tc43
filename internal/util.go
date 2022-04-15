package internal

func Contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
