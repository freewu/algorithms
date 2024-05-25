package main

// 359. Logger Rate Limiter
// Design a logger system that receives a stream of messages along with their timestamps. 
// Each unique message should only be printed at most every 10 seconds 
// (i.e. a message printed at timestamp t will prevent other identical messages from being printed until timestamp t + 10).

// All messages will come in chronological order. Several messages may arrive at the same timestamp.

// Implement the Logger class:
//     Logger() 
//         Initializes the logger object.
//     bool shouldPrintMessage(int timestamp, string message) 
//         Returns true if the message should be printed in the given timestamp, otherwise returns false.
    
// Example 1:
// Input
// ["Logger", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage"]
// [[], [1, "foo"], [2, "bar"], [3, "foo"], [8, "bar"], [10, "foo"], [11, "foo"]]
// Output
// [null, true, true, false, false, false, true]
// Explanation
// Logger logger = new Logger();
// logger.shouldPrintMessage(1, "foo");  // return true, next allowed timestamp for "foo" is 1 + 10 = 11
// logger.shouldPrintMessage(2, "bar");  // return true, next allowed timestamp for "bar" is 2 + 10 = 12
// logger.shouldPrintMessage(3, "foo");  // 3 < 11, return false
// logger.shouldPrintMessage(8, "bar");  // 8 < 12, return false
// logger.shouldPrintMessage(10, "foo"); // 10 < 11, return false
// logger.shouldPrintMessage(11, "foo"); // 11 >= 11, return true, next allowed timestamp for "foo" is 11 + 10 = 21
 
// Constraints:
//     0 <= timestamp <= 10^9
//     Every timestamp will be passed in non-decreasing order (chronological order).
//     1 <= message.length <= 30
//     At most 10^4 calls will be made to shouldPrintMessage.

import "fmt"

type Logger struct {
    data map[string]int
}

func Constructor() Logger {
    m := make(map[string]int)
    return Logger{data: m}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    if v, ok := this.data[message]; ok {
        if timestamp < v + 10 {
            return false
        }
    }
    this.data[message] = timestamp
    return true
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */

func main() {
// Logger logger = new Logger();
    obj := Constructor()
    fmt.Println(obj)
    // logger.shouldPrintMessage(1, "foo");  // return true, next allowed timestamp for "foo" is 1 + 10 = 11
    fmt.Println(obj.ShouldPrintMessage(1, "foo")) // true
    // logger.shouldPrintMessage(2, "bar");  // return true, next allowed timestamp for "bar" is 2 + 10 = 12
    fmt.Println(obj.ShouldPrintMessage(2, "bar")) // true
    // logger.shouldPrintMessage(3, "foo");  // 3 < 11, return false
    fmt.Println(obj.ShouldPrintMessage(3, "foo")) // false
    // logger.shouldPrintMessage(8, "bar");  // 8 < 12, return false
    fmt.Println(obj.ShouldPrintMessage(8, "bar")) // false
    // logger.shouldPrintMessage(10, "foo"); // 10 < 11, return false
    fmt.Println(obj.ShouldPrintMessage(10, "foo")) // false
    // logger.shouldPrintMessage(11, "foo"); // 11 >= 11, return true, next allowed timestamp for "foo" is 11 + 10 = 21
    fmt.Println(obj.ShouldPrintMessage(11, "foo")) // true
}