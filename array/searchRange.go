package main
func searchRange(nums []int, target int) []int {
    leftBorder := getLeft(nums, target)
    rightBorder := getRight(nums, target)
    // 情况一
    if leftBorder == -2 || rightBorder == -2 {
        return []int{-1, -1}
    }
    // 情况三
    if rightBorder - leftBorder > 1 {
        return []int{leftBorder + 1, rightBorder - 1}
    }
    // 情况二
    return []int{-1, -1}
}

func getLeft(nums []int, target int) int {
    left, right := 0, len(nums)-1
    border := -2 // 记录border没有被赋值的情况；这里不能赋值-1，target = num[0]时，会无法区分情况一和情况二
    for left <= right { // []闭区间
	mid := left + ((right - left) >> 1)
	if nums[mid] >= target { // 找到第一个等于target的位置
	    right = mid - 1
            border = right
	} else {
	    left =  mid + 1
	}
    }
    return border
}

func getRight(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    border := -2
    for left <= right {
	mid := left + ((right - left) >> 1)
	if nums[mid] > target { 
		right = mid - 1
	} else { // 找到第一个大于target的位置
	    left = mid + 1
        border = left
	}
    }
    return border

}