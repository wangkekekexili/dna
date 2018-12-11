package proximate

// EditDistance returns the edit distance between two slice of bytes.
func EditDistance(left, right []byte) int {
	if len(left) == 0 {
		return len(right)
	}
	if len(right) == 0 {
		return 0
	}

	previous := make([]int, len(right)+1)
	for j := 0; j <= len(right); j++ {
		previous[j] = j
	}
	for i := 1; i <= len(left); i++ {
		current := make([]int, len(right)+1)
		current[0] = i
		for j := 1; j <= len(right); j++ {
			if left[i-1] == right[j-1] {
				current[j] = previous[j-1]
			} else {
				current[j] = previous[j-1] + 1
			}
			if previous[j]+1 < current[j] {
				current[j] = previous[j] + 1
			}
			if current[j-1]+1 < current[j] {
				current[j] = current[j-1] + 1
			}
		}
		previous = current
	}

	return previous[len(right)]
}
