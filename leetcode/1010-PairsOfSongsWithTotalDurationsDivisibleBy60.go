package main

// 1010. Pairs of Songs With Total Durations Divisible by 60
// You are given a list of songs where the ith song has a duration of time[i] seconds.
// Return the number of pairs of songs for which their total duration in seconds is divisible by 60. 
// Formally, we want the number of indices i, j such that i < j with (time[i] + time[j]) % 60 == 0.

// Example 1:
// Input: time = [30,20,150,100,40]
// Output: 3
// Explanation: Three pairs have a total duration divisible by 60:
// (time[0] = 30, time[2] = 150): total duration 180
// (time[1] = 20, time[3] = 100): total duration 120
// (time[1] = 20, time[4] = 40): total duration 60

// Example 2:
// Input: time = [60,60,60]
// Output: 3
// Explanation: All three pairs have a total duration of 120, which is divisible by 60.

// Constraints:
//     1 <= time.length <= 6 * 10^4
//     1 <= time[i] <= 500

import "fmt"

// 暴力法 超出时间限制 33 / 35 
func numPairsDivisibleBy60(time []int) int {
    res, n := 0, len(time)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if (time[i] + time[j]) % 60 == 0 {
                res++
            }
        }
    }
    return res
}

func numPairsDivisibleBy601(time []int) int {
    res, freq := 0, make(map[int]int)
    for i := 0; i < len(time); i++ {
        mod := time[i] % 60
        freq[mod]++
    }
    for k, v := range freq {
        switch {
        case k == 0 || k == 30:
            count := freq[k]
            res += (count * (count - 1)) / 2
        case k <= 29:
            if _, ok := freq[60 - k]; ok {
                res += v * freq[60 - k]
            }
        }
    }
    return res
}

func numPairsDivisibleBy602(time []int) int {
    res, cnt := 0, [60]int{}
    for _, t := range time {
        res += cnt[(60 - t % 60) % 60]
        cnt[t % 60]++
    }
    return res
}

func main() {
    // Example 1:
    // Input: time = [30,20,150,100,40]
    // Output: 3
    // Explanation: Three pairs have a total duration divisible by 60:
    // (time[0] = 30, time[2] = 150): total duration 180
    // (time[1] = 20, time[3] = 100): total duration 120
    // (time[1] = 20, time[4] = 40): total duration 60
    fmt.Println(numPairsDivisibleBy60([]int{30,20,150,100,40})) // 3
    // Example 2:
    // Input: time = [60,60,60]
    // Output: 3
    // Explanation: All three pairs have a total duration of 120, which is divisible by 60.
    fmt.Println(numPairsDivisibleBy60([]int{60,60,60})) // 3

    fmt.Println(numPairsDivisibleBy601([]int{30,20,150,100,40})) // 3
    fmt.Println(numPairsDivisibleBy601([]int{60,60,60})) // 3

    fmt.Println(numPairsDivisibleBy602([]int{30,20,150,100,40})) // 3
    fmt.Println(numPairsDivisibleBy602([]int{60,60,60})) // 3
}