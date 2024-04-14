package main

// 403. Frog Jump
// A frog is crossing a river. The river is divided into some number of units, and at each unit, there may or may not exist a stone. 
// The frog can jump on a stone, but it must not jump into the water.

// Given a list of stones positions (in units) in sorted ascending order, determine if the frog can cross the river by landing on the last stone. 
// Initially, the frog is on the first stone and assumes the first jump must be 1 unit.

// If the frog's last jump was k units, its next jump must be either k - 1, k, or k + 1 units. 
// The frog can only jump in the forward direction.

// Example 1:
// Input: stones = [0,1,3,5,6,8,12,17]
// Output: true
// Explanation: The frog can jump to the last stone by jumping 1 unit to the 2nd stone, then 2 units to the 3rd stone, then 2 units to the 4th stone, then 3 units to the 6th stone, 4 units to the 7th stone, and 5 units to the 8th stone.

// Example 2:
// Input: stones = [0,1,2,3,4,8,9,11]
// Output: false
// Explanation: There is no way to jump to the last stone as the gap between the 5th and 6th stone is too large.
 
// Constraints:
//     2 <= stones.length <= 2000
//     0 <= stones[i] <= 2^31 - 1
//     stones[0] == 0
//     stones is sorted in a strictly increasing order.

import "fmt"

func canCross(stones []int) bool {
    if stones[1] != 1 {
        return false
    }
    dp := map[[2]int]bool{}
    var dfs func(i int, unit int) bool
    dfs = func(i int, unit int) bool {
        if i == len(stones) - 1 {
            return true
        }
        if v, f := dp[[2]int{i, unit}]; f {
            return v
        }
        res := false
        for j := i + 1; j < len(stones); j++ {
            newUnit := stones[j] - stones[i]
            if newUnit >= unit-1 && newUnit <= unit+1 && dfs(j, newUnit) {
                res = true
                break
            }
        }
        dp[[2]int{i, unit}] = res
        return dp[[2]int{i, unit}]
    }
    return dfs(1, 1)
}

func canCross1(stones []int) bool {
    cache := make([]map[int]bool, len(stones))
    var canCrossByPosition func(stones []int, k, idx int, repeated []map[int]bool) bool
    canCrossByPosition = func (stones []int, k, idx int, repeated []map[int]bool) bool {
        if v := repeated[idx]; v == nil {
            repeated[idx] = make(map[int]bool)
        }
        if idx == len(stones) - 1 {
            return true
        }
        if m, ok := repeated[idx][k]; ok {
            return m
        }
        for i := idx + 1; i < len(stones); i++ {
            delta := stones[i] - stones[idx]
            if delta == k + 1 {
                if canCrossByPosition(stones, k + 1, i, repeated) {
                    return true
                }
            } else if delta == k {
                if k == 0 {
                    continue
                }
                if canCrossByPosition(stones, k, i, repeated) {
                    return true
                }
            } else if delta == k - 1 {
                if k - 1 <= 0 {
                    continue
                }
                if canCrossByPosition(stones, k - 1, i, repeated) {
                    return true
                }
            }
            if delta > k + 1 {
                break
            }
        }
        repeated[idx][k] = false
        return false
    }
    return canCrossByPosition(stones, 0, 0, cache)
}

func main() {
    // Explanation: The frog can jump to the last stone by jumping 1 unit to the 2nd stone, then 2 units to the 3rd stone, then 2 units to the 4th stone, then 3 units to the 6th stone, 4 units to the 7th stone, and 5 units to the 8th stone.
    fmt.Println(canCross([]int{0,1,3,5,6,8,12,17})) // true
    // Explanation: There is no way to jump to the last stone as the gap between the 5th and 6th stone is too large.
    fmt.Println(canCross([]int{0,1,2,3,4,8,9,11})) // false

    fmt.Println(canCross1([]int{0,1,3,5,6,8,12,17})) // true
    fmt.Println(canCross1([]int{0,1,2,3,4,8,9,11})) // false
}