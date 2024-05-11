package main

// 729. My Calendar I
// You are implementing a program to use as your calendar. 
// We can add a new event if adding the event will not cause a double booking.

// A double booking happens when two events have some non-empty intersection (i.e., some moment is common to both events.).

// The event can be represented as a pair of integers start and end that represents a booking on the half-open interval [start, end), 
// the range of real numbers x such that start <= x < end.

// Implement the MyCalendar class:
//     MyCalendar() 
//         Initializes the calendar object.
//     boolean book(int start, int end) 
//         Returns true if the event can be added to the calendar successfully without causing a double booking. 
//         Otherwise, return false and do not add the event to the calendar.

// Example 1:
// Input
// ["MyCalendar", "book", "book", "book"]
// [[], [10, 20], [15, 25], [20, 30]]
// Output
// [null, true, false, true]
// Explanation
// MyCalendar myCalendar = new MyCalendar();
// myCalendar.book(10, 20); // return True
// myCalendar.book(15, 25); // return False, It can not be booked because time 15 is already booked by another event.
// myCalendar.book(20, 30); // return True, The event can be booked, as the first event takes every time less than 20, but not including 20.

// Constraints:
//     0 <= start < end <= 10^9
//     At most 1000 calls will be made to book.

import "fmt"

type MyCalendar struct {
    data [][]int 
}

func Constructor() MyCalendar {
    return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
    // 判断时间范围是否已占用
    for _, v := range this.data {
        if start >= v[0] && start < v[1] { // [10, 20] [15, 25]
            return false
        }
        if start <= v[0] && end > v[0] {  // [10, 20] [5, 25]
            return false
        }
    }
    this.data = append(this.data,[]int{ start, end })
    return true
}

type node struct {
    left, right *node
    start, end int
}

func NewNode(start, end int) *node {
    return &node{nil, nil, start, end}
}

type MyCalendar1 struct {
    root *node
}

func Constructor1() MyCalendar1 {
    return MyCalendar1{root: nil}
}


func (this *MyCalendar1) Book(start int, end int) bool {
    if this.root == nil {
        this.root = NewNode(start, end)
        return true
    }
    c := this.root
    for c != nil {
        if end <= c.start {
            if c.left == nil {
                c.left = NewNode(start, end)
                return true
            }
            c = c.left
        } else if start >= c.end {
            if c.right == nil {
                c.right = NewNode(start, end)
                return true
            }
            c = c.right
        } else {
            break
        }
    }
    return false
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
func main() {
    // MyCalendar myCalendar = new MyCalendar();
    obj := Constructor()
    // myCalendar.book(10, 20); // return True
    fmt.Println(obj.Book(10,20)) // true
    fmt.Println(obj)
    // myCalendar.book(15, 25); // return False, It can not be booked because time 15 is already booked by another event.
    fmt.Println(obj.Book(15,25)) // false
    fmt.Println(obj)
    // myCalendar.book(20, 30); // return True, The event can be booked, as the first event takes every time less than 20, but not including 20.
    fmt.Println(obj.Book(20,30)) // false
    fmt.Println(obj)

    fmt.Println(obj.Book(5,18)) // false
    fmt.Println(obj)


    // MyCalendar myCalendar = new MyCalendar();
    obj1 := Constructor1()
    // myCalendar.book(10, 20); // return True
    fmt.Println(obj1.Book(10,20)) // true
    fmt.Println(obj1)
    // myCalendar.book(15, 25); // return False, It can not be booked because time 15 is already booked by another event.
    fmt.Println(obj1.Book(15,25)) // false
    fmt.Println(obj1)
    // myCalendar.book(20, 30); // return True, The event can be booked, as the first event takes every time less than 20, but not including 20.
    fmt.Println(obj1.Book(20,30)) // false
    fmt.Println(obj1)

    fmt.Println(obj1.Book(5,18)) // false
    fmt.Println(obj1)
}