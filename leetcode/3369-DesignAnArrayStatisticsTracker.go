package main

// 3369. Design an Array Statistics Tracker
// Design a data structure that keeps track of the values in it and answers some queries regarding their mean, median, and mode.

// Implement the StatisticsTracker class.
//     StatisticsTracker(): 
//         Initialize the StatisticsTracker object with an empty array.
//     void addNumber(int number): 
//         Add number to the data structure.
//     void removeFirstAddedNumber(): 
//         Remove the earliest added number from the data structure.
//     int getMean(): 
//         Return the floored mean of the numbers in the data structure.
//     int getMedian(): 
//         Return the median of the numbers in the data structure.
//     int getMode(): 
//         Return the mode of the numbers in the data structure. 
//         If there are multiple modes, return the smallest one.

// Note:
//     The mean of an array is the sum of all the values divided by the number of values in the array.
//     The median of an array is the middle element of the array when it is sorted in non-decreasing order. If there are two choices for a median, the larger of the two values is taken.
//     The mode of an array is the element that appears most often in the array.

// Example 1:
// Input:
// ["StatisticsTracker", "addNumber", "addNumber", "addNumber", "addNumber", "getMean", "getMedian", "getMode", "removeFirstAddedNumber", "getMode"]
// [[], [4], [4], [2], [3], [], [], [], [], []]
// Output:
// [null, null, null, null, null, 3, 4, 4, null, 2]
// Explanation
// StatisticsTracker statisticsTracker = new StatisticsTracker();
// statisticsTracker.addNumber(4); // The data structure now contains [4]
// statisticsTracker.addNumber(4); // The data structure now contains [4, 4]
// statisticsTracker.addNumber(2); // The data structure now contains [4, 4, 2]
// statisticsTracker.addNumber(3); // The data structure now contains [4, 4, 2, 3]
// statisticsTracker.getMean(); // return 3
// statisticsTracker.getMedian(); // return 4
// statisticsTracker.getMode(); // return 4
// statisticsTracker.removeFirstAddedNumber(); // The data structure now contains [4, 2, 3]
// statisticsTracker.getMode(); // return 2

// Example 2:
// Input:
// ["StatisticsTracker", "addNumber", "addNumber", "getMean", "removeFirstAddedNumber", "addNumber", "addNumber", "removeFirstAddedNumber", "getMedian", "addNumber", "getMode"]
// [[], [9], [5], [], [], [5], [6], [], [], [8], []]
// Output:
// [null, null, null, 7, null, null, null, null, 6, null, 5]
// Explanation
// StatisticsTracker statisticsTracker = new StatisticsTracker();
// statisticsTracker.addNumber(9); // The data structure now contains [9]
// statisticsTracker.addNumber(5); // The data structure now contains [9, 5]
// statisticsTracker.getMean(); // return 7
// statisticsTracker.removeFirstAddedNumber(); // The data structure now contains [5]
// statisticsTracker.addNumber(5); // The data structure now contains [5, 5]
// statisticsTracker.addNumber(6); // The data structure now contains [5, 5, 6]
// statisticsTracker.removeFirstAddedNumber(); // The data structure now contains [5, 6]
// statisticsTracker.getMedian(); // return 6
// statisticsTracker.addNumber(8); // The data structure now contains [5, 6, 8]
// statisticsTracker.getMode(); // return 5

// Constraints:
//     1 <= number <= 10^9
//     At most, 10^5 calls will be made to addNumber, removeFirstAddedNumber, getMean, getMedian, and getMode in total.
//     removeFirstAddedNumber, getMean, getMedian, and getMode will be called only if there is at least one element in the data structure.

import "fmt"
// import "container/list"
// import "sort"

// // 超出时间限制 1242 / 1246
// type StatisticsTracker struct {
//     q    *list.List
//     s    int
//     cnt  map[int]int
//     nums []int
// }

// func Constructor() StatisticsTracker {
//     return StatisticsTracker{
//         q:    list.New(),
//         s:    0,
//         cnt:  make(map[int]int),
//         nums: make([]int, 0),
//     }
// }

// func (this *StatisticsTracker) AddNumber(number int) {
//     this.q.PushBack(number)
//     this.s += number
//     this.cnt[number]++
//     this.nums = append(this.nums, number)
//     sort.Ints(this.nums)
// }

// func (this *StatisticsTracker) RemoveFirstAddedNumber() {
//     if this.q.Len() == 0 { return }
//     number := this.q.Remove(this.q.Front()).(int)
//     this.s -= number
//     this.cnt[number]--
//     index := sort.SearchInts(this.nums, number)
//     this.nums = append(this.nums[:index], this.nums[index+1:]...)
// }

// func (this *StatisticsTracker) GetMean() int {
//     if this.q.Len() == 0 { return 0 }
//     return this.s / this.q.Len()
// }

// func (this *StatisticsTracker) GetMedian() int {
//     if this.q.Len() == 0 { return 0 }
//     return this.nums[len(this.nums)/2]
// }

