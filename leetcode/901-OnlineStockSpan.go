package main

// 901. Online Stock Span
// Design an algorithm that collects daily price quotes for some stock and returns the span of that stock's price for the current day.

// The span of the stock's price in one day is the maximum number of consecutive days (starting from that day and going backward) for which the stock price was less than or equal to the price of that day.
//     For example, if the prices of the stock in the last four days is [7,2,1,2] and the price of the stock today is 2, then the span of today is 4 because starting from today, the price of the stock was less than or equal 2 for 4 consecutive days.
//     Also, if the prices of the stock in the last four days is [7,34,1,2] and the price of the stock today is 8, then the span of today is 3 because starting from today, the price of the stock was less than or equal 8 for 3 consecutive days.

// Implement the StockSpanner class:
//     StockSpanner() Initializes the object of the class.
//     int next(int price) Returns the span of the stock's price given that today's price is price.
    
// Example 1:
// Input
// ["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
// [[], [100], [80], [60], [70], [60], [75], [85]]
// Output
// [null, 1, 1, 1, 2, 1, 4, 6]
// Explanation
// StockSpanner stockSpanner = new StockSpanner();
// stockSpanner.next(100); // return 1
// stockSpanner.next(80);  // return 1
// stockSpanner.next(60);  // return 1
// stockSpanner.next(70);  // return 2
// stockSpanner.next(60);  // return 1
// stockSpanner.next(75);  // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
// stockSpanner.next(85);  // return 6

// Constraints:
//     1 <= price <= 10^5
//     At most 10^4 calls will be made to next.

import "fmt"

type StockSpanner1 struct {
    stack [][2]int    
}

func Constructor1() StockSpanner1 {
    return StockSpanner1{[][2]int{}}
}

func (this *StockSpanner1) Next(price int) int {
    res := 1
    for l := len(this.stack)-1; l > -1 && this.stack[l][0] <= price; l-- {
        res += this.stack[l][1]
        this.stack = this.stack[:l]
    }
    this.stack = append(this.stack, [2]int{price, res})
    return res   
}

type StockSpanner struct {
    cur int
    stack  []struct{ idx, val int }
}

func Constructor() StockSpanner {
    return StockSpanner{0, []struct{ idx, val int }{}}
}

func (s *StockSpanner) Next(price int) int {
    res := 0
    for len(s.stack) > 0 && price >= s.stack[len(s.stack) - 1].val {
        s.stack = s.stack[:len(s.stack) - 1]
    }
    if len(s.stack) > 0 {
        res += s.cur - s.stack[len(s.stack) - 1].idx
    } else {
        res = s.cur + 1
    }
    s.stack = append(s.stack, struct { idx, val int }{s.cur, price})
    s.cur++
    return res
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

func main() {
    // StockSpanner stockSpanner = new StockSpanner();
    obj := Constructor()
    fmt.Println(obj)
    // stockSpanner.next(100); // return 1
    fmt.Println(obj.Next(100)) // 1
    fmt.Println(obj)
    // stockSpanner.next(80);  // return 1
    fmt.Println(obj.Next(80)) // 1
    fmt.Println(obj)
    // stockSpanner.next(60);  // return 1
    fmt.Println(obj.Next(60)) // 1
    fmt.Println(obj)
    // stockSpanner.next(70);  // return 2
    fmt.Println(obj.Next(70)) // 2
    fmt.Println(obj)
    // stockSpanner.next(60);  // return 1
    fmt.Println(obj.Next(60)) // 1
    fmt.Println(obj)
    // stockSpanner.next(75);  // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
    fmt.Println(obj.Next(75)) // 4
    fmt.Println(obj)
    // stockSpanner.next(85);  // return 6
    fmt.Println(obj.Next(85)) // 6
    fmt.Println(obj)

    obj1 := Constructor1()
    fmt.Println(obj1)
    // stockSpanner.next(100); // return 1
    fmt.Println(obj1.Next(100)) // 1
    fmt.Println(obj1)
    // stockSpanner.next(80);  // return 1
    fmt.Println(obj1.Next(80)) // 1
    fmt.Println(obj1)
    // stockSpanner.next(60);  // return 1
    fmt.Println(obj1.Next(60)) // 1
    fmt.Println(obj1)
    // stockSpanner.next(70);  // return 2
    fmt.Println(obj1.Next(70)) // 2
    fmt.Println(obj1)
    // stockSpanner.next(60);  // return 1
    fmt.Println(obj1.Next(60)) // 1
    fmt.Println(obj1)
    // stockSpanner.next(75);  // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
    fmt.Println(obj1.Next(75)) // 4
    fmt.Println(obj1)
    // stockSpanner.next(85);  // return 6
    fmt.Println(obj1.Next(85)) // 6
    fmt.Println(obj1)
}