package main

// 295. Find Median from Data Stream
// The median is the middle value in an ordered integer list. 
// If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.
//     For example, for arr = [2,3,4], the median is 3.
//     For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.

// Implement the MedianFinder class:
//     MedianFinder() initializes the MedianFinder object.
//     void addNum(int num) adds the integer num from the data stream to the data structure.
//     double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.
    
// Example 1:
// Input
// ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
// [[], [1], [2], [], [3], []]
// Output
// [null, null, null, 1.5, null, 2.0]
// Explanation
// MedianFinder medianFinder = new MedianFinder();
// medianFinder.addNum(1);    // arr = [1]
// medianFinder.addNum(2);    // arr = [1, 2]
// medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
// medianFinder.addNum(3);    // arr[1, 2, 3]
// medianFinder.findMedian(); // return 2.0
 
// Constraints:
//     -10^5 <= num <= 10^5
//     There will be at least one element in the data structure before calling findMedian.
//     At most 5 * 10^4 calls will be made to addNum and findMedian.

// Follow up:
//     If all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?
//     If 99% of all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?

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