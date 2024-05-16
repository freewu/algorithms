package main 

// 588. Design In-Memory File System
// Design a data structure that simulates an in-memory file system.
// Implement the FileSystem class:
//     FileSystem() 
//         Initializes the object of the system.
//     List<String> ls(String path)
//         If path is a file path, returns a list that only contains this file's name.
//         If path is a directory path, returns the list of file and directory names in this directory.
//         The answer should in lexicographic order.
//     void mkdir(String path) 
//         Makes a new directory according to the given path. 
//         The given directory path does not exist. 
//         If the middle directories in the path do not exist, you should create them as well.
//     void addContentToFile(String filePath, String content)
//         If filePath does not exist, creates that file containing given content.
//         If filePath already exists, appends the given content to original content.
//     String readContentFromFile(String filePath) 
//         Returns the content in the file at filePath.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/28/filesystem.png" />
// Input
// ["FileSystem", "ls", "mkdir", "addContentToFile", "ls", "readContentFromFile"]
// [[], ["/"], ["/a/b/c"], ["/a/b/c/d", "hello"], ["/"], ["/a/b/c/d"]]
// Output
// [null, [], null, null, ["a"], "hello"]
// Explanation
// FileSystem fileSystem = new FileSystem();
// fileSystem.ls("/");                         // return []
// fileSystem.mkdir("/a/b/c");
// fileSystem.addContentToFile("/a/b/c/d", "hello");
// fileSystem.ls("/");                         // return ["a"]
// fileSystem.readContentFromFile("/a/b/c/d"); // return "hello"
 
// Constraints:
//     1 <= path.length, filePath.length <= 100
//     path and filePath are absolute paths which begin with '/' and do not end with '/' except that the path is just "/".
//     You can assume that all directory names and file names only contain lowercase letters, and the same names will not exist in the same directory.
//     You can assume that all operations will be passed valid parameters, and users will not attempt to retrieve file content or list a directory or file that does not exist.
//     1 <= content.length <= 50
//     At most 300 calls will be made to ls, mkdir, addContentToFile, and readContentFromFile.

import "fmt"
// import "sort"
// import "strings"

// type Directory struct {
//     All []string
//     Dirs map[string]*Directory // filename -> directory
//     Files map[string]string // filename -> content
// }

// type FileSystem struct {
//     root *Directory
// }

// func Constructor() FileSystem {
//     return FileSystem{root: &Directory{
//         All: make([]string, 0),
//         Dirs: make(map[string]*Directory),
//         Files: make(map[string]string),
//     }}
// }

// func (this *FileSystem) Ls(path string) []string {
//     if path == "/" {
//         return this.root.All
//     }
//     s := this.Split(path)
//     n := len(s)
//     p := this.root
//     for i := 0; i < n-1; i++ {
//         p = p.Dirs[s[i]]
//     }
//     if dir, ok := p.Dirs[s[n-1]]; ok { // 找到最后一个节点
//         return dir.All // 若是目录，则返回目录下文件及目录的字典序排列
//     }
//     if _, ok := p.Files[s[n-1]]; ok {
//         return []string{s[n-1]} // 若是文件，则返回文件名
//     }
//     return []string{}
// }

// func (this *FileSystem) Mkdir(path string)  {
//     s := this.Split(path)
//     p := this.root
//     for _, name := range s {
//         if _, ok := p.Dirs[name]; !ok {
//             // 按照字典序将name插入到p.All中
//             index := sort.SearchStrings(p.All, name)
//             if index == 0 {
//                 p.All = append([]string{name}, p.All...)
//             } else if index == len(p.All) {
//                 p.All = append(p.All, name)
//             } else {
//                 p.All = append(p.All[:index], append([]string{name}, p.All[index:]...)...)
//             }
//             // 在p.Dirs中记录目录信息
//             p.Dirs[name] = &Directory{
//                 All: []string{},
//                 Dirs: make(map[string]*Directory),
//                 Files: make(map[string]string),
//             } 
//         }
//         p = p.Dirs[name]
//     }
// }

