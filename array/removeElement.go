package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))

}

// 移除元素
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
