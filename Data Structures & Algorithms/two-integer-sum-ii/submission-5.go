func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < r{
		diff := target - numbers[l]
		if diff < numbers[r]{
			r--
		}else if diff > numbers[r]{
			l++
		}else{
			return []int{l+1, r+1}
		}
	}
	return []int{}
}
