package main

// 1993. Operations on Tree
// You are given a tree with n nodes numbered from 0 to n - 1 in the form of a parent array parent where parent[i] is the parent of the ith node. 
// The root of the tree is node 0, so parent[0] = -1 since it has no parent. 
// You want to design a data structure that allows users to lock, unlock, and upgrade nodes in the tree.

// The data structure should support the following functions:
//     Lock: 
//         Locks the given node for the given user and prevents other users from locking the same node. 
//         You may only lock a node using this function if the node is unlocked.
//     Unlock: 
//         Unlocks the given node for the given user. 
//         You may only unlock a node using this function if it is currently locked by the same user.
//     Upgrade: 
//         Locks the given node for the given user and unlocks all of its descendants regardless of who locked it. 
//         You may only upgrade a node if all 3 conditions are true:
//             The node is unlocked,
//             It has at least one locked descendant (by any user), and
//             It does not have any locked ancestors.

// Implement the LockingTree class:
//     LockingTree(int[] parent) 
//         initializes the data structure with the parent array.
//     lock(int num, int user) 
//         returns true if it is possible for the user with id user to lock the node num, or false otherwise. 
//         If it is possible, the node num will become locked by the user with id user.
//     unlock(int num, int user) 
//         returns true if it is possible for the user with id user to unlock the node num, or false otherwise. 
//         If it is possible, the node num will become unlocked.
//     upgrade(int num, int user) 
//         returns true if it is possible for the user with id user to upgrade the node num, or false otherwise. 
//         If it is possible, the node num will be upgraded.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/29/untitled.png" />
// Input
// ["LockingTree", "lock", "unlock", "unlock", "lock", "upgrade", "lock"]
// [[[-1, 0, 0, 1, 1, 2, 2]], [2, 2], [2, 3], [2, 2], [4, 5], [0, 1], [0, 1]]
// Output
// [null, true, false, true, true, true, false]
// Explanation
// LockingTree lockingTree = new LockingTree([-1, 0, 0, 1, 1, 2, 2]);
// lockingTree.lock(2, 2);    // return true because node 2 is unlocked.
//                            // Node 2 will now be locked by user 2.
// lockingTree.unlock(2, 3);  // return false because user 3 cannot unlock a node locked by user 2.
// lockingTree.unlock(2, 2);  // return true because node 2 was previously locked by user 2.
//                            // Node 2 will now be unlocked.
// lockingTree.lock(4, 5);    // return true because node 4 is unlocked.
//                            // Node 4 will now be locked by user 5.
// lockingTree.upgrade(0, 1); // return true because node 0 is unlocked and has at least one locked descendant (node 4).
//                            // Node 0 will now be locked by user 1 and node 4 will now be unlocked.
// lockingTree.lock(0, 1);    // return false because node 0 is already locked.

// Constraints:
//     n == parent.length
//     2 <= n <= 2000
//     0 <= parent[i] <= n - 1 for i != 0
//     parent[0] == -1
//     0 <= num <= n - 1
//     1 <= user <= 10^4
//     parent represents a valid tree.
//     At most 2000 calls in total will be made to lock, unlock, and upgrade.

import "fmt"

type LockingTree struct {
    lock       map[int]bool
    lockByUser map[int]int
    childs     map[int][]int
    parents    []int
}

func Constructor(parent []int) LockingTree {
    childs := map[int][]int{}
    for i := 0; i < len(parent); i++ {
        if _, f := childs[parent[i]]; !f {
            childs[parent[i]] = []int{}
        }
        childs[parent[i]] = append(childs[parent[i]], i)
    }
    return LockingTree{
        lock:       map[int]bool{},
        lockByUser: map[int]int{},
        parents:    parent,
        childs:     childs,
    }
}

func (this *LockingTree) Lock(num int, user int) bool {
    if v, f := this.lock[num]; !f || !v {
        this.lock[num] = true
        this.lockByUser[num] = user
        return true
    }
    return false
}

func (this *LockingTree) Unlock(num int, user int) bool {
    locked, f := this.lock[num]
    owner, _ := this.lockByUser[num]
    if f && locked && owner == user {
        this.lock[num] = false
        return true
    }
    return false
}

func (this *LockingTree) HasParentLock(num int) bool {
    parentNum := this.parents[num]
    for parentNum != -1 && !this.lock[parentNum] {
        parentNum = this.parents[parentNum]
    }
    return parentNum != -1
}

func (this *LockingTree) HasChildLock(parent int) bool {
    lock := false
    if numChilds, f := this.childs[parent]; f {
        for _, child := range numChilds {
            if this.lock[child] {
                this.lock[child] = false
                lock = true
            }
            if this.HasChildLock(child) {
                lock = true
            }
        }
    }
    return lock
}

