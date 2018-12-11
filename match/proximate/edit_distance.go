package proximate

// EditDistance returns the edit distance between two slice of bytes.
func EditDistance(left, right []byte) int {
	if len(left) == 0 {
		return len(right)
	}
	if len(right) == 0 {
		return 0
	}

	var distance [][]int
	for i := 0; i <= len(left); i++ {
		distance = append(distance, make([]int, len(right)+1))
		distance[i][0] = i
	}
	for j := 0; j <= len(right); j++ {
		distance[0][j] = j
	}
	for i := 1; i <= len(left); i++ {
		for j := 1; j <= len(right); j++ {
			if left[i-1] == right[j-1] {
				distance[i][j] = distance[i-1][j-1]
			} else {
				distance[i][j] = distance[i-1][j-1] + 1
			}
			if distance[i][j-1]+1 < distance[i][j] {
				distance[i][j] = distance[i][j-1] + 1
			}
			if distance[i-1][j]+1 < distance[i][j] {
				distance[i][j] = distance[i-1][j] + 1
			}
		}
	}

	return distance[len(left)][len(right)]
}