// func (this *StatisticsTracker) GetMode() int {
//     if this.q.Len() == 0 { return 0 }
//     mx, mode := 0, 0
//     for num, cnt := range this.cnt {
//         if cnt > mx || (cnt == mx && num < mode) {
//             mx, mode = cnt, num
//         }
//     }
//     return mode
// }

import "container/heap"
import "container/list"
import "sort"

// MedianFinder 用于维护中位数
type MedianFinder struct {
    small *MaxHeap // 大根堆，存储较小的一半
    large *MinHeap // 小根堆，存储较大的一半
}

func NewMedianFinder() *MedianFinder {
    return &MedianFinder{
        small: &MaxHeap{},
        large: &MinHeap{},
    }
}

func (mf *MedianFinder) AddNum(num int) {
    if mf.small.Len() == 0 || num <= mf.small.Peek() {
        heap.Push(mf.small, num)
    } else {
        heap.Push(mf.large, num)
    }
    mf.Rebalance()
}

func (mf *MedianFinder) RemoveNum(num int) {
    if num <= mf.small.Peek() {
        mf.small.Remove(num)
    } else {
        mf.large.Remove(num)
    }
    mf.Rebalance()
}

func (mf *MedianFinder) FindMedian() int {
    if mf.small.Len() > mf.large.Len() {
        return mf.small.Peek()
    }
    return mf.large.Peek()
}

func (mf *MedianFinder) Rebalance() {
    if mf.small.Len() > mf.large.Len()+1 {
        heap.Push(mf.large, heap.Pop(mf.small))
    } else if mf.small.Len() < mf.large.Len() {
        heap.Push(mf.small, heap.Pop(mf.large))
    }
}

// MaxHeap 实现大根堆
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
    *h = old[:n-1]
    return x
}
func (h *MaxHeap) Peek() int {
    return (*h)[0]
}
func (h *MaxHeap) Remove(num int) {
    for i := 0; i < len(*h); i++ {
        if (*h)[i] == num {
            heap.Remove(h, i)
            break
        }
    }
}

// MinHeap 实现小根堆
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
    *h = old[:n-1]
    return x
}
func (h *MinHeap) Peek() int {
    return (*h)[0]
}
func (h *MinHeap) Remove(num int) {
    for i := 0; i < len(*h); i++ {
        if (*h)[i] == num {
            heap.Remove(h, i)
            break
        }
    }
}

// StatisticsTracker 实现统计功能
type StatisticsTracker struct {
    q            *list.List
    s            int
    cnt          map[int]int
    medianFinder *MedianFinder
    ts           *SortedSet
}

func Constructor() StatisticsTracker {
    return StatisticsTracker{
        q:            list.New(),
        s:            0,
        cnt:          make(map[int]int),
        medianFinder: NewMedianFinder(),
        ts:           NewSortedSet(),
    }
}

func (st *StatisticsTracker) AddNumber(number int) {
    st.q.PushBack(number)
    st.s += number
    st.ts.Remove([]int{-st.cnt[number], number})
    st.cnt[number]++
    st.medianFinder.AddNum(number)
    st.ts.Add([]int{-st.cnt[number], number})
}

func (st *StatisticsTracker) RemoveFirstAddedNumber() {
    if st.q.Len() == 0 { return }
    number := st.q.Remove(st.q.Front()).(int)
    st.s -= number
    st.ts.Remove([]int{-st.cnt[number], number})
    st.cnt[number]--
    if st.cnt[number] > 0 {
        st.ts.Add([]int{-st.cnt[number], number})
    }
    st.medianFinder.RemoveNum(number)
}

func (st *StatisticsTracker) GetMean() int {
    if st.q.Len() == 0 { return 0 }
    return st.s / st.q.Len()
}

func (st *StatisticsTracker) GetMedian() int {
    return st.medianFinder.FindMedian()
}

func (st *StatisticsTracker) GetMode() int {
    if st.ts.Len() == 0 { return 0 }
    return st.ts.Get(0)[1]
}

// SortedSet 实现有序集合
type SortedSet struct {
    data [][]int
}

func NewSortedSet() *SortedSet {
    return &SortedSet{data: make([][]int, 0)}
}

func (ss *SortedSet) Add(item []int) {
    index := sort.Search(len(ss.data), func(i int) bool {
        return ss.data[i][0] > item[0] || (ss.data[i][0] == item[0] && ss.data[i][1] >= item[1])
    })
    if index < len(ss.data) && ss.data[index][0] == item[0] && ss.data[index][1] == item[1] {
        return
    }
    ss.data = append(ss.data, item)
    sort.Slice(ss.data, func(i, j int) bool {
        return ss.data[i][0] < ss.data[j][0] || (ss.data[i][0] == ss.data[j][0] && ss.data[i][1] < ss.data[j][1])
    })
}

