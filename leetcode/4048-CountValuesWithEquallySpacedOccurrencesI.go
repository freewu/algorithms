package main

// 4048. Count Values With Equally Spaced Occurrences I
// You are given an integer array nums.

// An integer x is called special if:
//     1. x appears exactly three times in nums.
//     2. All three occurrences of x are equally spaced in nums. 
//        In other words, if all occurrences of x are at indices i1 < i2 < i3, then i2 - i1 = i3 - i2.

// Return the number of distinct special integers in nums.

// Example 1:
// Input: nums = [1,8,1,5,1,5,8,5]
// Output: 2
// Explanation:
// 1 is special because it occurs exactly three times at equally spaced indices 0, 2, and 4.
// 5 is special because it occurs exactly three times at equally spaced indices 3, 5, and 7.
// 8 is not special because it occurs only twice.
// Therefore, the answer is 2.

// Example 2:
// Input: nums = [8,8,8,8]
// Output: 0
// Explanation:
// 8 is not special because it does not occur exactly three times. Therefore, the answer is 0.

// Example 3:
// Input: nums = [8,6,6,8,8]
// Output: 0
// Explanation:
// 8 occurs at indices 0, 3, and 4, which are not equally spaced. 6 occurs only twice. Therefore, no integer is special.

// Constraints:
//     3 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"

func countSpecialIntegers(nums []int) int {
    res, mp := 0, make(map[int][]int)
    for i,v := range nums {
        mp[v] = append(mp[v], i)
    }
    for _,w := range mp {
        if len(w) == 3 {
            if w[1] - w[0] == w[2] - w[1] {
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,8,1,5,1,5,8,5]
    // Output: 2
    // Explanation:
    // 1 is special because it occurs exactly three times at equally spaced indices 0, 2, and 4.
    // 5 is special because it occurs exactly three times at equally spaced indices 3, 5, and 7.
    // 8 is not special because it occurs only twice.
    // Therefore, the answer is 2.
    fmt.Println(countSpecialIntegers([]int{1,8,1,5,1,5,8,5})) // 2
    // Example 2:
    // Input: nums = [8,8,8,8]
    // Output: 0
    // Explanation:
    // 8 is not special because it does not occur exactly three times. Therefore, the answer is 0.
    fmt.Println(countSpecialIntegers([]int{8,8,8,8})) // 0
    // Example 3:
    // Input: nums = [8,6,6,8,8]
    // Output: 0
    // Explanation:
    // 8 occurs at indices 0, 3, and 4, which are not equally spaced. 6 occurs only twice. Therefore, no integer is special.
    fmt.Println(countSpecialIntegers([]int{8,6,6,8,8})) // 0

    fmt.Println(countSpecialIntegers([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(countSpecialIntegers([]int{9,8,7,6,5,4,3,2,1})) // 0
}