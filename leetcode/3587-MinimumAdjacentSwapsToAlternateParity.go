package main

// 3587. Minimum Adjacent Swaps to Alternate Parity
// You are given an array nums of distinct integers.

// In one operation, you can swap any two adjacent elements in the array.

// An arrangement of the array is considered valid if the parity of adjacent elements alternates, meaning every pair of neighboring elements consists of one even and one odd number.

// Return the minimum number of adjacent swaps required to transform nums into any valid arrangement.

// If it is impossible to rearrange nums such that no two adjacent elements have the same parity, return -1.

// Example 1:
// Input: nums = [2,4,6,5,7]
// Output: 3
// Explanation:
// Swapping 5 and 6, the array becomes [2,4,5,6,7]
// Swapping 5 and 4, the array becomes [2,5,4,6,7]
// Swapping 6 and 7, the array becomes [2,5,4,7,6]. The array is now a valid arrangement. Thus, the answer is 3.

// Example 2:
// Input: nums = [2,4,5,7]
// Output: 1
// Explanation:
// By swapping 4 and 5, the array becomes [2,5,4,7], which is a valid arrangement. Thus, the answer is 1.

// Example 3:
// Input: nums = [1,2,3]
// Output: 0
// Explanation:
// The array is already a valid arrangement. Thus, no operations are needed.

// Example 4:
// Input: nums = [4,5,6,8]
// Output: -1
// Explanation:
// No valid arrangement is possible. Thus, the answer is -1.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     All elements in nums are distinct.

import "fmt"
import "sort"

func minSwaps(nums []int) int {
    odd, even := make([]int, 0), make([]int, 0)
    for i := 0; i < len(nums); i++ {
        nums[i] = nums[i] & 1
        if nums[i]&1 == 1 {
            odd = append(odd, i)
        } else {
            even = append(even, i)
        }
    }
    calc := func(odd, even []int, flag bool) int {
        res, neven, nodd := 0, make([]int, len(even)), make([]int, len(odd))
        copy(neven, even)
        copy(nodd, odd)
        for i := 0; i < len(nums); i++ {
            if flag && len(neven) == 0 {
                return 1 << 31
            }
            if !flag && len(nodd) == 0 {
                return 1 << 31
            }
            if flag {
                t := neven[0]
                index := sort.SearchInts(nodd, t)
                res += index
                neven = neven[1:]
            } else {
                t := nodd[0]
                index := sort.SearchInts(neven, t)
                res += index
                nodd = nodd[1:]
            }
            flag = !flag
        }
        return res
    }
    res := min(calc(odd, even, false), calc(odd, even, true))
    if res == 1 << 31 { return -1 } 
    return res
}

func minSwaps1(nums []int) int {
    pos := make([][]int,2)
    for i := range nums {
        pos[nums[i] % 2] = append(pos[nums[i] % 2],i)
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res, inf := 1 << 31, 1 << 31
    calc := func(x int) int {
        res,turn, index := 0, x, [2]int{}
        for i := range nums {
            if index[turn] < len(pos[turn]) {
                res += abs(pos[turn][index[turn]] - i)
                index[turn]++
            } else {
                return inf
            }
            turn ^= 1
        }    
        return res / 2
    }
    res = min(calc(0),calc(1))
    if res == inf {
        res = -1
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,4,6,5,7]
    // Output: 3
    // Explanation:
    // Swapping 5 and 6, the array becomes [2,4,5,6,7]
    // Swapping 5 and 4, the array becomes [2,5,4,6,7]
    // Swapping 6 and 7, the array becomes [2,5,4,7,6]. The array is now a valid arrangement. Thus, the answer is 3.
    fmt.Println(minSwaps([]int{2,4,6,5,7})) // 3
    // Example 2:
    // Input: nums = [2,4,5,7]
    // Output: 1
    // Explanation:
    // By swapping 4 and 5, the array becomes [2,5,4,7], which is a valid arrangement. Thus, the answer is 1.
    fmt.Println(minSwaps([]int{2,4,5,7})) // 1
    // Example 3:
    // Input: nums = [1,2,3]
    // Output: 0
    // Explanation:
    // The array is already a valid arrangement. Thus, no operations are needed.
    fmt.Println(minSwaps([]int{1,2,3})) // 0
    // Example 4:
    // Input: nums = [4,5,6,8]
    // Output: -1
    // Explanation:
    // No valid arrangement is possible. Thus, the answer is -1.
    fmt.Println(minSwaps([]int{4,5,6,8})) // -1

    fmt.Println(minSwaps([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minSwaps([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(minSwaps1([]int{2,4,6,5,7})) // 3
    fmt.Println(minSwaps1([]int{2,4,5,7})) // 1
    fmt.Println(minSwaps1([]int{1,2,3})) // 0
    fmt.Println(minSwaps1([]int{4,5,6,8})) // -1
    fmt.Println(minSwaps1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minSwaps1([]int{9,8,7,6,5,4,3,2,1})) // 0
}