func (this *LockingTree) Upgrade(num int, user int) bool {
    if v, f := this.lock[num]; !f || !v {
        parentLock := this.HasParentLock(num)
        if !parentLock {
            childsLock := this.HasChildLock(num)
            if childsLock {
                this.lock[num] = true
                this.lockByUser[num] = user
                return true
            }
        }
    }
    return false
}


type LockingTree1 struct {
    parent []int
    child [][]int
    locked []int
}


func Constructor1(parent []int) LockingTree1 {
    n := len(parent)
    child := make([][]int, n)
    for i := 1; i < n; i++ {
        child[parent[i]] = append(child[parent[i]], i)
    }
    return LockingTree1{parent, child, make([]int, len(parent))}
}


func (this *LockingTree1) Lock(num int, user int) bool {
    if this.locked[num] > 0 {
        return false
    }
    this.locked[num] = user
    return true
}


func (this *LockingTree1) Unlock(num int, user int) bool {
    if this.locked[num] != user { return false }
    this.locked[num]=0
    return true
}


func (this *LockingTree1) Upgrade(num int, user int) bool {
    curr := num
    for curr != -1 {
        if this.locked[curr] > 0 {  return false }
        curr = this.parent[curr]
    }
    lockedChild := false
    var dfs func(curr int)
    dfs = func(curr int) {
        if this.locked[curr] > 0 {
            lockedChild = true
            this.locked[curr] = 0
        }
        for _,c := range this.child[curr] {
            dfs(c)
        }
    }
    dfs(num)
    if lockedChild {
        this.locked[num] = user
    }
    return lockedChild
}

/**
 * Your LockingTree object will be instantiated and called as such:
 * obj := Constructor(parent);
 * param_1 := obj.Lock(num,user);
 * param_2 := obj.Unlock(num,user);
 * param_3 := obj.Upgrade(num,user);
 */

func main() {
    // LockingTree lockingTree = new LockingTree([-1, 0, 0, 1, 1, 2, 2]);
    obj := Constructor([]int{-1, 0, 0, 1, 1, 2, 2})
    fmt.Println(obj)
    // lockingTree.lock(2, 2);    // return true because node 2 is unlocked.
    //                            // Node 2 will now be locked by user 2.
    fmt.Println(obj.Lock(2, 2)) // true
    fmt.Println(obj)
    // lockingTree.unlock(2, 3);  // return false because user 3 cannot unlock a node locked by user 2.
    fmt.Println(obj.Unlock(2, 3)) // false
    fmt.Println(obj)
    // lockingTree.unlock(2, 2);  // return true because node 2 was previously locked by user 2.
    //                            // Node 2 will now be unlocked.
    fmt.Println(obj.Unlock(2, 2)) // true
    fmt.Println(obj)
    // lockingTree.lock(4, 5);    // return true because node 4 is unlocked.
    //                            // Node 4 will now be locked by user 5.
    fmt.Println(obj.Lock(4, 5)) // true
    fmt.Println(obj)
    // lockingTree.upgrade(0, 1); // return true because node 0 is unlocked and has at least one locked descendant (node 4).
    //                            // Node 0 will now be locked by user 1 and node 4 will now be unlocked.
    fmt.Println(obj.Upgrade(0, 1)) // true
    fmt.Println(obj)
    // lockingTree.lock(0, 1);    // return false because node 0 is already locked.
    fmt.Println(obj.Lock(0, 1)) // false
    fmt.Println(obj)

    // LockingTree lockingTree = new LockingTree([-1, 0, 0, 1, 1, 2, 2]);
    obj1 := Constructor1([]int{-1, 0, 0, 1, 1, 2, 2})
    fmt.Println(obj1)
    // lockingTree.lock(2, 2);    // return true because node 2 is unlocked.
    //                            // Node 2 will now be locked by user 2.
    fmt.Println(obj1.Lock(2, 2)) // true
    fmt.Println(obj1)
    // lockingTree.unlock(2, 3);  // return false because user 3 cannot unlock a node locked by user 2.
    fmt.Println(obj1.Unlock(2, 3)) // false
    fmt.Println(obj1)
    // lockingTree.unlock(2, 2);  // return true because node 2 was previously locked by user 2.
    //                            // Node 2 will now be unlocked.
    fmt.Println(obj1.Unlock(2, 2)) // true
    fmt.Println(obj1)
    // lockingTree.lock(4, 5);    // return true because node 4 is unlocked.
    //                            // Node 4 will now be locked by user 5.
    fmt.Println(obj1.Lock(4, 5)) // true
    fmt.Println(obj1)
    // lockingTree.upgrade(0, 1); // return true because node 0 is unlocked and has at least one locked descendant (node 4).
    //                            // Node 0 will now be locked by user 1 and node 4 will now be unlocked.
    fmt.Println(obj1.Upgrade(0, 1)) // true
    fmt.Println(obj1)
    // lockingTree.lock(0, 1);    // return false because node 0 is already locked.
    fmt.Println(obj1.Lock(0, 1)) // false
    fmt.Println(obj1)
}