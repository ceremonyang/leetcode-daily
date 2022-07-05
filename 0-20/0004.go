package main

// 4. 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var num []int
	i, j := 0, 0
	for {
		if len(nums1) == 0 {
			num = nums2
			break
		}
		if len(nums2) == 0 {
			num = nums1
			break
		}
		if nums1[i] > nums2[j] {
			num = append(num, nums2[j])
			j++
		} else {
			num = append(num, nums1[i])
			i++
		}
		if i == len(nums1) {
			num = append(num, nums2[j:]...)
			break
		}
		if j == len(nums2) {
			num = append(num, nums1[i:]...)
			break
		}
	}
	mid := int(len(num) / 2)
	if (len(num) % 2) == 0 {
		return float64(num[mid]+num[mid-1]) / 2
	}
	return float64(num[mid])
}
