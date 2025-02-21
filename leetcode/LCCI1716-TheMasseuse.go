package main

// 面试题 17.16. The Masseuse LCCI
// A popular masseuse receives a sequence of back-to-back appointment requests and is debating which ones to accept. 
// She needs a break between appointments and therefore she cannot accept any adjacent requests. 
// Given a sequence of back-to-back appoint­ ment requests, find the optimal (highest total booked minutes) set the masseuse can honor. 
// Return the number of minutes.

// Note: This problem is slightly different from the original one in the book.

// Example 1:
// Input:  [1,2,3,1]
// Output:  4
// Explanation:  Accept request 1 and 3, total minutes = 1 + 3 = 4

// Example 2:
// Input:  [2,7,9,3,1]
// Output:  12
// Explanation:  Accept request 1, 3 and 5, total minutes = 2 + 9 + 1 = 12

// Example 3:
// Input:  [2,1,4,5,3,1,1,3]
// Output:  12
// Explanation:  Accept request 1, 3, 5 and 8, total minutes = 2 + 4 + 3 + 3 = 12

import "fmt"

func massage(nums []int) int {
    if len(nums) == 0 { return 0 }
    if len(nums) == 1 { return nums[0] }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dp1, dp2 := max(nums[0], nums[1]), nums[0]
    for i := 2; i < len(nums); i++ {
        dp1, dp2 = max(dp1, dp2 + nums[i]), dp1
    }
    return dp1
}

func massage1(nums []int) int {
    x, y := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        y, x = max(y, x + v), y
    }
    return y
}

func main() {
    // Example 1:
    // Input:  [1,2,3,1]
    // Output:  4
    // Explanation:  Accept request 1 and 3, total minutes = 1 + 3 = 4
    fmt.Println(massage([]int{1,2,3,1})) // 4
    // Example 2:
    // Input:  [2,7,9,3,1]
    // Output:  12
    // Explanation:  Accept request 1, 3 and 5, total minutes = 2 + 9 + 1 = 12
    fmt.Println(massage([]int{2,7,9,3,1})) // 12
    // Example 3:
    // Input:  [2,1,4,5,3,1,1,3]
    // Output:  12
    // Explanation:  Accept request 1, 3, 5 and 8, total minutes = 2 + 4 + 3 + 3 = 12
    fmt.Println(massage([]int{2,1,4,5,3,1,1,3})) // 12

    fmt.Println(massage([]int{1,2,3,4,5,6,7,8,9})) // 25
    fmt.Println(massage([]int{9,8,7,6,5,4,3,2,1})) // 25

    fmt.Println(massage1([]int{1,2,3,1})) // 4
    fmt.Println(massage1([]int{2,7,9,3,1})) // 12
    fmt.Println(massage1([]int{2,1,4,5,3,1,1,3})) // 12
    fmt.Println(massage1([]int{1,2,3,4,5,6,7,8,9})) // 25
    fmt.Println(massage1([]int{9,8,7,6,5,4,3,2,1})) // 25
}