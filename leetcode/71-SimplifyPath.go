package main

// 71. Simplify Path
// Given a string path, which is an absolute path (starting with a slash '/') to a file or directory in a Unix-style file system, 
// convert it to the simplified canonical path.

// In a Unix-style file system, 
//     a period '.' refers to the current directory, 
//     a double period '..' refers to the directory up a level, 
//     and any multiple consecutive slashes (i.e. '//') are treated as a single slash '/'. 
//         For this problem, any other format of periods such as '...' are treated as file/directory names.

// The canonical path should have the following format:
//     The path starts with a single slash '/'.
//     Any two directories are separated by a single slash '/'.
//     The path does not end with a trailing '/'.
//     The path only contains the directories on the path from the root directory to the target file or directory (i.e., no period '.' or double period '..')
//     Return the simplified canonical path.

// Example 1:
// Input: path = "/home/"
// Output: "/home"
// Explanation: Note that there is no trailing slash after the last directory name.

// Example 2:
// Input: path = "/../"
// Output: "/"
// Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.

// Example 3:
// Input: path = "/home//foo/"
// Output: "/home/foo"
// Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.

// Constraints:
//     1 <= path.length <= 3000
//     path consists of English letters, digits, period '.', slash '/' or '_'.
//     path is a valid absolute Unix path.

import "fmt"
import "path/filepath"
import "strings"

// 
func simplifyPath(path string) string {
    arr := strings.Split(path, "/") // 使用 / 将字符串分割顾数组
    stack := make([]string, 0)
    var res string
    for i := 0; i < len(arr); i++ {
        cur := arr[i]
        //cur := strings.TrimSpace(arr[i]) 更加严谨的做法应该还要去掉末尾的空格
        if cur == ".." { // 如是 .. 返回上一层
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        } else if cur != "." && len(cur) > 0 { // 一般的目录
            stack = append(stack, arr[i])
        }
    }
    if len(stack) == 0 {
        return "/"
    }
    res = strings.Join(stack, "/")
    return "/" + res // 都以 / 开头
}

//  golang 的官方库 API
func simplifyPath1(path string) string {
    return filepath.Clean(path)
}

// best solution stack
func simplifyPath2(path string) string {
    paths, res := strings.Split(path, "/"), []string{}
    for _, p := range paths{
        if len(p) == 0 || p == "." {
            continue
        }
        if p == ".." { // 遇见 .. Pop
            if len(res) > 0 {
                res = res[:len(res)-1]
            }
            // else, skip the ..
        } else {
            res = append(res, p) // 入栈
        }
    }
    if len(res) == 0{
        return "/"
    }
    return fmt.Sprintf("/%s", strings.Join(res, "/"))
}

func simplifyPath3(path string) string {
    stack := []string{}
    for _, v := range strings.Split(path, "/") {
        if v == ".." {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        } else if v != "." && v != "" {
            stack = append(stack, v)
        }
    }
    return "/" + strings.Join(stack, "/")
}

func main() {
    fmt.Printf("simplifyPath(\"/home/\") = %v\n",simplifyPath("/home/")) // "/home"
    fmt.Printf("simplifyPath(\"/../\") = %v\n",simplifyPath("/../")) // "/"
    fmt.Printf("simplifyPath(\"/home//foo/\") = %v\n",simplifyPath("/home//foo/")) // "/home/foo"
    fmt.Printf("simplifyPath(\"/a/./b/../../c/\") = %v\n",simplifyPath("/a/./b/../../c/")) // "/c"
    fmt.Printf("simplifyPath(\"/a/../../b/../c//.//\") = %v\n",simplifyPath("/a/../../b/../c//.//")) // "/c"
    fmt.Printf("simplifyPath(\"/a//b////c/d//././/..\") = %v\n",simplifyPath("//a//b////c/d//././/..")) // "/a/b/c"

    fmt.Printf("simplifyPath1(\"/home/\") = %v\n",simplifyPath1("/home/")) // "/home"
    fmt.Printf("simplifyPath1(\"/../\") = %v\n",simplifyPath1("/../")) // "/"
    fmt.Printf("simplifyPath1(\"/home//foo/\") = %v\n",simplifyPath1("/home//foo/")) // "/home/foo"
    fmt.Printf("simplifyPath1(\"/a/./b/../../c/\") = %v\n",simplifyPath1("/a/./b/../../c/")) // "/c"
    fmt.Printf("simplifyPath1(\"/a/../../b/../c//.//\") = %v\n",simplifyPath1("/a/../../b/../c//.//")) // "/c"
    fmt.Printf("simplifyPath1(\"/a//b////c/d//././/..\") = %v\n",simplifyPath1("//a//b////c/d//././/..")) // "/a/b/c"

    fmt.Printf("simplifyPath2(\"/home/\") = %v\n",simplifyPath2("/home/")) // "/home"
    fmt.Printf("simplifyPath2(\"/../\") = %v\n",simplifyPath2("/../")) // "/"
    fmt.Printf("simplifyPath2(\"/home//foo/\") = %v\n",simplifyPath2("/home//foo/")) // "/home/foo"
    fmt.Printf("simplifyPath2(\"/a/./b/../../c/\") = %v\n",simplifyPath2("/a/./b/../../c/")) // "/c"
    fmt.Printf("simplifyPath2(\"/a/../../b/../c//.//\") = %v\n",simplifyPath2("/a/../../b/../c//.//")) // "/c"
    fmt.Printf("simplifyPath2(\"/a//b////c/d//././/..\") = %v\n",simplifyPath2("//a//b////c/d//././/..")) // "/a/b/c"

    fmt.Printf("simplifyPath3(\"/home/\") = %v\n",simplifyPath3("/home/")) // "/home"
    fmt.Printf("simplifyPath3(\"/../\") = %v\n",simplifyPath3("/../")) // "/"
    fmt.Printf("simplifyPath3(\"/home//foo/\") = %v\n",simplifyPath3("/home//foo/")) // "/home/foo"
    fmt.Printf("simplifyPath3(\"/a/./b/../../c/\") = %v\n",simplifyPath3("/a/./b/../../c/")) // "/c"
    fmt.Printf("simplifyPath3(\"/a/../../b/../c//.//\") = %v\n",simplifyPath3("/a/../../b/../c//.//")) // "/c"
    fmt.Printf("simplifyPath3(\"/a//b////c/d//././/..\") = %v\n",simplifyPath3("//a//b////c/d//././/..")) // "/a/b/c"

}
