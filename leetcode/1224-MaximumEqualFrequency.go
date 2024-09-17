package main

// 1224. Maximum Equal Frequency
// Given an array nums of positive integers, 
// return the longest possible length of an array prefix of nums, 
// such that it is possible to remove exactly one element from this prefix so 
// that every number that has appeared in it will have the same number of occurrences.

// If after removing one element there are no remaining elements, 
// it's still considered that every appeared number has the same number of ocurrences (0).

// Example 1:
// Input: nums = [2,2,1,1,5,3,3,5]
// Output: 7
// Explanation: For the subarray [2,2,1,1,5,3,3] of length 7, if we remove nums[4] = 5, we will get [2,2,1,1,3,3], so that each number will appear exactly twice.

// Example 2:
// Input: nums = [1,1,1,2,2,2,3,3,3,4,4,4,5]
// Output: 13

// Constraints:
//     2 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func maxEqualFreq(nums []int) int {
    mx, n := 0, len(nums)
    for _, v := range nums { // Find the maximum value in nums
        if v > mx {
            mx = v
        }
    }
    mapFreq := make([]int, n + 1) // mapping freq -> no. of elements
    a := make([]int, mx + 1)          // mapping element -> freq
    res, maxFreq, minFreq, count := 0, 0, 1 << 31 - 1, 0
    for i := 0; i < n; i++ {
        freq := a[nums[i]] + 1
        a[nums[i]] = freq
        if mapFreq[freq] == 0 { // adding current frequency
            count++
        }
        mapFreq[freq]++
        if freq > 1 && mapFreq[freq-1] == 1 { // remove previous frequency
            count--
            mapFreq[freq-1]--
            if freq-1 == minFreq { // updating minFreq
                minFreq++
            }
        } else if freq > 1 {
            mapFreq[freq-1]--
        }
        if freq > maxFreq { maxFreq = freq }
        if freq < minFreq { minFreq = freq }
        if count == 1 && (maxFreq == 1 || mapFreq[maxFreq] == 1) { // unique frequencies == 1 && (maxFreq==1 || number of elements at maxFreq == 1)
            if i + 1 > res { res = i + 1 }
        }
        if count == 2 { // unique frequencies == 2
            // element to be removed has (freq 1 || freq same as others+1)
            if mapFreq[1] == 1 || (mapFreq[maxFreq] == 1 && maxFreq == minFreq+1) {
                if i+1 > res { res = i + 1 }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,2,1,1,5,3,3,5]
    // Output: 7
    // Explanation: For the subarray [2,2,1,1,5,3,3] of length 7, if we remove nums[4] = 5, we will get [2,2,1,1,3,3], so that each number will appear exactly twice.
    fmt.Println(maxEqualFreq([]int{2,2,1,1,5,3,3,5})) // 7
    // Example 2:
    // Input: nums = [1,1,1,2,2,2,3,3,3,4,4,4,5]
    // Output: 13
    fmt.Println(maxEqualFreq([]int{1,1,1,2,2,2,3,3,3,4,4,4,5})) // 13
}