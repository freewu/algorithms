package main

// 732. My Calendar III
// A k-booking happens when k events have some non-empty intersection 
// (i.e., there is some time that is common to all k events.)

// You are given some events [startTime, endTime), 
// after each given event, return an integer k representing the maximum k-booking between all the previous events.

// Implement the MyCalendarThree class:
//     MyCalendarThree() 
//         Initializes the object.
//     int book(int startTime, int endTime) 
//         Returns an integer k representing the largest integer such that there exists a k-booking in the calendar.

// Example 1:
// Input
// ["MyCalendarThree", "book", "book", "book", "book", "book", "book"]
// [[], [10, 20], [50, 60], [10, 40], [5, 15], [5, 10], [25, 55]]
// Output
// [null, 1, 1, 2, 3, 3, 3]
// Explanation
// MyCalendarThree myCalendarThree = new MyCalendarThree();
// myCalendarThree.book(10, 20); // return 1
// myCalendarThree.book(50, 60); // return 1
// myCalendarThree.book(10, 40); // return 2
// myCalendarThree.book(5, 15); // return 3
// myCalendarThree.book(5, 10); // return 3
// myCalendarThree.book(25, 55); // return 3

// Constraints:
//     0 <= startTime < endTime <= 10^9
//     At most 400 calls will be made to book.

import "fmt"

type MyCalendarThree struct {
    Bookings [][2]int
    Intersection [][2]int
}

func Constructor() MyCalendarThree {
    return MyCalendarThree{}
}

func (this *MyCalendarThree) Book(start int, end int) int {
    current := [2]int{start, end-1}
    currentIntersection := [2]int{1, 1}
    res := 0
    for i, booking := range this.Bookings {
        for j := 0; j <= 1; j++ {
            if isIntersect(booking[j], current[0], current[1]) {
                this.Intersection[i][j]++
            }
            res = max(res, this.Intersection[i][j])
            if isIntersect(current[j], booking[0], booking[1]) {
                currentIntersection[j]++
            }
        }
    }
    this.Bookings = append(this.Bookings, current)
    this.Intersection = append(this.Intersection, currentIntersection)
    return max(res, max(currentIntersection[0], currentIntersection[1]))
}

func isIntersect(a, b1, b2 int) bool {
    return a >= b1 && a <= b2
}

func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */

func main() {
    // MyCalendarThree myCalendarThree = new MyCalendarThree();
    obj := Constructor()
    fmt.Println(obj)
    // myCalendarThree.book(10, 20); // return 1
    fmt.Println(obj.Book(10,20)) // 1
    fmt.Println(obj)
    // myCalendarThree.book(50, 60); // return 1
    fmt.Println(obj.Book(50, 60)) // 1
    fmt.Println(obj)
    // myCalendarThree.book(10, 40); // return 2
    fmt.Println(obj.Book(10, 40)) // 2
    fmt.Println(obj)
    // myCalendarThree.book(5, 15); // return 3
    fmt.Println(obj.Book(5, 15)) // 3
    fmt.Println(obj)
    // myCalendarThree.book(5, 10); // return 3
    fmt.Println(obj.Book(5, 10)) // 3
    fmt.Println(obj)
    // myCalendarThree.book(25, 55); // return 3
    fmt.Println(obj.Book(25, 55)) // 3
    fmt.Println(obj)
}