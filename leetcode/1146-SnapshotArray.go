package main

// 1146. Snapshot Array
// Implement a SnapshotArray that supports the following interface:
//     SnapshotArray(int length) initializes an array-like data structure with the given length. Initially, each element equals 0.
//     void set(index, val) sets the element at the given index to be equal to val.
//     int snap() takes a snapshot of the array and returns the snap_id: the total number of times we called snap() minus 1.
//     int get(index, snap_id) returns the value at the given index, at the time we took the snapshot with the given snap_id
    
// Example 1:
// Input: ["SnapshotArray","set","snap","set","get"]
// [[3],[0,5],[],[0,6],[0,0]]
// Output: [null,null,0,null,5]
// Explanation: 
// SnapshotArray snapshotArr = new SnapshotArray(3); // set the length to be 3
// snapshotArr.set(0,5);  // Set array[0] = 5
// snapshotArr.snap();  // Take a snapshot, return snap_id = 0
// snapshotArr.set(0,6);
// snapshotArr.get(0,0);  // Get the value of array[0] with snap_id = 0, return 5
 
// Constraints:
//     1 <= length <= 5 * 10^4
//     0 <= index < length
//     0 <= val <= 10^9
//     0 <= snap_id < (the total number of times we call snap())
//     At most 5 * 10^4 calls will be made to set, snap, and get.

import "fmt"

type Elem struct {
    val int
    version int
}

type SnapshotArray struct {
    data [][]Elem
    version int
}

func Constructor(length int) SnapshotArray {
    return SnapshotArray{ make([][]Elem, length), 0 }
}

func (this *SnapshotArray) Set(index int, val int)  {
    // Update value if we are at the same version
    if len(this.data[index]) > 0 {
        arr := this.data[index]
        if arr[len(arr)-1].version == this.version {
            this.data[index][len(arr)-1].val = val
            return
        }
    }
    // Add new record if we are at the higher version
    this.data[index] = append(this.data[index], Elem{val, this.version})
}

func (this *SnapshotArray) Snap() int {
    old := this.version
    this.version++
    return int(old)
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
    // Binary search for the exact snap_id or the first below it
    arr := this.data[index]
    l, r := 0, len(arr) - 1
    for l <= r {
        m := (l + r) / 2
        if arr[m].version == snap_id {
            return arr[m].val
        } else if arr[m].version < snap_id {
            l, r = m + 1, r
        } else {
            l, r = l, m - 1
        }
    }
    if l - 1 <= len(arr) && l - 1 >= 0 {
        return arr[l-1].val
    }
    return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */

func main() {
    // SnapshotArray snapshotArr = new SnapshotArray(3); // set the length to be 3
    obj := Constructor(3)
    // snapshotArr.set(0,5);  // Set array[0] = 5
    obj.Set(0,5)
    fmt.Println(obj)
    // snapshotArr.snap();  // Take a snapshot, return snap_id = 0
    fmt.Println(obj.Snap()) // 0
    // snapshotArr.set(0,6);
    obj.Set(0,6)
    fmt.Println(obj)
    // snapshotArr.get(0,0);  // Get the value of array[0] with snap_id = 0, return 5
    fmt.Println(obj.Get(0,0)) // 5
}