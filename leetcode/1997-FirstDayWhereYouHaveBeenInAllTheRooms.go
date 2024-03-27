package main

// 1997. First Day Where You Have Been in All the Rooms
// There are n rooms you need to visit, labeled from 0 to n - 1. 
// Each day is labeled, starting from 0. You will go in and visit one room a day.

// Initially on day 0, you visit room 0. 
// The order you visit the rooms for the coming days is determined by the following rules and a given 0-indexed array nextVisit of length n:
//     Assuming that on a day, you visit room i,
//     if you have been in room i an odd number of times (including the current visit), on the next day you will visit a room with a lower or equal room number specified by nextVisit[i] where 0 <= nextVisit[i] <= i;
//     if you have been in room i an even number of times (including the current visit), on the next day you will visit room (i + 1) mod n.

// Return the label of the first day where you have been in all the rooms. 
// It can be shown that such a day exists. Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: nextVisit = [0,0]
// Output: 2
// Explanation:
// - On day 0, you visit room 0. The total times you have been in room 0 is 1, which is odd.
//   On the next day you will visit room nextVisit[0] = 0
// - On day 1, you visit room 0, The total times you have been in room 0 is 2, which is even.
//   On the next day you will visit room (0 + 1) mod 2 = 1
// - On day 2, you visit room 1. This is the first day where you have been in all the rooms.

// Example 2:
// Input: nextVisit = [0,0,2]
// Output: 6
// Explanation:
// Your room visiting order for each day is: [0,0,1,0,0,1,2,...].
// Day 6 is the first day where you have been in all the rooms.

// Example 3:
// Input: nextVisit = [0,1,2,0]
// Output: 6
// Explanation:
// Your room visiting order for each day is: [0,0,1,1,2,2,3,...].
// Day 6 is the first day where you have been in all the rooms.
 
// Constraints:
//     n == nextVisit.length
//     2 <= n <= 10^5
//     0 <= nextVisit[i] <= i

import "fmt"

// dp
func firstDayBeenInAllRooms(nextVisit []int) int {
    // reached i+1, reach all elem from 0 to i even times
    dp := make([]int, len(nextVisit))
    dp[1] = 2
    day := 2
    for i := 2; i < len(nextVisit); i++ {
        // we are in room[i-1] and one day passed
        day++
        curRoom := nextVisit[i-1]
        // back to room[i-1], cost dp[i-1] - dp[curRoom] day, next day, we will arrival root[i] first time
        // dp[i] may be less than dp[j] while j < i, so we add another 1000000007
        day = (day + dp[i-1] - dp[curRoom] + 1 + 1000000007) % 1000000007
        dp[i] = day
    }
    return dp[len(nextVisit) - 1]
}

func firstDayBeenInAllRooms1(nextVisit []int) int {
    n := len(nextVisit)
    // 0 - n 消耗
    dp := make([]int, n)
    mod := 1000000007
    for i, j := range nextVisit[:n-1] {
        // i 是当前坐标，j 是回退坐标
        // 能跑到 i + 1，第一次是 dp[i]跑过来的，第二次是 j -> i 跑过来的
        // 第二次跑过来的消耗
        v := (dp[i] - dp[j] + mod) % mod
        dp[i+1] = (dp[i] + v + 2) % mod
    }
    return dp[n-1]
}

func main() {
    // - On day 0, you visit room 0. The total times you have been in room 0 is 1, which is odd.
    //   On the next day you will visit room nextVisit[0] = 0
    // - On day 1, you visit room 0, The total times you have been in room 0 is 2, which is even.
    //   On the next day you will visit room (0 + 1) mod 2 = 1
    // - On day 2, you visit room 1. This is the first day where you have been in all the rooms.
    fmt.Println(firstDayBeenInAllRooms([]int{0,0})) // 2

    // Your room visiting order for each day is: [0,0,1,0,0,1,2,...].
    // Day 6 is the first day where you have been in all the rooms.
    fmt.Println(firstDayBeenInAllRooms([]int{0,0,2})) // 6

    // Your room visiting order for each day is: [0,0,1,1,2,2,3,...].
    // Day 6 is the first day where you have been in all the rooms.
    fmt.Println(firstDayBeenInAllRooms([]int{0,1,2,0})) // 6

    fmt.Println(firstDayBeenInAllRooms1([]int{0,0})) // 2
    fmt.Println(firstDayBeenInAllRooms1([]int{0,0,2})) // 6
    fmt.Println(firstDayBeenInAllRooms1([]int{0,1,2,0})) // 6
}