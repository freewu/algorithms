package main

// 1552. Magnetic Force Between Two Balls
// In the universe Earth C-137, 
// Rick discovered a special form of magnetic force between two balls if they are put in his new invented basket. 
// Rick has n empty baskets, the ith basket is at position[i], 
// Morty has m balls and needs to distribute the balls into the baskets such that the minimum magnetic force between any two balls is maximum.

// Rick stated that magnetic force between two different balls at positions x and y is |x - y|.
// Given the integer array position and the integer m. Return the required force.

// Example 1:
// Input: position = [1,2,3,4,7], m = 3
// Output: 3
// Explanation: Distributing the 3 balls into baskets 1, 4 and 7 will make the magnetic force between ball pairs [3, 3, 6]. The minimum magnetic force is 3. We cannot achieve a larger minimum magnetic force than 3.

// Example 2:
// Input: position = [5,4,3,2,1,1000000000], m = 2
// Output: 999999999
// Explanation: We can use baskets 1 and 1000000000.

// Constraints:
//     n == position.length
//     2 <= n <= 10^5
//     1 <= position[i] <= 10^9
//     All integers in position are distinct.
//     2 <= m <= position.length

import "fmt"
import "sort"

func maxDistance(position []int, m int) int {
    sort.Ints(position)
    low, high := 1, position[len(position)-1]
    canPlaceBalls := func(position []int, m, dist int) bool {
        n, count, lastBall := len(position), 1, position[0]
        for i := 1; i < n; i++ {
            if position[i] - lastBall >= dist {
                count += 1
                lastBall = position[i]
                if count == m { // 能放置 m 个球
                    return true
                }
            }
        }
        return false
    }
    for low <= high {
        mid := low + (high - low) / 2
        if canPlaceBalls(position, m, mid) {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return high
}

func maxDistance1(position []int, m int) int {
    sort.Ints(position)
    check := func (position []int, m int, mid int) bool {
        count, pre := 1, position[0]
        for _, v := range position[1:] {
            if v - pre >= mid {
                count++
                pre = v
            }
        }
        return count >= m
    }
    left, right := 0, position[len(position)-1] - position[0] + 1
    for left + 1 < right {
        mid := left + ((right - left) >> 1)
        if check(position, m, mid) {
            left = mid
        } else {
            right = mid
        }
    }
    return left
}

func maxDistance2(position []int, m int) int {
    sort.Ints(position)
    n := len(position)
    return sort.Search((position[n - 1] - position[0]) / (m - 1), func (target int) bool {
        target++
        count, pre := 1, position[0]
        for _, v := range position[1:] {
            if v >= pre + target {
                count++
                pre = v
            }
        }
        return count < m
    })
}

func main() {
    // Example 1:
    // Input: position = [1,2,3,4,7], m = 3
    // Output: 3
    // Explanation: Distributing the 3 balls into baskets 1, 4 and 7 will make the magnetic force between ball pairs [3, 3, 6]. The minimum magnetic force is 3. We cannot achieve a larger minimum magnetic force than 3.
    fmt.Println(maxDistance([]int{1,2,3,4,7}, 3)) // 3
    // Example 2:
    // Input: position = [5,4,3,2,1,1000000000], m = 2
    // Output: 999999999
    // Explanation: We can use baskets 1 and 1000000000.
    fmt.Println(maxDistance([]int{5,4,3,2,1,1000000000}, 2)) // 999999999

    fmt.Println(maxDistance([]int{1,2,3,4,5,6,7,8,9}, 3)) // 4
    fmt.Println(maxDistance([]int{9,8,7,6,5,4,3,2,1}, 3)) // 4

    fmt.Println(maxDistance1([]int{1,2,3,4,7}, 3)) // 3
    fmt.Println(maxDistance1([]int{5,4,3,2,1,1000000000}, 2)) // 999999999
    fmt.Println(maxDistance1([]int{1,2,3,4,5,6,7,8,9}, 3)) // 4
    fmt.Println(maxDistance1([]int{9,8,7,6,5,4,3,2,1}, 3)) // 4

    fmt.Println(maxDistance2([]int{1,2,3,4,7}, 3)) // 3
    fmt.Println(maxDistance2([]int{5,4,3,2,1,1000000000}, 2)) // 999999999
    fmt.Println(maxDistance2([]int{1,2,3,4,5,6,7,8,9}, 3)) // 4
    fmt.Println(maxDistance2([]int{9,8,7,6,5,4,3,2,1}, 3)) // 4
}