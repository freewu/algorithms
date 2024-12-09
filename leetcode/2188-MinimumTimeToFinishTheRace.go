package main

// 2188. Minimum Time to Finish the Race
// You are given a 0-indexed 2D integer array tires where tires[i] = [fi, ri] indicates 
// that the ith tire can finish its xth successive lap in fi * ri(x-1) seconds.
//     For example, if fi = 3 and ri = 2, then the tire would finish its 1st lap in 3 seconds, 
//     its 2nd lap in 3 * 2 = 6 seconds, its 3rd lap in 3 * 22 = 12 seconds, etc.

// You are also given an integer changeTime and an integer numLaps.

// The race consists of numLaps laps and you may start the race with any tire. 
// You have an unlimited supply of each tire and after every lap, 
// you may change to any given tire (including the current tire type) if you wait changeTime seconds.

// Return the minimum time to finish the race.

// Example 1:
// Input: tires = [[2,3],[3,4]], changeTime = 5, numLaps = 4
// Output: 21
// Explanation: 
// Lap 1: Start with tire 0 and finish the lap in 2 seconds.
// Lap 2: Continue with tire 0 and finish the lap in 2 * 3 = 6 seconds.
// Lap 3: Change tires to a new tire 0 for 5 seconds and then finish the lap in another 2 seconds.
// Lap 4: Continue with tire 0 and finish the lap in 2 * 3 = 6 seconds.
// Total time = 2 + 6 + 5 + 2 + 6 = 21 seconds.
// The minimum time to complete the race is 21 seconds.

// Example 2:
// Input: tires = [[1,10],[2,2],[3,4]], changeTime = 6, numLaps = 5
// Output: 25
// Explanation: 
// Lap 1: Start with tire 1 and finish the lap in 2 seconds.
// Lap 2: Continue with tire 1 and finish the lap in 2 * 2 = 4 seconds.
// Lap 3: Change tires to a new tire 1 for 6 seconds and then finish the lap in another 2 seconds.
// Lap 4: Continue with tire 1 and finish the lap in 2 * 2 = 4 seconds.
// Lap 5: Change tires to tire 0 for 6 seconds then finish the lap in another 1 second.
// Total time = 2 + 4 + 6 + 2 + 4 + 6 + 1 = 25 seconds.
// The minimum time to complete the race is 25 seconds. 

// Constraints:
//     1 <= tires.length <= 10^5
//     tires[i].length == 2
//     1 <= fi, changeTime <= 10^5
//     2 <= ri <= 10^5
//     1 <= numLaps <= 1000

import "fmt"

func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
    inf := 1 << 31
    cost := [18]int{}
    for i := range cost {
        cost[i] = inf
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, v := range tires {
        f, r, s, t := v[0], v[1], 0, v[0]
        for i := 1; t <= changeTime + f; i++ {
            s += t
            cost[i] = min(cost[i], s)
            t *= r
        }
    }
    dp := make([]int, numLaps + 1)
    for i := range dp {
        dp[i] = inf
    }
    dp[0] = -changeTime
    for i := 1; i <= numLaps; i++ {
        for j := 1; j < min(18, i + 1); j++ {
            dp[i] = min(dp[i], dp[i-j] + cost[j])
        }
        dp[i] += changeTime
    }
    return dp[numLaps]
}

func minimumFinishTime1(tires [][]int, changeTime int, numLaps int) int {
    // 根据题目的数据范围一个轮胎最多跑17圈，所以设置一个上限17
    // 首先预处理出连续使用同一个轮胎跑 x 圈的最小耗时，记作 minSec[x]，这可以通过遍历每个轮胎计算出来。
    // minSec[i] 表示只使用一种轮胎（不知道是那种轮胎），最少需要多少秒
    minSec := make([]int, 18) 
    inf := 1 << 30
    for i := range minSec {
        minSec[i] = inf
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, ch := range tires {
        f, r := ch[0], ch[1]
        x, ti, sum := 1, f, 0
        for ti <= changeTime+f {
            sum += ti
            minSec[x] = min(minSec[x], sum)
            ti = ti * r
            x++
        }
    }
    dp := make([]int, numLaps + 1)
    dp[0] = -changeTime // 这里怎么想,这里只能赋这个值
    for i := 1; i <= numLaps; i++ {
        dp[i] = inf
        for j := 1; j <= 17 && j <= i; j++ {
            // 当i==j时，也就是从第0个开始里，这里还是加了 changeTime 所以初值得用-changeTime
            dp[i] = min(dp[i], changeTime + dp[i-j] + minSec[j])
        }
    }
    return dp[numLaps]
}

func main() {
    // Example 1:
    // Input: tires = [[2,3],[3,4]], changeTime = 5, numLaps = 4
    // Output: 21
    // Explanation: 
    // Lap 1: Start with tire 0 and finish the lap in 2 seconds.
    // Lap 2: Continue with tire 0 and finish the lap in 2 * 3 = 6 seconds.
    // Lap 3: Change tires to a new tire 0 for 5 seconds and then finish the lap in another 2 seconds.
    // Lap 4: Continue with tire 0 and finish the lap in 2 * 3 = 6 seconds.
    // Total time = 2 + 6 + 5 + 2 + 6 = 21 seconds.
    // The minimum time to complete the race is 21 seconds.
    fmt.Println(minimumFinishTime([][]int{{2,3},{3,4}}, 5, 4)) // 21
    // Example 2:
    // Input: tires = [[1,10],[2,2],[3,4]], changeTime = 6, numLaps = 5
    // Output: 25
    // Explanation: 
    // Lap 1: Start with tire 1 and finish the lap in 2 seconds.
    // Lap 2: Continue with tire 1 and finish the lap in 2 * 2 = 4 seconds.
    // Lap 3: Change tires to a new tire 1 for 6 seconds and then finish the lap in another 2 seconds.
    // Lap 4: Continue with tire 1 and finish the lap in 2 * 2 = 4 seconds.
    // Lap 5: Change tires to tire 0 for 6 seconds then finish the lap in another 1 second.
    // Total time = 2 + 4 + 6 + 2 + 4 + 6 + 1 = 25 seconds.
    // The minimum time to complete the race is 25 seconds. 
    fmt.Println(minimumFinishTime([][]int{{1,10},{2,2},{3,4}}, 6, 5)) // 25

    fmt.Println(minimumFinishTime1([][]int{{2,3},{3,4}}, 5, 4)) // 21
    fmt.Println(minimumFinishTime1([][]int{{1,10},{2,2},{3,4}}, 6, 5)) // 25
}