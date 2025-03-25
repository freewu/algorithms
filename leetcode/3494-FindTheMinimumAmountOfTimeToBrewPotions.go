package main

// 3494. Find the Minimum Amount of Time to Brew Potions
// You are given two integer arrays, skill and mana, of length n and m, respectively.

// In a laboratory, n wizards must brew m potions in order. 
// Each potion has a mana capacity mana[j] and must pass through all the wizards sequentially to be brewed properly. 
// The time taken by the ith wizard on the jth potion is timeij = skill[i] * mana[j].

// Since the brewing process is delicate, a potion must be passed to the next wizard immediately after the current wizard completes their work. 
// This means the timing must be synchronized so that each wizard begins working on a potion exactly when it arrives. ​

// Return the minimum amount of time required for the potions to be brewed properly.

// Example 1:
// Input: skill = [1,5,2,4], mana = [5,1,4,2]
// Output: 110
// Explanation:
// Potion Number | Start time | Wizard 0 done by | Wizard 1 done by | Wizard 2 done by | Wizard 3 done by
// 0                   0               5                   30                  40              60
// 1                   52              53                  58                  60              64
// 2                   54              58                  78                  86              102
// 3                   86              88                  98                  102             110
// As an example for why wizard 0 cannot start working on the 1st potion before time t = 52, consider the case where the wizards started preparing the 1st potion at time t = 50. At time t = 58, wizard 2 is done with the 1st potion, but wizard 3 will still be working on the 0th potion till time t = 60.

// Example 2:
// Input: skill = [1,1,1], mana = [1,1,1]
// Output: 5
// Explanation:
// Preparation of the 0th potion begins at time t = 0, and is completed by time t = 3.
// Preparation of the 1st potion begins at time t = 1, and is completed by time t = 4.
// Preparation of the 2nd potion begins at time t = 2, and is completed by time t = 5.

// Example 3:
// Input: skill = [1,2,3,4], mana = [1,2]
// Output: 21

// Constraints:
//     n == skill.length
//     m == mana.length
//     1 <= n, m <= 5000
//     1 <= mana[i], skill[i] <= 5000

import "fmt"

func minTime(skill []int, mana []int) int64 {
    n, m := len(skill), len(mana)
    dp := make([]int64, n)
    max := func (x, y int64) int64 { if x > y { return x; }; return y; }
    for j := 0; j < m; j++ {
        for i := 0; i < n; i++ {
            if (i > 0) {
                dp[i] = max(dp[i-1], dp[i])
            }
            dp[i] += int64(skill[i] * mana[j])
        }
        for i := n-1; i > 0; i-- {
            dp[i-1] = dp[i] - int64(skill[i] * mana[j])
        }
    }
    return dp[n-1]
}

func minTime1(skill, mana []int) int64 {
    n, m := len(skill), len(mana)
    s := make([]int, n + 1)
    for i, v := range skill {
        s[i + 1] = s[i] + v
    }
    suf := []int{n - 1}
    for i := n - 2; i >= 0; i-- {
        if skill[i] > skill[suf[len(suf) - 1]] {
            suf = append(suf, i)
        }
    }
    pre := []int{0}
    for i := 1; i < n; i++ {
        if skill[i] > skill[pre[len(pre) - 1]] {
            pre = append(pre, i)
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    start := 0
    for j := 1; j < m; j++ {
        record := suf
        if mana[j - 1] < mana[j] {
            record = pre
        }
        mx := 0
        for _, i := range record {
            mx = max(mx, mana[j - 1] * s[i + 1] - mana[j] * s[i])
        }
        start += mx
    }
    return int64(start + mana[m - 1] * s[n])
}

func main() {
    // Example 1:
    // Input: skill = [1,5,2,4], mana = [5,1,4,2]
    // Output: 110
    // Explanation:
    // Potion Number | Start time | Wizard 0 done by | Wizard 1 done by | Wizard 2 done by | Wizard 3 done by
    // 0                   0               5                   30                  40              60
    // 1                   52              53                  58                  60              64
    // 2                   54              58                  78                  86              102
    // 3                   86              88                  98                  102             110
    // As an example for why wizard 0 cannot start working on the 1st potion before time t = 52, consider the case where the wizards started preparing the 1st potion at time t = 50. At time t = 58, wizard 2 is done with the 1st potion, but wizard 3 will still be working on the 0th potion till time t = 60.
    fmt.Println(minTime([]int{1,5,2,4}, []int{5,1,4,2})) // 110
    // Example 2:
    // Input: skill = [1,1,1], mana = [1,1,1]
    // Output: 5
    // Explanation:
    // Preparation of the 0th potion begins at time t = 0, and is completed by time t = 3.
    // Preparation of the 1st potion begins at time t = 1, and is completed by time t = 4.
    // Preparation of the 2nd potion begins at time t = 2, and is completed by time t = 5.
    fmt.Println(minTime([]int{1,1,1}, []int{1,1,1})) // 5
    // Example 3:
    // Input: skill = [1,2,3,4], mana = [1,2]
    // Output: 21
    fmt.Println(minTime([]int{1,2,3,4}, []int{1,2})) // 21

    fmt.Println(minTime([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 525
    fmt.Println(minTime([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 729
    fmt.Println(minTime([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 729
    fmt.Println(minTime([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 525

    fmt.Println(minTime1([]int{1,5,2,4}, []int{5,1,4,2})) // 110
    fmt.Println(minTime1([]int{1,1,1}, []int{1,1,1})) // 5
    fmt.Println(minTime1([]int{1,2,3,4}, []int{1,2})) // 21
    fmt.Println(minTime1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 525
    fmt.Println(minTime1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 729
    fmt.Println(minTime1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 729
    fmt.Println(minTime1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 525
}