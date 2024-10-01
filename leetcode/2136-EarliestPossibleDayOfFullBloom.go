package main

// 2136. Earliest Possible Day of Full Bloom
// You have n flower seeds. Every seed must be planted first before it can begin to grow, then bloom. 
// Planting a seed takes time and so does the growth of a seed. 
// You are given two 0-indexed integer arrays plantTime and growTime, of length n each:
//     1. plantTime[i] is the number of full days it takes you to plant the ith seed. 
//        Every day, you can work on planting exactly one seed. 
//        You do not have to work on planting the same seed on consecutive days, 
//        but the planting of a seed is not complete until you have worked plantTime[i] days on planting it in total.
//     2. growTime[i] is the number of full days it takes the ith seed to grow after being completely planted. 
//        After the last day of its growth, the flower blooms and stays bloomed forever.

// From the beginning of day 0, you can plant the seeds in any order.

// Return the earliest possible day where all seeds are blooming.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/12/21/1.png" />
// Input: plantTime = [1,4,3], growTime = [2,3,1]
// Output: 9
// Explanation: The grayed out pots represent planting days, colored pots represent growing days, and the flower represents the day it blooms.
// One optimal way is:
// On day 0, plant the 0th seed. The seed grows for 2 full days and blooms on day 3.
// On days 1, 2, 3, and 4, plant the 1st seed. The seed grows for 3 full days and blooms on day 8.
// On days 5, 6, and 7, plant the 2nd seed. The seed grows for 1 full day and blooms on day 9.
// Thus, on day 9, all the seeds are blooming.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/12/21/2.png" />
// Input: plantTime = [1,2,3,2], growTime = [2,1,2,1]
// Output: 9
// Explanation: The grayed out pots represent planting days, colored pots represent growing days, and the flower represents the day it blooms.
// One optimal way is:
// On day 1, plant the 0th seed. The seed grows for 2 full days and blooms on day 4.
// On days 0 and 3, plant the 1st seed. The seed grows for 1 full day and blooms on day 5.
// On days 2, 4, and 5, plant the 2nd seed. The seed grows for 2 full days and blooms on day 8.
// On days 6 and 7, plant the 3rd seed. The seed grows for 1 full day and blooms on day 9.
// Thus, on day 9, all the seeds are blooming.

// Example 3:
// Input: plantTime = [1], growTime = [1]
// Output: 2
// Explanation: On day 0, plant the 0th seed. The seed grows for 1 full day and blooms on day 2.
// Thus, on day 2, all the seeds are blooming.

// Constraints:
//     n == plantTime.length == growTime.length
//     1 <= n <= 10^5
//     1 <= plantTime[i], growTime[i] <= 10^4

import "fmt"
import "sort"

func earliestFullBloom(plantTime []int, growTime []int) int {
    type Plant struct {
        plantTime int // 种植时间
        growTime int  // 生长时间
    }
    plants := make([]Plant, len(plantTime))
    for i, v := range plantTime {
        plants[i] = Plant{v, growTime[i]}
    }
    sort.Slice(plants, func (i int,j int) bool { 
        return plants[i].growTime > plants[j].growTime // 生长时间最越大越靠前
    })
    res, day := -1, -1
    for i := 0; i < len(plants); i++ {
        day += plants[i].plantTime // 种植时间需要累加
        bloomDay := day + plants[i].growTime + 1 // 加上生长时间(可以并行)
        if bloomDay > res {
            res = bloomDay
        }
    }
    return res
}

type IntsPair [2][]int

func (p IntsPair) Len() int           { return len(p[0]) }
func (p IntsPair) Less(i, j int) bool { return p[0][i] > p[0][j] }
func (p IntsPair) Swap(i, j int) {
    p[0][i], p[0][j] = p[0][j], p[0][i]
    p[1][i], p[1][j] = p[1][j], p[1][i]
}

func earliestFullBloom1(plantTime []int, growTime []int) int {
    pair := IntsPair{growTime, plantTime}
    sort.Sort(pair)
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for plantTimeSum, i := 0, 0; i < len(plantTime); i++ {
        plantTimeSum += plantTime[i]
        res = max(res, plantTimeSum+growTime[i])
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/12/21/1.png" />
    // Input: plantTime = [1,4,3], growTime = [2,3,1]
    // Output: 9
    // Explanation: The grayed out pots represent planting days, colored pots represent growing days, and the flower represents the day it blooms.
    // One optimal way is:
    // On day 0, plant the 0th seed. The seed grows for 2 full days and blooms on day 3.
    // On days 1, 2, 3, and 4, plant the 1st seed. The seed grows for 3 full days and blooms on day 8.
    // On days 5, 6, and 7, plant the 2nd seed. The seed grows for 1 full day and blooms on day 9.
    // Thus, on day 9, all the seeds are blooming.
    fmt.Println(earliestFullBloom([]int{1,4,3}, []int{2,3,1})) // 9
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/12/21/2.png" />
    // Input: plantTime = [1,2,3,2], growTime = [2,1,2,1]
    // Output: 9
    // Explanation: The grayed out pots represent planting days, colored pots represent growing days, and the flower represents the day it blooms.
    // One optimal way is:
    // On day 1, plant the 0th seed. The seed grows for 2 full days and blooms on day 4.
    // On days 0 and 3, plant the 1st seed. The seed grows for 1 full day and blooms on day 5.
    // On days 2, 4, and 5, plant the 2nd seed. The seed grows for 2 full days and blooms on day 8.
    // On days 6 and 7, plant the 3rd seed. The seed grows for 1 full day and blooms on day 9.
    // Thus, on day 9, all the seeds are blooming.
    fmt.Println(earliestFullBloom([]int{1,2,3,2}, []int{2,1,2,1})) // 9
    // Example 3:
    // Input: plantTime = [1], growTime = [1]
    // Output: 2
    // Explanation: On day 0, plant the 0th seed. The seed grows for 1 full day and blooms on day 2.
    // Thus, on day 2, all the seeds are blooming.
    fmt.Println(earliestFullBloom([]int{1}, []int{1})) // 2

    fmt.Println(earliestFullBloom1([]int{1,4,3}, []int{2,3,1})) // 9
    fmt.Println(earliestFullBloom1([]int{1,2,3,2}, []int{2,1,2,1})) // 9
    fmt.Println(earliestFullBloom1([]int{1}, []int{1})) // 2
}