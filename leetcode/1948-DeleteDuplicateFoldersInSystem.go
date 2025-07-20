package main

// 1948. Delete Duplicate Folders in System
// Due to a bug, there are many duplicate folders in a file system. 
// You are given a 2D array paths, where paths[i] is an array representing an absolute path to the ith folder in the file system.
//     For example, ["one", "two", "three"] represents the path "/one/two/three".

// Two folders (not necessarily on the same level) are identical if they contain the same non-empty set of identical subfolders and underlying subfolder structure. 
// The folders do not need to be at the root level to be identical. 
// If two or more folders are identical, then mark the folders as well as all their subfolders.
//     1. For example, folders "/a" and "/b" in the file structure below are identical. 
//        They (as well as their subfolders) should all be marked:
//             /a
//             /a/x
//             /a/x/y
//             /a/z
//             /b
//             /b/x
//             /b/x/y
//             /b/z
//     2. However, if the file structure also included the path "/b/w", then the folders "/a" and "/b" would not be identical. 
//        Note that "/a/x" and "/b/x" would still be considered identical even with the added folder.

// Once all the identical folders and their subfolders have been marked, the file system will delete all of them. 
// The file system only runs the deletion once, so any folders that become identical after the initial deletion are not deleted.

// Return the 2D array ans containing the paths of the remaining folders after deleting all the marked folders. 
// The paths may be returned in any order.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/19/lc-dupfolder1.jpg" />
// Input: paths = [["a"],["c"],["d"],["a","b"],["c","b"],["d","a"]]
// Output: [["d"],["d","a"]]
// Explanation: The file structure is as shown.
// Folders "/a" and "/c" (and their subfolders) are marked for deletion because they both contain an empty
// folder named "b".

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/07/19/lc-dupfolder2.jpg" />
// Input: paths = [["a"],["c"],["a","b"],["c","b"],["a","b","x"],["a","b","x","y"],["w"],["w","y"]]
// Output: [["c"],["c","b"],["a"],["a","b"]]
// Explanation: The file structure is as shown. 
// Folders "/a/b/x" and "/w" (and their subfolders) are marked for deletion because they both contain an empty folder named "y".
// Note that folders "/a" and "/c" are identical after the deletion, but they are not deleted because they were not marked beforehand.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/07/19/lc-dupfolder3.jpg" />
// Input: paths = [["a","b"],["c","d"],["c"],["a"]]
// Output: [["c"],["c","d"],["a"],["a","b"]]
// Explanation: All folders are unique in the file system.
// Note that the returned array can be in a different order as the order does not matter.

// Constraints:
//     1 <= paths.length <= 2 * 10^4
//     1 <= paths[i].length <= 500
//     1 <= paths[i][j].length <= 10
//     1 <= sum(paths[i][j].length) <= 2 * 10^5
//     path[i][j] consists of lowercase English letters.
//     No two paths lead to the same folder.
//     For any folder not at the root level, its parent folder will also be in the input.

import "fmt"
import "sort"
import "strings"

func deleteDuplicateFolder(paths [][]string) [][]string {
    type Folder struct {
        children map[string]*Folder
        name string // 文件夹名称
        delete bool  // 删除标记
    }
    res, root := [][]string{}, &Folder{}
    for _, path := range paths {
        f := root // 将 path 加入字典树
        for _, s := range path {
            if f.children == nil {
                f.children = map[string]*Folder{}
            }
            if f.children[s] == nil {
                f.children[s] = &Folder{}
            }
            f = f.children[s]
            f.name = s
        }
    }
    folders := map[string][]*Folder{} // 存储括号表达式及其对应的文件夹节点列表
    var dfs func(*Folder) string
    dfs = func(f *Folder) string {
        if f.children == nil { return "(" + f.name + ")" }
        expr := make([]string, 0, len(f.children))
        for _, child := range f.children {
            expr = append(expr, dfs(child))
        }
        sort.Strings(expr)
        subTreeExpr := strings.Join(expr, "") // 按字典序拼接所有子树
        folders[subTreeExpr] = append(folders[subTreeExpr], f)
        return "(" + f.name + subTreeExpr + ")"
    }
    dfs(root)
    for _, fs := range folders {
        if len(fs) > 1 { // 将括号表达式对应的节点个数大于 1 的节点全部删除
            for _, f := range fs {
                f.delete = true
            }
        }
    }
    // 再次 DFS 这颗字典树，仅访问未被删除的节点，并将路径记录到答案中
    path := []string{}
    var dfs2 func(*Folder)
    dfs2 = func(f *Folder) {
        if f.delete { return }
        path = append(path, f.name)
        res = append(res, append([]string(nil), path...))
        for _, child := range f.children {
            dfs2(child)
        }
        path = path[:len(path)-1]
    }
    for _, child := range root.children {
        dfs2(child)
    }
    return res
}