func (ss *SortedSet) Remove(item []int) {
    index := sort.Search(len(ss.data), func(i int) bool {
        return ss.data[i][0] > item[0] || (ss.data[i][0] == item[0] && ss.data[i][1] >= item[1])
    })
    if index < len(ss.data) && ss.data[index][0] == item[0] && ss.data[index][1] == item[1] {
        ss.data = append(ss.data[:index], ss.data[index+1:]...)
    }
}

func (ss *SortedSet) Len() int {
    return len(ss.data)
}

func (ss *SortedSet) Get(index int) []int {
    return ss.data[index]
}

/**
 * Your StatisticsTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNumber(number);
 * obj.RemoveFirstAddedNumber();
 * param_3 := obj.GetMean();
 * param_4 := obj.GetMedian();
 * param_5 := obj.GetMode();
 */

// 太难了 用 python 来 close 这个问题吧
// from sortedcontainers import SortedList

// class StatisticsTracker:
//     def __init__(self):
//         self.q = deque()
//         self.s = 0
//         self.cnt = defaultdict(int)
//         self.sl = SortedList()
//         self.sl2 = SortedList(key=lambda x: (-x[1], x[0]))

//     def addNumber(self, number: int) -> None:
//         self.q.append(number)
//         self.sl.add(number)
//         self.sl2.discard((number, self.cnt[number]))
//         self.cnt[number] += 1
//         self.sl2.add((number, self.cnt[number]))
//         self.s += number

//     def removeFirstAddedNumber(self) -> None:
//         number = self.q.popleft()
//         self.sl.remove(number)
//         self.sl2.discard((number, self.cnt[number]))
//         self.cnt[number] -= 1
//         self.sl2.add((number, self.cnt[number]))
//         self.s -= number

//     def getMean(self) -> int:
//         return self.s // len(self.q)

//     def getMedian(self) -> int:
//         return self.sl[len(self.q) // 2]

//     def getMode(self) -> int:
//         return self.sl2[0][0]

func main() {
    // Example 1:
    // StatisticsTracker statisticsTracker = new StatisticsTracker();
    obj1 := Constructor()
    fmt.Println(obj1)
    // statisticsTracker.addNumber(4); // The data structure now contains [4]
    obj1.AddNumber(4)
    fmt.Println(obj1)
    // statisticsTracker.addNumber(4); // The data structure now contains [4, 4]
    obj1.AddNumber(4)
    fmt.Println(obj1)
    // statisticsTracker.addNumber(2); // The data structure now contains [4, 4, 2]
    obj1.AddNumber(2)
    fmt.Println(obj1)
    // statisticsTracker.addNumber(3); // The data structure now contains [4, 4, 2, 3]
    obj1.AddNumber(3)
    fmt.Println(obj1)
    // statisticsTracker.getMean(); // return 3
    fmt.Println(obj1.GetMean()) // 3
    // statisticsTracker.getMedian(); // return 4
    fmt.Println(obj1.GetMedian()) // 4
    // statisticsTracker.getMode(); // return 4
    fmt.Println(obj1.GetMode()) // 4
    // statisticsTracker.removeFirstAddedNumber(); // The data structure now contains [4, 2, 3]
    obj1.RemoveFirstAddedNumber()
    fmt.Println(obj1) // 4
    // statisticsTracker.getMode(); // return 2
    fmt.Println(obj1.GetMode()) // 2

    // Example 2:
    // StatisticsTracker statisticsTracker = new StatisticsTracker();
    obj2 := Constructor()
    fmt.Println(obj2)
    // statisticsTracker.addNumber(9); // The data structure now contains [9]
    obj2.AddNumber(9)
    fmt.Println(obj2)
    // statisticsTracker.addNumber(5); // The data structure now contains [9, 5]
    obj2.AddNumber(5)
    fmt.Println(obj2)
    // statisticsTracker.getMean(); // return 7
    fmt.Println(obj2.GetMean()) // 7
    // statisticsTracker.removeFirstAddedNumber(); // The data structure now contains [5]
    obj2.RemoveFirstAddedNumber()
    fmt.Println(obj2)
    // statisticsTracker.addNumber(5); // The data structure now contains [5, 5]
    obj2.AddNumber(5)
    fmt.Println(obj2)
    // statisticsTracker.addNumber(6); // The data structure now contains [5, 5, 6]
    obj2.AddNumber(6)
    fmt.Println(obj2)
    // statisticsTracker.removeFirstAddedNumber(); // The data structure now contains [5, 6]
    obj2.RemoveFirstAddedNumber()
    fmt.Println(obj2)
    // statisticsTracker.getMedian(); // return 6
    fmt.Println(obj2.GetMedian()) // 6
    // statisticsTracker.addNumber(8); // The data structure now contains [5, 6, 8]
    obj2.AddNumber(8)
    fmt.Println(obj2)
    // statisticsTracker.getMode(); // return 5
    fmt.Println(obj2.GetMode()) // 5
}