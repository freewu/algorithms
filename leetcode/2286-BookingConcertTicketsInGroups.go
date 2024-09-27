package main

// 2286. Booking Concert Tickets in Groups
// A concert hall has n rows numbered from 0 to n - 1, each with m seats, numbered from 0 to m - 1. 
// You need to design a ticketing system that can allocate seats in the following cases:
//     1. If a group of k spectators can sit together in a row.
//     2. If every member of a group of k spectators can get a seat. They may or may not sit together.

// Note that the spectators are very picky. Hence:
//     1. They will book seats only if each member of their group can get a seat with row number less than or equal to maxRow.
//        maxRow can vary from group to group.
//     2. In case there are multiple rows to choose from, the row with the smallest number is chosen. 
//        If there are multiple seats to choose in the same row, the seat with the smallest number is chosen.

// Implement the BookMyShow class:
//     BookMyShow(int n, int m) 
//         Initializes the object with n as number of rows and m as number of seats per row.
//     int[] gather(int k, int maxRow) 
//         Returns an array of length 2 denoting the row and seat number (respectively) of the first seat being allocated to the k members of the group, who must sit together. 
//         In other words, it returns the smallest possible r and c such that all [c, c + k - 1] seats are valid and empty in row r, and r <= maxRow. 
//         Returns [] in case it is not possible to allocate seats to the group.
//     boolean scatter(int k, int maxRow) 
//         Returns true if all k members of the group can be allocated seats in rows 0 to maxRow, who may or may not sit together. 
//         If the seats can be allocated, it allocates k seats to the group with the smallest row numbers, and the smallest possible seat numbers in each row. 
//         Otherwise, returns false.

// Example 1:
// Input
// ["BookMyShow", "gather", "gather", "scatter", "scatter"]
// [[2, 5], [4, 0], [2, 0], [5, 1], [5, 1]]
// Output
// [null, [0, 0], [], true, false]
// Explanation
// BookMyShow bms = new BookMyShow(2, 5); // There are 2 rows with 5 seats each 
// bms.gather(4, 0); // return [0, 0]
//                   // The group books seats [0, 3] of row 0. 
// bms.gather(2, 0); // return []
//                   // There is only 1 seat left in row 0,
//                   // so it is not possible to book 2 consecutive seats. 
// bms.scatter(5, 1); // return True
//                    // The group books seat 4 of row 0 and seats [0, 3] of row 1. 
// bms.scatter(5, 1); // return False
//                    // There is only one seat left in the hall.

// Constraints:
//     1 <= n <= 5 * 10^4
//     1 <= m, k <= 10^9
//     0 <= maxRow <= n - 1
//     At most 5 * 10^4 calls in total will be made to gather and scatter.

import "fmt"

// type SegTree struct {
//     sum int64
//     posMax int
//     l, r int
//     left, right *SegTree
// }

// func fix(v *SegTree, remain []int64) {
//     v.sum = v.left.sum + v.right.sum
//     if remain[v.left.posMax] >= remain[v.right.posMax] {
//         v.posMax = v.left.posMax
//     } else {
//         v.posMax = v.right.posMax
//     }
// }

// func build(remain []int64, l, r int) *SegTree {
//     var v SegTree
//     v.l = l
//     v.r = r
//     if l == r {
//         v.sum = remain[l]
//         v.posMax = l
//         return &v
//     }
//     mid := (l+r)/2
//     v.left = build(remain, l, mid)
//     v.right = build(remain, mid+1, r)
//     fix(&v, remain)
//     return &v
// }

// func update(v *SegTree, remain []int64, i int) {
//     if v.l == v.r {
//         v.sum = remain[i]
//         return
//     }
//     mid := (v.l + v.r)/2
//     if i <= mid {
//         update(v.left, remain, i)
//     } else {
//         update(v.right, remain, i)
//     }
//     fix(v, remain)
// }

// func getGather(v *SegTree, remain []int64, l, r, k int) (int, bool) {
//     if remain[v.posMax] < int64(k) {
//         return 0, false
//     }
//     if v.l == v.r {
//         return v.posMax, true
//     }
//     mid := (v.l + v.r)/2
//     if r <= mid {
//         return getGather(v.left, remain, l, r, k)
//     }
//     if l >= mid+1 {
//         return getGather(v.right, remain, l, r, k)
//     }
//     left, okLeft := getGather(v.left, remain, l, mid, k)
//     if !okLeft {
//         return getGather(v.right, remain, mid+1, r, k)
//     }
//     return left, okLeft
// }

// func getScatter(v *SegTree, l, r int) int64 {
//     if v.l == l && v.r == r {
//         return v.sum
//     }
//     mid := (v.l + v.r)/2
//     if r <= mid {
//         return getScatter(v.left, l, r)
//     }
//     if l >= mid+1 {
//         return getScatter(v.right, l, r)
//     }
//     return getScatter(v.left, l, mid) + getScatter(v.right, mid+1, r)
// }

// type BookMyShow struct {
//     m int64
//     remain []int64
//     root *SegTree
// }

// func Constructor(n int, m int) BookMyShow {
//     remain := make([]int64, n)
//     for i := 0; i < n; i++ {
//         remain[i] = int64(m)
//     }
//     return BookMyShow{int64(m), remain, build(remain, 0, n-1)}
// }


// func (this *BookMyShow) Gather(k int, maxRow int) []int {
//     i, ok := getGather(this.root, this.remain, 0, maxRow, k)
//     if !ok {
//         return []int{}
//     }
//     res := []int{i, int(this.m - this.remain[i])}
//     this.remain[i] -= int64(k)
//     update(this.root, this.remain, i)
//     return res
// }

