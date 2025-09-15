package main

// 3679. Minimum Discards to Balance Inventory
// You are given two integers w and m, and an integer array arrivals, 
// where arrivals[i] is the type of item arriving on day i (days are 1-indexed).

// Items are managed according to the following rules:
//     1. Each arrival may be kept or discarded; an item may only be discarded on its arrival day.
//     2. For each day i, consider the window of days [max(1, i - w + 1), i] (the w most recent days up to day i):
//         2.1 For any such window, each item type may appear at most m times among kept arrivals whose arrival day lies in that window.
//         2.2 If keeping the arrival on day i would cause its type to appear more than m times in the window, 
//             that arrival must be discarded.

// Return the minimum number of arrivals to be discarded so that every w-day window contains at most m occurrences of each type.

// Example 1:
// Input: arrivals = [1,2,1,3,1], w = 4, m = 2
// Output: 0
// Explanation:
// On day 1, Item 1 arrives; the window contains no more than m occurrences of this type, so we keep it.
// On day 2, Item 2 arrives; the window of days 1 - 2 is fine.
// On day 3, Item 1 arrives, window [1, 2, 1] has item 1 twice, within limit.
// On day 4, Item 3 arrives, window [1, 2, 1, 3] has item 1 twice, allowed.
// On day 5, Item 1 arrives, window [2, 1, 3, 1] has item 1 twice, still valid.
// There are no discarded items, so return 0.

// Example 2:
// Input: arrivals = [1,2,3,3,3,4], w = 3, m = 2
// Output: 1
// Explanation:
// On day 1, Item 1 arrives. We keep it.
// On day 2, Item 2 arrives, window [1, 2] is fine.
// On day 3, Item 3 arrives, window [1, 2, 3] has item 3 once.
// On day 4, Item 3 arrives, window [2, 3, 3] has item 3 twice, allowed.
// On day 5, Item 3 arrives, window [3, 3, 3] has item 3 three times, exceeds limit, so the arrival must be discarded.
// On day 6, Item 4 arrives, window [3, 4] is fine.
// Item 3 on day 5 is discarded, and this is the minimum number of arrivals to discard, so return 1.

// Constraints:
//     1 <= arrivals.length <= 10^5
//     1 <= arrivals[i] <= 10^5
//     1 <= w <= arrivals.length
//     1 <= m <= w

import "fmt"

func minArrivalsToDiscard(arrivals []int, w int, m int) int {
    res, count := 0, make(map[int]int)
    for i, v := range arrivals {
        // v 进入窗口
        if count[v] == m { // v 的个数已达上限
            // 注意 v 在未来要离开窗口，但由于已经丢弃，不能计入
            // 这里直接置为 0，未来离开窗口就是 count[0]--，不影响答案
            arrivals[i] = 0
            res++
        } else {
            count[v]++
        }
        // 左端点元素离开窗口，为下一个循环做准备
        left := i + 1 - w
        if left >= 0 {
            count[arrivals[left]]--
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: arrivals = [1,2,1,3,1], w = 4, m = 2
    // Output: 0
    // Explanation:
    // On day 1, Item 1 arrives; the window contains no more than m occurrences of this type, so we keep it.
    // On day 2, Item 2 arrives; the window of days 1 - 2 is fine.
    // On day 3, Item 1 arrives, window [1, 2, 1] has item 1 twice, within limit.
    // On day 4, Item 3 arrives, window [1, 2, 1, 3] has item 1 twice, allowed.
    // On day 5, Item 1 arrives, window [2, 1, 3, 1] has item 1 twice, still valid.
    // There are no discarded items, so return 0.
    fmt.Println(minArrivalsToDiscard([]int{1,2,1,3,1}, 4, 2)) // 0
    // Example 2:
    // Input: arrivals = [1,2,3,3,3,4], w = 3, m = 2
    // Output: 1
    // Explanation:
    // On day 1, Item 1 arrives. We keep it.
    // On day 2, Item 2 arrives, window [1, 2] is fine.
    // On day 3, Item 3 arrives, window [1, 2, 3] has item 3 once.
    // On day 4, Item 3 arrives, window [2, 3, 3] has item 3 twice, allowed.
    // On day 5, Item 3 arrives, window [3, 3, 3] has item 3 three times, exceeds limit, so the arrival must be discarded.
    // On day 6, Item 4 arrives, window [3, 4] is fine.
    // Item 3 on day 5 is discarded, and this is the minimum number of arrivals to discard, so return 1.
    fmt.Println(minArrivalsToDiscard([]int{1,2,3,3,3,4}, 3, 2)) // 1

    fmt.Println(minArrivalsToDiscard([]int{1,2,3,4,5,6,7,8,9}, 3, 2)) // 0
    fmt.Println(minArrivalsToDiscard([]int{9,8,7,6,5,4,3,2,1}, 3, 2)) // 0
}