package main

// 2562. Find the Array Concatenation Value
// You are given a 0-indexed integer array nums.

// The concatenation of two numbers is the number formed by concatenating their numerals.
//     For example, the concatenation of 15, 49 is 1549.

// The concatenation value of nums is initially equal to 0. 
// Perform this operation until nums becomes empty:
//     1. If there exists more than one number in nums, 
//        pick the first element and last element in nums respectively and add the value of their concatenation to the concatenation value of nums, 
//        then delete the first and last element from nums.
//     2. If one element exists, add its value to the concatenation value of nums, then delete it.

// Return the concatenation value of the nums.

// Example 1:
// Input: nums = [7,52,2,4]
// Output: 596
// Explanation: Before performing any operation, nums is [7,52,2,4] and concatenation value is 0.
//  - In the first operation:
// We pick the first element, 7, and the last element, 4.
// Their concatenation is 74, and we add it to the concatenation value, so it becomes equal to 74.
// Then we delete them from nums, so nums becomes equal to [52,2].
//  - In the second operation:
// We pick the first element, 52, and the last element, 2.
// Their concatenation is 522, and we add it to the concatenation value, so it becomes equal to 596.
// Then we delete them from the nums, so nums becomes empty.
// Since the concatenation value is 596 so the answer is 596.

// Example 2:
// Input: nums = [5,14,13,8,12]
// Output: 673
// Explanation: Before performing any operation, nums is [5,14,13,8,12] and concatenation value is 0.
//  - In the first operation:
// We pick the first element, 5, and the last element, 12.
// Their concatenation is 512, and we add it to the concatenation value, so it becomes equal to 512.
// Then we delete them from the nums, so nums becomes equal to [14,13,8].
//  - In the second operation:
// We pick the first element, 14, and the last element, 8.
// Their concatenation is 148, and we add it to the concatenation value, so it becomes equal to 660.
// Then we delete them from the nums, so nums becomes equal to [13].
//  - In the third operation:
// nums has only one element, so we pick 13 and add it to the concatenation value, so it becomes equal to 673.
// Then we delete it from nums, so nums become empty.
// Since the concatenation value is 673 so the answer is 673.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 10^4

import "fmt"
import "strconv"

func findTheArrayConcVal2(nums []int) int64 {
    res, n := int64(0), len(nums)
    concatValues := func(v1, v2 int) int64 {
        res, _ := strconv.ParseInt(strconv.Itoa(v1) + strconv.Itoa(v2), 10, 64)
        return res
    }
    for i := 0; i < n / 2; i++ {
        res += concatValues(nums[i], nums[n - i - 1])
    }
    if n % 2 != 0 {
        res += int64(nums[n / 2])
    }
    return res
}

func findTheArrayConcVal(nums []int) int64 {
    res, n := 0, len(nums)
    for i := 0; i < (n + 1) / 2; i++ {
        if i == n - i - 1 {
            res += nums[i]
            break
        }
        a, b := nums[i], nums[n - i - 1]
        c := b
        for c > 0 {
            a *= 10
            c /= 10
        }
        res += (a + b)
    }
    return int64(res)
}

func findTheArrayConcVal1(nums []int) int64 {
    res, left, right := 0, 0, len(nums) - 1 
    for left <= right {
        if left == right {
            res += nums[left]
            left++
        } else {
            mul, temp := 1, nums[right]
            for temp != 0 {
                mul *= 10
                temp /= 10
            }
            res += (nums[left] * mul + nums[right])
            left++
            right--
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [7,52,2,4]
    // Output: 596
    // Explanation: Before performing any operation, nums is [7,52,2,4] and concatenation value is 0.
    //  - In the first operation:
    // We pick the first element, 7, and the last element, 4.
    // Their concatenation is 74, and we add it to the concatenation value, so it becomes equal to 74.
    // Then we delete them from nums, so nums becomes equal to [52,2].
    //  - In the second operation:
    // We pick the first element, 52, and the last element, 2.
    // Their concatenation is 522, and we add it to the concatenation value, so it becomes equal to 596.
    // Then we delete them from the nums, so nums becomes empty.
    // Since the concatenation value is 596 so the answer is 596.
    fmt.Println(findTheArrayConcVal([]int{7,52,2,4})) // 594
    // Example 2:
    // Input: nums = [5,14,13,8,12]
    // Output: 673
    // Explanation: Before performing any operation, nums is [5,14,13,8,12] and concatenation value is 0.
    //  - In the first operation:
    // We pick the first element, 5, and the last element, 12.
    // Their concatenation is 512, and we add it to the concatenation value, so it becomes equal to 512.
    // Then we delete them from the nums, so nums becomes equal to [14,13,8].
    //  - In the second operation:
    // We pick the first element, 14, and the last element, 8.
    // Their concatenation is 148, and we add it to the concatenation value, so it becomes equal to 660.
    // Then we delete them from the nums, so nums becomes equal to [13].
    //  - In the third operation:
    // nums has only one element, so we pick 13 and add it to the concatenation value, so it becomes equal to 673.
    // Then we delete it from nums, so nums become empty.
    // Since the concatenation value is 673 so the answer is 673.
    fmt.Println(findTheArrayConcVal([]int{5,14,13,8,12})) // 673

    fmt.Println(findTheArrayConcVal([]int{1,2,3,4,5,6,7,8,9})) // 135
    fmt.Println(findTheArrayConcVal([]int{9,8,7,6,5,4,3,2,1})) // 315

    fmt.Println(findTheArrayConcVal1([]int{7,52,2,4})) // 594
    fmt.Println(findTheArrayConcVal1([]int{5,14,13,8,12})) // 673
    fmt.Println(findTheArrayConcVal1([]int{1,2,3,4,5,6,7,8,9})) // 135
    fmt.Println(findTheArrayConcVal1([]int{9,8,7,6,5,4,3,2,1})) // 315

    fmt.Println(findTheArrayConcVal2([]int{7,52,2,4})) // 594
    fmt.Println(findTheArrayConcVal2([]int{5,14,13,8,12})) // 673
    fmt.Println(findTheArrayConcVal2([]int{1,2,3,4,5,6,7,8,9})) // 135
    fmt.Println(findTheArrayConcVal2([]int{9,8,7,6,5,4,3,2,1})) // 315
}