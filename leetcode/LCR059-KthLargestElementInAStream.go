package main

// LCR 059. 数据流中的第 K 大元素
// 设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。

// 请实现 KthLargest 类：
//     KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
//     int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。

// 示例：
// 输入：
// ["KthLargest", "add", "add", "add", "add", "add"]
// [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
// 输出：
// [null, 4, 5, 5, 8, 8]
// 解释：
// KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
// kthLargest.add(3);   // return 4
// kthLargest.add(5);   // return 5
// kthLargest.add(10);  // return 5
// kthLargest.add(9);   // return 8
// kthLargest.add(4);   // return 8

// 提示：
//     1 <= k <= 10^4
//     0 <= nums.length <= 10^4
//     -10^4 <= nums[i] <= 10^4
//     -10^4 <= val <= 10^4
//     最多调用 add 方法 10^4 次
//     题目数据保证，在查找第 k 大元素时，数组中至少有 k 个元素

import "fmt"
import "container/heap"

type MinHeap []int
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// real thing starts here
type KthLargest struct {
    data *MinHeap
    k    int
}

func Constructor(k int, nums []int) KthLargest {
    h := &MinHeap{}
    heap.Init(h)
    // push all elements to the heap
    for i := 0; i < len(nums); i++ {
        heap.Push(h, nums[i])
    }
    // remove the smaller elements from the heap such that only the k largest elements are in the heap
    for len(*h) > k {
        heap.Pop(h)
    }
    return KthLargest{h, k}
}

func (this *KthLargest) Add(val int) int {
    if len(*this.data) < this.k {
        heap.Push(this.data, val)
    } else if val > (*this.data)[0] {
        heap.Push(this.data, val)
        heap.Pop(this.data)
    }
    return (*this.data)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {
    // KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
    obj := Constructor(3,[]int{4, 5, 8, 2})
    fmt.Println(obj)
    // kthLargest.add(3);   // return 4
    fmt.Println(obj.Add(3)) // 4
    fmt.Println(obj)
    // kthLargest.add(5);   // return 5
    fmt.Println(obj.Add(5)) // 5
    fmt.Println(obj)
    // kthLargest.add(10);  // return 5
    fmt.Println(obj.Add(10)) // 5
    fmt.Println(obj)
    // kthLargest.add(9);   // return 8
    fmt.Println(obj.Add(9)) // 8
    fmt.Println(obj)
    // kthLargest.add(4);   // return 8
    fmt.Println(obj.Add(4)) // 8
    fmt.Println(obj)
}