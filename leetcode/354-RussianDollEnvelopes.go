package main

// 354. Russian Doll Envelopes
// You are given a 2D array of integers envelopes where envelopes[i] = [wi, hi] represents the width and the height of an envelope.
// One envelope can fit into another if and only if both the width and height of one envelope are greater than the other envelope's width and height.
// Return the maximum number of envelopes you can Russian doll (i.e., put one inside the other).

// Note: You cannot rotate an envelope.

// Example 1:
// Input: envelopes = [[5,4],[6,4],[6,7],[2,3]]
// Output: 3
// Explanation: The maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).

// Example 2:
// Input: envelopes = [[1,1],[1,1],[1,1]]
// Output: 1
 
// Constraints:
//     1 <= envelopes.length <= 10^5
//     envelopes[i].length == 2
//     1 <= wi, hi <= 10^5

import "fmt"
import "sort"

// 二分法
func maxEnvelopes(envelopes [][]int) int {
    // 排序
    sort.Slice(envelopes, func(i, j int) bool {
        if envelopes[i][0] == envelopes[j][0] {
            return envelopes[i][1] > envelopes[j][1]
        }
        return envelopes[i][0] < envelopes[j][0]
    })
    dp := []int{}
    for _, e := range envelopes {
        low, high := 0, len(dp)
        for low < high {
            mid := low + (high-low)>>1
            if dp[mid] >= e[1] {
                high = mid
            } else {
                low = mid + 1
            }
        }
        if low == len(dp) {
            dp = append(dp, e[1])
        } else {
            dp[low] = e[1]
        }
    }
    return len(dp)
}

func maxEnvelopes1(envelopes [][]int) int {
    sort.Slice( envelopes, func( i, j int ) bool {
        if envelopes[i][0] != envelopes[j][0] {
            return envelopes[i][0] < envelopes[j][0]
        } else {
            return envelopes[i][1] > envelopes[j][1]
        }
    } )
    res := []int{}
    for i := 0; i < len(envelopes); i++ {
        if len(res) == 0 || envelopes[i][1] > res[len(res)-1] {
            res = append(res, envelopes[i][1] )
        } else {
            index := sort.SearchInts( res, envelopes[i][1] )
            res[index] = envelopes[i][1]
        }
    }
    return len(res)
}

func main() {
    // Explanation: The maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).
    fmt.Println(maxEnvelopes([][]int{{5,4},{6,4},{6,7},{2,3}})) // 3
    fmt.Println(maxEnvelopes([][]int{{1,1},{1,1},{1,1},{1,1}})) // 1

    fmt.Println(maxEnvelopes1([][]int{{5,4},{6,4},{6,7},{2,3}})) // 3
    fmt.Println(maxEnvelopes1([][]int{{1,1},{1,1},{1,1},{1,1}})) // 1
}