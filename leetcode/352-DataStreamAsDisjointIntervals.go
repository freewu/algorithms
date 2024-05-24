package main

// 352. Data Stream as Disjoint Intervals
// Given a data stream input of non-negative integers a1, a2, ..., an, 
// summarize the numbers seen so far as a list of disjoint intervals.

// Implement the SummaryRanges class:
//     SummaryRanges() 
//         Initializes the object with an empty stream.
//     void addNum(int value) 
//         Adds the integer value to the stream.
//     int[][] getIntervals() 
//         Returns a summary of the integers in the stream currently as a list of disjoint intervals [starti, endi]. 
//         The answer should be sorted by starti.

// Example 1:
// Input
// ["SummaryRanges", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals"]
// [[], [1], [], [3], [], [7], [], [2], [], [6], []]
// Output
// [null, null, [[1, 1]], null, [[1, 1], [3, 3]], null, [[1, 1], [3, 3], [7, 7]], null, [[1, 3], [7, 7]], null, [[1, 3], [6, 7]]]
// Explanation
// SummaryRanges summaryRanges = new SummaryRanges();
// summaryRanges.addNum(1);      // arr = [1]
// summaryRanges.getIntervals(); // return [[1, 1]]
// summaryRanges.addNum(3);      // arr = [1, 3]
// summaryRanges.getIntervals(); // return [[1, 1], [3, 3]]
// summaryRanges.addNum(7);      // arr = [1, 3, 7]
// summaryRanges.getIntervals(); // return [[1, 1], [3, 3], [7, 7]]
// summaryRanges.addNum(2);      // arr = [1, 2, 3, 7]
// summaryRanges.getIntervals(); // return [[1, 3], [7, 7]]
// summaryRanges.addNum(6);      // arr = [1, 2, 3, 6, 7]
// summaryRanges.getIntervals(); // return [[1, 3], [6, 7]]
 
// Constraints:
//     0 <= value <= 10^4
//     At most 3 * 10^4 calls will be made to addNum and getIntervals.
//     At most 10^2 calls will be made to getIntervals.
 
// Follow up: What if there are lots of merges and the number of disjoint intervals is small compared to the size of the data stream?

import "fmt"
import "sort"

type SummaryRanges struct {
    cache map[int]bool
}

func Constructor() SummaryRanges {
    return SummaryRanges{map[int]bool{}}
}

func (this *SummaryRanges) AddNum(value int)  {
    this.cache[value] = true   
}

func (this *SummaryRanges) GetIntervals() [][]int {
    res, keys := [][]int{}, []int{}
    for k := range this.cache{
        keys = append(keys,k)
    }
    sort.Slice(keys,func(i,j int)bool {
        return keys[i] < keys[j]
    })
    for _,k := range keys{
        if len(res) > 0  && res[len(res) - 1][1] + 1 == k {
            res[len(res) - 1][1] = k
        } else {
            res = append(res , []int{ k, k })
        }
    }
    return res
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */

func main() {
    // SummaryRanges summaryRanges = new SummaryRanges();
    obj := Constructor() 
    fmt.Println(obj)
    // summaryRanges.addNum(1);      // arr = [1]
    obj.AddNum(1)
    fmt.Println(obj)
    // summaryRanges.getIntervals(); // return [[1, 1]]
    fmt.Println(obj.GetIntervals()) // [[1, 1]]
    // summaryRanges.addNum(3);      // arr = [1, 3]
    obj.AddNum(3)
    fmt.Println(obj)
    // summaryRanges.getIntervals(); // return [[1, 1], [3, 3]]
    fmt.Println(obj.GetIntervals()) // [[1, 1], [3, 3]]
    // summaryRanges.addNum(7);      // arr = [1, 3, 7]
    obj.AddNum(7)
    fmt.Println(obj)
    // summaryRanges.getIntervals(); // return [[1, 1], [3, 3], [7, 7]]
    fmt.Println(obj.GetIntervals()) // [[1, 1], [3, 3], [7, 7]]
    // summaryRanges.addNum(2);      // arr = [1, 2, 3, 7]
    obj.AddNum(2)
    fmt.Println(obj)
    // summaryRanges.getIntervals(); // return [[1, 3], [7, 7]]
    fmt.Println(obj.GetIntervals()) // [[1, 3], [7, 7]]
    // summaryRanges.addNum(6);      // arr = [1, 2, 3, 6, 7]
    obj.AddNum(6)
    fmt.Println(obj)
    // summaryRanges.getIntervals(); // return [[1, 3], [6, 7]]
    fmt.Println(obj.GetIntervals()) // [[1, 3], [6, 7]]
}