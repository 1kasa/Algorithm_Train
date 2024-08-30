package leetcode

func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if val != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func removeElementOther(nums []int,val int) int {
	leftIndex,rightIndex := 0,len(nums)-1
	for leftIndex<=rightIndex {
		for leftIndex<=rightIndex && nums[leftIndex]!=val {
			leftIndex++
		}
		for leftIndex<=rightIndex && nums[rightIndex]==val {
			rightIndex--
		}
		if leftIndex<rightIndex {
			nums[leftIndex]=nums[rightIndex]
			leftIndex++
			rightIndex--
		}
	
	}
	return leftIndex
}