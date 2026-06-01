func twoSum(numbers []int, target int) []int {
    n := len(numbers)
    i := 0
    j := n-1
    for _, _ = range numbers{
        if numbers[i] + numbers[j] > target{
            j--
        }else if numbers[i] + numbers[j] < target{
            i++
        }else{
            return []int{i+1,j+1}
        }
    }
    return []int{}

}
