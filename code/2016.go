package code

func maximumDifference(nums []int) int {
	a := 0
	diff := nums[1] - nums[0]
	for i := 1; i < len(nums); i++ {
		tmp := nums[i] - nums[a]
		if tmp < 0 {
			a = i
		}
		if tmp > diff {
			diff = tmp
		}
	}
	if diff > 0 {
		return diff
	}
	return -1
}
