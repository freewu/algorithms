package main

// 3987. Minimum Total Cost to Process All Elements
// You are given an integer array nums and an integer k.

// Initially, you have k units of resources.

// You must process the elements of nums from left to right. 
// To process the ith element, you need nums[i] resources.

// If your available resources are less than nums[i], you may perform an operation that increases your available resources by k. 
// The value of k is fixed and does not change throughout the process. 
// The first such operation incurs a cost of 1, the second incurs a cost of 2, and so on.

// After processing the ith element, your available resources decrease by nums[i].

// Return an integer denoting the minimum total cost required to process all elements. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [1,2,3,4], k = 4
// Output: 3
// Explanation:
// After processing nums[0], we have 4 - 1 = 3 units of resources left.
// After processing nums[1], we have 3 - 2 = 1 unit of resources left.
// Since nums[2] = 3 and only 1 unit of resources is available, we perform the first operation costing 1. After processing nums[2], we have 1 + 4 - 3 = 2 units of resources left.
// Since nums[3] = 4 and only 2 units of resources are available, we perform the second operation costing 2, to have 2 + 4 = 6 units of resources, which is enough to process nums[3].
// Thus, the total cost is 1 + 2 = 3.

// Example 2:
// Input: nums = [1,1,7,14], k = 4
// Output: 15
// Explanation:
// After processing nums[0], we have 4 - 1 = 3 units of resources left.
// After processing nums[1], we have 3 - 1 = 2 units of resources left.
// Since nums[2] = 7 and only 2 units of resources are available, we perform two operations costing 1 + 2 = 3. After processing nums[2], we have 2 + 4 + 4 - 7 = 3 units of resources left.
// Since nums[3] = 14 and only 3 units of resources are available, we perform three operations costing 3 + 4 + 5 = 12, to have 3 + 4 + 4 + 4 = 15 units of resources, which is enough to process nums[3].
// Thus, the total cost is 3 + 12 = 15.

// Example 3:
// Input: nums = [1,2,3,4], k = 10
// Output: 0
// Explanation:
// To process all elements, we can use the initial 10 units of resources without performing any operations. Thus, the total cost required is 0.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= k <= 10^9

import "fmt"

func minimumCost(nums []int, k int) int {
    left, sum := k, 0 // sum 总操作次数
    for _, v := range nums {
        if left < v {
            op := (v - left - 1) / k + 1 // 把 left 增大到 >= v，至少操作 op 次
            sum += op
            left += op * k
        }
        left -= v
    }
    // 1 + 2 + ... + sum
    sum %= 1_000_000_007
    return sum * (sum + 1) / 2 % 1_000_000_007
}

func minimumCost1(nums []int, k int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    count := (sum - 1) / k
    a, b := count, count+1
    if a % 2 == 0 {
        a /= 2
    }
    if b % 2 == 0 {
        b /= 2
    }
    a %= 1_000_000_007
    b %= 1_000_000_007
    res := a * b
    // res %= 1_000_000_007
    return res % 1_000_000_007
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4], k = 4
    // Output: 3
    // Explanation:
    // After processing nums[0], we have 4 - 1 = 3 units of resources left.
    // After processing nums[1], we have 3 - 2 = 1 unit of resources left.
    // Since nums[2] = 3 and only 1 unit of resources is available, we perform the first operation costing 1. After processing nums[2], we have 1 + 4 - 3 = 2 units of resources left.
    // Since nums[3] = 4 and only 2 units of resources are available, we perform the second operation costing 2, to have 2 + 4 = 6 units of resources, which is enough to process nums[3].
    // Thus, the total cost is 1 + 2 = 3.
    fmt.Println(minimumCost([]int{1,2,3,4}, 4)) // 3
    // Example 2:
    // Input: nums = [1,1,7,14], k = 4
    // Output: 15
    // Explanation:
    // After processing nums[0], we have 4 - 1 = 3 units of resources left.
    // After processing nums[1], we have 3 - 1 = 2 units of resources left.
    // Since nums[2] = 7 and only 2 units of resources are available, we perform two operations costing 1 + 2 = 3. After processing nums[2], we have 2 + 4 + 4 - 7 = 3 units of resources left.
    // Since nums[3] = 14 and only 3 units of resources are available, we perform three operations costing 3 + 4 + 5 = 12, to have 3 + 4 + 4 + 4 = 15 units of resources, which is enough to process nums[3].
    // Thus, the total cost is 3 + 12 = 15.
    fmt.Println(minimumCost([]int{1,1,7,14}, 4)) // 15
    // Example 3:
    // Input: nums = [1,2,3,4], k = 10
    // Output: 0
    // Explanation:
    // To process all elements, we can use the initial 10 units of resources without performing any operations. Thus, the total cost required is 0.
    fmt.Println(minimumCost([]int{1,2,3,4}, 10)) // 0

    fmt.Println(minimumCost([]int{1,2,3,4,5,6,7,8,9}, 2)) // 253
    fmt.Println(minimumCost([]int{9,8,7,6,5,4,3,2,1}, 2)) // 253

    fmt.Println(minimumCost1([]int{1,2,3,4}, 4)) // 3
    fmt.Println(minimumCost1([]int{1,1,7,14}, 4)) // 15
    fmt.Println(minimumCost1([]int{1,2,3,4}, 10)) // 0
    fmt.Println(minimumCost1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 253
    fmt.Println(minimumCost1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 253
}
