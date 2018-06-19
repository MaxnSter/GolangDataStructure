package disappeared

func findDisappearedNumbers(nums []int) []int {
	disappeare := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		idx := absInt(nums[i]) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			disappeare = append(disappeare, i+1)
		}
	}

	return disappeare
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
