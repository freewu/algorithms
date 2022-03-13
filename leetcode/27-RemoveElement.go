package main

/*
27 Remove Element

Given an array and a value, remove all instances of that value in-place and return the new length.
Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
The order of elements can be changed. It doesn't matter what you leave beyond the new length.

Example 1:

	Given nums = [3,2,2,3], val = 3, [2,2]
	Your function should return length = 2, with the first two elements of nums being 2.

Example 2:

	Given nums = [0,1,2,2,3,0,4,2], val = 2, [0,1,3,0,4]
	Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4.
	Note that the order of those five elements can be arbitrary.
	It doesn't matter what values are set beyond the returned length.

解题思路
	O(n)
	使用一个变量来记录机数组下标，如果数组下标对应的值 ==  要移除的值  蹼过变量的累加
*/

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	var t = 0
	for i := 0; i < len(nums); i++ {
		nums[t] = nums[i] // 需要前 t 个元素是不包含  val 的 如  nums: [0,1,2,2,3,0,4,2] val: 2 [0,1,3,0,4,0,4,2]
		if nums[i] != val {
			t++
		}
		fmt.Printf("round: %v\n",i )
		fmt.Printf("nums: %v \n",nums)
		fmt.Printf("t: %v \n",t)
	}
    return t
}

func main() {
	fmt.Printf("removeElement([]int{3,2,2,3},3) = %v\n",removeElement([]int{3,2,2,3},3))
	fmt.Printf("removeElement([]int{3,2,2,3},3) = %v\n",removeElement([]int{3,2,2,3},3))
	fmt.Printf("removeElement([]int{0,1,2,2,3,0,4,2},2) = %v\n",removeElement([]int{0,1,2,2,3,0,4,2},2))
}