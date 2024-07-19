package main

// 759. Employee Free Time
// We are given a list schedule of employees, which represents the working time for each employee.
// Each employee has a list of non-overlapping Intervals, and these intervals are in sorted order.
// Return the list of finite intervals representing common, positive-length free time for all employees, also in sorted order.
// (Even though we are representing Intervals in the form [x, y], the objects inside are Intervals, not lists or arrays.
//     For example, schedule[0][0].start = 1, schedule[0][0].end = 2, and schedule[0][0][0] is not defined).  
//     Also, we wouldn't include intervals like [5, 5] in our answer, as they have zero length.

// Example 1:
// Input: schedule = [[[1,2],[5,6]],[[1,3]],[[4,10]]]
// Output: [[3,4]]
// Explanation: There are a total of three employees, and all common
// free time intervals would be [-inf, 1], [3, 4], [10, inf].
// We discard any intervals that contain inf as they aren't finite.

// Example 2:
// Input: schedule = [[[1,3],[6,7]],[[2,4]],[[2,5],[9,12]]]
// Output: [[5,6],[7,9]]

// Constraints:
//     1 <= schedule.length , schedule[i].length <= 50
//     0 <= schedule[i].start < schedule[i].end <= 10^8

import "fmt"
import "sort"

type Interval struct {
    Start int
    End   int
}

/**
 * Definition for an Interval.
 * type Interval struct {
 *     Start int
 *     End   int
 * }
 */
func employeeFreeTime(schedule [][]*Interval) []*Interval {
    res,s := []*Interval{}, []*Interval{}
    for i:= range schedule {
        s = append(s, schedule[i]...)
    }
    sort.Slice(s, func(i, j int) bool {
        return s[i].Start < s[j].Start || (s[i].Start == s[j].Start && s[i].End > s[j].End)
    })
    max := func (x, y int) int { if x > y { return x; }; return y; }
    i1, i2 := s[0].Start, s[0].End
    for i := 1; i<len(s); i++ {
        if s[i].Start == i1 {
            continue
        }
        if s[i].Start > i2 {
            res = append(res, &Interval{i2, s[i].Start})
        }
        i1, i2 = s[i].Start, max(i2, s[i].End)
    }
    return res
}

func main() {
    // Example 1:
    // Input: schedule = [[[1,2],[5,6]],[[1,3]],[[4,10]]]
    // Output: [[3,4]]
    // Explanation: There are a total of three employees, and all common
    // free time intervals would be [-inf, 1], [3, 4], [10, inf].
    // We discard any intervals that contain inf as they aren't finite.
    schedule1 := [][]*Interval{
        []*Interval{ &Interval{1,2},  &Interval{5,6}, },
        []*Interval{ &Interval{1,3},  &Interval{4,10}, },
    }
    fmt.Println(employeeFreeTime(schedule1)) // [[3,4]]
    // Example 2:
    // Input: schedule = [[[1,3],[6,7]],[[2,4]],[[2,5],[9,12]]]
    // Output: [[5,6],[7,9]]
    schedule2 := [][]*Interval{
        []*Interval{ &Interval{1,3},  &Interval{6,7}, },
        []*Interval{ &Interval{2,4},  &Interval{2,5}, &Interval{9,12}, },
    }
    fmt.Println(employeeFreeTime(schedule2)) // [[5,6],[7,9]]
}