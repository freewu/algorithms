package main

// 1467. Probability of a Two Boxes Having The Same Number of Distinct Balls
// Given 2n balls of k distinct colors. 
// You will be given an integer array balls of size k where balls[i] is the number of balls of color i.

// All the balls will be shuffled uniformly at random, 
// then we will distribute the first n balls to the first box and the remaining n balls to the other box (Please read the explanation of the second example carefully).

// Please note that the two boxes are considered different. 
// For example, if we have two balls of colors a and b, and two boxes [] and (), 
// then the distribution [a] (b) is considered different than the distribution [b] (a) (Please read the explanation of the first example carefully).

// Return the probability that the two boxes have the same number of distinct balls. 
// Answers within 10^-5 of the actual value will be accepted as correct.

// Example 1:
// Input: balls = [1,1]
// Output: 1.00000
// Explanation: Only 2 ways to divide the balls equally:
// - A ball of color 1 to box 1 and a ball of color 2 to box 2
// - A ball of color 2 to box 1 and a ball of color 1 to box 2
// In both ways, the number of distinct colors in each box is equal. The probability is 2/2 = 1

// Example 2:
// Input: balls = [2,1,1]
// Output: 0.66667
// Explanation: We have the set of balls [1, 1, 2, 3]
// This set of balls will be shuffled randomly and we may have one of the 12 distinct shuffles with equal probability (i.e. 1/12):
// [1,1 / 2,3], [1,1 / 3,2], [1,2 / 1,3], [1,2 / 3,1], [1,3 / 1,2], [1,3 / 2,1], [2,1 / 1,3], [2,1 / 3,1], [2,3 / 1,1], [3,1 / 1,2], [3,1 / 2,1], [3,2 / 1,1]
// After that, we add the first two balls to the first box and the second two balls to the second box.
// We can see that 8 of these 12 possible random distributions have the same number of distinct colors of balls in each box.
// Probability is 8/12 = 0.66667

// Example 3:
// Input: balls = [1,2,1,2]
// Output: 0.60000
// Explanation: The set of balls is [1, 2, 2, 3, 4, 4]. It is hard to display all the 180 possible random shuffles of this set but it is easy to check that 108 of them will have the same number of distinct colors in each box.
// Probability = 108 / 180 = 0.6

// Constraints:
//     1 <= balls.length <= 8
//     1 <= balls[i] <= 6
//     sum(balls) is even.

import "fmt"

func getProbability(balls []int) float64 {
    cache := map[int]int{}
    totalBalls := 0
    for _, ball := range balls {
        totalBalls += ball
    }
    combine := func (n, k int) int {
        if k > n-k {
            k = n - k
        }
        res := 1
        for i := 1; i <= k; i++ {
            res *= n - k + i
            res /= i
        }
        return res
    }
    var dfs func (i, c, t int) int 
    dfs = func (i, c, t int) int {
        if i == len(balls) {
            if t == 0 && c == totalBalls / 2 {
                return 1
            }
            return 0
        }
        key := i * 10000 + c * 100 + t
        if val, found := cache[key]; found {
            return val
        }
        res := dfs(i+1, c, t+1) + dfs(i+1, c+balls[i], t-1)
        for j := 1; j < balls[i]; j++ {
            res += dfs(i+1, c+j, t) * combine(balls[i], j)
        }
        cache[key] = res
        return res
    }
    return float64(dfs(0, 0, 0)) / float64(combine(totalBalls, totalBalls / 2))
}

func main() {
    // Example 1:
    // Input: balls = [1,1]
    // Output: 1.00000
    // Explanation: Only 2 ways to divide the balls equally:
    // - A ball of color 1 to box 1 and a ball of color 2 to box 2
    // - A ball of color 2 to box 1 and a ball of color 1 to box 2
    // In both ways, the number of distinct colors in each box is equal. The probability is 2/2 = 1
    fmt.Println(getProbability([]int{1,1})) // 1.00000
    // Example 2: 
    // Input: balls = [2,1,1]
    // Output: 0.66667
    // Explanation: We have the set of balls [1, 1, 2, 3]
    // This set of balls will be shuffled randomly and we may have one of the 12 distinct shuffles with equal probability (i.e. 1/12):
    // [1,1 / 2,3], [1,1 / 3,2], [1,2 / 1,3], [1,2 / 3,1], [1,3 / 1,2], [1,3 / 2,1], [2,1 / 1,3], [2,1 / 3,1], [2,3 / 1,1], [3,1 / 1,2], [3,1 / 2,1], [3,2 / 1,1]
    // After that, we add the first two balls to the first box and the second two balls to the second box.
    // We can see that 8 of these 12 possible random distributions have the same number of distinct colors of balls in each box.
    // Probability is 8/12 = 0.66667
    fmt.Println(getProbability([]int{2,1,1})) // 0.66667
    // Example 3:
    // Input: balls = [1,2,1,2]
    // Output: 0.60000
    // Explanation: The set of balls is [1, 2, 2, 3, 4, 4]. It is hard to display all the 180 possible random shuffles of this set but it is easy to check that 108 of them will have the same number of distinct colors in each box.
    // Probability = 108 / 180 = 0.6
    fmt.Println(getProbability([]int{1,2,1,2})) // 0.60000
}