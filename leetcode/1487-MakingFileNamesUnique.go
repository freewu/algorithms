package main

// 1487. Making File Names Unique
// Given an array of strings names of size n. 
// You will create n folders in your file system such that, 
// at the ith minute, you will create a folder with the name names[i].

// Since two files cannot have the same name, 
// if you enter a folder name that was previously used, 
// the system will have a suffix addition to its name in the form of (k), 
// where, k is the smallest positive integer such that the obtained name remains unique.

// Return an array of strings of length n where ans[i] is the actual name the system will assign to the ith folder when you create it.

// Example 1:
// Input: names = ["pes","fifa","gta","pes(2019)"]
// Output: ["pes","fifa","gta","pes(2019)"]
// Explanation: Let's see how the file system creates folder names:
// "pes" --> not assigned before, remains "pes"
// "fifa" --> not assigned before, remains "fifa"
// "gta" --> not assigned before, remains "gta"
// "pes(2019)" --> not assigned before, remains "pes(2019)"

// Example 2:
// Input: names = ["gta","gta(1)","gta","avalon"]
// Output: ["gta","gta(1)","gta(2)","avalon"]
// Explanation: Let's see how the file system creates folder names:
// "gta" --> not assigned before, remains "gta"
// "gta(1)" --> not assigned before, remains "gta(1)"
// "gta" --> the name is reserved, system adds (k), since "gta(1)" is also reserved, systems put k = 2. it becomes "gta(2)"
// "avalon" --> not assigned before, remains "avalon"

// Example 3:
// Input: names = ["onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece"]
// Output: ["onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece(4)"]
// Explanation: When the last folder is created, the smallest positive valid k is 4, and it becomes "onepiece(4)".

// Constraints:
//     1 <= names.length <= 5 * 10^4
//     1 <= names[i].length <= 20
//     names[i] consists of lowercase English letters, digits, and/or round brackets.

import "fmt"
import "strconv"

func getFolderNames(names []string) []string {
    res, mp := []string{}, make(map[string]int)
    for _, n := range names {
        if mp[n] > 0 {
            mx, ts := 0, ""
            for {
                mx = mp[n]
                mp[n]++
                ts = fmt.Sprintf("%s(%d)", n, mx)
                if mp[ts] == 0 { break }
            }
            mp[ts]++
            res = append(res, ts)
        } else {
            mp[n]++
            res = append(res, n)
        }
    }
    return res
}

func getFolderNames1(names []string) []string {
    n := len(names)
    res, mp := make([]string, n), make(map[string]int)
    for i := 0; i < n; i++ {
        if mp[names[i]] == 0 { // 不存在同名
            res[i] = names[i]
            mp[names[i]] = 1
        } else { // 存在同名
            index := mp[names[i]]
            for mp[names[i]+"("+strconv.Itoa(index)+")"] > 0 { // 取到最大文件名
                index++
            }
            res[i] = names[i]+"(" + strconv.Itoa(index) + ")"
            mp[res[i]] = 1
            mp[names[i]] = index + 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: names = ["pes","fifa","gta","pes(2019)"]
    // Output: ["pes","fifa","gta","pes(2019)"]
    // Explanation: Let's see how the file system creates folder names:
    // "pes" --> not assigned before, remains "pes"
    // "fifa" --> not assigned before, remains "fifa"
    // "gta" --> not assigned before, remains "gta"
    // "pes(2019)" --> not assigned before, remains "pes(2019)"
    fmt.Println(getFolderNames([]string{"pes","fifa","gta","pes(2019)"})) // ["pes","fifa","gta","pes(2019)"]
    // Example 2:
    // Input: names = ["gta","gta(1)","gta","avalon"]
    // Output: ["gta","gta(1)","gta(2)","avalon"]
    // Explanation: Let's see how the file system creates folder names:
    // "gta" --> not assigned before, remains "gta"
    // "gta(1)" --> not assigned before, remains "gta(1)"
    // "gta" --> the name is reserved, system adds (k), since "gta(1)" is also reserved, systems put k = 2. it becomes "gta(2)"
    // "avalon" --> not assigned before, remains "avalon"
    fmt.Println(getFolderNames([]string{"gta","gta(1)","gta","avalon"})) // ["gta","gta(1)","gta(2)","avalon"]
    // Example 3:
    // Input: names = ["onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece"]
    // Output: ["onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece(4)"]
    // Explanation: When the last folder is created, the smallest positive valid k is 4, and it becomes "onepiece(4)".
    fmt.Println(getFolderNames([]string{"onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece"})) // ["onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece(4)"]
    
    fmt.Println(getFolderNames1([]string{"pes","fifa","gta","pes(2019)"})) // ["pes","fifa","gta","pes(2019)"]
    fmt.Println(getFolderNames1([]string{"gta","gta(1)","gta","avalon"})) // ["gta","gta(1)","gta(2)","avalon"]
    fmt.Println(getFolderNames1([]string{"onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece"})) // ["onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece(4)"]
}