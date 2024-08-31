package main

// LCR 058. 我的日程安排表 I
// 请实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内没有其他安排，则可以存储这个新的日程安排。

// MyCalendar 有一个 book(int start, int end)方法。它意味着在 start 到 end 时间内增加一个日程安排，注意，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x < end。

// 当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生重复预订。

// 每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true。
// 否则，返回 false 并且不要将该日程安排添加到日历中。

// 请按照以下步骤调用 MyCalendar 类: MyCalendar cal = new MyCalendar(); MyCalendar.book(start, end)

// 示例:
// 输入:
// ["MyCalendar","book","book","book"]
// [[],[10,20],[15,25],[20,30]]
// 输出: [null,true,false,true]
// 解释: 
// MyCalendar myCalendar = new MyCalendar();
// MyCalendar.book(10, 20); // returns true 
// MyCalendar.book(15, 25); // returns false ，第二个日程安排不能添加到日历中，因为时间 15 已经被第一个日程安排预定了
// MyCalendar.book(20, 30); // returns true ，第三个日程安排可以添加到日历中，因为第一个日程安排并不包含时间 20 

// 提示：
//     每个测试用例，调用 MyCalendar.book 函数最多不超过 1000次。
//     0 <= start < end <= 10^9

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