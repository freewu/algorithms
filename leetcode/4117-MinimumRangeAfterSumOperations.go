package main

// 4117. Minimum Range After Sum Operations
// You are given three integers a, b, and c.

// You may perform the following operation any number of times (including zero):
//     1. Choose one of the three integers and replace it with the sum of the other two integers. 
//        The other two integers remain unchanged.

// The range of three integers is defined as the difference between the maximum value and the minimum value.

// Return the minimum range that can be obtained after performing the operation any number of times.

// Example 1:
// Input: a = 9, b = 2, c = 5
// Output: 5
// Explanation:
// One sequence of operations which results in the minimum range is as follows:
//     Set a to b + c = 2 + 5 = 7. The three integers are now a = 7, b = 2, c = 5.
//     The range of the three integers is max(7, 2, 5) - min(7, 2, 5) = 7 - 2 = 5.

// Example 2:
// Input: a = 5, b = 7, c = 6
// Output: 2
// Explanation:
// Without doing any operations, the range of the three integers is max(5, 7, 6) - min(5, 7, 6) = 7 - 5 = 2.

// Example 3:
// Input: a = 3, b = 3, c = 3
// Output: 0
// Explanation:
// Without doing any operations, the range of the three integers is max(3, 3, 3) - min(3, 3, 3) = 3 - 3 = 0.

// Constraints:
//     1 <= a, b, c <= 100

import "fmt"
import "sort"

func minRange(a int, b int, c int) int {
    // 把 a 变成 b + c，三个数从小到大分别为 b,c,b + c，极差为 c。
    // 把 b 变成 a + c，三个数从小到大分别为 a,c,a + c，极差为 c。
    // 把 c 变成 a + b，三个数从小到大分别为 a,b,a + b，极差为 b。
    nums := []int{a, b, c}
    sort.Ints(nums)
    return min(nums[2] - nums[0], nums[1])
}

func minRange1(a int, b int, c int) int {
    if a > b {
        a, b = b, a
    }
    if a > c {
        a, c = c, a
    }
    if b > c {
        b, c = c, b
    }
    return min(c - a, b)
}

func main() {
    // Example 1:
    // Input: a = 9, b = 2, c = 5
    // Output: 5
    // Explanation:
    // One sequence of operations which results in the minimum range is as follows:
    //     Set a to b + c = 2 + 5 = 7. The three integers are now a = 7, b = 2, c = 5.
    //     The range of the three integers is max(7, 2, 5) - min(7, 2, 5) = 7 - 2 = 5.
    fmt.Println(minRange(9, 2, 5)) // 5 
    // Example 2:
    // Input: a = 5, b = 7, c = 6
    // Output: 2
    // Explanation:
    // Without doing any operations, the range of the three integers is max(5, 7, 6) - min(5, 7, 6) = 7 - 5 = 2.
    fmt.Println(minRange(5, 7, 6)) // 2 
    // Example 3:
    // Input: a = 3, b = 3, c = 3
    // Output: 0
    // Explanation:
    // Without doing any operations, the range of the three integers is max(3, 3, 3) - min(3, 3, 3) = 3 - 3 = 0.
    fmt.Println(minRange(3, 3, 3)) // 0 

    fmt.Println(minRange(1, 1, 1)) // 0
    fmt.Println(minRange(100, 1, 1)) // 1 
    fmt.Println(minRange(1, 100, 1)) // 1
    fmt.Println(minRange(1, 1, 100)) // 1 
    fmt.Println(minRange(100, 100, 1)) // 99 
    fmt.Println(minRange(1, 100, 100)) // 99
    fmt.Println(minRange(100, 1, 100)) // 99  
    fmt.Println(minRange(100, 100, 100)) // 0 

    fmt.Println(minRange1(9, 2, 5)) // 5 
    fmt.Println(minRange1(5, 7, 6)) // 2 
    fmt.Println(minRange1(3, 3, 3)) // 0 
    fmt.Println(minRange1(1, 1, 1)) // 0
    fmt.Println(minRange1(100, 1, 1)) // 1 
    fmt.Println(minRange1(1, 100, 1)) // 1
    fmt.Println(minRange1(1, 1, 100)) // 1 
    fmt.Println(minRange1(100, 100, 1)) // 99 
    fmt.Println(minRange1(1, 100, 100)) // 99
    fmt.Println(minRange1(100, 1, 100)) // 99  
    fmt.Println(minRange1(100, 100, 100)) // 0 
}