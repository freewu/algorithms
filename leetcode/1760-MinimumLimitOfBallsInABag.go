package main

// 1760. Minimum Limit of Balls in a Bag
// You are given an integer array nums where the ith bag contains nums[i] balls. 
// You are also given an integer maxOperations.

// You can perform the following operation at most maxOperations times:
//     Take any bag of balls and divide it into two new bags with a positive number of balls.
//         For example, a bag of 5 balls can become two new bags of 1 and 4 balls, or two new bags of 2 and 3 balls.

// Your penalty is the maximum number of balls in a bag. 
// You want to minimize your penalty after the operations.

// Return the minimum possible penalty after performing the operations.

// Example 1:
// Input: nums = [9], maxOperations = 2
// Output: 3
// Explanation: 
// - Divide the bag with 9 balls into two bags of sizes 6 and 3. [9] -> [6,3].
// - Divide the bag with 6 balls into two bags of sizes 3 and 3. [6,3] -> [3,3,3].
// The bag with the most number of balls has 3 balls, so your penalty is 3 and you should return 3.

// Example 2:
// Input: nums = [2,4,8,2], maxOperations = 4
// Output: 2
// Explanation:
// - Divide the bag with 8 balls into two bags of sizes 4 and 4. [2,4,8,2] -> [2,4,4,4,2].
// - Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,4,4,4,2] -> [2,2,2,4,4,2].
// - Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,2,2,4,4,2] -> [2,2,2,2,2,4,2].
// - Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,2,2,2,2,4,2] -> [2,2,2,2,2,2,2,2].
// The bag with the most number of balls has 2 balls, so your penalty is 2, and you should return 2.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= maxOperations, nums[i] <= 10^9

import "fmt"
import "sort"

func minimumSize(nums []int, maxOperations int) int {
    mn, mx := 1, 1
    for _, v := range nums {
        if v > mx {
            mx = v
        }
    }
    isAchievable := func(nums []int, maxOperations int, target int) bool {
        operations := 0
        for _, balls := range nums {
            bags := balls / target 
            if balls % target != 0 {
                bags++
            }
            operations += bags - 1;
            if operations > maxOperations { return false }
        }
        return true
    }
    for mn < mx {
        mid := (mx + mn) / 2
        if isAchievable(nums, maxOperations, mid) {
            mx = mid
        } else {
            mn = mid + 1
        }        
    }
    return mn
}

func minimumSize1(nums []int, maxOperations int) int {
    left, right := 1, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, num := range nums {
        right = max(right, num)
    }
    check := func(nums []int, cost int, maxOperations int) bool {
        count := 0
        for _, v := range nums {
            count += (v - 1) / cost
            if count > maxOperations { return false }
        }
        return true
    }
    for left <= right {
        mid := (left + right) / 2
        if check(nums, mid, maxOperations) {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return left
}

func minimumSize2(nums []int, maxOperations int) int {
    return sort.Search(1_000_000_000, func(i int) bool {
        if i == 0 { return false }
        count := 0
        for _, v := range nums {
            count += (v + i - 1) / i - 1
        }
        return count <= maxOperations
    })
}

func main() {
    // Example 1:
    // Input: nums = [9], maxOperations = 2
    // Output: 3
    // Explanation: 
    // - Divide the bag with 9 balls into two bags of sizes 6 and 3. [9] -> [6,3].
    // - Divide the bag with 6 balls into two bags of sizes 3 and 3. [6,3] -> [3,3,3].
    // The bag with the most number of balls has 3 balls, so your penalty is 3 and you should return 3.
    fmt.Println(minimumSize([]int{9}, 2)) // 3
    // Example 2: 
    // Input: nums = [2,4,8,2], maxOperations = 4
    // Output: 2
    // Explanation:
    // - Divide the bag with 8 balls into two bags of sizes 4 and 4. [2,4,8,2] -> [2,4,4,4,2].
    // - Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,4,4,4,2] -> [2,2,2,4,4,2].
    // - Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,2,2,4,4,2] -> [2,2,2,2,2,4,2].
    // - Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,2,2,2,2,4,2] -> [2,2,2,2,2,2,2,2].
    // The bag with the most number of balls has 2 balls, so your penalty is 2, and you should return 2.
    fmt.Println(minimumSize([]int{2,4,8,2}, 4)) // 2

    fmt.Println(minimumSize([]int{1,2,3,4,5,6,7,8,9}, 4)) // 5
    fmt.Println(minimumSize([]int{9,8,7,6,5,4,3,2,1}, 4)) // 5

    fmt.Println(minimumSize1([]int{9}, 2)) // 3
    fmt.Println(minimumSize1([]int{2,4,8,2}, 4)) // 2
    fmt.Println(minimumSize1([]int{1,2,3,4,5,6,7,8,9}, 4)) // 5
    fmt.Println(minimumSize1([]int{9,8,7,6,5,4,3,2,1}, 4)) // 5

    fmt.Println(minimumSize2([]int{9}, 2)) // 3
    fmt.Println(minimumSize2([]int{2,4,8,2}, 4)) // 2
    fmt.Println(minimumSize2([]int{1,2,3,4,5,6,7,8,9}, 4)) // 5
    fmt.Println(minimumSize2([]int{9,8,7,6,5,4,3,2,1}, 4)) // 5
}