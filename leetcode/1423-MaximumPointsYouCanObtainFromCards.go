package main

// 1423. Maximum Points You Can Obtain from Cards
// There are several cards arranged in a row, and each card has an associated number of points. 
// The points are given in the integer array cardPoints.

// In one step, you can take one card from the beginning or from the end of the row. 
// You have to take exactly k cards.

// Your score is the sum of the points of the cards you have taken.

// Given the integer array cardPoints and the integer k, return the maximum score you can obtain.

// Example 1:
// Input: cardPoints = [1,2,3,4,5,6,1], k = 3
// Output: 12
// Explanation: After the first step, your score will always be 1. However, choosing the rightmost card first will maximize your total score. The optimal strategy is to take the three cards on the right, giving a final score of 1 + 6 + 5 = 12.

// Example 2:
// Input: cardPoints = [2,2,2], k = 2
// Output: 4
// Explanation: Regardless of which two cards you take, your score will always be 4.

// Example 3:
// Input: cardPoints = [9,7,7,9,7,7,9], k = 7
// Output: 55
// Explanation: You have to take all the cards. Your score is the sum of points of all cards.

// Constraints:
//     1 <= cardPoints.length <= 10^5
//     1 <= cardPoints[i] <= 10^4
//     1 <= k <= cardPoints.length

import "fmt"

func maxScore(cardPoints []int, k int) int {
    left, sum, res, n := 0, 0, 0, len(cardPoints)
    for i := 0; i < k; i++ {
        cardPoints = append(cardPoints, cardPoints[i])
    }
    for right := 0; right < len(cardPoints); right++ {
        sum += cardPoints[right]
        for right - left + 1 > k {
            sum -= cardPoints[left]
            left++
        }
        if right - left + 1 == k && left >= n - k {
            if sum > res {
                res = sum
            }
        }
    }
    return res
}

func maxScore1(cardPoints []int, k int) int {
    total := 0
    for _, v := range cardPoints {
        total += v
    }
    k = len(cardPoints) - k
    if k == 0 {
        return total
    }
    res, sum := total, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, p := range cardPoints {
        sum += p
        if i < k-1 { continue }
        res = min(res, sum)
        sum -= cardPoints[i-k+1]
    }
    return total - res
}

func maxScore2(cardPoints []int, k int) int {
    lsum, rsum, res := 0, 0, 0
    for i := 0; i < k; i++ {
        lsum = lsum + cardPoints[i]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res = max(res,lsum)
    rightIndex := len(cardPoints) - 1
    for i := k - 1; i >= 0; i-- {
        lsum, rsum = lsum - cardPoints[i], rsum + cardPoints[rightIndex]
        res = max(res,lsum+rsum)
        rightIndex--
    }
    return res
}

func main() {
    // Example 1:
    // Input: cardPoints = [1,2,3,4,5,6,1], k = 3
    // Output: 12
    // Explanation: After the first step, your score will always be 1. However, choosing the rightmost card first will maximize your total score. The optimal strategy is to take the three cards on the right, giving a final score of 1 + 6 + 5 = 12.
    fmt.Println(maxScore([]int{1,2,3,4,5,6,1}, 3)) // 12
    // Example 2:
    // Input: cardPoints = [2,2,2], k = 2
    // Output: 4
    // Explanation: Regardless of which two cards you take, your score will always be 4.
    fmt.Println(maxScore([]int{2,2,2}, 2)) // 4
    // Example 3:
    // Input: cardPoints = [9,7,7,9,7,7,9], k = 7
    // Output: 55
    // Explanation: You have to take all the cards. Your score is the sum of points of all cards.
    fmt.Println(maxScore([]int{9,7,7,9,7,7,9}, 7)) // 55

    fmt.Println(maxScore([]int{100,40,17,9,73,75}, 3)) // 248

    fmt.Println(maxScore1([]int{1,2,3,4,5,6,1}, 3)) // 12
    fmt.Println(maxScore1([]int{2,2,2}, 2)) // 4
    fmt.Println(maxScore1([]int{9,7,7,9,7,7,9}, 7)) // 55
    fmt.Println(maxScore1([]int{100,40,17,9,73,75}, 3)) // 248

    fmt.Println(maxScore2([]int{1,2,3,4,5,6,1}, 3)) // 12
    fmt.Println(maxScore2([]int{2,2,2}, 2)) // 4
    fmt.Println(maxScore2([]int{9,7,7,9,7,7,9}, 7)) // 55
    fmt.Println(maxScore2([]int{100,40,17,9,73,75}, 3)) // 248
}