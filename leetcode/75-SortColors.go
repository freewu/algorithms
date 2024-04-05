package main

// 75. Sort Colors
// Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
// You must solve this problem without using the library's sort function.

// Example 1:
// Input: nums = [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]

// Example 2:
// Input: nums = [2,0,1]
// Output: [0,1,2]

// Constraints:
//     n == nums.length
//     1 <= n <= 300
//     nums[i] is either 0, 1, or 2.

// Follow up: Could you come up with a one-pass algorithm using only constant extra space?

import "fmt"

func sortColors(nums []int) {
    zero, one := 0, 0
    for i, n := range nums {
        nums[i] = 2 // 每次都设置成 2,后面会根据 zero & one 再重写
        if n <= 1 { // <= 1  这步才会执行
            nums[one] = 1
            one++
        }
        if n == 0 { // 只有 n = 0 时候才处理
            nums[zero] = 0
            zero++
        }
        //fmt.Printf("round %v,nums = %v\n",i,nums)
    }
}

func main() {
	nums1 := []int{2,0,2,1,1,0}
	fmt.Printf("before nums1 = %v\n",nums1) // [2 0 2 1 1 0]
	sortColors(nums1)
	fmt.Printf("after nums1 = %v\n",nums1) // [0 0 1 1 2 2]

	nums2 := []int{2,0,1}
	fmt.Printf("before nums2 = %v\n",nums2) // [2 0 1]
	sortColors(nums2)
	fmt.Printf("after nums2 = %v\n",nums2) // [0 1 2]
}
