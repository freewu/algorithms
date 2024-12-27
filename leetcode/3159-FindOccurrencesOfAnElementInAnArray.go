package main

// 3159. Find Occurrences of an Element in an Array
// You are given an integer array nums, an integer array queries, and an integer x.

// For each queries[i], you need to find the index of the queries[i]th occurrence of x in the nums array. 
// If there are fewer than queries[i] occurrences of x, the answer should be -1 for that query.

// Return an integer array answer containing the answers to all queries.

// Example 1:
// Input: nums = [1,3,1,7], queries = [1,3,2,4], x = 1
// Output: [0,-1,2,-1]
// Explanation:
// For the 1st query, the first occurrence of 1 is at index 0.
// For the 2nd query, there are only two occurrences of 1 in nums, so the answer is -1.
// For the 3rd query, the second occurrence of 1 is at index 2.
// For the 4th query, there are only two occurrences of 1 in nums, so the answer is -1.

// Example 2:
// Input: nums = [1,2,3], queries = [10], x = 5
// Output: [-1]
// Explanation:
// For the 1st query, 5 doesn't exist in nums, so the answer is -1.

// Constraints:
//     1 <= nums.length, queries.length <= 10^5
//     1 <= queries[i] <= 10^5
//     1 <= nums[i], x <= 10^4

import "fmt"

func occurrencesOfElement(nums []int, queries []int, x int) []int {
    pos, res := []int{}, make([]int, len(queries))
    for i, v := range nums {
        if v == x {
            pos = append(pos, i)
        }
    }
    for i, index := range queries {
        if index - 1 < len(pos) {
            res[i] = pos[index - 1]
        } else {
            res[i] = -1
        }
    }
    return res
}

func occurrencesOfElement1(nums []int, queries []int, x int) []int {
    i := 0
    for j, v := range nums {
        if v == x {
            nums[i] = j
            i++
        }
    }
    for j, v := range queries {
        if v <= i {
            queries[j] = nums[v-1]
        } else {
            queries[j] = -1
        }
    }
    return queries
}

func main() {
    // Example 1:
    // Input: nums = [1,3,1,7], queries = [1,3,2,4], x = 1
    // Output: [0,-1,2,-1]
    // Explanation:
    // For the 1st query, the first occurrence of 1 is at index 0.
    // For the 2nd query, there are only two occurrences of 1 in nums, so the answer is -1.
    // For the 3rd query, the second occurrence of 1 is at index 2.
    // For the 4th query, there are only two occurrences of 1 in nums, so the answer is -1.
    fmt.Println(occurrencesOfElement([]int{1,3,1,7}, []int{1,3,2,4}, 1)) // [0,-1,2,-1]
    // Example 2:
    // Input: nums = [1,2,3], queries = [10], x = 5
    // Output: [-1]
    // Explanation:
    // For the 1st query, 5 doesn't exist in nums, so the answer is -1.
    fmt.Println(occurrencesOfElement([]int{1,2,3}, []int{10}, 5)) // [-1]

    fmt.Println(occurrencesOfElement1([]int{1,3,1,7}, []int{1,3,2,4}, 1)) // [0,-1,2,-1]
    fmt.Println(occurrencesOfElement1([]int{1,2,3}, []int{10}, 5)) // [-1]
}