package main

// 3769. Sort Integers by Binary Reflection
// You are given an integer array nums.

// The binary reflection of a positive integer is defined as the number obtained by reversing the order of its binary digits (ignoring any leading zeros) and interpreting the resulting binary number as a decimal.

// Sort the array in ascending order based on the binary reflection of each element. 
// If two different numbers have the same binary reflection, the smaller original number should appear first.

// Return the resulting sorted array.

// Example 1:
// Input: nums = [4,5,4]
// Output: [4,4,5]
// Explanation:
// Binary reflections are:
// 4 -> (binary) 100 -> (reversed) 001 -> 1
// 5 -> (binary) 101 -> (reversed) 101 -> 5
// 4 -> (binary) 100 -> (reversed) 001 -> 1
// Sorting by the reflected values gives [4, 4, 5].

// Example 2:
// Input: nums = [3,6,5,8]
// Output: [8,3,6,5]
// Explanation:
// Binary reflections are:
// 3 -> (binary) 11 -> (reversed) 11 -> 3
// 6 -> (binary) 110 -> (reversed) 011 -> 3
// 5 -> (binary) 101 -> (reversed) 101 -> 5
// 8 -> (binary) 1000 -> (reversed) 0001 -> 1
// Sorting by the reflected values gives [8, 3, 6, 5].
// Note that 3 and 6 have the same reflection, so we arrange them in increasing order of original value.
 
// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 10^9

import "fmt"
import "sort"
import "strconv"

func sortByReflection(nums []int) []int {
    sort.Slice(nums, func(i, j int) bool {
        x, y := strconv.FormatInt(int64(nums[i]), 2), strconv.FormatInt(int64(nums[j]), 2)
        nX, nY := "", ""
        for i := len(x) - 1; i >= 0; i-- {
            nX += string(x[i])
        }
        for i := len(y) - 1; i >= 0; i-- {
            nY += string(y[i])
        }
        xN, _ := strconv.ParseInt(nX, 2, 64)
        yN, _ := strconv.ParseInt(nY, 2, 64)
        if xN == yN {
            return  nums[i] < nums[j]
        }
        return xN < yN
    })
    return nums
}

func sortByReflection1(nums []int) []int {
    binaryReflection := func(x int) int {
        if x <= 0 { return 0 }
        reflected := 0
        for x > 0 {
            reflected = (reflected << 1) | (x & 1)
            x >>= 1
        }
        return reflected
    }
    sort.Slice(nums, func(i, j int) bool {
        ri, rj := binaryReflection(nums[i]), binaryReflection(nums[j])
        if ri != rj {
            return ri < rj
        }
        return nums[i] < nums[j]
    })
    return nums
}

func sortByReflection2(nums []int) []int {
    binaryReflection := func(num int) int {
        res := 0
        for num > 0 {
            res <<= 1
            res += num & 0b01
            num >>= 1
        }
        return res
    }
    sort.Slice(nums, func(i, j int) bool {
        c1, c2 := binaryReflection(nums[i]), binaryReflection(nums[j])
        if c1 != c2 {
            return c1 < c2
        } else {
            return nums[i] < nums[j]
        }
    })
    return nums
}

func main() {
    // Example 1:
    // Input: nums = [4,5,4]
    // Output: [4,4,5]
    // Explanation:
    // Binary reflections are:
    // 4 -> (binary) 100 -> (reversed) 001 -> 1
    // 5 -> (binary) 101 -> (reversed) 101 -> 5
    // 4 -> (binary) 100 -> (reversed) 001 -> 1
    // Sorting by the reflected values gives [4, 4, 5].
    fmt.Println(sortByReflection([]int{4,5,4})) // [4,4,5]
    // Example 2:
    // Input: nums = [3,6,5,8]
    // Output: [8,3,6,5]
    // Explanation:
    // Binary reflections are:
    // 3 -> (binary) 11 -> (reversed) 11 -> 3
    // 6 -> (binary) 110 -> (reversed) 011 -> 3
    // 5 -> (binary) 101 -> (reversed) 101 -> 5
    // 8 -> (binary) 1000 -> (reversed) 0001 -> 1
    // Sorting by the reflected values gives [8, 3, 6, 5].
    // Note that 3 and 6 have the same reflection, so we arrange them in increasing order of original value.
    fmt.Println(sortByReflection([]int{3,6,5,8})) // [8,3,6,5]

    fmt.Println(sortByReflection([]int{1,2,3,4,5,6,7,8,9})) // [1 2 4 8 3 6 5 7 9]
    fmt.Println(sortByReflection([]int{9,8,7,6,5,4,3,2,1})) // [1 2 4 8 3 6 5 7 9]

    fmt.Println(sortByReflection1([]int{4,5,4})) // [4,4,5]
    fmt.Println(sortByReflection1([]int{3,6,5,8})) // [8,3,6,5]
    fmt.Println(sortByReflection1([]int{1,2,3,4,5,6,7,8,9})) // [1 2 4 8 3 6 5 7 9]
    fmt.Println(sortByReflection1([]int{9,8,7,6,5,4,3,2,1})) // [1 2 4 8 3 6 5 7 9]

    fmt.Println(sortByReflection2([]int{4,5,4})) // [4,4,5]
    fmt.Println(sortByReflection2([]int{3,6,5,8})) // [8,3,6,5]
    fmt.Println(sortByReflection2([]int{1,2,3,4,5,6,7,8,9})) // [1 2 4 8 3 6 5 7 9]
    fmt.Println(sortByReflection2([]int{9,8,7,6,5,4,3,2,1})) // [1 2 4 8 3 6 5 7 9]
}