func deleteDuplicateFolder1(paths [][]string) [][]string {
    type Trie struct {
        serial   string           // current node structure's serialized representation
        children map[string]*Trie // current node's child nodes
    }
    root := &Trie{children: make(map[string]*Trie)} // root node
    // build a trie tree
    for _, path := range paths {
        cur := root
        for _, node := range path {
            if _, ok := cur.children[node]; !ok {
                cur.children[node] = &Trie{children: make(map[string]*Trie)}
            }
            cur = cur.children[node]
        }
    }
    freq := make(map[string]int) // hash table records the occurrence times of each serialized representation
    // post-order traversal based on depth-first search, calculate the serialized representation of each node structure
    var construct func(node *Trie)
    construct = func(node *Trie) {
        if len(node.children) == 0 { return } // if it is a leaf node, no operation is needed.
        v := make([]string, 0, len(node.children))
        for folder, child := range node.children {
            construct(child)
            v = append(v, folder+"("+child.serial+")")
        }
        sort.Strings(v)
        node.serial = strings.Join(v, "")
        freq[node.serial]++
    }
    construct(root)
    res, path:= make([][]string, 0), make([]string, 0)
    // operate the trie, delete duplicate folders
    var operate func(*Trie)
    operate = func(node *Trie) {
        if freq[node.serial] > 1 { return } // if the serialization representation appears more than once, it needs to be deleted
        if len(path) > 0 {
            tmp := make([]string, len(path))
            copy(tmp, path)
            res = append(res, tmp)
        }
        for folder, child := range node.children {
            path = append(path, folder)
            operate(child)
            path = path[:len(path)-1]
        }
    }
    operate(root)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/07/19/lc-dupfolder1.jpg" />
    // Input: paths = [["a"],["c"],["d"],["a","b"],["c","b"],["d","a"]]
    // Output: [["d"],["d","a"]]
    // Explanation: The file structure is as shown.
    // Folders "/a" and "/c" (and their subfolders) are marked for deletion because they both contain an empty
    // folder named "b".
    fmt.Println(deleteDuplicateFolder([][]string{{"a"},{"c"},{"d"},{"a","b"},{"c","b"},{"d","a"}})) // [["d"],["d","a"]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/07/19/lc-dupfolder2.jpg" />
    // Input: paths = [["a"],["c"],["a","b"],["c","b"],["a","b","x"],["a","b","x","y"],["w"],["w","y"]]
    // Output: [["c"],["c","b"],["a"],["a","b"]]
    // Explanation: The file structure is as shown. 
    // Folders "/a/b/x" and "/w" (and their subfolders) are marked for deletion because they both contain an empty folder named "y".
    // Note that folders "/a" and "/c" are identical after the deletion, but they are not deleted because they were not marked beforehand.
    fmt.Println(deleteDuplicateFolder([][]string{{"a"},{"c"},{"a","b"},{"c","b"},{"a","b","x"},{"a","b","x","y"},{"w"},{"w","y"}})) // [["c"],["c","b"],["a"],["a","b"]]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/07/19/lc-dupfolder3.jpg" />
    // Input: paths = [["a","b"],["c","d"],["c"],["a"]]
    // Output: [["c"],["c","d"],["a"],["a","b"]]
    // Explanation: All folders are unique in the file system.
    // Note that the returned array can be in a different order as the order does not matter.
    fmt.Println(deleteDuplicateFolder([][]string{{"a","b"},{"c","d"},{"c"},{"a"}})) // [["c"],["c","d"],["a"],["a","b"]]

    fmt.Println(deleteDuplicateFolder1([][]string{{"a"},{"c"},{"d"},{"a","b"},{"c","b"},{"d","a"}})) // [["d"],["d","a"]]
    fmt.Println(deleteDuplicateFolder1([][]string{{"a"},{"c"},{"a","b"},{"c","b"},{"a","b","x"},{"a","b","x","y"},{"w"},{"w","y"}})) // [["c"],["c","b"],["a"],["a","b"]]
    fmt.Println(deleteDuplicateFolder1([][]string{{"a","b"},{"c","d"},{"c"},{"a"}})) // [["c"],["c","d"],["a"],["a","b"]]
}