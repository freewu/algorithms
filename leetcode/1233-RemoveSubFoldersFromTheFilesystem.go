package main

// 1233. Remove Sub-Folders from the Filesystem
// Given a list of folders folder, return the folders after removing all sub-folders in those folders. 
// You may return the answer in any order.

// If a folder[i] is located within another folder[j], it is called a sub-folder of it.

// The format of a path is one or more concatenated strings of the form: '/' followed by one or more lowercase English letters.
//     For example, "/leetcode" and "/leetcode/problems" are valid paths while an empty string and "/" are not.

// Example 1:
// Input: folder = ["/a","/a/b","/c/d","/c/d/e","/c/f"]
// Output: ["/a","/c/d","/c/f"]
// Explanation: Folders "/a/b" is a subfolder of "/a" and "/c/d/e" is inside of folder "/c/d" in our filesystem.

// Example 2:
// Input: folder = ["/a","/a/b/c","/a/b/d"]
// Output: ["/a"]
// Explanation: Folders "/a/b/c" and "/a/b/d" will be removed because they are subfolders of "/a".

// Example 3:
// Input: folder = ["/a/b/c","/a/b/ca","/a/b/d"]
// Output: ["/a/b/c","/a/b/ca","/a/b/d"]

// Constraints:
//     1 <= folder.length <= 4 * 10^4
//     2 <= folder[i].length <= 100
//     folder[i] contains only lowercase letters and '/'.
//     folder[i] always starts with the character '/'.
//     Each folder name is unique.

import "fmt"
import "sort"
import "strings"

func removeSubfolders(folder []string) []string {
    sort.Slice(folder, func(i, j int) bool { // 从短到到长排列
        return len(folder[i]) < len(folder[j])
    })
    for i := 0; i < len(folder); i++ {
        for j := i + 1; j < len(folder); j++ {
            if strings.HasPrefix(folder[j], folder[i]) {
                if strings.Contains(folder[j], folder[i]+"/") {
                    left, right := folder[:j], folder[j+1:]
                    folder = append(left, right...)
                    j = len(left) - 1
                }
            }
        }
    }
    return folder
}

func removeSubfolders1(folder []string) []string {
    sort.Strings(folder)
    res := []string{ folder[0] }
    for _, f := range folder[1:] {
        last := res[len(res)-1]
        if !strings.HasPrefix(f, last) || f[len(last)] != '/' {
            res = append(res, f)
        }
    }
    return res
}

func removeSubfolders2(folder []string) []string {
    sort.Strings(folder)
    res := []string{folder[0]}
    for i := 1; i < len(folder); i++ {
        m, n := len(res[len(res) - 1]), len(folder[i])
        if m >= n || !(res[len(res) - 1] == folder[i][:m] && folder[i][m] == '/') {
            res = append(res, folder[i])
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: folder = ["/a","/a/b","/c/d","/c/d/e","/c/f"]
    // Output: ["/a","/c/d","/c/f"]
    // Explanation: Folders "/a/b" is a subfolder of "/a" and "/c/d/e" is inside of folder "/c/d" in our filesystem.
    fmt.Println(removeSubfolders([]string{"/a","/a/b","/c/d","/c/d/e","/c/f"})) // ["/a","/c/d","/c/f"]
    // Example 2:
    // Input: folder = ["/a","/a/b/c","/a/b/d"]
    // Output: ["/a"]
    // Explanation: Folders "/a/b/c" and "/a/b/d" will be removed because they are subfolders of "/a".
    fmt.Println(removeSubfolders([]string{"/a","/a/b/c","/a/b/d"})) // ["/a"]
    // Example 3:
    // Input: folder = ["/a/b/c","/a/b/ca","/a/b/d"]
    // Output: ["/a/b/c","/a/b/ca","/a/b/d"]
    fmt.Println(removeSubfolders([]string{"/a/b/c","/a/b/ca","/a/b/d"})) // ["/a/b/c","/a/b/ca","/a/b/d"]

    fmt.Println(removeSubfolders1([]string{"/a","/a/b","/c/d","/c/d/e","/c/f"})) // ["/a","/c/d","/c/f"]
    fmt.Println(removeSubfolders1([]string{"/a","/a/b/c","/a/b/d"})) // ["/a"]
    fmt.Println(removeSubfolders1([]string{"/a/b/c","/a/b/ca","/a/b/d"})) // ["/a/b/c","/a/b/ca","/a/b/d"]

    fmt.Println(removeSubfolders2([]string{"/a","/a/b","/c/d","/c/d/e","/c/f"})) // ["/a","/c/d","/c/f"]
    fmt.Println(removeSubfolders2([]string{"/a","/a/b/c","/a/b/d"})) // ["/a"]
    fmt.Println(removeSubfolders2([]string{"/a/b/c","/a/b/ca","/a/b/d"})) // ["/a/b/c","/a/b/ca","/a/b/d"]
}