package main

// 630. Course Schedule III
// There are n different online courses numbered from 1 to n. 
// You are given an array courses where courses[i] = [durationi, lastDayi] indicate 
// that the ith course should be taken continuously for durationi days and must be finished before or on lastDayi.

// You will start on the 1st day and you cannot take two or more courses simultaneously.
// Return the maximum number of courses that you can take.

// Example 1:
// Input: courses = [[100,200],[200,1300],[1000,1250],[2000,3200]]
// Output: 3
// Explanation: 
// There are totally 4 courses, but you can take 3 courses at most:
// First, take the 1st course, it costs 100 days so you will finish it on the 100th day, and ready to take the next course on the 101st day.
// Second, take the 3rd course, it costs 1000 days so you will finish it on the 1100th day, and ready to take the next course on the 1101st day. 
// Third, take the 2nd course, it costs 200 days so you will finish it on the 1300th day. 
// The 4th course cannot be taken now, since you will finish it on the 3300th day, which exceeds the closed date.

// Example 2:
// Input: courses = [[1,2]]
// Output: 1

// Example 3:
// Input: courses = [[3,2],[4,3]]
// Output: 0

// Constraints:
//     1 <= courses.length <= 10^4
//     1 <= durationi, lastDayi <= 10^4

import "fmt"
import "sort"
import "container/heap"
import "slices"

type MaxHeap []int
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Pop() (v interface{}) {
    *h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
    return v
}
func (h *MaxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }

func scheduleCourse(courses [][]int) int {
    day, maxHeap := 0, &MaxHeap{}
    sort.Slice(courses, func(i, j int) bool {
        return courses[i][1] < courses[j][1]
    })
    for _, course := range courses {
        heap.Push(maxHeap, course[0])
        day += course[0]
        if day > course[1] { day -= heap.Pop(maxHeap).(int); }
        if day > course[1] { return maxHeap.Len() }
    }
    return maxHeap.Len()
}


func scheduleCourse1(courses [][]int) int {
    // 贪心的思想,紧着先截止的考虑,
    // 如果有一门截止日期在后面的课能使得当前的花费时间减少,就替换.(因为修完一门课的收益只是1,但需要优先考虑那些先结束的课)
    hp := &CourseIntHeap{make(sort.IntSlice, 0, 1<<8)}
    heap.Init(hp)
    slices.SortFunc(courses, func(a, b []int) int { // 按照截止日期在前排序
        return a[1] - b[1]
    })
    end, res := 0, 0
    for _, course := range courses {
        if end+course[0] <= course[1] {
            res++
            heap.Push(hp, course[0]) // 以持续时间入堆,看后续的课程能否优化持续时间
            end += course[0]
        } else {
            // 贪心!! 当前的课程无法再结束时间前修完了,如果花费时间短,就让整体花费时间更优化,
            // 这里有个潜台词,这个课程一定是可以替换的(也就是替换完后,这个课程的结束时间一定在修改后的end前)
            if hp.Len() > 0 && hp.IntSlice[0] > course[0] {
                end -= hp.IntSlice[0] - course[0]
                // 正常应该先pop,再push,使用替换对顶然后Fix的方式,减少一次堆调整
                hp.IntSlice[0] = course[0]
                heap.Fix(hp, 0)
            }
        }
    }
    return res
}

type CourseIntHeap struct{ sort.IntSlice }
func (c CourseIntHeap) Less(i, j int) bool { return c.IntSlice[i] > c.IntSlice[j] }
func (c *CourseIntHeap) Push(x any)        { (*c).IntSlice = append((*c).IntSlice, x.(int)) }
func (c *CourseIntHeap) Pop() any {
    end := c.IntSlice.Len() - 1
    res := c.IntSlice[end]
    c.IntSlice = c.IntSlice[:end]
    return res
}

func main() {
    // Example 1:
    // Input: courses = [[100,200],[200,1300],[1000,1250],[2000,3200]]
    // Output: 3
    // Explanation: 
    // There are totally 4 courses, but you can take 3 courses at most:
    // First, take the 1st course, it costs 100 days so you will finish it on the 100th day, and ready to take the next course on the 101st day.
    // Second, take the 3rd course, it costs 1000 days so you will finish it on the 1100th day, and ready to take the next course on the 1101st day. 
    // Third, take the 2nd course, it costs 200 days so you will finish it on the 1300th day. 
    // The 4th course cannot be taken now, since you will finish it on the 3300th day, which exceeds the closed date.
    fmt.Println(scheduleCourse([][]int{{100,200},{200,1300},{1000,1250},{2000,3200}})) // 3
    // Example 2:
    // Input: courses = [[1,2]]
    // Output: 1
    fmt.Println(scheduleCourse([][]int{{1,2}})) // 1
    // Example 3:
    // Input: courses = [[3,2],[4,3]]
    // Output: 0
    fmt.Println(scheduleCourse([][]int{{3,2},{4,3}})) // 0

    fmt.Println(scheduleCourse1([][]int{{100,200},{200,1300},{1000,1250},{2000,3200}})) // 3
    fmt.Println(scheduleCourse1([][]int{{1,2}})) // 1
    fmt.Println(scheduleCourse1([][]int{{3,2},{4,3}})) // 0
}