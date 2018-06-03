package twosum

// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/description/
func TwoSumSorted(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		complement := target - numbers[i]
		if idx := bSearch(numbers[i+1:], complement); idx != -1 {
			idx += i + 1
			return []int{i + 1, idx + 1}
		}
	}

	return nil
}

func bSearch(numbers []int, finder int) (idx int) {
	begin := 0
	end := len(numbers) - 1

	for begin <= end {
		mid := begin + (end-begin)/2

		if numbers[mid] > finder {
			end = mid - 1
		} else if numbers[mid] < finder {
			begin = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
