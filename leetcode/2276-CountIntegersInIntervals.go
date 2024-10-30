package main

// 2276. Count Integers in Intervals
// Given an empty set of intervals, implement a data structure that can:
//     Add an interval to the set of intervals.
//     Count the number of integers that are present in at least one interval.

// Implement the CountIntervals class:
//     CountIntervals() Initializes the object with an empty set of intervals.
//     void add(int left, int right) Adds the interval [left, right] to the set of intervals.
//     int count() Returns the number of integers that are present in at least one interval.

// Note that an interval [left, right] denotes all the integers x where left <= x <= right.

// Example 1:
// Input
// ["CountIntervals", "add", "add", "count", "add", "count"]
// [[], [2, 3], [7, 10], [], [5, 8], []]
// Output
// [null, null, null, 6, null, 8]
// Explanation
// CountIntervals countIntervals = new CountIntervals(); // initialize the object with an empty set of intervals. 
// countIntervals.add(2, 3);  // add [2, 3] to the set of intervals.
// countIntervals.add(7, 10); // add [7, 10] to the set of intervals.
// countIntervals.count();    // return 6
//                            // the integers 2 and 3 are present in the interval [2, 3].
//                            // the integers 7, 8, 9, and 10 are present in the interval [7, 10].
// countIntervals.add(5, 8);  // add [5, 8] to the set of intervals.
// countIntervals.count();    // return 8
//                            // the integers 2 and 3 are present in the interval [2, 3].
//                            // the integers 5 and 6 are present in the interval [5, 8].
//                            // the integers 7 and 8 are present in the intervals [5, 8] and [7, 10].
//                            // the integers 9 and 10 are present in the interval [7, 10].

// Constraints:
//     1 <= left <= right <= 10^9
//     At most 10^5 calls in total will be made to add and count.
//     At least one call will be made to count.

import "fmt"
import "sort"

type CountIntervals struct {
    intervals [][]int
    isCounted bool
    count     int
}

// Constructor starts counting Intervals
func Constructor() CountIntervals {
    return CountIntervals{ make([][]int, 0), false, 0, }
}

// Add appends an interval to the list of Intervals
func (this *CountIntervals) Add(left int, right int) {
    this.intervals = append(this.intervals, []int{left, right})
    this.isCounted = false
}

// Count will provide the total integers in b/w the intervals
func (this *CountIntervals) Count() int {
    if this.isCounted { return this.count } // 处理多次 count 直接返回结果
    this.intervals = this.merge()
    this.isCounted = true
    this.count = 0
    for _, i := range this.intervals {
        this.count += ((i[1] - i[0]) + 1)
    }
    return this.count
}

// merge will combine all the overlapping intervals
func (this *CountIntervals) merge() [][]int {
    res := make([][]int, 0)
    // edge case
    if len(this.intervals) <= 1 { return this.intervals }
    sort.Slice(this.intervals, func(i, j int) bool {
        return this.intervals[i][0] < this.intervals[j][0]
    })
    // base case
    hand := make([]int, 0)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, i := range this.intervals {
        if len(hand) == 0 { // first interval
            hand = i
            continue
        }
        if hand[1] >= i[0] { // overlap found
            hand = []int{hand[0], max(hand[1], i[1])}
            continue
        }
        // overlap or type [i, i+n] [i+n+1, i+n+1+m] found
        // can be combined to [i, i+n+1+m]
        if hand[1]+1 == i[0] {
            hand = []int{hand[0], max(hand[1], i[1])}
            continue
        }
        res = append(res, hand) // interval is unique
        hand = i
    }
    res = append(res, hand) // last interval left in hand
    return res
}

/**
 * Your CountIntervals object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(left,right);
 * param_2 := obj.Count();
 */

func main() {
    // CountIntervals countIntervals = new CountIntervals(); // initialize the object with an empty set of intervals. 
    obj := Constructor()
    fmt.Println(obj)
    // countIntervals.add(2, 3);  // add [2, 3] to the set of intervals.
    obj.Add(2,3)
    fmt.Println(obj)
    // countIntervals.add(7, 10); // add [7, 10] to the set of intervals.
    obj.Add(7,10)
    fmt.Println(obj)
    // countIntervals.count();    // return 6
    //                            // the integers 2 and 3 are present in the interval [2, 3].
    //                            // the integers 7, 8, 9, and 10 are present in the interval [7, 10].
    fmt.Println(obj.Count()) // 6
    // countIntervals.add(5, 8);  // add [5, 8] to the set of intervals.
    obj.Add(5, 8)
    fmt.Println(obj)
    // countIntervals.count();    // return 8
    //                            // the integers 2 and 3 are present in the interval [2, 3].
    //                            // the integers 5 and 6 are present in the interval [5, 8].
    //                            // the integers 7 and 8 are present in the intervals [5, 8] and [7, 10].
    //                            // the integers 9 and 10 are present in the interval [7, 10].
    fmt.Println(obj.Count()) // 8
}