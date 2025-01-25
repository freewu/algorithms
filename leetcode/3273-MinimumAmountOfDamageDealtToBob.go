package main

// 3273. Minimum Amount of Damage Dealt to Bob
// You are given an integer power and two integer arrays damage and health, both having length n.

// Bob has n enemies, where enemy i will deal Bob damage[i] points of damage per second while they are alive (i.e. health[i] > 0).

// Every second, after the enemies deal damage to Bob, he chooses one of the enemies that is still alive and deals power points of damage to them.

// Determine the minimum total amount of damage points that will be dealt to Bob before all n enemies are dead.

// Example 1:
// Input: power = 4, damage = [1,2,3,4], health = [4,5,6,8]
// Output: 39
// Explanation:
// Attack enemy 3 in the first two seconds, after which enemy 3 will go down, the number of damage points dealt to Bob is 10 + 10 = 20 points.
// Attack enemy 2 in the next two seconds, after which enemy 2 will go down, the number of damage points dealt to Bob is 6 + 6 = 12 points.
// Attack enemy 0 in the next second, after which enemy 0 will go down, the number of damage points dealt to Bob is 3 points.
// Attack enemy 1 in the next two seconds, after which enemy 1 will go down, the number of damage points dealt to Bob is 2 + 2 = 4 points.

// Example 2:
// Input: power = 1, damage = [1,1,1,1], health = [1,2,3,4]
// Output: 20
// Explanation:
// Attack enemy 0 in the first second, after which enemy 0 will go down, the number of damage points dealt to Bob is 4 points.
// Attack enemy 1 in the next two seconds, after which enemy 1 will go down, the number of damage points dealt to Bob is 3 + 3 = 6 points.
// Attack enemy 2 in the next three seconds, after which enemy 2 will go down, the number of damage points dealt to Bob is 2 + 2 + 2 = 6 points.
// Attack enemy 3 in the next four seconds, after which enemy 3 will go down, the number of damage points dealt to Bob is 1 + 1 + 1 + 1 = 4 points.

// Example 3:
// Input: power = 8, damage = [40], health = [59]
// Output: 320

// Constraints:
//     1 <= power <= 10^4
//     1 <= n == damage.length == health.length <= 10^5
//     1 <= damage[i], health[i] <= 10^4

import "fmt"
import "math"
import "sort"
import "slices"

func minDamage(power int, damage []int, health []int) int64 {
    res, n := 0, len(health)
    ceil := func (n int, divisor int) int { return int(math.Ceil(float64(n) / float64(divisor))) }
    pairs := make([][2]int, n)
    for i := 0;i < n; i++ {
        health[i] = ceil(health[i], power)
        pairs[i] = [2]int{ damage[i], health[i] }
    }
    sort.Slice(pairs, func(i,j int) bool {
        if float64(pairs[i][0]) / float64(pairs[i][1]) == float64(pairs[j][0]) / float64(pairs[j][1]) { 
            return pairs[i][1] < pairs[j][1]
        }
        return float64(pairs[i][0]) / float64(pairs[i][1]) > float64(pairs[j][0]) / float64(pairs[j][1])
    })
    suffix := make([]int, n)
    suffix[n - 1] = pairs[n - 1][0]
    for i := n - 2; i >= 0; i-- {
        suffix[i] = pairs[i][0] + suffix[i + 1]
    }
    for i := 0; i < n; i++ {
        res += (suffix[i] * pairs[i][1])
    }
    return int64(res)
}

func minDamage1(power int, damage, health []int) int64 {
    type Pair struct{ k, d int }
    arr := make([]Pair, len(health))
    for i, v := range health {
        arr[i] = Pair{ (v - 1) / power + 1, damage[i] }
    }
    slices.SortFunc(arr, func(p, q Pair) int { 
        return p.k * q.d - q.k * p.d 
    })
    res, sum := 0, 0
    for _, v := range arr {
        sum += v.k
        res += (sum * v.d)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: power = 4, damage = [1,2,3,4], health = [4,5,6,8]
    // Output: 39
    // Explanation:
    // Attack enemy 3 in the first two seconds, after which enemy 3 will go down, the number of damage points dealt to Bob is 10 + 10 = 20 points.
    // Attack enemy 2 in the next two seconds, after which enemy 2 will go down, the number of damage points dealt to Bob is 6 + 6 = 12 points.
    // Attack enemy 0 in the next second, after which enemy 0 will go down, the number of damage points dealt to Bob is 3 points.
    // Attack enemy 1 in the next two seconds, after which enemy 1 will go down, the number of damage points dealt to Bob is 2 + 2 = 4 points.
    fmt.Println(minDamage(4, []int{1,2,3,4}, []int{4,5,6,8})) // 39
    // Example 2:
    // Input: power = 1, damage = [1,1,1,1], health = [1,2,3,4]
    // Output: 20
    // Explanation:
    // Attack enemy 0 in the first second, after which enemy 0 will go down, the number of damage points dealt to Bob is 4 points.
    // Attack enemy 1 in the next two seconds, after which enemy 1 will go down, the number of damage points dealt to Bob is 3 + 3 = 6 points.
    // Attack enemy 2 in the next three seconds, after which enemy 2 will go down, the number of damage points dealt to Bob is 2 + 2 + 2 = 6 points.
    // Attack enemy 3 in the next four seconds, after which enemy 3 will go down, the number of damage points dealt to Bob is 1 + 1 + 1 + 1 = 4 points.
    fmt.Println(minDamage(1, []int{1,1,1,1}, []int{1,2,3,4})) // 20
    // Example 3:
    // Input: power = 8, damage = [40], health = [59]
    // Output: 320
    fmt.Println(minDamage(8, []int{40}, []int{59})) // 320

    fmt.Println(minDamage1(4, []int{1,2,3,4}, []int{4,5,6,8})) // 39
    fmt.Println(minDamage1(1, []int{1,1,1,1}, []int{1,2,3,4})) // 20
    fmt.Println(minDamage1(8, []int{40}, []int{59})) // 320
}