// func min(x, y int64) int64 {
//     if x < y {
//         return x
//     }
//     return y
// }

// func (this *BookMyShow) Scatter(k int, maxRow int) bool {
//     got := getScatter(this.root, 0, maxRow)
//     if got < int64(k) {
//         return false
//     }
//     kk := int64(k)
//     for i := 0; i <= maxRow; i++ {
//         d := min(kk, this.remain[i])
//         this.remain[i] -= d
//         kk -= d
//         update(this.root, this.remain, i)
//     }
//     return true
// }

type BookMyShow struct {
    rows, seats int
    stree       []SegNode
}

type SegNode struct {
    max int
    sum int
}

func Constructor(rows int, seats int) BookMyShow {
    size := 1
    for size < rows*2 {
        size = size * 2
    }
    obj := BookMyShow{rows, seats, make([]SegNode, size)}
    obj.build(0, 0, rows-1)
    return obj
}

func (this *BookMyShow) build(idx, ldx, rdx int) {
    if ldx == rdx {
        this.stree[idx] = SegNode{this.seats, this.seats}
        return
    }
    mdx := (ldx + rdx) / 2
    this.stree[idx] = SegNode{this.seats, (rdx - ldx + 1) * this.seats}
    this.build(2*idx+1, ldx, mdx)
    this.build(2*idx+2, mdx+1, rdx)
}

func (this *BookMyShow) queryMax(idx, ldx, rdx, k, maxRow int) []int {
    if ldx > maxRow { return []int{} }
    if this.stree[idx].max < k { return []int{} }
    if ldx == rdx { return []int{ldx, this.seats - this.stree[idx].max} }
    mdx := (ldx + rdx) / 2
    result := this.queryMax(2*idx+1, ldx, mdx, k, maxRow)
    if 0 != len(result) { return result }
    return this.queryMax(2*idx+2, mdx+1, rdx, k, maxRow)
}

func (this *BookMyShow) decreaseMax(idx, ldx, rdx, row, diff int) {
    if ldx > row || rdx < row { return }
    if ldx == rdx {
        this.stree[idx].max -= diff
        this.stree[idx].sum -= diff
        return
    }
    mdx := (ldx + rdx) / 2
    this.stree[idx].sum -= diff
    this.decreaseMax(2*idx+1, ldx, mdx, row, diff)
    this.decreaseMax(2*idx+2, mdx+1, rdx, row, diff)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    this.stree[idx].max = max(this.stree[2*idx+1].max, this.stree[2*idx+2].max)
}

func (this *BookMyShow) querySum(idx, ldx, rdx, maxRow int) int {
    if ldx > maxRow { return 0 }
    if rdx <= maxRow { return this.stree[idx].sum }
    mdx := (ldx + rdx) / 2
    return this.querySum(2*idx+1, ldx, mdx, maxRow) + this.querySum(2*idx+2, mdx+1, rdx, maxRow)
}

func (this *BookMyShow) decreaseSum(idx, ldx, rdx, diff, maxRow int) {
    if ldx > maxRow { return }
    if ldx == rdx {
        this.stree[idx].max -= diff
        this.stree[idx].sum -= diff
        return
    }
    mdx := (ldx + rdx) / 2
    this.stree[idx].sum -= diff
    if mdx+1 > maxRow || this.stree[2*idx+1].sum >= diff {
        this.decreaseSum(2*idx+1, ldx, mdx, diff, maxRow)
    } else {
        diff -= this.stree[2*idx+1].sum
        this.decreaseSum(2*idx+1, ldx, mdx, this.stree[2*idx+1].sum, maxRow)
        this.decreaseSum(2*idx+2, mdx+1, rdx, diff, maxRow)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    this.stree[idx].max = max(this.stree[2*idx+1].max, this.stree[2*idx+2].max)
}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
    result := this.queryMax(0, 0, this.rows-1, k, maxRow)
    if 0 != len(result) {
        this.decreaseMax(0, 0, this.rows-1, result[0], k)
    }
    return result
}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
    cnt := this.querySum(0, 0, this.rows-1, maxRow)
    result := cnt >= k
    if result {
        this.decreaseSum(0, 0, this.rows-1, k, maxRow)
    }
    return result
}

/**
 * Your BookMyShow object will be instantiated and called as such:
 * obj := Constructor(n, m);
 * param_1 := obj.Gather(k,maxRow);
 * param_2 := obj.Scatter(k,maxRow);
 */

func main() {
    // BookMyShow bms = new BookMyShow(2, 5); // There are 2 rows with 5 seats each 
    obj := Constructor(2,5)
    // bms.gather(4, 0); // return [0, 0]
    //                   // The group books seats [0, 3] of row 0. 
    fmt.Println(obj.Gather(4, 0)) // [0, 0]
    // bms.gather(2, 0); // return []
    //                   // There is only 1 seat left in row 0,
    //                   // so it is not possible to book 2 consecutive seats. 
    fmt.Println(obj.Gather(2, 0)) // []
    // bms.scatter(5, 1); // return True
    //                    // The group books seat 4 of row 0 and seats [0, 3] of row 1. 
    fmt.Println(obj.Scatter(5, 1)) // True
    // bms.scatter(5, 1); // return False
    //                    // There is only one seat left in the hall.
    fmt.Println(obj.Scatter(5, 1)) // False
}