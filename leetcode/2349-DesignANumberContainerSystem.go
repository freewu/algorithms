package main

// 2349. Design a Number Container System
// Design a number container system that can do the following:
//     Insert or Replace a number at the given index in the system.
//     Return the smallest index for the given number in the system.

// Implement the NumberContainers class:
//     NumberContainers() 
//         Initializes the number container system.
//     void change(int index, int number) 
//         Fills the container at index with the number. 
//         If there is already a number at that index, replace it.
//     int find(int number) 
//         Returns the smallest index for the given number, 
//         or -1 if there is no index that is filled by number in the system.

// Example 1:
// Input
// ["NumberContainers", "find", "change", "change", "change", "change", "find", "change", "find"]
// [[], [10], [2, 10], [1, 10], [3, 10], [5, 10], [10], [1, 20], [10]]
// Output
// [null, -1, null, null, null, null, 1, null, 2]
// Explanation
// NumberContainers nc = new NumberContainers();
// nc.find(10); // There is no index that is filled with number 10. Therefore, we return -1.
// nc.change(2, 10); // Your container at index 2 will be filled with number 10.
// nc.change(1, 10); // Your container at index 1 will be filled with number 10.
// nc.change(3, 10); // Your container at index 3 will be filled with number 10.
// nc.change(5, 10); // Your container at index 5 will be filled with number 10.
// nc.find(10); // Number 10 is at the indices 1, 2, 3, and 5. Since the smallest index that is filled with 10 is 1, we return 1.
// nc.change(1, 20); // Your container at index 1 will be filled with number 20. Note that index 1 was filled with 10 and then replaced with 20. 
// nc.find(10); // Number 10 is at the indices 2, 3, and 5. The smallest index that is filled with 10 is 2. Therefore, we return 2.

// Constraints:
//     1 <= index, number <= 10^9
//     At most 10^5 calls will be made in total to change and find.

import "fmt"
import "container/heap"

type Entry struct {
    Index       int
    Value       int
    HeapIndex   int
}

type EntryHeap []*Entry

func (this EntryHeap) Len() int { return len(this) }
func (this EntryHeap) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
    this[i].HeapIndex, this[j].HeapIndex = i, j
}
func (this EntryHeap) Less(i, j int) bool {
    if this[i].Value == this[j].Value { return this[i].Index < this[j].Index; }
    return this[i].Value < this[j].Value
}
func (this *EntryHeap) Push(v interface{}) {
    *this = append(*this, v.(*Entry))
}

func (this *EntryHeap) Pop() interface{} {
    n := len(*this)
    v := (*this)[n-1]
    *this = (*this)[:n-1]
    return v
}

type NumberContainers struct {
    Entries map[int]*Entry
    Values  map[int]*EntryHeap
}

func Constructor() NumberContainers {
    return NumberContainers{ make(map[int]*Entry), make(map[int]*EntryHeap)}
}

func (this *NumberContainers) Change(index int, number int) {
    if _, ok := this.Values[number]; !ok {
        this.Values[number] = &EntryHeap{}
    }
    if v, ok := this.Entries[index]; ok {
        heap.Remove(this.Values[v.Value], v.HeapIndex)
        n := this.Values[number].Len()
        this.Entries[index].HeapIndex = n
        this.Entries[index].Value = number
    } else {
        n := this.Values[number].Len()
        this.Entries[index] = &Entry{
            Index:     index,
            Value:     number,
            HeapIndex: n,
        }
    }
    heap.Push(this.Values[number], this.Entries[index])
}

func (this *NumberContainers) Find(number int) int {
    if _, ok := this.Values[number]; !ok || this.Values[number].Len() == 0 {
        return -1
    }
    return (*this.Values[number])[0].Index
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */

func main() {
    // NumberContainers nc = new NumberContainers();
    obj := Constructor()
    fmt.Println(obj)
    // nc.find(10); // There is no index that is filled with number 10. Therefore, we return -1.
    fmt.Println(obj.Find(10)) // -1
    // nc.change(2, 10); // Your container at index 2 will be filled with number 10.
    obj.Change(2, 10)
    fmt.Println(obj)
    // nc.change(1, 10); // Your container at index 1 will be filled with number 10.
    obj.Change(1, 10)
    fmt.Println(obj)
    // nc.change(3, 10); // Your container at index 3 will be filled with number 10.
    obj.Change(3, 10)
    fmt.Println(obj)
    // nc.change(5, 10); // Your container at index 5 will be filled with number 10.
    obj.Change(5, 10)
    fmt.Println(obj)
    // nc.find(10); // Number 10 is at the indices 1, 2, 3, and 5. Since the smallest index that is filled with 10 is 1, we return 1.
    fmt.Println(obj.Find(10)) // 1
    // nc.change(1, 20); // Your container at index 1 will be filled with number 20. Note that index 1 was filled with 10 and then replaced with 20. 
    obj.Change(1, 20)
    fmt.Println(obj)
    // nc.find(10); // Number 10 is at the indices 2, 3, and 5. The smallest index that is filled with 10 is 2. Therefore, we return 2.
    fmt.Println(obj.Find(10)) // 2
}