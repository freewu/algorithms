package main

// 1966. Binary Searchable Numbers in an Unsorted Array
// Consider a function that implements an algorithm similar to Binary Search. 
// The function has two input parameters: sequence is a sequence of integers, and target is an integer value. 
// The purpose of the function is to find if the target exists in the sequence.

// The pseudocode of the function is as follows:
//     func(sequence, target)
//         while sequence is not empty
//             randomly choose an element from sequence as the pivot
//             if pivot = target, return true
//             else if pivot < target, remove pivot and all elements to its left from the sequence
//             else, remove pivot and all elements to its right from the sequence
//         end while
//         return false

// When the sequence is sorted, the function works correctly for all values. 
// When the sequence is not sorted, the function does not work for all values, but may still work for some values.

// Given an integer array nums, representing the sequence, that contains unique numbers and may or may not be sorted, 
// return the number of values that are guaranteed to be found using the function, for every possible pivot selection.

// Example 1:
// Input: nums = [7]
// Output: 1
// Explanation: 
// Searching for value 7 is guaranteed to be found.
// Since the sequence has only one element, 7 will be chosen as the pivot. Because the pivot equals the target, the function will return true.

// Example 2:
// Input: nums = [-1,5,2]
// Output: 1
// Explanation: 
// Searching for value -1 is guaranteed to be found.
// If -1 was chosen as the pivot, the function would return true.
// If 5 was chosen as the pivot, 5 and 2 would be removed. In the next loop, the sequence would have only -1 and the function would return true.
// If 2 was chosen as the pivot, 2 would be removed. In the next loop, the sequence would have -1 and 5. No matter which number was chosen as the next pivot, the function would find -1 and return true.
// Searching for value 5 is NOT guaranteed to be found.
// If 2 was chosen as the pivot, -1, 5 and 2 would be removed. The sequence would be empty and the function would return false.
// Searching for value 2 is NOT guaranteed to be found.
// If 5 was chosen as the pivot, 5 and 2 would be removed. In the next loop, the sequence would have only -1 and the function would return false.
// Because only -1 is guaranteed to be found, you should return 1.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^5 <= nums[i] <= 10^5
//     All the values of nums are unique.

// Follow-up: If nums has duplicates, would you modify your algorithm? If so, how?

import "fmt"

func binarySearchableNumbers(nums []int) int {
    start, n, mx := 0, len(nums), nums[0]
    incArr := make([]int, n)
    incArr[0] = nums[0]
    for i := 1; i < n; i++ {
        if nums[i] > mx { // 比左边所有数都大，放入数组
            mx = nums[i]
            start++
            incArr[start] = nums[i]
        } else { // 比左边数小，把大于这个数的都删掉
            for start >= 0 && incArr[start] > nums[i] {
                start--
            }
        }
    }
    return start + 1
}

func binarySearchableNumbers1(nums []int) int {
    res, n, inf := 0, len(nums), 1 << 31
    dp := make([]bool, n)
    mx, mn := -inf, inf
    for i, v := range nums {
        if v > mx {
            dp[i] = true
            mx = v
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
        if nums[i] < mn && dp[i] {
            res++
        }
        mn = min(mn, nums[i])
    }
    return res
}

// 单调栈
// 利用单调栈, 寻找i位置右侧比[i]小的,如果找不到(会保留在栈)中 => stk剩余的元素即满足 [i]<min([i:end])
// 同理反向扫描,寻找i左侧比[i]大的,如果找不到,也会保留在栈中,两者取交集
// 优化下, 一侧靠max取看是否满足条件
func binarySearchableNumbers2(nums []int) int {
    inf := -1 << 31
    stack := []int{ inf } // 保留值即可, 添加哨兵,避免len判断
    mx := inf
    for _, v := range nums {
        for stack[len(stack)-1] > v {
            stack = stack[:len(stack) - 1]
        }
        if v > mx { // v > preMax,这样满足左侧没有比v大的约束
            stack = append(stack, v)
            mx = v
        }
    }
    return len(stack) - 1 // 哨兵不算, 留在栈中的,代表寻找不到右侧有比cur小的元素
}

func main() {
    // Example 1:
    // Input: nums = [7]
    // Output: 1
    // Explanation: 
    // Searching for value 7 is guaranteed to be found.
    // Since the sequence has only one element, 7 will be chosen as the pivot. Because the pivot equals the target, the function will return true.
    fmt.Println(binarySearchableNumbers([]int{7})) // 1
    // Example 2:
    // Input: nums = [-1,5,2]
    // Output: 1
    // Explanation: 
    // Searching for value -1 is guaranteed to be found.
    // If -1 was chosen as the pivot, the function would return true.
    // If 5 was chosen as the pivot, 5 and 2 would be removed. In the next loop, the sequence would have only -1 and the function would return true.
    // If 2 was chosen as the pivot, 2 would be removed. In the next loop, the sequence would have -1 and 5. No matter which number was chosen as the next pivot, the function would find -1 and return true.
    // Searching for value 5 is NOT guaranteed to be found.
    // If 2 was chosen as the pivot, -1, 5 and 2 would be removed. The sequence would be empty and the function would return false.
    // Searching for value 2 is NOT guaranteed to be found.
    // If 5 was chosen as the pivot, 5 and 2 would be removed. In the next loop, the sequence would have only -1 and the function would return false.
    // Because only -1 is guaranteed to be found, you should return 1.
    fmt.Println(binarySearchableNumbers([]int{-1,5,2})) // 1

    fmt.Println(binarySearchableNumbers1([]int{7})) // 1
    fmt.Println(binarySearchableNumbers1([]int{-1,5,2})) // 1

    fmt.Println(binarySearchableNumbers2([]int{7})) // 1
    fmt.Println(binarySearchableNumbers2([]int{-1,5,2})) // 1
}