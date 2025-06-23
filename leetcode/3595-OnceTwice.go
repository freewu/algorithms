package main

// 3595. Once Twice
// You are given an integer array nums. In this array:
//     1. Exactly one element appears once.
//     2. Exactly one element appears twice.
//     3. All other elements appear exactly three times.

// Return an integer array of length 2, where the first element is the one that appears once, and the second is the one that appears twice.

// Your solution must run in O(n) time and O(1) space.

// Example 1:
// Input: nums = [2,2,3,2,5,5,5,7,7]
// Output: [3,7]
// Explanation:
// The element 3 appears once, and the element 7 appears twice. The remaining elements each appear three times.

// Example 2:
// Input: nums = [4,4,6,4,9,9,9,6,8]
// Output: [8,6]
// Explanation:
// The element 8 appears once, and the element 6 appears twice. The remaining elements each appear three times.

// Constraints:
//     3 <= nums.length <= 10^5
//     -2^31 <= nums[i] <= 2^31 - 1
//     nums.length is a multiple of 3.
//     Exactly one element appears once, one element appears twice, and all other elements appear three times.

import "fmt"

func onceTwice(nums []int) []int {
    count := make([]int, 32)
    maskA, maskB, valA, valB := 0, 0, 0, 0
    for _, v := range nums {
        for i := 0; i < 32; i++ {
            count[i] += (v >> i) & 1
            count[i] %= 3
        }
    }
    for i := 0; i < 32; i++ {
        if count[i] == 1 {
            maskA |= 1 << i
            valA |= 1 << i
            maskB |= 1 << i
        } else if count[i] == 2 {
            maskA |= 1 << i
            valB |= 1 << i
            maskB |= 1 << i
        }
    }
    onesA, twosA, onesB, twosB := 0, 0, 0, 0
    for _, v := range nums {
        if (v & maskA) == valA {
            twosA |= onesA & v
            onesA ^= v
            common := onesA & twosA
            onesA &^= common
            twosA &^= common
        }
        if (v & maskB) == valB {
            twosB |= onesB & v
            onesB ^= v
            common := onesB & twosB
            onesB &^= common
            twosB &^= common
        }
    }
    return []int{ onesA, twosB }
}

func main() {
    // Example 1:
    // Input: nums = [2,2,3,2,5,5,5,7,7]
    // Output: [3,7]
    // Explanation:
    // The element 3 appears once, and the element 7 appears twice. The remaining elements each appear three times.
    fmt.Println(onceTwice([]int{2,2,3,2,5,5,5,7,7})) // [3,7]
    // Example 2:
    // Input: nums = [4,4,6,4,9,9,9,6,8]
    // Output: [8,6]
    // Explanation:
    // The element 8 appears once, and the element 6 appears twice. The remaining elements each appear three times.
    fmt.Println(onceTwice([]int{4,4,6,4,9,9,9,6,8})) // [8,6]

    fmt.Println(onceTwice([]int{1,2,3,4,5,6,7,8,9})) // [6 0]
    fmt.Println(onceTwice([]int{9,8,7,6,5,4,3,2,1})) // [6 0]
}