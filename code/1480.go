package code

// stupid
func runningSum(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	var res = make([]int, len(nums))
	i := 1
	res[0] = nums[0]
	for {
		res[i] = res[i-1] + nums[i]
		i += 1
		if i == len(nums) {
			return res
		}
	}
}

// update
func runningSum2(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums

}
