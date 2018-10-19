package main

import "fmt"

func main() {
	for i:=0;i<5;i++ {
		nums := make([]int, 0, i+1)
		for j:=0;j<i+1;j++{
			nums = append(nums, j+1)
		}
		fmt.Println(nums)
		fmt.Println(permute(nums))
	}
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}
	var result [][]int
	for i:=0;i< len(nums);i++ {
		nums[0], nums[i] = nums[i], nums[0]
		next := permute(nums[1:])
		for _, n := range next {
			p := make([]int, len(n)+1)
			p[0] = nums[0]
			copy(p[1:], n)
			result = append(result, p)
		}
		nums[0], nums[i] = nums[i], nums[0]
	}
	return result
}