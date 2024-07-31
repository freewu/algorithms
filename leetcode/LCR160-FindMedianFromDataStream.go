package main

// LCR 160. 数据流中的中位数
// 中位数 是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。
// 例如，
//     [2,3,4] 的中位数是 3
//     [2,3] 的中位数是 (2 + 3) / 2 = 2.

// 设计一个支持以下两种操作的数据结构：
//     void addNum(int num) - 从数据流中添加一个整数到数据结构中。
//     double findMedian() - 返回目前所有元素的中位数。

// 示例 1：
// 输入：
// ["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
// [[],[1],[2],[],[3],[]]
// 输出：[null,null,null,1.50000,null,2.00000]

// 示例 2：
// 输入：
// ["MedianFinder","addNum","findMedian","addNum","findMedian"]
// [[],[2],[],[3],[]]
// 输出：[null,null,2.00000,null,2.50000]

// 提示：
//     最多会对 addNum、findMedian 进行 50000 次调用。

import "fmt"
import "container/heap"

// type MedianFinder struct {
//     data []int
//     sum int // 保存累加值
// }

// func Constructor() MedianFinder {
//     return MedianFinder{ data: []int{}, sum: 0, }
// }

// func (this *MedianFinder) AddNum(num int)  {
//     this.data = append(this.data,num)
//     this.sum += num
// }

// func (this *MedianFinder) FindMedian() float64 {
//     return float64(this.sum) / float64(len(this.data))
// }

type MinHeap []int
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type MaxHeap []int
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type MedianFinder struct {
    ha *MinHeap //存更大的一半
    hb *MaxHeap //存更小的一半 a的元素个数最多比b多1
}

func Constructor() MedianFinder {
    return MedianFinder{ha: &MinHeap{}, hb: &MaxHeap{}}
}

func (this *MedianFinder) AddNum(num int) {
    //实际上次序不重要，只需保证ha的元素个数最多比hb多1即可
    //注意是调用heap.push，而不是自己写的push，因为heap.push会自动维护堆的性质(up/down)
    if this.ha.Len() == this.hb.Len() {
        //先往大根堆存，再往小根堆存
        heap.Push(this.hb, num)
        heap.Push(this.ha, heap.Pop(this.hb))
    } else {
        //先往小根堆存，再往大根堆存
        heap.Push(this.ha, num)
        heap.Push(this.hb, heap.Pop(this.ha))
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if this.ha.Len() == this.hb.Len() {
        return float64((*this.ha)[0]+(*this.hb)[0]) / 2
    } else {
        return float64((*this.ha)[0])
    }
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
    // MedianFinder medianFinder = new MedianFinder();
    obj := Constructor()
    // medianFinder.addNum(1);    // arr = [1]
    obj.AddNum(1)
    fmt.Println(obj)
    // medianFinder.addNum(2);    // arr = [1, 2]
    obj.AddNum(2)
    fmt.Println(obj)
    // medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
    fmt.Println(obj.FindMedian()) // 1.5
    // medianFinder.addNum(3);    // arr[1, 2, 3]
    obj.AddNum(3)
    fmt.Println(obj)
    // medianFinder.findMedian(); // return 2.0
    fmt.Println(obj.FindMedian()) // 2.0
}