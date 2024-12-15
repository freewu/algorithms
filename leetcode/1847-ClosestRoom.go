package main

// 1847. Closest Room
// There is a hotel with n rooms. 
// The rooms are represented by a 2D integer array rooms where rooms[i] = [roomIdi, sizei] denotes 
// that there is a room with room number roomIdi and size equal to sizei. 
// Each roomIdi is guaranteed to be unique.

// You are also given k queries in a 2D array queries where queries[j] = [preferredj, minSizej]. 
// The answer to the jth query is the room number id of a room such that:
//     The room has a size of at least minSizej, and
//     abs(id - preferredj) is minimized, where abs(x) is the absolute value of x.

// If there is a tie in the absolute difference, then use the room with the smallest such id. 
// If there is no such room, the answer is -1.

// Return an array answer of length k where answer[j] contains the answer to the jth query.

// Example 1:
// Input: rooms = [[2,2],[1,2],[3,2]], queries = [[3,1],[3,3],[5,2]]
// Output: [3,-1,3]
// Explanation: The answers to the queries are as follows:
// Query = [3,1]: Room number 3 is the closest as abs(3 - 3) = 0, and its size of 2 is at least 1. The answer is 3.
// Query = [3,3]: There are no rooms with a size of at least 3, so the answer is -1.
// Query = [5,2]: Room number 3 is the closest as abs(3 - 5) = 2, and its size of 2 is at least 2. The answer is 3.

// Example 2:
// Input: rooms = [[1,4],[2,3],[3,5],[4,1],[5,2]], queries = [[2,3],[2,4],[2,5]]
// Output: [2,1,3]
// Explanation: The answers to the queries are as follows:
// Query = [2,3]: Room number 2 is the closest as abs(2 - 2) = 0, and its size of 3 is at least 3. The answer is 2.
// Query = [2,4]: Room numbers 1 and 3 both have sizes of at least 4. The answer is 1 since it is smaller.
// Query = [2,5]: Room number 3 is the only room with a size of at least 5. The answer is 3.

// Constraints:
//     n == rooms.length
//     1 <= n <= 10^5
//     k == queries.length
//     1 <= k <= 10^4
//     1 <= roomIdi, preferredj <= 10^7
//     1 <= sizei, minSizej <= 10^7

import "fmt"
import "sort"
import "math/rand"
import "slices"

type node struct {
    Ch       [2]*node
    priority int
    Val      int
}

func (o *node) cmp(b int) int {
    switch {
    case b < o.Val:
        return 0
    case b > o.Val:
        return 1
    default:
        return -1
    }
}

func (o *node) rotate(d int) *node {
    x := o.Ch[d^1]
    o.Ch[d^1] = x.Ch[d]
    x.Ch[d] = o
    return x
}

type Treap struct {
    root *node
}

func (t *Treap) _put(o *node, val int) *node {
    if o == nil {
        return &node{priority: rand.Int(), Val: val}
    }
    d := o.cmp(val)
    o.Ch[d] = t._put(o.Ch[d], val)
    if o.Ch[d].priority > o.priority {
        o = o.rotate(d ^ 1)
    }
    return o
}

func (t *Treap) Put(val int) {
    t.root = t._put(t.root, val)
}

func (t *Treap) _delete(o *node, val int) *node {
    if d := o.cmp(val); d >= 0 {
        o.Ch[d] = t._delete(o.Ch[d], val)
        return o
    }
    if o.Ch[1] == nil {
        return o.Ch[0]
    }
    if o.Ch[0] == nil {
        return o.Ch[1]
    }
    d := 0
    if o.Ch[0].priority > o.Ch[1].priority {
        d = 1
    }
    o = o.rotate(d)
    o.Ch[d] = t._delete(o.Ch[d], val)
    return o
}

func (t *Treap) Delete(val int) {
    t.root = t._delete(t.root, val)
}

func (t *Treap) LowerBound(val int) (lb *node) {
    for o := t.root; o != nil; {
        switch c := o.cmp(val); {
        case c == 0:
            lb = o
            o = o.Ch[0]
        case c > 0:
            o = o.Ch[1]
        default:
            return o
        }
    }
    return
}
func (t *Treap) HighBound(val int) (lb *node) {
    for o := t.root; o != nil; {
        switch c := o.cmp(val); {
        case c == 0:
            o = o.Ch[0]
        case c > 0:
            lb = o
            o = o.Ch[1]
        default:
            return o
        }
    }
    return
}

// 平衡二叉树
func closestRoom(rooms [][]int, queries [][]int) []int {
    res, q := make([]int, len(queries)), make([][]int, len(queries))
    set := &Treap{}
    for i := 0; i < len(queries); i++ {
        q[i] = append(queries[i], i)
    }
    sort.Slice(rooms, func(i, j int) bool { return rooms[i][1] > rooms[j][1] })
    sort.Slice(q, func(i, j int) bool {return q[i][1] > q[j][1] })
    j := 0
    for i := 0; i < len(q); i++ {
        for j< len(rooms) && rooms[j][1] >= q[i][1] {
            set.Put(rooms[j][0])
            j++
        }
        if j == 0 {
            res[q[i][2]] = -1
            continue
        }
        hb := set.LowerBound(q[i][0])
        if hb != nil && hb.Val == q[i][0] {
            res[q[i][2]] = hb.Val
            continue
        }
        lb := set.HighBound(q[i][0])
        if lb == nil || (hb != nil && q[i][0] - lb.Val > hb.Val - q[i][0]) {
            res[q[i][2]] = hb.Val
        } else {
            res[q[i][2]] = lb.Val
        }
    }
    return res
}

