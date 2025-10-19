package main

// 3717. Minimum Operations to Make the Array Beautiful
// You are given an integer array nums.

// An array is called beautiful if for every index i > 0, the value at nums[i] is divisible by nums[i - 1].

// In one operation, you may increment any element of nums by 1.

// Return the minimum number of operations required to make the array beautiful.

// Example 1:
// Input: nums = [3,7,9]
// Output: 2
// Explanation:
// Applying the operation twice on nums[1] makes the array beautiful: [3,9,9]

// Example 2:
// Input: nums = [1,1,1]
// Output: 0
// Explanation:
// The given array is already beautiful.

// Example 3:
// Input: nums = [4]
// Output: 0
// Explanation:
// The array has only one element, so it's already beautiful.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

// 美丽数组的定义：对于数组中的每个元素（除了第一个元素），该元素必须能被前一个元素整除
// 解题思路：
//     1. 从数组的第二个元素开始遍历，计算当前元素与前一个元素的余数
//     2. 如果余数不为 0，计算需要增加的值（使得当前元素能被前一个元素整除），并将该值累加到总操作次数中
//     3. 更新当前元素的值（增加后的值），并将其设为下一次迭代的前一个元素
//     4. 遍历完成后，返回总操作次数
func minOperations(nums []int) int64 {
    res, n, prev := 0, len(nums), nums[0]
    if n <= 1 { return 0 } // 如果数组长度小于等于 1，数组已经是美丽数组，直接返回 0
    for i := 1; i < n; i++ {
        current := nums[i]
        remainder := current % prev // 计算当前元素与前一个元素的余数
        if remainder != 0 { // 如果余数不为 0，计算需要增加的值（使得当前元素能被前一个元素整除），并将该值累加到总操作次数中。
            add := prev - remainder
            res += add
            current += add
        }
        prev = current // 更新当前元素的值（增加后的值），并将其设为下一次迭代的前一个元素
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [3,7,9]
    // Output: 2
    // Explanation:
    // Applying the operation twice on nums[1] makes the array beautiful: [3,9,9]
    fmt.Println(minOperations([]int{3,7,9})) // 2
    // Example 2:
    // Input: nums = [1,1,1]
    // Output: 0
    // Explanation:
    // The given array is already beautiful.
    fmt.Println(minOperations([]int{1,1,1})) // 0
    // Example 3:
    // Input: nums = [4]
    // Output: 0
    // Explanation:
    // The array has only one element, so it's already beautiful.
    fmt.Println(minOperations([]int{4})) // 0

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9})) // 14
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1})) // 36
}