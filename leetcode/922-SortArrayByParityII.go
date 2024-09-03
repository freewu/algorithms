package main

// 922. Sort Array By Parity II
// Given an array of integers nums, half of the integers in nums are odd, and the other half are even.
// Sort the array so that whenever nums[i] is odd, i is odd, and whenever nums[i] is even, i is even.
// Return any answer array that satisfies this condition.

// Example 1:
// Input: nums = [4,2,5,7]
// Output: [4,5,2,7]
// Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.

// Example 2:
// Input: nums = [2,3]
// Output: [2,3]

// Constraints:
//     2 <= nums.length <= 2 * 10^4
//     nums.length is even.
//     Half of the integers in nums are even.
//     0 <= nums[i] <= 1000

// Follow Up: Could you solve it in-place?

import "fmt"

func sortArrayByParityII(nums []int) []int {
    res, even, odd :=  []int{}, []int{}, []int{}
    for _, v := range nums {
        if v % 2 == 0 {
            even = append(even, v)
        } else {
            odd = append(odd, v)
        }
    }
    for i := 0; i < len(nums) / 2; i++ {
        res = append(res, even[i])
        res = append(res, odd[i])
    }
    return res
}

// 快慢指针
func sortArrayByParityII1(nums []int) []int {
    slow, fast, n := 0, 0, len(nums)
    for slow < n {
        if slow % 2 == 0 {
            for fast < n && nums[fast] % 2 != 0 {
                fast++
            }
            if slow != fast {
                nums[slow], nums[fast] = nums[fast], nums[slow]
            }
            
        } else {
            for fast < n && nums[fast] % 2 == 0 {
                fast++
            }
            if slow != fast {
                nums[slow], nums[fast] = nums[fast], nums[slow]
            }
        }
        slow++
        fast = slow
    }
    return nums
}

func main() {
    // Example 1:
    // Input: nums = [4,2,5,7]
    // Output: [4,5,2,7]
    // Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.
    fmt.Println(sortArrayByParityII([]int{4,2,5,7})) // [4,5,2,7]
    // Example 2:
    // Input: nums = [2,3]
    // Output: [2,3]
    fmt.Println(sortArrayByParityII([]int{2,3})) // [2,3]

    fmt.Println(sortArrayByParityII1([]int{4,2,5,7})) // [4,5,2,7]
    fmt.Println(sortArrayByParityII1([]int{2,3})) // [2,3]
}