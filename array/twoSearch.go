package main

import "fmt"

//二分查找
func Search(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			fmt.Println("found at index", mid)
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func searchInsert(nums []int, target int) int {
    left := 0
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        }else if nums[mid] > target{
            right = mid -1
        }else {
            left  = mid + 1
        }
    }
    return left

}

func main() {
	nums := []int{1,3,5,6}
	result := searchInsert(nums, 5)
	fmt.Println(result)
}
