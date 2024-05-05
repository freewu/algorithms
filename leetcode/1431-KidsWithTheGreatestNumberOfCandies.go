package main

// 1431. Kids With the Greatest Number of Candies
// There are n kids with candies.
// You are given an integer array candies, 
// where each candies[i] represents the number of candies the ith kid has, 
// and an integer extraCandies, denoting the number of extra candies that you have.

// Return a boolean array result of length n, where result[i] is true if, 
// after giving the ith kid all the extraCandies, 
// they will have the greatest number of candies among all the kids, or false otherwise.

// Note that multiple kids can have the greatest number of candies.

// Example 1:
// Input: candies = [2,3,5,1,3], extraCandies = 3
// Output: [true,true,true,false,true] 
// Explanation: If you give all extraCandies to:
// - Kid 1, they will have 2 + 3 = 5 candies, which is the greatest among the kids.
// - Kid 2, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
// - Kid 3, they will have 5 + 3 = 8 candies, which is the greatest among the kids.
// - Kid 4, they will have 1 + 3 = 4 candies, which is not the greatest among the kids.
// - Kid 5, they will have 3 + 3 = 6 candies, which is the greatest among the kids.

// Example 2:
// Input: candies = [4,2,1,1,2], extraCandies = 1
// Output: [true,false,false,false,false] 
// Explanation: There is only 1 extra candy.
// Kid 1 will always have the greatest number of candies, even if a different kid is given the extra candy.

// Example 3:
// Input: candies = [12,1,12], extraCandies = 10
// Output: [true,false,true]

// Constraints:
//     n == candies.length
//     2 <= n <= 100
//     1 <= candies[i] <= 100
//     1 <= extraCandies <= 50

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
    res, mx := make([]bool,len(candies)), 0
    // 获得最多糖果的数量
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _,v := range candies {
        mx = max(mx, v)
    }
    for i,v := range candies {
        if v + extraCandies >= mx { // 将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有 最多 的糖果。注意，允许有多个孩子同时拥有 最多 的糖果数目。
            res[i] = true
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: candies = [2,3,5,1,3], extraCandies = 3
    // Output: [true,true,true,false,true] 
    // Explanation: If you give all extraCandies to:
    // - Kid 1, they will have 2 + 3 = 5 candies, which is the greatest among the kids.
    // - Kid 2, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
    // - Kid 3, they will have 5 + 3 = 8 candies, which is the greatest among the kids.
    // - Kid 4, they will have 1 + 3 = 4 candies, which is not the greatest among the kids.
    // - Kid 5, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
    fmt.Println(kidsWithCandies([]int{2,3,5,1,3}, 3)) //  [true,true,true,false,true] 
    // Example 2:
    // Input: candies = [4,2,1,1,2], extraCandies = 1
    // Output: [true,false,false,false,false] 
    // Explanation: There is only 1 extra candy.
    // Kid 1 will always have the greatest number of candies, even if a different kid is given the extra candy.
    fmt.Println(kidsWithCandies([]int{4,2,1,1,2}, 1)) // [true,false,false,false,false] 
    // Example 3:
    // Input: candies = [12,1,12], extraCandies = 10
    // Output: [true,false,true]
    fmt.Println(kidsWithCandies([]int{12,1,12}, 10)) // [true,false,true]
}