func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ind1, ind2 := 0, 0
	mergedNums := make([]int, 0)
	for ind1 < len(nums1) && ind2 < len(nums2){
		if nums1[ind1] < nums2[ind2]{
			mergedNums = append(mergedNums, nums1[ind1])
			ind1++
		}else{
			mergedNums = append(mergedNums, nums2[ind2])
			ind2++
		}
	}
	for ind1 < len(nums1){
		mergedNums = append(mergedNums, nums1[ind1])
		ind1++
	}
	for ind2 < len(nums2){
		mergedNums = append(mergedNums, nums2[ind2])
		ind2++
	}
	totalLen := len(nums1) + len(nums2)
	mid := 0
	if totalLen % 2 != 0{
		mid = totalLen/2
		return float64(mergedNums[mid])
	}else{
		mid = totalLen/2
		return (float64(mergedNums[mid-1])+float64(mergedNums[mid]))/2
	}
}
