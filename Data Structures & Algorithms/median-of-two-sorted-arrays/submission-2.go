func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	m, n := len(nums1), len(nums2)
	med1, med2 := 0, 0
	for count:=0; count < (m + n) / 2 + 1; count++{
		med2 = med1
		if i < m && j < n{
			if nums1[i] < nums2[j]{
				med1 = nums1[i]
				i++
			}else{
				med1 = nums2[j]
				j++
			}
		}else if i < m{
			med1 = nums1[i]
			i++
		}else {
			med1 = nums2[j]
			j++
		}

	}
	if (m+n) % 2 == 1{
		return float64(med1)
	}
	return float64(med1+med2) / 2.0
}