func closestRoom1(rooms [][]int, queries [][]int) []int {
    sort.Slice(rooms, func(i, j int) bool { // Sort rooms by decreasing size
        return rooms[i][1] > rooms[j][1]
    })
    type Query struct { // Prepare queries with their original indices
        Preferred, MinSize, Index int
    }
    queryList := make([]Query, len(queries))
    for i, q := range queries {
        queryList[i] = Query{ q[0], q[1], i }
    }
    sort.Slice(queryList, func(i, j int) bool { // Sort queries by decreasing MinSize
        return queryList[i].MinSize > queryList[j].MinSize
    })
    res, availableRooms := make([]int, len(queries)), []int{} // Available rooms IDs in sorted order
    j := 0 // Index for rooms
    findClosest := func(availableRooms []int, preferred int) int { // binary search to find the closest room ID
        index := sort.Search(len(availableRooms), func(i int) bool { // Perform binary search to find the closest room to the preferred room
            return availableRooms[i] >= preferred
        })
        res := -1
        if index < len(availableRooms) {
            res = availableRooms[index]
        }
        if index > 0 && (res == -1 || preferred - availableRooms[index - 1] <= res - preferred) {
            res = availableRooms[index - 1]
        }
        return res
    }
    for _, q := range queryList { // Iterate over queries and rooms
        
        for j < len(rooms) && rooms[j][1] >= q.MinSize { // Add rooms that meet the minimum size requirement for the current query
            availableRooms = append(availableRooms, rooms[j][0]) // Add the room ID and keep it sorted
            sort.Ints(availableRooms)
            j++
        }
        if len(availableRooms) > 0 { // Find the closest room ID to the preferred room
            res[q.Index] = findClosest(availableRooms, q.Preferred)
        } else {
            res[q.Index] = -1 // No room meets the size requirement
        }
    }
    return res
}

// import "github.com/emirpasic/gods/v2/trees/redblacktree"
// import "slices"

// func closestRoom(rooms [][]int, queries [][]int) []int {
//     slices.SortFunc(rooms, func(a, b []int) int { return b[1] - a[1] })  // 按照 size 从大到小排序
//     n := len(queries)
//     queryIds := make([]int, n)
//     for i := range queryIds {
//         queryIds[i] = i
//     }
//     slices.SortFunc(queryIds, func(i, j int) int { return queries[j][1] - queries[i][1] }) // 按照 minSize 从大到小排序
//     res := make([]int, n)
//     for i := range res {
//         res[i] = -1
//     }
//     roomIds := redblacktree.New[int, struct{}]() // import "github.com/emirpasic/gods/v2/trees/redblacktree"
//     j := 0
//     for _, i := range queryIds {
//         preferredId, minSize := queries[i][0], queries[i][1]
//         for j < len(rooms) && rooms[j][1] >= minSize {
//             roomIds.Put(rooms[j][0], struct{}{})
//             j++
//         }
//         diff := math.MaxInt
//         // 左边的差
//         if node, ok := roomIds.Floor(preferredId); ok {
//             diff = preferredId - node.Key
//             res[i] = node.Key
//         }
//         // 右边的差
//         if node, ok := roomIds.Ceiling(preferredId); ok && node.Key-preferredId < diff {
//             res[i] = node.Key
//         }
//     }
//     return res
// }

func main() {
    // Example 1:
    // Input: rooms = [[2,2],[1,2],[3,2]], queries = [[3,1],[3,3],[5,2]]
    // Output: [3,-1,3]
    // Explanation: The answers to the queries are as follows:
    // Query = [3,1]: Room number 3 is the closest as abs(3 - 3) = 0, and its size of 2 is at least 1. The answer is 3.
    // Query = [3,3]: There are no rooms with a size of at least 3, so the answer is -1.
    // Query = [5,2]: Room number 3 is the closest as abs(3 - 5) = 2, and its size of 2 is at least 2. The answer is 3.
    fmt.Println(closestRoom([][]int{{2,2},{1,2},{3,2}}, [][]int{{3,1},{3,3},{5,2}})) // [3,-1,3]
    // Example 2:
    // Input: rooms = [[1,4],[2,3],[3,5],[4,1],[5,2]], queries = [[2,3],[2,4],[2,5]]
    // Output: [2,1,3]
    // Explanation: The answers to the queries are as follows:
    // Query = [2,3]: Room number 2 is the closest as abs(2 - 2) = 0, and its size of 3 is at least 3. The answer is 2.
    // Query = [2,4]: Room numbers 1 and 3 both have sizes of at least 4. The answer is 1 since it is smaller.
    // Query = [2,5]: Room number 3 is the only room with a size of at least 5. The answer is 3.
    fmt.Println(closestRoom([][]int{{1,4},{2,3},{3,5},{4,1},{5,2}}, [][]int{{2,3},{2,4},{2,5}})) // [2,1,3]

    fmt.Println(closestRoom1([][]int{{2,2},{1,2},{3,2}}, [][]int{{3,1},{3,3},{5,2}})) // [3,-1,3]
    fmt.Println(closestRoom1([][]int{{1,4},{2,3},{3,5},{4,1},{5,2}}, [][]int{{2,3},{2,4},{2,5}})) // [2,1,3]

    fmt.Println(closestRoom2([][]int{{2,2},{1,2},{3,2}}, [][]int{{3,1},{3,3},{5,2}})) // [3,-1,3]
    fmt.Println(closestRoom2([][]int{{1,4},{2,3},{3,5},{4,1},{5,2}}, [][]int{{2,3},{2,4},{2,5}})) // [2,1,3]
}