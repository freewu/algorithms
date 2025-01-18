package main

// 2560. House Robber IV
// There are several consecutive houses along a street, each of which has some money inside. 
// There is also a robber, who wants to steal money from the homes, but he refuses to steal from adjacent homes.

// The capability of the robber is the maximum amount of money he steals from one house of all the houses he robbed.

// You are given an integer array nums representing how much money is stashed in each house. 
// More formally, the ith house from the left has nums[i] dollars.

// You are also given an integer k, representing the minimum number of houses the robber will steal from. 
// It is always possible to steal at least k houses.

// Return the minimum capability of the robber out of all the possible ways to steal at least k houses.

// Example 1:
// Input: nums = [2,3,5,9], k = 2
// Output: 5
// Explanation: 
// There are three ways to rob at least 2 houses:
// - Rob the houses at indices 0 and 2. Capability is max(nums[0], nums[2]) = 5.
// - Rob the houses at indices 0 and 3. Capability is max(nums[0], nums[3]) = 9.
// - Rob the houses at indices 1 and 3. Capability is max(nums[1], nums[3]) = 9.
// Therefore, we return min(5, 9, 9) = 5.

// Example 2:
// Input: nums = [2,7,9,3,1], k = 2
// Output: 2
// Explanation: There are 7 ways to rob the houses. The way which leads to minimum capability is to rob the house at index 0 and 4. Return max(nums[0], nums[4]) = 2.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= k <= (nums.length + 1)/2

import "fmt"
import "slices"
import "sort"

func minCapability1(nums []int, k int) int {
    mx, n := slices.Max(nums), len(nums)
    return sort.Search(mx, func(p int) bool {
        s := 0
        for i := 0; i < n; i++ {
            if nums[i] <= p {
                s++
                i++
            }
        }
        return s >= k
    })
}

func minCapability(nums []int, k int) int {
    left, right := 1, 1000000000
    for left <= right {
        v, mid := 0, (left + right) >> 1
        for i := 0; i < len(nums); {
            if nums[i] <= mid {
                i += 2
                v++
                if v >= k { break }
            } else {
                i++
            }
        }
        if v >= k {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return right + 1
}

func minCapability2(nums []int, k int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    left, right := nums[0], nums[0]
    for _, v := range nums {
        left, right = min(left, v), max(right, v)
    }
    check := func(x int) bool {
        count, flag := 0, false
        for _, v := range nums {
            if flag { // reset
                flag = false
                continue
            }
            if v > x { continue }
            flag = true
            count++
        }
        return count >= k
    }
    for left < right {
        mid := (left + right) / 2
        if check(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func main() {
    // Example 1:
    // Input: nums = [2,3,5,9], k = 2
    // Output: 5
    // Explanation: 
    // There are three ways to rob at least 2 houses:
    // - Rob the houses at indices 0 and 2. Capability is max(nums[0], nums[2]) = 5.
    // - Rob the houses at indices 0 and 3. Capability is max(nums[0], nums[3]) = 9.
    // - Rob the houses at indices 1 and 3. Capability is max(nums[1], nums[3]) = 9.
    // Therefore, we return min(5, 9, 9) = 5.
    fmt.Println(minCapability([]int{2,3,5,9}, 2)) // 5
    // Example 2:
    // Input: nums = [2,7,9,3,1], k = 2
    // Output: 2
    // Explanation: There are 7 ways to rob the houses. The way which leads to minimum capability is to rob the house at index 0 and 4. Return max(nums[0], nums[4]) = 2.
    fmt.Println(minCapability([]int{2,7,9,3,1}, 2)) // 2

    fmt.Println(minCapability([]int{1,2,3,4,5,6,7,8,9}, 2)) // 3
    fmt.Println(minCapability([]int{9,8,7,6,5,4,3,2,1}, 2)) // 3

    fmt.Println(minCapability1([]int{2,3,5,9}, 2)) // 5
    fmt.Println(minCapability1([]int{2,7,9,3,1}, 2)) // 2
    fmt.Println(minCapability1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 3
    fmt.Println(minCapability1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 3

    fmt.Println(minCapability2([]int{2,3,5,9}, 2)) // 5
    fmt.Println(minCapability2([]int{2,7,9,3,1}, 2)) // 2
    fmt.Println(minCapability2([]int{1,2,3,4,5,6,7,8,9}, 2)) // 3
    fmt.Println(minCapability2([]int{9,8,7,6,5,4,3,2,1}, 2)) // 3
}