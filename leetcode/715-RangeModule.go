package main

// 715. Range Module
// A Range Module is a module that tracks ranges of numbers. 
// Design a data structure to track the ranges represented as half-open intervals and query about them.

// A half-open interval [left, right) denotes all the real numbers x where left <= x < right.

// Implement the RangeModule class:
//     RangeModule() Initializes the object of the data structure.
//     void addRange(int left, int right) Adds the half-open interval [left, right), tracking every real number in that interval. Adding an interval that partially overlaps with currently tracked numbers should add any numbers in the interval [left, right) that are not already tracked.
//     boolean queryRange(int left, int right) Returns true if every real number in the interval [left, right) is currently being tracked, and false otherwise.
//     void removeRange(int left, int right) Stops tracking every real number currently being tracked in the half-open interval [left, right).

// Example 1:
// Input
// ["RangeModule", "addRange", "removeRange", "queryRange", "queryRange", "queryRange"]
// [[], [10, 20], [14, 16], [10, 14], [13, 15], [16, 17]]
// Output
// [null, null, null, true, false, true]
// Explanation
// RangeModule rangeModule = new RangeModule();
// rangeModule.addRange(10, 20);
// rangeModule.removeRange(14, 16);
// rangeModule.queryRange(10, 14); // return True,(Every number in [10, 14) is being tracked)
// rangeModule.queryRange(13, 15); // return False,(Numbers like 14, 14.03, 14.17 in [13, 15) are not being tracked)
// rangeModule.queryRange(16, 17); // return True, (The number 16 in [16, 17) is still being tracked, despite the remove operation)

// Constraints:
//     1 <= left < right <= 10^9
//     At most 10^4 calls will be made to addRange, queryRange, and removeRange.

import "fmt"

type RangeModule struct {
    intervals [][2]int
}

func Constructor() RangeModule {
    return RangeModule{}
}

func (this *RangeModule) AddRange(left int, right int)  {
    if len(this.intervals) == 0 {
        this.intervals = append(this.intervals, [2]int{left, right})
        return
    }
    if right < this.intervals[0][0] {
        this.intervals = append(this.intervals, [2]int{})
        copy(this.intervals[1:], this.intervals)
        this.intervals[0] = [2]int{left, right}
        return
    }
    least, greatest := left, right
    updatedIntervals, added := [][2]int{}, false
    for i, interval := range this.intervals {
        if left > interval[1] {
            updatedIntervals = append(updatedIntervals, interval)
        } else if right < interval[0] {
            updatedIntervals = append(updatedIntervals, [2]int{least, greatest})
            updatedIntervals = append(updatedIntervals, this.intervals[i:]...)
            added = true
            break
        }
        if (left >= interval[0] && left <= interval[1]) || (right >= interval[0] && right <= interval[1]) || (left <= interval[0] && right >= interval[1]) {
            if interval[0] <= least {
                least = interval[0]
            }
            if interval[1] >= greatest {
                greatest = interval[1]
            }
        }
    }
    if !added {
        updatedIntervals = append(updatedIntervals, [2]int{least, greatest})
    }
    this.intervals = updatedIntervals
    return
}


func (this *RangeModule) QueryRange(left int, right int) bool {
    l, r := 0, len(this.intervals) - 1
    for l <= r {
        mid := (l + r) / 2
        interval := this.intervals[mid]
        if left >= interval[0] && right <= interval[1] {
            return true
        } else if left >= interval[0] && left <= interval[1] && right > interval[1] {
            return false
        } else if right <= interval[1] && right >= interval[0] && left < interval[0] {
            return false
        } else if left < interval[0] {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return false
}


func (this *RangeModule) RemoveRange(left int, right int)  {
    updatedIntervals := [][2]int{}
    for i, interval := range this.intervals {
        if right < interval[0] {
            updatedIntervals = append(updatedIntervals, this.intervals[i:]...)
            break
        }
        if left >= interval[0] && left <= interval[1] {
            if left != interval[0] {
                updatedIntervals = append(updatedIntervals, [2]int{interval[0], left})
            }
            if right < interval[1] { 
                updatedIntervals = append(updatedIntervals, [2]int{right, interval[1]})
                continue
            }
        } else if left <= interval[0] && right > interval[1] {
            continue
        } else if right >= interval[0] && right < interval[1] {
            updatedIntervals = append(updatedIntervals, [2]int{right, interval[1]})
        } else if right < interval[0] || left > interval[1] {
            updatedIntervals = append(updatedIntervals, interval)
        }
    } 
    this.intervals = updatedIntervals
    return
}

func main() {
    // RangeModule rangeModule = new RangeModule();
    obj := Constructor()
    fmt.Println(obj)
    // rangeModule.addRange(10, 20);
    obj.AddRange(10, 20)
    fmt.Println(obj)
    // rangeModule.removeRange(14, 16);
    obj.RemoveRange(14, 16)
    fmt.Println(obj)
    // rangeModule.queryRange(10, 14); // return True,(Every number in [10, 14) is being tracked)
    fmt.Println(obj.QueryRange(10, 14)) // true
    // rangeModule.queryRange(13, 15); // return False,(Numbers like 14, 14.03, 14.17 in [13, 15) are not being tracked)
    fmt.Println(obj.QueryRange(13, 15)) // false
    // rangeModule.queryRange(16, 17); // return True, (The number 16 in [16, 17) is still being tracked, despite the remove operation)
    fmt.Println(obj.QueryRange(16, 17)) // true
}