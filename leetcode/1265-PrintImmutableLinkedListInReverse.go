package main

// 1265. Print Immutable Linked List in Reverse
// You are given an immutable linked list, print out all values of each node in reverse with the help of the following interface:
//     ImmutableListNode: An interface of immutable linked list, you are given the head of the list.

// You need to use the following functions to access the linked list (you can't access the ImmutableListNode directly):
//     ImmutableListNode.printValue(): Print value of the current node.
//     ImmutableListNode.getNext(): Return the next node.

// The input is only given to initialize the linked list internally. 
// You must solve this problem without modifying the linked list. 
// In other words, you must operate the linked list using only the mentioned APIs.

// Example 1:
// Input: head = [1,2,3,4]
// Output: [4,3,2,1]

// Example 2:
// Input: head = [0,-4,-1,3,-5]
// Output: [-5,3,-1,-4,0]

// Example 3:
// Input: head = [-2,0,6,4,4,-6]
// Output: [-6,4,4,6,0,-2]

// Constraints:
//     The length of the linked list is between [1, 1000].
//     The value of each node in the linked list is between [-1000, 1000].

// Follow up:
//     Could you solve this problem in:
//         Constant space complexity?
//         Linear time complexity and less than linear space complexity?

import "fmt"

/*   Below is the interface for ImmutableListNode, which is already defined for you.
 *
 *   type ImmutableListNode struct {
 *       
 *   }
 *
 *   func (this *ImmutableListNode) getNext() ImmutableListNode {
 *		// return the next node.
 *   }
 *
 *   func (this *ImmutableListNode) printValue() {
 *		// print the value of this node.
 *   }
 */
func printLinkedListInReverse(head ImmutableListNode) {
    // 常量空间消耗的话只能时间n^2了？每次从头扫描到最后然后打印
    // 时间O(n)且复杂小于O(n)，可以考虑二分，不过时间做不到O(n)而是nlogn
    var dfs func(first, last ImmutableListNode) 
    dfs = func(first, last ImmutableListNode) {
        if first == nil {
            return 
        }
        if first == last || first.getNext() == nil {
            first.printValue()
            return 
        } 
        p,pp := first, first.getNext()
        for ; pp!=last && pp.getNext()!=last; p,pp = p.getNext(), pp.getNext().getNext() {}
        dfs(p.getNext(), last)
        dfs(first, p)
    }
    dfs(head, nil)
}

func main() {

}