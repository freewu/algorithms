package main

// 731. My Calendar II
// You are implementing a program to use as your calendar. 
// We can add a new event if adding the event will not cause a triple booking.

// A triple booking happens when three events have some non-empty intersection (i.e., some moment is common to all the three events.).

// The event can be represented as a pair of integers start and end that represents a booking on the half-open interval [start, end), 
// the range of real numbers x such that start <= x < end.

// Implement the MyCalendarTwo class:
//     MyCalendarTwo() 
//         Initializes the calendar object.
//     boolean book(int start, int end) 
//         Returns true if the event can be added to the calendar successfully without causing a triple booking. 
//         Otherwise, return false and do not add the event to the calendar.

// Example 1:
// Input
// ["MyCalendarTwo", "book", "book", "book", "book", "book", "book"]
// [[], [10, 20], [50, 60], [10, 40], [5, 15], [5, 10], [25, 55]]
// Output
// [null, true, true, true, false, true, true]
// Explanation
// MyCalendarTwo myCalendarTwo = new MyCalendarTwo();
// myCalendarTwo.book(10, 20); // return True, The event can be booked. 
// myCalendarTwo.book(50, 60); // return True, The event can be booked. 
// myCalendarTwo.book(10, 40); // return True, The event can be double booked. 
// myCalendarTwo.book(5, 15);  // return False, The event cannot be booked, because it would result in a triple booking.
// myCalendarTwo.book(5, 10); // return True, The event can be booked, as it does not use time 10 which is already double booked.
// myCalendarTwo.book(25, 55); // return True, The event can be booked, as the time in [25, 40) will be double booked with the third event, the time [40, 50) will be single booked, and the time [50, 55) will be double booked with the second event.

// Constraints:
//     0 <= start < end <= 10^9
//     At most 1000 calls will be made to book.

import "fmt"

type pair struct{
    st int  //start time
    et int  //end time
}

type MyCalendarTwo struct {
    events []pair
    doubleBooked []pair
}

func Constructor() MyCalendarTwo {
    return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
    for _, v := range this.doubleBooked {
        if v.et > start && v.st < end {
            return false
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range this.events {
        if v.et > start && v.st < end {
            this.doubleBooked = append(this.doubleBooked, pair{max(v.st, start), min(v.et, end)})
        }
    }
    this.events = append(this.events, pair{start,end})
    return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

func main() {
    // MyCalendarTwo myCalendarTwo = new MyCalendarTwo();
    obj := Constructor()
    fmt.Println(obj)
    // myCalendarTwo.book(10, 20); // return True, The event can be booked. 
    fmt.Println(obj.Book(10,20)) // true
    fmt.Println(obj)
    // myCalendarTwo.book(50, 60); // return True, The event can be booked. 
    fmt.Println(obj.Book(50, 60)) // true
    fmt.Println(obj)
    // myCalendarTwo.book(10, 40); // return True, The event can be double booked. 
    fmt.Println(obj.Book(10, 40)) // true
    fmt.Println(obj)
    // myCalendarTwo.book(5, 15);  // return False, The event cannot be booked, because it would result in a triple booking.
    fmt.Println(obj.Book(5, 15)) // false
    fmt.Println(obj)
    // myCalendarTwo.book(5, 10); // return True, The event can be booked, as it does not use time 10 which is already double booked.
    fmt.Println(obj.Book(5, 10)) // true
    fmt.Println(obj)
    // myCalendarTwo.book(25, 55); // return True, The event can be booked, as the time in [25, 40) will be double booked with the third event, the time [40, 50) will be single booked, and the time [50, 55) will be double booked with the second event.
    fmt.Println(obj.Book(25, 55)) // true
    fmt.Println(obj)
}