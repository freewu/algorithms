package main

// 659. Split Array into Consecutive Subsequences
// You are given an integer array nums that is sorted in non-decreasing order.
// Determine if it is possible to split nums into one or more subsequences such that both of the following conditions are true:
//     Each subsequence is a consecutive increasing sequence (i.e. each integer is exactly one more than the previous integer).
//     All subsequences have a length of 3 or more.
    
// Return true if you can split nums according to the above conditions, or false otherwise.
// A subsequence of an array is a new array that is formed from the original array by deleting some (can be none) of the elements without disturbing the relative positions of the remaining elements. (i.e., [1,3,5] is a subsequence of [1,2,3,4,5] while [1,3,2] is not).

// Example 1:
// Input: nums = [1,2,3,3,4,5]
// Output: true
// Explanation: nums can be split into the following subsequences:
// [1,2,3,3,4,5] --> 1, 2, 3
// [1,2,3,3,4,5] --> 3, 4, 5

// Example 2:
// Input: nums = [1,2,3,3,4,4,5,5]
// Output: true
// Explanation: nums can be split into the following subsequences:
// [1,2,3,3,4,4,5,5] --> 1, 2, 3, 4, 5
// [1,2,3,3,4,4,5,5] --> 3, 4, 5

// Example 3:
// Input: nums = [1,2,3,4,4,5]
// Output: false
// Explanation: It is impossible to split nums into consecutive increasing subsequences of length 3 or more.
 
// Constraints:
//     1 <= nums.length <= 10^4
//     -1000 <= nums[i] <= 1000
//     nums is sorted in non-decreasing order.

import "fmt"

func isPossible(nums []int) bool {
    cnt, need := make(map[int]int), make(map[int]int)
    for _, v := range nums {
        cnt[v]++
    }
    for _, v := range nums {
        if cnt[v] == 0 {
            continue
        }
        cnt[v]--
        if need[v-1] > 0 {
            need[v-1]--
            need[v]++
        } else if cnt[v+1] > 0 && cnt[v+2] > 0 {
            cnt[v+1]--
            cnt[v+2]--
            need[v+2]++
        } else {
            return false
        }
    }
    return true
}

func isPossible1(nums []int) bool {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    pre, state := -1 << 32 - 1, [3]int{}
    for len(nums) > 0 {
        num := nums[0]
        if pre + 1 != num {
            if state[1] != 0 || state[2] != 0 {
                return false
            }
            state[0] = 0
        }
        count := 0
        for len(nums) > 0 && nums[0] == num {
            count++
            nums = nums[1:]
        }
        count -= state[1]+state[2]
        if count < 0 {
            return false
        }
        state[0], state[1], state[2] = state[1] + min(count, state[0]), state[2], max(0, count-state[0])
        pre = num
    }
    return state[1] == 0 && state[2] == 0
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,3,4,5]
    // Output: true
    // Explanation: nums can be split into the following subsequences:
    // [1,2,3,3,4,5] --> 1, 2, 3
    // [1,2,3,3,4,5] --> 3, 4, 5
    fmt.Println(isPossible([]int{1,2,3,3,4,5})) // true
    // Example 2:
    // Input: nums = [1,2,3,3,4,4,5,5]
    // Output: true
    // Explanation: nums can be split into the following subsequences:
    // [1,2,3,3,4,4,5,5] --> 1, 2, 3, 4, 5
    // [1,2,3,3,4,4,5,5] --> 3, 4, 5
    fmt.Println(isPossible([]int{1,2,3,3,4,4,5,5})) // true
    // Example 3:
    // Input: nums = [1,2,3,4,4,5]
    // Output: false
    // Explanation: It is impossible to split nums into consecutive increasing subsequences of length 3 or more.
    fmt.Println(isPossible([]int{1,2,3,4,4,5})) // false

    fmt.Println(isPossible1([]int{1,2,3,3,4,5})) // true
    fmt.Println(isPossible1([]int{1,2,3,3,4,4,5,5})) // true
    fmt.Println(isPossible1([]int{1,2,3,4,4,5})) // false
}
