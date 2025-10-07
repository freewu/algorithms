package main

// 1488. Avoid Flood in The City
// Your country has an infinite number of lakes. 
// Initially, all the lakes are empty, but when it rains over the nth lake, the nth lake becomes full of water. 
// If it rains over a lake that is full of water, there will be a flood. 
// Your goal is to avoid floods in any lake.

// Given an integer array rains where:
//     rains[i] > 0 means there will be rains over the rains[i] lake.
//     rains[i] == 0 means there are no rains this day and you can choose one lake this day and dry it.

// Return an array ans where:
//     ans.length == rains.length
//     ans[i] == -1 if rains[i] > 0.
//     ans[i] is the lake you choose to dry in the ith day if rains[i] == 0.

// If there are multiple valid answers return any of them. 
// If it is impossible to avoid flood return an empty array.

// Notice that if you chose to dry a full lake, it becomes empty, but if you chose to dry an empty lake, nothing changes.

// Example 1:
// Input: rains = [1,2,3,4]
// Output: [-1,-1,-1,-1]
// Explanation: After the first day full lakes are [1]
// After the second day full lakes are [1,2]
// After the third day full lakes are [1,2,3]
// After the fourth day full lakes are [1,2,3,4]
// There's no day to dry any lake and there is no flood in any lake.

// Example 2:
// Input: rains = [1,2,0,0,2,1]
// Output: [-1,-1,2,1,-1,-1]
// Explanation: After the first day full lakes are [1]
// After the second day full lakes are [1,2]
// After the third day, we dry lake 2. Full lakes are [1]
// After the fourth day, we dry lake 1. There is no full lakes.
// After the fifth day, full lakes are [2].
// After the sixth day, full lakes are [1,2].
// It is easy that this scenario is flood-free. [-1,-1,1,2,-1,-1] is another acceptable scenario.

// Example 3:
// Input: rains = [1,2,0,1,2]
// Output: []
// Explanation: After the second day, full lakes are  [1,2]. We have to dry one lake in the third day.
// After that, it will rain over lakes [1,2]. It's easy to prove that no matter which lake you choose to dry in the 3rd day, the other one will flood.

// Constraints:
//     1 <= rains.length <= 10^5
//     0 <= rains[i] <= 10^9

import "fmt"

// greedy
func avoidFlood(rains []int) []int {
    res, mp, zeros := make([]int, len(rains)), make(map[int]int), []int{}
    for i, r := range rains {
        if r == 0 {
            zeros = append(zeros, i)
            continue
        }
        if _, ok := mp[r]; ok {
            zi := -1
            // find the closest index of zeros that could dry the current fulled lake
            //  O(NlogN) if using binary search to find the closest index
            for ti, tzi := range zeros {
                if tzi > mp[r] {
                    zi = tzi
                    zeros = append(zeros[:ti], zeros[ti+1:] ...)
                    break
                }
            }
            if zi == -1 { // 防不住
                return []int{}
            }
            res[zi] = r
        } 
        mp[r] = i
        res[i] = -1 // 抽干 这个湖泊的水
    }
    for i := range res { // fill the unused zeros index
        if res[i] == 0 {
            res[i] = 1
        }
    }
    return res
}

// 贪心 +  区间并查集 O(n*α(n))
func avoidFlood1(rains []int) []int {
    // 并查集的元素是天数
    // - 如果是连续的雨天或者是已经用过的"晴天"就合并(一个晴天是否被用过可以通过ans中设置的值判定)
    // - 当一个湖下雨时它已经满水了,则从上一次"这个湖"下雨往后找一个晴天
    n := len(rains)
    parent := make([]int, n)
    for i := range parent {
        parent[i] = i
    }
    find := func(x int) int {
        res := x
        for parent[res] != res {
            res = parent[res]
        }
        for parent[x] != res {
            parent[x], x = res, parent[x]
        }
        return res
    }
    lastRain := map[int]int{} // 标记一个lake上次下雨的时间, k:lakeId  v:day
    res := make([]int, n)
    for cur, rain := range rains {
        if rain == 0 { continue }
        res[cur] = -1
        lake := rain
        if last, ok := lastRain[lake]; ok { // 之前下过雨(代表湖里有水,每次新下雨都会抽干旧的,留下最后一次雨水)
            j := find(last + 1)
            for ; j < cur && res[j] != 0; j = find(j + 1) { // 一直往后寻找,直到越界或者找到一个没用过的晴天(res[j]==0)
                parent[j] = parent[j+1] // 因为不知道最终会合并到哪个上,选择合并到下一个即可(当然还是要找它的头部)
            }
            if j == cur { // 找到当天也没找到,必然洪水
                return []int{}
            } else {
                res[j] = lake
                // 因为每次去下一个区间都是通过 find(j+1)跳的,这个到了j就往回跳了. 另外可以不用merge j, 因为标记j是否被使用是用 res 的值去标记的!!
                parent[j-1] = parent[j]
            }
        }
        lastRain[lake] = cur
    }
    for i, x := range res { // 题目要求"不能闲着",必须找个湖抽水
        if x == 0 {
            res[i] = 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: rains = [1,2,3,4]
    // Output: [-1,-1,-1,-1]
    // Explanation: After the first day full lakes are [1]
    // After the second day full lakes are [1,2]
    // After the third day full lakes are [1,2,3]
    // After the fourth day full lakes are [1,2,3,4]
    // There's no day to dry any lake and there is no flood in any lake.
    fmt.Println(avoidFlood([]int{1,2,3,4})) // [-1,-1,-1,-1]
    // Example 2:
    // Input: rains = [1,2,0,0,2,1]
    // Output: [-1,-1,2,1,-1,-1]
    // Explanation: After the first day full lakes are [1]
    // After the second day full lakes are [1,2]
    // After the third day, we dry lake 2. Full lakes are [1]
    // After the fourth day, we dry lake 1. There is no full lakes.
    // After the fifth day, full lakes are [2].
    // After the sixth day, full lakes are [1,2].
    // It is easy that this scenario is flood-free. [-1,-1,1,2,-1,-1] is another acceptable scenario.
    fmt.Println(avoidFlood([]int{1,2,0,0,2,1})) // [-1,-1,2,1,-1,-1]
    // Example 3:
    // Input: rains = [1,2,0,1,2]
    // Output: []
    // Explanation: After the second day, full lakes are  [1,2]. We have to dry one lake in the third day.
    // After that, it will rain over lakes [1,2]. It's easy to prove that no matter which lake you choose to dry in the 3rd day, the other one will flood.
    fmt.Println(avoidFlood([]int{1,2,0,1,2})) // []

    fmt.Println(avoidFlood([]int{1,2,3,4,5,6,7,8,9})) // [-1 -1 -1 -1 -1 -1 -1 -1 -1]
    fmt.Println(avoidFlood([]int{9,8,7,6,5,4,3,2,1})) // [-1 -1 -1 -1 -1 -1 -1 -1 -1]

    fmt.Println(avoidFlood1([]int{1,2,3,4})) // [-1,-1,-1,-1]
    fmt.Println(avoidFlood1([]int{1,2,0,0,2,1})) // [-1,-1,2,1,-1,-1]
    fmt.Println(avoidFlood1([]int{1,2,0,1,2})) // []
    fmt.Println(avoidFlood1([]int{1,2,3,4,5,6,7,8,9})) // [-1 -1 -1 -1 -1 -1 -1 -1 -1]
    fmt.Println(avoidFlood1([]int{9,8,7,6,5,4,3,2,1})) // [-1 -1 -1 -1 -1 -1 -1 -1 -1]
}