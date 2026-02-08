package main

// 3834. Merge Adjacent Equal Elements
// You are given an integer array nums.

// You must repeatedly apply the following merge operation until no more changes can be made:
//     1. If any two adjacent elements are equal, 
//        choose the leftmost such adjacent pair in the current array and replace them with a single element equal to their sum.

// After each merge operation, the array size decreases by 1. 
// Repeat the process on the updated array until no more changes can be made.

// Return the final array after all possible merge operations.

// Example 1:
// Input: nums = [3,1,1,2]
// Output: [3,4]
// Explanation:
// The middle two elements are equal and merged into 1 + 1 = 2, resulting in [3, 2, 2].
// The last two elements are equal and merged into 2 + 2 = 4, resulting in [3, 4].
// No adjacent equal elements remain. Thus, the answer is [3, 4].

// Example 2:
// Input: nums = [2,2,4]
// Output: [8]
// Explanation:
// The first two elements are equal and merged into 2 + 2 = 4, resulting in [4, 4].
// The first two elements are equal and merged into 4 + 4 = 8, resulting in [8].

// Example 3:
// Input: nums = [3,7,5]
// Output: [3,7,5]
// Explanation:
// There are no adjacent equal elements in the array, so no operations are performed.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5​​​​​​​

import "fmt"

func mergeAdjacent(nums []int) []int64 {
    res := make([]int64,0)
    for _, v := range nums {
        res = append(res, int64(v))
        for len(res) >= 2 &&
            res[len(res) - 2] == res[len(res) - 1] {
            last := res[len(res) - 1]
            res = res[:len(res) - 1]
            res[len(res) - 1] += last
        }
    }
    return res  
}

func mergeAdjacent1(nums []int) []int64 {
    stack := []int64{}
    for _, v := range nums {
        stack = append(stack, int64(v))
        for len(stack) > 1 {
            num1 := stack[len(stack) - 1]
            num2 := stack[len(stack) - 2]
            if num1 != num2 { break }
            stack = stack[:len(stack) - 2]
            stack = append(stack, num1 * 2)
        }
        
    }
    return stack
}

func mergeAdjacent2(nums []int) []int64 {
    if nums == nil || len(nums) == 0 { return []int64{} }
    res := make([]int64, 0, len(nums))
    for i := 0; i < len(nums); i++ {
        v := int64(nums[i])
        for {
            n := len(res)
            if n == 0 {
                break
            }
            last := res[n-1]
            if last == v {
                v += last
                res = res[:n-1]     
            } else {
                break
            }
        }
        res = append(res, v)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,1,1,2]
    // Output: [3,4]
    // Explanation:
    // The middle two elements are equal and merged into 1 + 1 = 2, resulting in [3, 2, 2].
    // The last two elements are equal and merged into 2 + 2 = 4, resulting in [3, 4].
    // No adjacent equal elements remain. Thus, the answer is [3, 4].
    fmt.Println(mergeAdjacent([]int{3,1,1,2})) // [3,4]
    // Example 2:
    // Input: nums = [2,2,4]
    // Output: [8]
    // Explanation:
    // The first two elements are equal and merged into 2 + 2 = 4, resulting in [4, 4].
    // The first two elements are equal and merged into 4 + 4 = 8, resulting in [8].
    fmt.Println(mergeAdjacent([]int{2,2,4})) // [8]
    // Example 3:
    // Input: nums = [3,7,5]
    // Output: [3,7,5]
    // Explanation:
    // There are no adjacent equal elements in the array, so no operations are performed. 
    fmt.Println(mergeAdjacent([]int{3,7,5})) // [3,7,5]

    fmt.Println(mergeAdjacent([]int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(mergeAdjacent([]int{9,8,7,6,5,4,3,2,1})) // [9 8 7 6 5 4 3 2 1]

    fmt.Println(mergeAdjacent1([]int{3,1,1,2})) // [3,4]
    fmt.Println(mergeAdjacent1([]int{2,2,4})) // [8]
    fmt.Println(mergeAdjacent1([]int{3,7,5})) // [3,7,5]
    fmt.Println(mergeAdjacent1([]int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(mergeAdjacent1([]int{9,8,7,6,5,4,3,2,1})) // [9 8 7 6 5 4 3 2 1]

    fmt.Println(mergeAdjacent2([]int{3,1,1,2})) // [3,4]
    fmt.Println(mergeAdjacent2([]int{2,2,4})) // [8]
    fmt.Println(mergeAdjacent2([]int{3,7,5})) // [3,7,5]
    fmt.Println(mergeAdjacent2([]int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(mergeAdjacent2([]int{9,8,7,6,5,4,3,2,1})) // [9 8 7 6 5 4 3 2 1]
}