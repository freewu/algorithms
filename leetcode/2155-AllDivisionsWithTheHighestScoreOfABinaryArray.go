package main

// 2155. All Divisions With the Highest Score of a Binary Array
// You are given a 0-indexed binary array nums of length n. 
// nums can be divided at index i (where 0 <= i <= n) into two arrays (possibly empty) numsleft and numsright:
//     1. numsleft has all the elements of nums between index 0 and i - 1 (inclusive), 
//        while numsright has all the elements of nums between index i and n - 1 (inclusive).
//     2. If i == 0, numsleft is empty, while numsright has all the elements of nums.
//     3. If i == n, numsleft has all the elements of nums, while numsright is empty.

// The division score of an index i is the sum of the number of 0's in numsleft and the number of 1's in numsright.

// Return all distinct indices that have the highest possible division score. 
// You may return the answer in any order.

// Example 1:
// Input: nums = [0,0,1,0]
// Output: [2,4]
// Explanation: Division at index
// - 0: numsleft is []. numsright is [0,0,1,0]. The score is 0 + 1 = 1.
// - 1: numsleft is [0]. numsright is [0,1,0]. The score is 1 + 1 = 2.
// - 2: numsleft is [0,0]. numsright is [1,0]. The score is 2 + 1 = 3.
// - 3: numsleft is [0,0,1]. numsright is [0]. The score is 2 + 0 = 2.
// - 4: numsleft is [0,0,1,0]. numsright is []. The score is 3 + 0 = 3.
// Indices 2 and 4 both have the highest possible division score 3.
// Note the answer [4,2] would also be accepted.

// Example 2:
// Input: nums = [0,0,0]
// Output: [3]
// Explanation: Division at index
// - 0: numsleft is []. numsright is [0,0,0]. The score is 0 + 0 = 0.
// - 1: numsleft is [0]. numsright is [0,0]. The score is 1 + 0 = 1.
// - 2: numsleft is [0,0]. numsright is [0]. The score is 2 + 0 = 2.
// - 3: numsleft is [0,0,0]. numsright is []. The score is 3 + 0 = 3.
// Only index 3 has the highest possible division score 3.

// Example 3:
// Input: nums = [1,1]
// Output: [0]
// Explanation: Division at index
// - 0: numsleft is []. numsright is [1,1]. The score is 0 + 2 = 2.
// - 1: numsleft is [1]. numsright is [1]. The score is 0 + 1 = 1.
// - 2: numsleft is [1,1]. numsright is []. The score is 0 + 0 = 0.
// Only index 0 has the highest possible division score 2.

// Constraints:
//     n == nums.length
//     1 <= n <= 10^5
//     nums[i] is either 0 or 1.

import "fmt"

func maxScoreIndices(nums []int) []int {
    res, mx, zero, one := []int{}, -1 << 31, 0, 0
    for _, v := range nums {
        if v != 0 {
            one++
        } else {
            zero++
        }
    }
    ones, zeros := 0, zero
    for i, v := range nums {
        sum := (zero - zeros) + (one - ones)
        if v == 0 {
            zeros--
        } else {
            ones++
        }
        if sum > mx {
            mx, res = sum, []int{i}
        } else if sum == mx {
            mx, res = sum, append(res, i)
        }
    }
    sum := (zero - zeros) + (one - ones)
    if sum > mx {
        res = []int{ len(nums) }
    } else if sum == mx {
        res = append(res, len(nums))
    }
    return res
}

func maxScoreIndices1(nums []int) []int {
    res, score, mx := make([]int, 0, len(nums)), 0, 0
    res = append(res, 0)
    for i, v := range nums {
        if v == 1 {
            score--
            continue
        }
        score++
        if score == mx {
            res = append(res, i + 1)
        } else if score > mx {
            mx, res = score, res[:0]
            res = append(res, i + 1)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,0,1,0]
    // Output: [2,4]
    // Explanation: Division at index
    // - 0: numsleft is []. numsright is [0,0,1,0]. The score is 0 + 1 = 1.
    // - 1: numsleft is [0]. numsright is [0,1,0]. The score is 1 + 1 = 2.
    // - 2: numsleft is [0,0]. numsright is [1,0]. The score is 2 + 1 = 3.
    // - 3: numsleft is [0,0,1]. numsright is [0]. The score is 2 + 0 = 2.
    // - 4: numsleft is [0,0,1,0]. numsright is []. The score is 3 + 0 = 3.
    // Indices 2 and 4 both have the highest possible division score 3.
    // Note the answer [4,2] would also be accepted.
    fmt.Println(maxScoreIndices([]int{0,0,1,0})) // [2,4]
    // Example 2:
    // Input: nums = [0,0,0]
    // Output: [3]
    // Explanation: Division at index
    // - 0: numsleft is []. numsright is [0,0,0]. The score is 0 + 0 = 0.
    // - 1: numsleft is [0]. numsright is [0,0]. The score is 1 + 0 = 1.
    // - 2: numsleft is [0,0]. numsright is [0]. The score is 2 + 0 = 2.
    // - 3: numsleft is [0,0,0]. numsright is []. The score is 3 + 0 = 3.
    // Only index 3 has the highest possible division score 3.
    fmt.Println(maxScoreIndices([]int{0,0,0})) // [3]
    // Example 3:
    // Input: nums = [1,1]
    // Output: [0]
    // Explanation: Division at index
    // - 0: numsleft is []. numsright is [1,1]. The score is 0 + 2 = 2.
    // - 1: numsleft is [1]. numsright is [1]. The score is 0 + 1 = 1.
    // - 2: numsleft is [1,1]. numsright is []. The score is 0 + 0 = 0.
    // Only index 0 has the highest possible division score 2.
    fmt.Println(maxScoreIndices([]int{1,1})) // [0]

    fmt.Println(maxScoreIndices1([]int{0,0,1,0})) // [2,4]
    fmt.Println(maxScoreIndices1([]int{0,0,0})) // [3]
    fmt.Println(maxScoreIndices1([]int{1,1})) // [0]
}