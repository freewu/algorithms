package main

// 4121. Minimum Cost to Make All Array Elements Equal
// You are given an integer array nums of length n.

// You may perform any of the following operations any number of times, in any order.
//     1. Prefix operation: Choose an index i such that 0 < i < n - 1, and set all elements from index 0 to i - 1 to nums[i]. 
//        The cost of this operation is i * nums[i].
//     2. Suffix operation: Choose an index i such that 0 < i < n - 1, and set all elements from index i + 1 to n - 1 to nums[i]. 
//        The cost of this operation is (n - i - 1) * nums[i].

// The cost is based on the current value of nums[i] and includes every element in the chosen prefix or suffix, 
// even if it already equals nums[i].

// Return the minimum total cost required to make all elements of nums equal.

// Example 1:
// Input: nums = [1,2,3]
// Output: 4
// Explanation:
// One sequence of operations that achieves the minimum cost is as follows:
//  nums      | Operation | i | Modified nums | Cost
//  [1, 2, 3] | Suffix    | 1 | [1, 2, 2]     | (3 - 1 - 1) * 2 = 2
//  [1, 2, 2] | Prefix    | 1 | [2, 2, 2]     | 1 * 2 = 2
// Hence, the total cost is 2 + 2 = 4.

// Example 2:
// Input: nums = [3,3,3,6,3,3]
// Output: 9
// Explanation:
// One sequence of operations that achieves the minimum cost is as follows:
//  nums                  | Operation | i | Modified nums          | Cost
//  [3, 3, 3, 6, 3, 3]    | Suffix    | 2 | [3, 3, 3, 3, 3, 3]     | (6 - 2 - 1) * 3 = 9
// Hence, the total cost is 9.

// Example 3:
// Input: nums = [3,4,4,5,4,4,3]
// Output: 20
// Explanation:
// One sequence the minimum cost is the minimum cost is as follows:
//  nums                  | Operation | i | Modified nums         | Cost
//  [3, 4, 4, 5, 4, 4, 3] | Prefix    | 4 | [4, 4, 4, 4, 4, 4, 3] | 4 * 4 = 16
//  [4, 4, 4, 4, 4, 4, 3] | Suffix    | 5 | [4, 4, 4, 4, 4, 4, 4] | (7 - 5 - 1) * 4 = 4
// Hence, the total cost is 16 + 4 = 20.

// Constraints:
//     3 <= n == nums.length <= 10^{5}
//     1 <= nums[i] <= 10^{9}

import "fmt"

func minCost(nums []int) int64 {
    n := len(nums)
    flag := true
    for i := 1; i < n; i++ {
        if nums[i] != nums[0] {
            flag = false
            break
        }
    }
    if flag {
        return 0
    }
    res, l := int64(1 << 61), 0
    for l < n {
        r := l
        for r + 1 < n && nums[r+1] == nums[l] { // 找到连续相等的元素
            r++
        }
        if r >= 1 && l <= n - 2 { // 避免边界情况
            runLen := int64(r - l + 1)
            cost := (int64(n) - runLen) * int64(nums[l])
            if cost < res {
                res = cost
            }
        }
        l = r + 1
    }
    return res
}

func minCost1(nums []int) int64 {
    res := 1 << 61
    for i := 0; i < len(nums); i++ {
        j := i
        for j < len(nums) && nums[j] == nums[i] { // 找到连续相等的元素
            j++
        }
        if i == j-1 && (i == 0 || i == len(nums)-1) { // 避免边界情况
            continue
        }
        res = min(res, nums[i]*(len(nums)-(j-i))) // 计算当前元素的最小成本
        i = j - 1
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 4
    // Explanation:
    // One sequence of operations that achieves the minimum cost is as follows:
    //  nums      | Operation | i | Modified nums | Cost
    //  [1, 2, 3] | Suffix    | 1 | [1, 2, 2]     | (3 - 1 - 1) * 2 = 2
    //  [1, 2, 2] | Prefix    | 1 | [2, 2, 2]     | 1 * 2 = 2
    // Hence, the total cost is 2 + 2 = 4.
    fmt.Println(minCost([]int{1,2,3})) // 4
    // Example 2:
    // Input: nums = [3,3,3,6,3,3]
    // Output: 9
    // Explanation:
    // One sequence of operations that achieves the minimum cost is as follows:
    //  nums                  | Operation | i | Modified nums          | Cost
    //  [3, 3, 3, 6, 3, 3]    | Suffix    | 2 | [3, 3, 3, 3, 3, 3]     | (6 - 2 - 1) * 3 = 9
    // Hence, the total cost is 9.
    fmt.Println(minCost([]int{3,3,3,6,3,3})) // 9
    // Example 3:
    // Input: nums = [3,4,4,5,4,4,3]
    // Output: 20
    // Explanation:
    // One sequence the minimum cost is the minimum cost is as follows:
    //  nums                  | Operation | i | Modified nums         | Cost
    //  [3, 4, 4, 5, 4, 4, 3] | Prefix    | 4 | [4, 4, 4, 4, 4, 4, 3] | 4 * 4 = 16
    //  [4, 4, 4, 4, 4, 4, 3] | Suffix    | 5 | [4, 4, 4, 4, 4, 4, 4] | (7 - 5 - 1) * 4 = 4
    // Hence, the total cost is 16 + 4 = 20.
    fmt.Println(minCost([]int{3,4,4,5,4,4,3})) // 20

    fmt.Println(minCost([]int{1,2,3,4,5,6,7,8,9})) // 16
    fmt.Println(minCost([]int{9,8,7,6,5,4,3,2,1})) // 16

    fmt.Println(minCost1([]int{1,2,3})) // 4
    fmt.Println(minCost1([]int{3,3,3,6,3,3})) // 9
    fmt.Println(minCost1([]int{3,4,4,5,4,4,3})) // 20
    fmt.Println(minCost1([]int{1,2,3,4,5,6,7,8,9})) // 16
    fmt.Println(minCost1([]int{9,8,7,6,5,4,3,2,1})) // 16
}