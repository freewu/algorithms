package main

// 283. Move Zeroes
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.

// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]

// Example 2:
// Input: nums = [0]
// Output: [0]

// Constraints:
//     1 <= nums.length <= 10^4
//     -2^31 <= nums[i] <= 2^31 - 1

import "fmt"

func moveZeroes(nums []int) {
    if len(nums) == 0 {
        return
    }
    j := 0 // 先默认第一个是0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 { // 如果是 0 就跳过
            // fmt.Printf("i = %v,nums[i] = %v,j = %v,nums = %v\n",i,nums[i],j,nums)
            if i != j {
                nums[i], nums[j] = nums[j], nums[i]
            }
            j++
        }
    }
}

func moveZeroes1(nums []int)  {
    for pos, cur := 0, 0; cur < len(nums); cur++ {
        if nums[cur] != 0 {
            nums[pos], nums[cur] =  nums[cur] ,nums[pos]
            pos++
        }
    }
}

func main() {
	nums := []int{ 0,1,0,3,12 }
	fmt.Printf("before: nums = %v\n",nums)
	moveZeroes(nums)
	fmt.Printf("after: nums = %v\n",nums)

	nums1 := []int{1,0,2,3,4,5,6,0 }
	fmt.Printf("before: nums1 = %v\n",nums1)
	moveZeroes(nums1)
	fmt.Printf("after: nums1 = %v\n",nums1)

    nums11 := []int{ 0,1,0,3,12 }
	fmt.Printf("before: nums11 = %v\n",nums11)
	moveZeroes1(nums11)
	fmt.Printf("after: nums11 = %v\n",nums11)

	nums12 := []int{1,0,2,3,4,5,6,0 }
	fmt.Printf("before: nums12 = %v\n",nums12)
	moveZeroes1(nums12)
	fmt.Printf("after: nums12 = %v\n",nums12)
}