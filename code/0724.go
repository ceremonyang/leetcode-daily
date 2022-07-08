package code

func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		left := sum(nums[:i])
		right := sum(nums[i+1:])
		if left == right {
			return i
		}
	}
	return -1

}

func sum(nums []int) int {
	s := 0
	for n := range nums {
		s += nums[n]
	}
	return s
}

// update
func pivotIndex2(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	left := 0
	right := sum
	for j := 0; j < len(nums); j++ {
		right -= nums[j]
		if left == right {
			return j
		}
		left += nums[j]

	}
	return -1
}

/*
* sum(nums) = total
* left = x
* right = total - x - nums[i]
* if left == right
* then x  == total - x - nums[i]
 */

func pivotIndex3(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	left := 0
	for j := 0; j < len(nums); j++ {
		if 2*left+nums[j] == sum {
			return j
		}
		left += nums[j]

	}
	return -1
}
