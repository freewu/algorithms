package main

// 4086. Minimum Interval Selection to Cover Target Intervals
// You are given two 2D integer arrays, availableIntervals and targetIntervals, each representing a list of closed intervals.

// Select a subset of intervals from availableIntervals so that every real point contained in any interval in targetIntervals belongs to at least one selected interval.

// Points outside the target intervals do not need to be covered. 
// Selected intervals may overlap or extend beyond the target intervals.

// Return the minimum number of intervals you need to select. 
// If it is impossible to cover all the target intervals, return -1.

// Example 1:
// Input: availableIntervals = [[1,4],[1,7],[3,9],[5,10],[10,15]], targetIntervals = [[1,3],[2,5],[7,8],[10,13],[11,15],[12,13]]
// Output: 3
// Explanation:
// The target intervals cover the ranges [1, 5], [7, 8], and [10, 15]. 
// By selecting [1, 4], [3, 9], and [10, 15] from availableIntervals, all target intervals can be fully covered with three intervals.

// Example 2:
// Input: availableIntervals = [[4,6],[1,4],[2,5]], targetIntervals = [[3,4],[2,5],[6,7]]
// Output: -1
// Explanation:
// It is impossible to select intervals from availableIntervals to cover all the intervals in targetIntervals.

// Example 3:
// Input: availableIntervals = [[1,2],[3,4]], targetIntervals = [[1,4]]
// Output: -1
// Explanation:
// Selecting both available intervals leaves every real point strictly between 2 and 3 uncovered. 
// For example, 2.5 belongs to the target interval but to neither available interval.

// Constraints:
//     - 1 <= availableIntervals.length, targetIntervals.length <= 5 * 10^{4}
//     - availableIntervals[i].length == targetIntervals[j].length == 2
//     - 1 <= availableIntervals[i][0] <= availableIntervals[i][1] <= 10^{5}
//     - 1 <= targetIntervals[j][0] <= targetIntervals[j][1] <= 10^{5}

import "fmt"
import "sort"

func minSelection(availableIntervals, targetIntervals [][]int) int {
    sort.Slice(availableIntervals, func(a, b int) bool {  // 按左端点排序
        return availableIntervals[a][0] < availableIntervals[b][0]
    })
    sort.Slice(targetIntervals, func(a, b int) bool { // 按左端点排序
        return targetIntervals[a][0] < targetIntervals[b][0]
    })
    res, tarL, tarR, currEnd, nextEnd := 0, 1 << 61, -1<< 61 -1, 0, 0  // 已建造的桥的右端点, 下一座桥的右端点的最大值
    j := 0
    for i, p := range targetIntervals {
        tarL = min(tarL, p[0])
        tarR = max(tarR, p[1])
        // 下一个区间与 [tarL, tarR] 相交，继续合并
        if i < len(targetIntervals)-1 && targetIntervals[i+1][0] <= tarR {
            continue
        }
        for currEnd < tarR {
            // 当前最靠左的未覆盖位置
            uncover := max(tarL, currEnd)
            // 寻找下一座桥的最远点
            for j < len(availableIntervals) && availableIntervals[j][0] <= uncover {
                nextEnd = max(nextEnd, availableIntervals[j][1])
                j++
            }
            // 无法覆盖 uncover，或者无法继续向右延伸
            if nextEnd < uncover || nextEnd == uncover && uncover < tarR {
                return -1
            }
            currEnd = nextEnd // 建桥后，最远可以到达 nextEnd
            res++
        }
        tarL = 1 << 61
    }
    return res
}

var avmx, tgmx [100_001]uint32
var gen uint32

func minSelection1(availableIntervals, targetIntervals [][]int) int {
    //avmx, tgmx := make([]uint32, 100_001), make([]uint32, 100_001)
    //mx, gen := 0, uint32(1 << 31)
    gen += 1 << 18
    mx := 0
    for _, a := range availableIntervals {
        mx = max(mx, a[1])
        avmx[a[0]] = max(avmx[a[0]], uint32(a[1])|gen)
    }
    for _, t := range targetIntervals {
        mx = max(mx, t[1])
        tgmx[t[0]] = max(tgmx[t[0]], uint32(t[1])|gen)
    }
    for i := 1; i <= mx; i++ {
        avmx[i] = max(avmx[i], avmx[i-1])
    }
    res, reach, g := 0, -1, int(gen)
    for left, e := range tgmx[:mx + 1] { 
        if e < gen {
            continue
        }
        for r := int(e) - g; reach < r; res++ {
            b := int(avmx[max(reach, left)]) - g    
            if b < left || b <= reach {
                return -1
            }
            reach = b
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: availableIntervals = [[1,4],[1,7],[3,9],[5,10],[10,15]], targetIntervals = [[1,3],[2,5],[7,8],[10,13],[11,15],[12,13]]
    // Output: 3
    // Explanation:
    // The target intervals cover the ranges [1, 5], [7, 8], and [10, 15]. 
    // By selecting [1, 4], [3, 9], and [10, 15] from availableIntervals, all target intervals can be fully covered with three intervals.
    fmt.Println(minSelection([][]int{{1,4},{1,7},{3,9},{5,10},{10,15}}, [][]int{{1,3},{2,5},{7,8},{10,13},{11,15},{12,13}})) // 3 
    // Example 2:
    // Input: availableIntervals = [[4,6],[1,4],[2,5]], targetIntervals = [[3,4],[2,5],[6,7]]
    // Output: -1
    // Explanation:
    // It is impossible to select intervals from availableIntervals to cover all the intervals in targetIntervals.
    fmt.Println(minSelection([][]int{{4,6},{1,4},{2,5}}, [][]int{{3,4},{2,5},{6,7}})) // -1 
    // Example 3:
    // Input: availableIntervals = [[1,2],[3,4]], targetIntervals = [[1,4]]
    // Output: -1
    // Explanation:
    // Selecting both available intervals leaves every real point strictly between 2 and 3 uncovered. 
    // For example, 2.5 belongs to the target interval but to neither available interval.
    fmt.Println(minSelection([][]int{{1,2},{3,4}}, [][]int{{1,4}})) // -1 

    fmt.Println(minSelection1([][]int{{1,4},{1,7},{3,9},{5,10},{10,15}}, [][]int{{1,3},{2,5},{7,8},{10,13},{11,15},{12,13}})) // 3 
    fmt.Println(minSelection1([][]int{{4,6},{1,4},{2,5}}, [][]int{{3,4},{2,5},{6,7}})) // -1 
    fmt.Println(minSelection1([][]int{{1,2},{3,4}}, [][]int{{1,4}})) // -1 
}