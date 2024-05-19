package main

// 428. Serialize and Deserialize N-ary Tree
// Serialization is the process of converting a data structure 
// or object into a sequence of bits so that it can be stored in a file or memory buffer, 
// or transmitted across a network connection link to be reconstructed later in the same 
// or another computer environment.

// Design an algorithm to serialize and deserialize an N-ary tree. 
// An N-ary tree is a rooted tree in which each node has no more than N children. 
// There is no restriction on how your serialization/deserialization algorithm should work. 
// You just need to ensure that an N-ary tree can be serialized to a string 
// and this string can be deserialized to the original tree structure.

// For example, you may serialize the following 3-ary tree
// <img src="https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png" />

// as [1 [3[5 6] 2 4]]. Note that this is just an example, you do not necessarily need to follow this format.
// Or you can follow LeetCode's level order traversal serialization format, where each group of children is separated by the null value.

// For example, the above tree may be serialized as [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14].
// <img src="https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png" />

// You do not necessarily need to follow the above-suggested formats, 
// there are many more different formats that work so please be creative and come up with different approaches yourself.

// Example 1:
// Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]

// Example 2:
// Input: root = [1,null,3,2,4,null,5,6]
// Output: [1,null,3,2,4,null,5,6]

// Example 3:
// Input: root = []
// Output: []

// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     0 <= Node.val <= 10^4
//     The height of the n-ary tree is less than or equal to 1000
//     Do not use class member/global/static variables to store states. Your encode and decode algorithms should be stateless.

import "fmt"
import "strconv"
import "strings"

type Node struct {
    Val int
    Children []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Codec struct {

}

func Constructor() *Codec {
    return &Codec{}
}

func (this *Codec) serialize(root *Node) string {
    if root == nil {
        return ""
    }
    res, q := make([]string, 0), make([]*Node, 0)
    q = append(q, root)
    for ;len(q) > 0; {
        nxtLevel := make([]*Node, 0)
        for ; len(q) > 0; {
            cn := q[0]
            q = q[1:]
            res = append(res, strconv.Itoa(cn.Val))
            if cn.Children == nil || len(cn.Children) == 0 {
                res = append(res, "0")
            } else {
                res = append(res, strconv.Itoa(len(cn.Children)))
            }

            for _, cd := range cn.Children {
                nxtLevel = append(nxtLevel, cd)
            }
        }
        q = nxtLevel
    }
    return strings.Join(res, ",")
}

func (this *Codec) deserialize(data string) *Node {
    if len(data) == 0 {
        return nil
    }
    res := strings.Split(data, ",")
    i2n := make(map[int]*Node)
    fi, ci := 0, 2
    for ; fi < len(res); fi += 2{
        if _, e := i2n[fi]; !e {
            nv, _ := strconv.Atoi(res[fi])
            i2n[fi] = &Node{Val: nv}
            i2n[fi].Children = make([]*Node, 0)
        }
        nc, _ := strconv.Atoi(res[fi + 1])
        bdr := ci + nc * 2
        for ; ci < bdr; ci += 2 {
            cv, _ := strconv.Atoi(res[ci])
            cn := &Node{Val: cv}
            cn.Children = make([]*Node, 0)
            i2n[ci] = cn
            i2n[fi].Children = append(i2n[fi].Children, cn)
        }
    }
    return i2n[0]
}
 
/**
  * Your Codec object will be instantiated and called as such:
  * obj := Constructor();
  * data := obj.serialize(root);
  * ans := obj.deserialize(data);
  */

func main() {
    obj := Constructor()
    // Example 1:
    // Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
    // Output: [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
    tree1 := &Node{
        1,
        []*Node {
            &Node{2, nil },
            &Node{3, []*Node{ &Node{6, nil }, &Node{7, []*Node{ &Node{ 11, []*Node{ &Node{14, nil }} }, } }, }  },
            &Node{4, []*Node{ &Node{8, []*Node{ &Node{12, nil }}  }, } },
            &Node{5, []*Node{ &Node{9, []*Node{ &Node{13, nil }}  }, &Node{10, nil } } },
        },
    }
    data1 := obj.serialize(tree1)
    fmt.Println(data1)
    t1 := obj.deserialize(data1)
    fmt.Println(t1.Val)
    fmt.Println(len(t1.Children))
    // Example 2:
    // Input: root = [1,null,3,2,4,null,5,6]
    // Output: [1,null,3,2,4,null,5,6]
    tree2 := &Node{
        1,
        []*Node {
            &Node{3, []*Node{ &Node{5, nil }, &Node{6, nil }, }  },
            &Node{2, nil },
            &Node{4, nil },
        },
    }
    data2 := obj.serialize(tree2)
    fmt.Println(data2)
    t2 := obj.deserialize(data2)
    fmt.Println(t2.Val)
    fmt.Println(len(t2.Children))
    // Example 3:
    // Input: root = []
    // Output: []
    data3 := obj.serialize(nil)
    fmt.Println(data3)
    t3 := obj.deserialize(data3)
    fmt.Println(t3)
}