// func (this *FileSystem) AddContentToFile(filePath string, content string)  {
//     s := this.Split(filePath)
//     n := len(s)
//     p := this.root
//     for i := 0; i < n-1; i++ {
//         p = p.Dirs[s[i]]
//     }
//     filename := s[n-1]
//     if f, ok := p.Files[filename]; ok {
//         // 文件已经存在，则进行追加操作
//         b := []byte(f)
//         b = append(b, []byte(content)...)
//         p.Files[filename] = string(b)
//     } else {
//         // 文件不存在，则创建文件
//         p.Files[filename] = content
//         // 并将新的文件插入到All中，和插入目录一个方式。
//         index := sort.SearchStrings(p.All, filename)
//         if index == 0 {
//             p.All = append([]string{filename}, p.All...)
//         } else if index == len(p.All) {
//             p.All = append(p.All, filename)
//         } else {
//             p.All = append(p.All[:index], append([]string{filename}, p.All[index:]...)...)
//         }
//     }
// }

// func (this *FileSystem) ReadContentFromFile(filePath string) string {
//     s := this.Split(filePath)
//     n := len(s)
//     p := this.root
//     for i := 0; i < n-1; i++ {
//         p = p.Dirs[s[i]]
//     }
//     return p.Files[s[n-1]]
// }

// func (this *FileSystem) Split(path string) []string {
//     s := strings.Split(path, "/")
//     return s[1:]
// }

import "slices"
import "strings"

type Directory struct {
    subdir map[string]*Directory
    file map[string]string
}

func newDirectory() *Directory {
    return &Directory{
        subdir: make(map[string]*Directory),
        file: make(map[string]string),
    }
}

type FileSystem struct {
    root *Directory
}

func Constructor() FileSystem {
    return FileSystem{
        root: newDirectory(),
    }
}

func (this *FileSystem) Ls(path string) []string {
    paths := this.paths(path)
    dir := this.root
    for _, p := range paths {
        var ok bool
        dir, ok = dir.subdir[p]
        if !ok {
            return []string{p}
        }
    }
    entries := []string{}
    for subdir := range dir.subdir {
        entries = append(entries, subdir)
    }
    for file := range dir.file {
        entries = append(entries, file)
    }
    slices.Sort(entries)
    return entries
}

func (this *FileSystem) paths(path string) []string{
    if path == "/" {
        return nil
    }
    return strings.Split(path[1:], "/")
}

func (this *FileSystem) Mkdir(path string)  {
    paths := this.paths(path)
    dir := this.root
    for _, p := range paths {
        if _, ok := dir.subdir[p]; !ok {
            dir.subdir[p] = newDirectory()
        }
        dir = dir.subdir[p]
    }
}

func (this *FileSystem) AddContentToFile(filePath string, content string)  {
    paths := this.paths(filePath)
    dir := this.root
    for _, p := range paths[:len(paths)-1] {
        dir = dir.subdir[p]
    }
    dir.file[paths[len(paths)-1]] += content
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
        paths := this.paths(filePath)
    dir := this.root
    for _, p := range paths[:len(paths)-1] {
        dir = dir.subdir[p]
    }
    return dir.file[paths[len(paths)-1]]
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */

func main() {
    // FileSystem fileSystem = new FileSystem();
    obj := Constructor()
    fmt.Println(obj)
    // fileSystem.ls("/");                         // return []
    fmt.Println(obj.Ls("/")) // []
    // fileSystem.mkdir("/a/b/c");
    obj.Mkdir("/a/b/c")
    fmt.Println(obj)
    // fileSystem.addContentToFile("/a/b/c/d", "hello");
    obj.AddContentToFile("/a/b/c/d", "hello")
    fmt.Println(obj)
    // fileSystem.ls("/");                         // return ["a"]
    fmt.Println(obj.Ls("/")) // ["a"]
    // fileSystem.readContentFromFile("/a/b/c/d"); // return "hello"
    fmt.Println(obj.ReadContentFromFile("/a/b/c/d")) // "hello"
}