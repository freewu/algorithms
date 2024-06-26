package main

// 346. Moving Average from Data Stream
// Given a stream of integers and a window size, calculate the moving average of all integers in the sliding window.
// Implement the MovingAverage class:
//     MovingAverage(int size) Initializes the object with the size of the window size.
//     double next(int val) Returns the moving average of the last size values of the stream.
    
// Example 1:
// Input
// ["MovingAverage", "next", "next", "next", "next"]
// [[3], [1], [10], [3], [5]]
// Output
// [null, 1.0, 5.5, 4.66667, 6.0]
// Explanation
// MovingAverage movingAverage = new MovingAverage(3);
// movingAverage.next(1); // return 1.0 = 1 / 1
// movingAverage.next(10); // return 5.5 = (1 + 10) / 2
// movingAverage.next(3); // return 4.66667 = (1 + 10 + 3) / 3
// movingAverage.next(5); // return 6.0 = (10 + 3 + 5) / 3
 
// Constraints:
//     1 <= size <= 1000
//     -10^5 <= val <= 10^5
//     At most 10^4 calls will be made to next.

import "fmt"

type MovingAverage struct {
    size int
    data []int
}

func Constructor(size int) MovingAverage {
    return MovingAverage{size, []int{} }
}

func (this *MovingAverage) Next(val int) float64 {
    if len(this.data) < this.size {
        this.data = append(this.data, val)
    } else {
        this.data = append(this.data[1:this.size], val)
    }
    
    sum := 0 
    for _, v := range this.data {
        sum += v
    }
    return float64(sum) / float64(len(this.data))
}

type MovingAverage1 struct {
    nums []int
    size int
    sum  int
}

/** Initialize your data structure here. */
func Constructor1(size int) MovingAverage1 {
    return MovingAverage1{size: size}
}

func (this *MovingAverage1) Next(val int) float64 {
    if len(this.nums) >= this.size {
        this.sum -= this.nums[0]
        this.nums = this.nums[1:]
    }
    this.nums = append(this.nums, val)
    this.sum += val
	return float64(this.sum) / float64(len(this.nums))
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

func main() {
    // MovingAverage movingAverage = new MovingAverage(3);
    obj := Constructor(3)
    fmt.Println(obj)
    // movingAverage.next(1); // return 1.0 = 1 / 1
    fmt.Println(obj.Next(1)) // 1.0
    fmt.Println(obj)
    // movingAverage.next(10); // return 5.5 = (1 + 10) / 2
    fmt.Println(obj.Next(10)) // 5.5
    fmt.Println(obj)
    // movingAverage.next(3); // return 4.66667 = (1 + 10 + 3) / 3
    fmt.Println(obj.Next(3)) // 4.66667
    fmt.Println(obj)
    // movingAverage.next(5); // return 6.0 = (10 + 3 + 5) / 3
    fmt.Println(obj.Next(5)) // 6.0
    fmt.Println(obj)

    obj1 := Constructor(4)
    fmt.Println(obj1)
    // movingAverage.next(1); // return 1.0 = 1 / 1
    fmt.Println(obj1.Next(1)) // 1.0
    fmt.Println(obj1)
    // movingAverage.next(10); // return 5.5 = (1 + 10) / 2
    fmt.Println(obj1.Next(10)) // 5.5
    fmt.Println(obj1)
    // movingAverage.next(3); // return 4.66667 = (1 + 10 + 3) / 3
    fmt.Println(obj1.Next(3)) // 4.66667
    fmt.Println(obj1)
    // movingAverage.next(5); // return 4.75= (1 + 10 + 3 + 5) / 4
    fmt.Println(obj1.Next(5)) // 4.75
    fmt.Println(obj1)

    // movingAverage.next(5); // return 5.75 = (10 + 3 + 5 + 5) / 4
    fmt.Println(obj1.Next(5)) // 5.75
    fmt.Println(obj1)


    // MovingAverage movingAverage = new MovingAverage(3);
    obj11 := Constructor1(3)
    fmt.Println(obj11)
    // movingAverage.next(1); // return 1.0 = 1 / 1
    fmt.Println(obj11.Next(1)) // 1.0
    fmt.Println(obj11)
    // movingAverage.next(10); // return 5.5 = (1 + 10) / 2
    fmt.Println(obj11.Next(10)) // 5.5
    fmt.Println(obj11)
    // movingAverage.next(3); // return 4.66667 = (1 + 10 + 3) / 3
    fmt.Println(obj11.Next(3)) // 4.66667
    fmt.Println(obj11)
    // movingAverage.next(5); // return 6.0 = (10 + 3 + 5) / 3
    fmt.Println(obj11.Next(5)) // 6.0
    fmt.Println(obj11)

    obj12 := Constructor1(4)
    fmt.Println(obj12)
    // movingAverage.next(1); // return 1.0 = 1 / 1
    fmt.Println(obj12.Next(1)) // 1.0
    fmt.Println(obj12)
    // movingAverage.next(10); // return 5.5 = (1 + 10) / 2
    fmt.Println(obj12.Next(10)) // 5.5
    fmt.Println(obj12)
    // movingAverage.next(3); // return 4.66667 = (1 + 10 + 3) / 3
    fmt.Println(obj12.Next(3)) // 4.66667
    fmt.Println(obj12)
    // movingAverage.next(5); // return 4.75= (1 + 10 + 3 + 5) / 4
    fmt.Println(obj12.Next(5)) // 4.75
    fmt.Println(obj12)

    // movingAverage.next(5); // return 5.75 = (10 + 3 + 5 + 5) / 4
    fmt.Println(obj12.Next(5)) // 5.75
    fmt.Println(obj12)
}