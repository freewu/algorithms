package main

// 1656. Design an Ordered Stream
// There is a stream of n (idKey, value) pairs arriving in an arbitrary order, 
// where idKey is an integer between 1 and n and value is a string. No two pairs have the same id.

// Design a stream that returns the values in increasing order of their IDs by returning a chunk (list) of values after each insertion. 
// The concatenation of all the chunks should result in a list of the sorted values.

// Implement the OrderedStream class:
//     OrderedStream(int n) 
//         Constructs the stream to take n values.
//     String[] insert(int idKey, String value) 
//         Inserts the pair (idKey, value) into the stream, then returns the largest possible chunk of currently inserted values that appear next in the order.

// Example:
// <img src="https://assets.leetcode.com/uploads/2020/11/10/q1.gif" />
// Input
// ["OrderedStream", "insert", "insert", "insert", "insert", "insert"]
// [[5], [3, "ccccc"], [1, "aaaaa"], [2, "bbbbb"], [5, "eeeee"], [4, "ddddd"]]
// Output
// [null, [], ["aaaaa"], ["bbbbb", "ccccc"], [], ["ddddd", "eeeee"]]
// Explanation
// // Note that the values ordered by ID is ["aaaaa", "bbbbb", "ccccc", "ddddd", "eeeee"].
// OrderedStream os = new OrderedStream(5);
// os.insert(3, "ccccc"); // Inserts (3, "ccccc"), returns [].
// os.insert(1, "aaaaa"); // Inserts (1, "aaaaa"), returns ["aaaaa"].
// os.insert(2, "bbbbb"); // Inserts (2, "bbbbb"), returns ["bbbbb", "ccccc"].
// os.insert(5, "eeeee"); // Inserts (5, "eeeee"), returns [].
// os.insert(4, "ddddd"); // Inserts (4, "ddddd"), returns ["ddddd", "eeeee"].
// // Concatentating all the chunks returned:
// // [] + ["aaaaa"] + ["bbbbb", "ccccc"] + [] + ["ddddd", "eeeee"] = ["aaaaa", "bbbbb", "ccccc", "ddddd", "eeeee"]
// // The resulting order is the same as the order above.

// Constraints:
//     1 <= n <= 1000
//     1 <= id <= n
//     value.length == 5
//     value consists only of lowercase letters.
//     Each call to insert will have a unique id.
//     Exactly n calls will be made to insert.

import "fmt"

type OrderedStream struct {
    ptr int // The pointer starts at the first index [0] in this.list.
    list []string
}

func Constructor(n int) OrderedStream {
    return OrderedStream{ list : make([]string, n), ptr  : 0, }
}

func (this *OrderedStream) Insert(id int, value string) []string {
    this.list[id - 1] = value // Always insert. It is a 1-indexed id in a 0-indexed slice, so subtract 1
    // Only when a value at the first index [0] is added do we move this.ptr.
    // Until then, we return empty strings.
    if this.list[this.ptr] == "" {
        return []string{}
    }
    // At this point, we have a value at this.ptr.
    // We increment this.ptr until we find an empty value or we reach
    // the end of this.list.
    end := this.ptr
    for this.ptr < len(this.list) && this.list[this.ptr] != "" {
        this.ptr++
        end++
    }
    return this.list[id-1:end] // Return the sub-array of values.
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */

func main() {
    // OrderedStream os = new OrderedStream(5);
    obj := Constructor(5)
    fmt.Println(obj)
    // os.insert(3, "ccccc"); // Inserts (3, "ccccc"), returns [].
    fmt.Println(obj.Insert(3, "ccccc")) // []
    fmt.Println(obj)
    // os.insert(1, "aaaaa"); // Inserts (1, "aaaaa"), returns ["aaaaa"].
    fmt.Println(obj.Insert(1, "aaaaa")) // ["aaaaa"]
    fmt.Println(obj)
    // os.insert(2, "bbbbb"); // Inserts (2, "bbbbb"), returns ["bbbbb", "ccccc"].
    fmt.Println(obj.Insert(2, "bbbbb")) // ["bbbbb", "ccccc"]
    fmt.Println(obj)
    // os.insert(5, "eeeee"); // Inserts (5, "eeeee"), returns [].
    fmt.Println(obj.Insert(5, "eeeee")) // []
    fmt.Println(obj)
    // os.insert(4, "ddddd"); // Inserts (4, "ddddd"), returns ["ddddd", "eeeee"].
    // // Concatentating all the chunks returned:
    fmt.Println(obj.Insert(4, "ddddd")) // ["ddddd", "eeeee"]
    fmt.Println(obj)
}