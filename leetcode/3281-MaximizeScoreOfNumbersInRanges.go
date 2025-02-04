package main

// 3281. Maximize Score of Numbers in Ranges
// You are given an array of integers start and an integer d, representing n intervals [start[i], start[i] + d].

// You are asked to choose n integers where the ith integer must belong to the ith interval. 
// The score of the chosen integers is defined as the minimum absolute difference between any two integers that have been chosen.

// Return the maximum possible score of the chosen integers.

// Example 1:
// Input: start = [6,0,3], d = 2
// Output: 4
// Explanation:
// The maximum possible score can be obtained by choosing integers: 8, 0, and 4. 
// The score of these chosen integers is min(|8 - 0|, |8 - 4|, |0 - 4|) which equals 4.

// Example 2:
// Input: start = [2,6,13,13], d = 5
// Output: 5
// Explanation:
// The maximum possible score can be obtained by choosing integers: 2, 7, 13, and 18.
// The score of these chosen integers is min(|2 - 7|, |2 - 13|, |2 - 18|, |7 - 13|, |7 - 18|, |13 - 18|) which equals 5.

// Constraints:
//     2 <= start.length <= 10^5
//     0 <= start[i] <= 10^9
//     0 <= d <= 10^9

import "fmt"
import "sort"

func maxPossibleScore(start []int, d int) int {
    sort.Ints(start)
    res, left, right := 0, 0, start[len(start) - 1] - start[0] + d
    for left <= right {
        match, last, mid := true, start[0], (left + right) / 2
        for i := 1; i < len(start); i++ {
            if last + mid > start[i] + d {
                match = false
                break
            }
            if last + mid < start[i] {
                last = start[i]
            } else {
                last = last + mid
            }
        }
        if match {
            res, left = mid, mid + 1
        } else {
            right = mid - 1
        }
    }
    return res
}

func maxPossibleScore1(start []int, d int) int {
    sort.Ints(start)
    n := len(start)
    left, right := 0, (start[n-1] + d - start[0]) / (n - 1) + 1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    check := func(score int) bool {
        x := -1 << 31
        for _, v := range start {
            x = max(x + score, v)
            if x > v + d {
                return false
            }
        }
        return true
    }
    for left + 1 < right {
        mid := left + (right - left) / 2
        if check(mid) {
            left = mid
        } else {
            right = mid
        }
    }
    return left
}

func maxPossibleScore2(start []int, d int) int {
    sort.Ints(start)
    res, left, right := 0, 0, start[1] + d - start[0]
    max := func(x, y int) int { if x > y { return x; }; return y; }
    check := func(score int) bool {
        before := start[0]
        for i := 1; i < len(start); i++ {
            if before + score <= start[i] + d {
                before = max(start[i], before + score)
            } else {
                return false
            }
        }
        return true
    }
    for left <= right {
        mid := (left + right) / 2
        if check(mid) {
            left, res = mid + 1, max(res, mid)
        } else {
            right = mid - 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: start = [6,0,3], d = 2
    // Output: 4
    // Explanation:
    // The maximum possible score can be obtained by choosing integers: 8, 0, and 4. 
    // The score of these chosen integers is min(|8 - 0|, |8 - 4|, |0 - 4|) which equals 4.
    fmt.Println(maxPossibleScore([]int{6,0,3}, 2)) // 4
    // Example 2:
    // Input: start = [2,6,13,13], d = 5
    // Output: 5
    // Explanation:
    // The maximum possible score can be obtained by choosing integers: 2, 7, 13, and 18.
    // The score of these chosen integers is min(|2 - 7|, |2 - 13|, |2 - 18|, |7 - 13|, |7 - 18|, |13 - 18|) which equals 5.
    fmt.Println(maxPossibleScore([]int{2,6,13,13}, 5)) // 5

    fmt.Println(maxPossibleScore([]int{1,2,3,4,5,6,7,8,9}, 5)) // 1
    fmt.Println(maxPossibleScore([]int{9,8,7,6,5,4,3,2,1}, 5)) // 1

    fmt.Println(maxPossibleScore1([]int{6,0,3}, 2)) // 4
    fmt.Println(maxPossibleScore1([]int{2,6,13,13}, 5)) // 5
    fmt.Println(maxPossibleScore1([]int{1,2,3,4,5,6,7,8,9}, 5)) // 1
    fmt.Println(maxPossibleScore1([]int{9,8,7,6,5,4,3,2,1}, 5)) // 1

    fmt.Println(maxPossibleScore2([]int{6,0,3}, 2)) // 4
    fmt.Println(maxPossibleScore2([]int{2,6,13,13}, 5)) // 5
    fmt.Println(maxPossibleScore2([]int{1,2,3,4,5,6,7,8,9}, 5)) // 1
    fmt.Println(maxPossibleScore2([]int{9,8,7,6,5,4,3,2,1}, 5)) // 1
}