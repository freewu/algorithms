package main

// 2616. Minimize the Maximum Difference of Pairs
// You are given a 0-indexed integer array nums and an integer p. 
// Find p pairs of indices of nums such that the maximum difference amongst all the pairs is minimized. 
// Also, ensure no index appears more than once amongst the p pairs.

// Note that for a pair of elements at the index i and j, the difference of this pair is |nums[i] - nums[j]|, 
// where |x| represents the absolute value of x.

// Return the minimum maximum difference among all p pairs. We define the maximum of an empty set to be zero.

// Example 1:
// Input: nums = [10,1,2,7,1,3], p = 2
// Output: 1
// Explanation: The first pair is formed from the indices 1 and 4, and the second pair is formed from the indices 2 and 5. 
// The maximum difference is max(|nums[1] - nums[4]|, |nums[2] - nums[5]|) = max(0, 1) = 1.
// Therefore, we return 1.

// Example 2:
// Input: nums = [4,2,1,2], p = 1
// Output: 0
// Explanation: Let the indices 1 and 3 form a pair. 
// The difference of that pair is |2 - 2| = 0, which is the minimum we can attain.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9
//     0 <= p <= (nums.length)/2

import "fmt"
import "sort"
import "slices"

func minimizeMax(nums []int, p int) int {
    res, n:= 1 << 31,  len(nums)
    sort.Ints(nums)
    mx := nums[n - 1]
    if n > 1 {
        mx -= nums[0]
    }
    check := func(t int) bool {
        i, count := 1, 0
        for i < n {
            if nums[i] - nums[i-1] <= t {
                count++
                if count == p { return true }
                i += 2
            } else {
                i++
            }
        }
        return count >= p
    }
    left, right := 0, mx
    for left <= right {
        mid := left + (right - left) >> 1
        if check(mid) {
            res, right = min(res, mid), mid - 1
        } else {
            left = mid + 1
        }
    }
    return res
}

func minimizeMax1(nums []int, p int) int {
    sort.Ints(nums)
    check := func(target int)bool{
        count := 0
        for i := 0; i < len(nums)-1; i++ {
            if nums[i + 1] - nums[i] <= target { // 都选
                count++
                i++
            }
        }
        return count >= p
    }
    left, right := 0, slices.Max(nums) - 1
    for left <= right {
        mid := left + (right - left) >> 1
        if check(mid){
            right = mid-1
        }else{
            left = mid+1
        }
    }
    return left 
}

func main() {
    // Example 1:
    // Input: nums = [10,1,2,7,1,3], p = 2
    // Output: 1
    // Explanation: The first pair is formed from the indices 1 and 4, and the second pair is formed from the indices 2 and 5. 
    // The maximum difference is max(|nums[1] - nums[4]|, |nums[2] - nums[5]|) = max(0, 1) = 1.
    // Therefore, we return 1.
    fmt.Println(minimizeMax([]int{10,1,2,7,1,3}, 2)) // 1
    // Example 2:
    // Input: nums = [4,2,1,2], p = 1
    // Output: 0
    // Explanation: Let the indices 1 and 3 form a pair. 
    // The difference of that pair is |2 - 2| = 0, which is the minimum we can attain.
    fmt.Println(minimizeMax([]int{4,2,1,2}, 1)) // 0

    fmt.Println(minimizeMax([]int{1,2,3,4,5,6,7,8,9}, 1)) // 1
    fmt.Println(minimizeMax([]int{9,8,7,6,5,4,3,2,1}, 1)) // 1

    fmt.Println(minimizeMax1([]int{10,1,2,7,1,3}, 2)) // 1
    fmt.Println(minimizeMax1([]int{4,2,1,2}, 1)) // 0
    fmt.Println(minimizeMax1([]int{1,2,3,4,5,6,7,8,9}, 1)) // 1
    fmt.Println(minimizeMax1([]int{9,8,7,6,5,4,3,2,1}, 1)) // 1
}