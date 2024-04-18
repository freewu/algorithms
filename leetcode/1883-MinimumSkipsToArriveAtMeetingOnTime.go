package main

// 1883. Minimum Skips to Arrive at Meeting On Time
// You are given an integer hoursBefore, the number of hours you have to travel to your meeting. 
// To arrive at your meeting, you have to travel through n roads. 
// The road lengths are given as an integer array dist of length n, where dist[i] describes the length of the ith road in kilometers. 
// In addition, you are given an integer speed, which is the speed (in km/h) you will travel at.

// After you travel road i, you must rest and wait for the next integer hour before you can begin traveling on the next road. 
// Note that you do not have to rest after traveling the last road because you are already at the meeting.
//     For example, if traveling a road takes 1.4 hours, you must wait until the 2 hour mark before traveling the next road. 
//     If traveling a road takes exactly 2 hours, you do not need to wait.

// However, you are allowed to skip some rests to be able to arrive on time, meaning you do not need to wait for the next integer hour. 
// Note that this means you may finish traveling future roads at different hour marks.
//     For example, suppose traveling the first road takes 1.4 hours and traveling the second road takes 0.6 hours.
//     Skipping the rest after the first road will mean you finish traveling the second road right at the 2 hour mark, letting you start traveling the third road immediately.

// Return the minimum number of skips required to arrive at the meeting on time, or -1 if it is impossible.

// Example 1:
// Input: dist = [1,3,2], speed = 4, hoursBefore = 2
// Output: 1
// Explanation:
// Without skipping any rests, you will arrive in (1/4 + 3/4) + (3/4 + 1/4) + (2/4) = 2.5 hours.
// You can skip the first rest to arrive in ((1/4 + 0) + (3/4 + 0)) + (2/4) = 1.5 hours.
// Note that the second rest is shortened because you finish traveling the second road at an integer hour due to skipping the first rest.

// Example 2:
// Input: dist = [7,3,5,5], speed = 2, hoursBefore = 10
// Output: 2
// Explanation:
// Without skipping any rests, you will arrive in (7/2 + 1/2) + (3/2 + 1/2) + (5/2 + 1/2) + (5/2) = 11.5 hours.
// You can skip the first and third rest to arrive in ((7/2 + 0) + (3/2 + 0)) + ((5/2 + 0) + (5/2)) = 10 hours.

// Example 3:
// Input: dist = [7,3,5,5], speed = 1, hoursBefore = 10
// Output: -1
// Explanation: It is impossible to arrive at the meeting on time even if you skip all the rests.
 
// Constraints:
//     n == dist.length
//     1 <= n <= 1000
//     1 <= dist[i] <= 10^5
//     1 <= speed <= 10^6
//     1 <= hoursBefore <= 10^7

import "fmt"

func minSkips(dist []int, speed int, hoursBefore int) int {
    // subproblem is the minimum number of hours
    // to reach destination from here with at most n skips.
    // times are stored with a radix of speed to avoid dealing with floats or rationals.
    // Technically should use int64 since speed is upto 10^6 and dist[i] is upto 10^5.
    dp := make([][]int, len(dist), len(dist));
    for i := 0; i < len(dist); i++ {
        dp[i] = make([]int, len(dist), len(dist))
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    // Round up to the nearest integer using a speed radix.
    speedCeil := func (time int, speed int) int {
        return ( (time + speed - 1) / speed) * speed // round up to the nearest multiple of speed
    }
    // The final step needs not skip, per the spec.
    dp[len(dist)-1][0] = dist[len(dist)-1]
    for i := len(dist) - 2; i >= 0; i-- {
        // dp[i][0] is the time with no skips, so we always
        dp[i][0] = speedCeil(dp[i+1][0], speed) + dist[i]
        for skips := 1; skips < len(dist) - 1 - i; skips++ {
            // time := dist[i] / speed  (since we use speed radix this is a no-op)
            // Choose between an answer from the same point with lower skips, and an 
            // answer with no additional skips added here and an answer with one additional
            // skip added here.
            dp[i][skips] = min(
                dp[i][skips-1], 
                min(speedCeil(dp[i+1][skips], speed)+ dist[i], dp[i+1][skips-1] + dist[i]),
            )
        }
        skips := len(dist) - 1 - i
        // The case where there is a skip at every step. As this position can have one more 
        // skip than the previous this is the case where we only consider an answer from the
        // same point with lower skips, and an answer with one additional skip added here.
        dp[i][skips] = min(dp[i][skips-1], dp[i+1][skips-1]+dist[i])
    }
    needed := hoursBefore * speed
    for skips := 0; skips < len(dist); skips++ {
        if speedCeil(dp[0][skips], speed) <= needed {
            return skips 
        }
    }
    return -1
}

func minSkips1(dist []int, speed int, hoursBefore int) int {
    n, mx := len(dist), 1 << 32 -1
    f := make([][]int, n + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
        for j := range f[i] {
            f[i][j] = mx / 2
        }
    }
    f[0][0] = 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        for j := 0; j <= i; j++ {
            if j != i {
                f[i][j] = min(f[i][j], ((f[i - 1][j] + (dist[i - 1]) - 1) / (speed) + 1) * (speed))
            }
            if j != 0 {
                f[i][j] = min(f[i][j], f[i - 1][j - 1] + (dist[i - 1]))
            }
        }
    }
    for j := 0; j <= n; j++ {
        if f[n][j] <= (hoursBefore) * (speed) {
            return j
        }
    }
    return -1
}

func main() {
    // Without skipping any rests, you will arrive in (1/4 + 3/4) + (3/4 + 1/4) + (2/4) = 2.5 hours.
    // You can skip the first rest to arrive in ((1/4 + 0) + (3/4 + 0)) + (2/4) = 1.5 hours.
    // Note that the second rest is shortened because you finish traveling the second road at an integer hour due to skipping the first rest.
    fmt.Println(minSkips([]int{1,3,2},4,2)) // 1
    // Without skipping any rests, you will arrive in (7/2 + 1/2) + (3/2 + 1/2) + (5/2 + 1/2) + (5/2) = 11.5 hours.
    // You can skip the first and third rest to arrive in ((7/2 + 0) + (3/2 + 0)) + ((5/2 + 0) + (5/2)) = 10 hours.
    fmt.Println(minSkips([]int{7,3,5,5},2,10)) // 2
    // Explanation: It is impossible to arrive at the meeting on time even if you skip all the rests.
    fmt.Println(minSkips([]int{7,3,5,5},1,10)) // -1

    fmt.Println(minSkips1([]int{1,3,2},4,2)) // 1
    fmt.Println(minSkips1([]int{7,3,5,5},2,10)) // 2
    fmt.Println(minSkips1([]int{7,3,5,5},1,10)) // -1
}