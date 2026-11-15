package main

// 4113. Game of Divisible Sums
// Alice and Bob are playing a game with an integer array nums. Alice goes first.

// On each turn, the current player takes and removes either the leftmost element or, if at least two elements remain, the two leftmost elements.

// The sum of the elements taken in that turn determines who plays next:
//     1. If the sum is divisible by 3, the other player plays next.
//     2. Otherwise, the same player plays again.

// The score of a player is the sum of the values of all elements that player takes.

// Both players play optimally, with each player trying to maximize their own score.

// The game ends when all elements have been taken.

// Return Alice's score minus Bob's score.

// Example 1:
// Input: nums = [3,6,4]
// Output: 5
// Explanation:
// One sequence of optimal turns is as follows:
// 1. Alice takes the two leftmost elements. 
//    Their sum is 3 + 6 = 9, which is divisible by 3, so Bob plays next. 
//    Alice's score is now 9, and the remaining array is [4].
// 2. Bob takes the only remaining element. 
//    His score is 4, and the game ends.
// Alice's score minus Bob's score is 9 - 4 = 5, so the answer is 5.

// Example 2:
// Input: nums = [6,1,2,1,2]
// Output: 12
// Explanation:
// One sequence of optimal turns is as follows:
// 1. Alice takes the two leftmost elements. 
//    Their sum is 6 + 1 = 7, which is not divisible by 3, so Alice plays again. 
//    Her score is now 7, and the remaining array is [2, 1, 2].
// 2. Alice takes the leftmost element. 
//    Its value is 2, which is not divisible by 3, so Alice plays again. 
//    Her score is now 7 + 2 = 9, and the remaining array is [1, 2].
// 3. Alice takes both remaining elements. 
//    Her score is now 9 + 1 + 2 = 12, and the game ends.
// Bob takes no elements, so his score is 0. 
// Alice's score minus Bob's score is 12 - 0 = 12, so the answer is 12.

// Example 3:
// Input: nums = [3,3,4,5]
// Output: -3
// Explanation:
// One sequence of optimal turns is as follows:
// 1. Alice takes the two leftmost elements. 
//    Their sum is 3 + 3 = 6, which is divisible by 3, so Bob plays next. 
//    Alice's score is now 6, and the remaining array is [4, 5].
// 2. Bob takes both remaining elements. 
//    His score is 4 + 5 = 9, and the game ends.
// Alice's score minus Bob's score is 6 - 9 = -3, so the answer is -3.

// Constraints:
//     1 <= nums.length <= 10^{5}
//     1 <= nums[i] <= 10^{9}

import "fmt"

func maxDifference(nums []int) int64 {
    n := len(nums)
    dp := make([]int64, n+1)
    for i := n - 1; i >= 0; i-- {
        res1, s1 := int64(0), int64(nums[i])
        if s1 % 3 == 0 {
            res1 = s1 - dp[i+1]
        } else {
            res1 = s1 + dp[i+1]
        }
        if i + 1 < n {
            res2, s2 := int64(0), int64(nums[i]) + int64(nums[i+1])
            if s2 % 3 == 0 {
                res2 = s2 - dp[i+2]
            } else {
                res2 = s2 + dp[i+2]
            }
            if res1 > res2 {
                dp[i] = res1
            } else {
                dp[i] = res2
            }
        } else {
            dp[i] = res1
        }
    }
    return dp[0]
}

func maxDifference1(nums []int) int64 {
    score := func(sum, nxt int64) int64 {
        if sum % 3 == 0 {
            return sum - nxt
        }
        return sum + nxt
    }
    nxt1, nxt2 := int64(0), int64(0)
    for i := len(nums) - 1 ; i >= 0 ; i -- {
        sum := int64(nums[i])
        best := score(sum , nxt1)
        if i + 1 < len(nums) {
            curr := score(sum + int64(nums[i+1]), nxt2)
            if curr  > best {
                best =  curr
            }
        }
        nxt1, nxt2 = best, nxt1
    }
    return nxt1 
}

func maxDifference2(nums []int) int64 {
    next1, next2, n := int64(0), int64(0), len(nums)
    for i := n - 1; i >= 0; i-- {
        best, s1 := int64(0), int64(nums[i])
        if s1 % 3 == 0 {
            best = s1 - next1
        } else {
            best = s1 + next1
        }
        if i+1 < n {
            s2 := int64(nums[i]) + int64(nums[i+1])
            var cand int64
            if s2%3 == 0 {
                cand = s2 - next2
            } else {
                cand = s2 + next2
            }
            if cand > best {
                best = cand
            }
        }
        next2 = next1
        next1 = best
    }
    return next1
}

func main() {
    // Example 1:
    // Input: nums = [3,6,4]
    // Output: 5
    // Explanation:
    // One sequence of optimal turns is as follows:
    // 1. Alice takes the two leftmost elements. 
    //    Their sum is 3 + 6 = 9, which is divisible by 3, so Bob plays next. 
    //    Alice's score is now 9, and the remaining array is [4].
    // 2. Bob takes the only remaining element. 
    //    His score is 4, and the game ends.
    // Alice's score minus Bob's score is 9 - 4 = 5, so the answer is 5.
    fmt.Println(maxDifference([]int{3,6,4})) // 5
    // Example 2:
    // Input: nums = [6,1,2,1,2]
    // Output: 12
    // Explanation:
    // One sequence of optimal turns is as follows:
    // 1. Alice takes the two leftmost elements. 
    //    Their sum is 6 + 1 = 7, which is not divisible by 3, so Alice plays again. 
    //    Her score is now 7, and the remaining array is [2, 1, 2].
    // 2. Alice takes the leftmost element. 
    //    Its value is 2, which is not divisible by 3, so Alice plays again. 
    //    Her score is now 7 + 2 = 9, and the remaining array is [1, 2].
    // 3. Alice takes both remaining elements. 
    //    Her score is now 9 + 1 + 2 = 12, and the game ends.
    // Bob takes no elements, so his score is 0. 
    // Alice's score minus Bob's score is 12 - 0 = 12, so the answer is 12.
    fmt.Println(maxDifference([]int{6,1,2,1,2})) // 12
    // Example 3:
    // Input: nums = [3,3,4,5]
    // Output: -3
    // Explanation:
    // One sequence of optimal turns is as follows:
    // 1. Alice takes the two leftmost elements. 
    //    Their sum is 3 + 3 = 6, which is divisible by 3, so Bob plays next. 
    //    Alice's score is now 6, and the remaining array is [4, 5].
    // 2. Bob takes both remaining elements. 
    //    His score is 4 + 5 = 9, and the game ends.
    // Alice's score minus Bob's score is 6 - 9 = -3, so the answer is -3.
    fmt.Println(maxDifference([]int{3,3,4,5})) // -3

    fmt.Println(maxDifference([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxDifference([]int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(maxDifference1([]int{3,6,4})) // 5
    fmt.Println(maxDifference1([]int{6,1,2,1,2})) // 12
    fmt.Println(maxDifference1([]int{3,3,4,5})) // -3
    fmt.Println(maxDifference1([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxDifference1([]int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(maxDifference2([]int{3,6,4})) // 5
    fmt.Println(maxDifference2([]int{6,1,2,1,2})) // 12
    fmt.Println(maxDifference2([]int{3,3,4,5})) // -3
    fmt.Println(maxDifference2([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxDifference2([]int{9,8,7,6,5,4,3,2,1})) // 45
}