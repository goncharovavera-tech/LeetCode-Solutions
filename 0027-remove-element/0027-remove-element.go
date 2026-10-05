func removeElement(nums []int, val int) int {
	count := 1
    for i:=len(nums)-1; i >= 0; i--{
		if nums[i] == val{
			nums[i] = nums[len(nums) - count]
			count ++
		}
	}
	return len(nums) - (count - 1)
}