package main

// 2454. Next Greater Element IV
// You are given a 0-indexed array of non-negative integers nums. 
// For each integer in nums, you must find its respective second greater integer.

// The second greater integer of nums[i] is nums[j] such that:
//     1. j > i
//     2. nums[j] > nums[i]
//     3. There exists exactly one index k such that nums[k] > nums[i] and i < k < j.

// If there is no such nums[j], the second greater integer is considered to be -1.
//     For example, in the array [1, 2, 4, 3], the second greater integer of 1 is 4, 2 is 3, and that of 3 and 4 is -1.

// Return an integer array answer, where answer[i] is the second greater integer of nums[i].

// Example 1:
// Input: nums = [2,4,0,9,6]
// Output: [9,6,6,-1,-1]
// Explanation:
// 0th index: 4 is the first integer greater than 2, and 9 is the second integer greater than 2, to the right of 2.
// 1st index: 9 is the first, and 6 is the second integer greater than 4, to the right of 4.
// 2nd index: 9 is the first, and 6 is the second integer greater than 0, to the right of 0.
// 3rd index: There is no integer greater than 9 to its right, so the second greater integer is considered to be -1.
// 4th index: There is no integer greater than 6 to its right, so the second greater integer is considered to be -1.
// Thus, we return [9,6,6,-1,-1].

// Example 2:
// Input: nums = [3,3]
// Output: [-1,-1]
// Explanation:
// We return [-1,-1] since neither integer has any integer greater than it.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9

import "fmt"
import "sort"

func secondGreaterElement(nums []int) []int {
    s1, s2 := []int{}, []int{}
    for i, v := range nums {
        for len(s2) != 0 && nums[s2[len(s2) - 1]] < v {
            nums[s2[len(s2) - 1]] = v
            s2 = s2[:len(s2) - 1]
        }
        index := sort.Search(len(s1), func(i int) bool {
            return nums[s1[i]] < v
        })
        if index == len(s1) {
            s1 = append(s1, i)
        } else {
            s2 = append(s2, s1[index:]...)
            s1[index] = i
            s1 = s1[:index + 1]
        }
    }
    for _, s := range [...][]int{s1, s2} {
        for _, index := range s {
            nums[index] = -1
        }
    }
    return nums
}

func secondGreaterElement1(nums []int) []int {
    n, i1, i2 := len(nums), -1, -1
    res, s1, s2 := make([]int, n), make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        res[i] = -1
    }
    for i, v := range nums {
        for i2 >= 0 && v > nums[s2[i2]] {
            res[s2[i2]] = v
            i2--
        }
        j := i1
        for i1 >= 0 && v > nums[s1[i1]] {
            i1--
        }
        for k := i1 + 1; k <= j; k++ {
            i2++
            s2[i2] = s1[k]
        }
        i1++
        s1[i1] = i
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,4,0,9,6]
    // Output: [9,6,6,-1,-1]
    // Explanation:
    // 0th index: 4 is the first integer greater than 2, and 9 is the second integer greater than 2, to the right of 2.
    // 1st index: 9 is the first, and 6 is the second integer greater than 4, to the right of 4.
    // 2nd index: 9 is the first, and 6 is the second integer greater than 0, to the right of 0.
    // 3rd index: There is no integer greater than 9 to its right, so the second greater integer is considered to be -1.
    // 4th index: There is no integer greater than 6 to its right, so the second greater integer is considered to be -1.
    // Thus, we return [9,6,6,-1,-1].
    fmt.Println(secondGreaterElement([]int{2,4,0,9,6})) // [9,6,6,-1,-1]
    // Example 2:
    // Input: nums = [3,3]
    // Output: [-1,-1]
    // Explanation:
    // We return [-1,-1] since neither integer has any integer greater than it.
    fmt.Println(secondGreaterElement([]int{3,3})) // [-1,-1]

    fmt.Println(secondGreaterElement([]int{1,2,3,4,5,6,7,8,9})) // [3 4 5 6 7 8 9 -1 -1]
    fmt.Println(secondGreaterElement([]int{9,8,7,6,5,4,3,2,1})) // [-1 -1 -1 -1 -1 -1 -1 -1 -1]

    fmt.Println(secondGreaterElement1([]int{2,4,0,9,6})) // [9,6,6,-1,-1]
    fmt.Println(secondGreaterElement1([]int{3,3})) // [-1,-1]
    fmt.Println(secondGreaterElement1([]int{1,2,3,4,5,6,7,8,9})) // [3 4 5 6 7 8 9 -1 -1]
    fmt.Println(secondGreaterElement1([]int{9,8,7,6,5,4,3,2,1})) // [-1 -1 -1 -1 -1 -1 -1 -1 -1]
}