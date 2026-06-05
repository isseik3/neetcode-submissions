func searchMatrix(matrix [][]int, target int) bool {
	// binary search
	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, rows * cols - 1
	// treat the matrix as a flattened sorted array of length m times n, then binary search it.
	for left <= right {
		mid := left + (right - left)/2
		val := matrix[mid/cols][mid%cols]
		if val == target {
			return true
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
