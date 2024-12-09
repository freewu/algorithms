package main

// 2964. Number of Divisible Triplet Sums
// Given a 0-indexed integer array nums and an integer d, 
// return the number of triplets (i, j, k) such that i < j < k and (nums[i] + nums[j] + nums[k]) % d == 0.

// Example 1:
// Input: nums = [3,3,4,7,8], d = 5
// Output: 3
// Explanation: The triplets which are divisible by 5 are: (0, 1, 2), (0, 2, 4), (1, 2, 4).
// It can be shown that no other triplet is divisible by 5. Hence, the answer is 3.

// Example 2:
// Input: nums = [3,3,3,3], d = 3
// Output: 4
// Explanation: Any triplet chosen here has a sum of 9, which is divisible by 3. Hence, the answer is the total number of triplets which is 4.

// Example 3:
// Input: nums = [3,3,3,3], d = 6
// Output: 0
// Explanation: Any triplet chosen here has a sum of 9, which is not divisible by 6. Hence, the answer is 0.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 10^9
//     1 <= d <= 10^9

import "fmt"

func divisibleTripletCount(nums []int, d int) int {
    res, n := 0, len(nums)
    mp := make(map[int]int)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            res += mp[(d - (nums[i] + nums[j]) % d) % d]
        }
        mp[nums[i] % d]++
    }
    return res
}

func divisibleTripletCount1(nums []int, d int) int {
    res, n := 0, len(nums)
    count := make([]int, d)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            res += count[(d - (nums[i] + nums[j]) % d) % d]
        }
        count[nums[i] % d]++
    }
    return res
}

func divisibleTripletCount2(nums []int, d int) int {
    res, count := 0, make([]int, d)
    for i, v := range nums {
        res += count[(d - v % d) % d]
        for j := 0; j < i; j++ {
            count[(nums[j] + nums[i]) % d]++
        } 
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,3,4,7,8], d = 5
    // Output: 3
    // Explanation: The triplets which are divisible by 5 are: (0, 1, 2), (0, 2, 4), (1, 2, 4).
    // It can be shown that no other triplet is divisible by 5. Hence, the answer is 3.
    fmt.Println(divisibleTripletCount([]int{3,3,4,7,8}, 5)) // 3
    // Example 2:
    // Input: nums = [3,3,3,3], d = 3
    // Output: 4
    // Explanation: Any triplet chosen here has a sum of 9, which is divisible by 3. Hence, the answer is the total number of triplets which is 4.
    fmt.Println(divisibleTripletCount([]int{3,3,3,3}, 3)) // 4
    // Example 3:
    // Input: nums = [3,3,3,3], d = 6
    // Output: 0
    // Explanation: Any triplet chosen here has a sum of 9, which is not divisible by 6. Hence, the answer is 0.
    fmt.Println(divisibleTripletCount([]int{3,3,3,3}, 6)) // 0

    fmt.Println(divisibleTripletCount1([]int{3,3,4,7,8}, 5)) // 3
    fmt.Println(divisibleTripletCount1([]int{3,3,3,3}, 3)) // 4
    fmt.Println(divisibleTripletCount1([]int{3,3,3,3}, 6)) // 0

    fmt.Println(divisibleTripletCount2([]int{3,3,4,7,8}, 5)) // 3
    fmt.Println(divisibleTripletCount2([]int{3,3,3,3}, 3)) // 4
    fmt.Println(divisibleTripletCount2([]int{3,3,3,3}, 6)) // 0
}