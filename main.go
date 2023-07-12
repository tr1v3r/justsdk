// Package sdk provides ...
package sdk

// GetSum ...
func GetSum(a int, b int) int {
	return a + b
}

// GetAvg ...
func GetAvg(nums ...int) float64 {
	if len(nums) == 0 {
		return 0
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}

	return float64(sum) / float64(len(nums))
}
