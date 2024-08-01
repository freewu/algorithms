package main

// LCR 184. 设计自助结算系统
// 请设计一个自助结账系统，该系统需要通过一个队列来模拟顾客通过购物车的结算过程，需要实现的功能有：
//     get_max()：获取结算商品中的最高价格，如果队列为空，则返回 -1
//     add(value)：将价格为 value 的商品加入待结算商品队列的尾部
//     remove()：移除第一个待结算的商品价格，如果队列为空，则返回 -1

// 注意，为保证该系统运转高效性，以上函数的均摊时间复杂度均为 O(1)

// 示例 1：
// 输入: 
// ["Checkout","add","add","get_max","remove","get_max"]
// [[],[4],[7],[],[],[]]
// 输出: [null,null,null,7,4,7]

// 示例 2：
// 输入: 
// ["Checkout","remove","get_max"]
// [[],[],[]]
// 输出: [null,-1,-1]

// 提示：
//     1 <= get_max, add, remove 的总操作数 <= 10000
//     1 <= value <= 10^5

import "fmt"

type Checkout struct {
    queue []int
    deque []int // 单调递减，保持第0个始终为最大值
}

func Constructor() Checkout {
    return Checkout{}
}

func (this *Checkout) Get_max() int {
    if len(this.deque) == 0 {
        return -1
    }
    return this.deque[0]
}

func (this *Checkout) Add(value int) {
    this.queue = append(this.queue, value)
    for len(this.deque) > 0 && this.deque[len(this.deque)-1] < value {
        this.deque = this.deque[:len(this.deque)-1]
    }
    this.deque = append(this.deque, value)
}

func (this *Checkout) Remove() int {
    if len(this.queue) == 0 {
        return -1
    }
    x := this.queue[0]
    this.queue = this.queue[1:]
    if x == this.deque[0] {
        this.deque = this.deque[1:]
    }
    return x
}

type Checkout1 struct {
    maxPrice      int
    maxPriceIndex int
    queue         []int
}

func Constructor1() Checkout1 {
    queue := make([]int, 0)
    return Checkout1{
        maxPrice:      -1,
        maxPriceIndex: -1,
        queue:         queue,
    }
}

func (this *Checkout1) Get_max() int {
    if len(this.queue) == 0 {
        return -1
    }
    return this.queue[this.maxPriceIndex]
}

func (this *Checkout1) Add(value int) {
    if value > this.maxPrice {
        this.maxPrice = value
        this.maxPriceIndex = len(this.queue)
    }
    this.queue = append(this.queue, value)
}

func (this *Checkout1) Remove() int {
    if len(this.queue) == 0 {
        return -1
    }
    res := this.queue[0]
    this.queue = this.queue[1:]
    this.maxPriceIndex--
    if this.maxPriceIndex < 0 {
        if len(this.queue) > 0 {
            maxPriceIndex, maxPrice := 0, this.queue[0]
            for i := 1; i < len(this.queue); i++ {
                if this.queue[i] > maxPrice {
                    maxPrice = this.queue[i]
                    maxPriceIndex = i
                }
            }
            this.maxPrice = maxPrice
            this.maxPriceIndex = maxPriceIndex
        } else {
            this.maxPrice = -1
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入: 
    // ["Checkout","add","add","get_max","remove","get_max"]
    // [[],[4],[7],[],[],[]]
    // 输出: [null,null,null,7,4,7]
    obj1 := Constructor()
    obj1.Add(4)
    fmt.Println(obj1)
    obj1.Add(7)
    fmt.Println(obj1)
    fmt.Println(obj1.Get_max()) // 7
    fmt.Println(obj1.Remove()) // 4
    fmt.Println(obj1.Get_max()) // 7
    // 示例 2：
    // 输入: 
    // ["Checkout","remove","get_max"]
    // [[],[],[]]
    // 输出: [null,-1,-1]
    obj2 := Constructor()
    fmt.Println(obj2.Remove()) // -1
    fmt.Println(obj2.Get_max()) // -1

    obj11 := Constructor1()
    obj11.Add(4)
    fmt.Println(obj11)
    obj11.Add(7)
    fmt.Println(obj11)
    fmt.Println(obj11.Get_max()) // 7
    fmt.Println(obj11.Remove()) // 4
    fmt.Println(obj11.Get_max()) // 7
    obj12 := Constructor1()
    fmt.Println(obj12)
    fmt.Println(obj12.Remove()) // -1
    fmt.Println(obj12.Get_max()) // -1
}