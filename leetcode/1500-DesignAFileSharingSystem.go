package main

// 1500. Design a File Sharing System
// We will use a file-sharing system to share a very large file which consists of m small chunks with IDs from 1 to m.

// When users join the system, the system should assign a unique ID to them. 
// The unique ID should be used once for each user, but when a user leaves the system, the ID can be reused again.

// Users can request a certain chunk of the file, the system should return a list of IDs of all the users who own this chunk. 
// If the user receives a non-empty list of IDs, they receive the requested chunk successfully.

// Implement the FileSharing class:
//     FileSharing(int m) 
//         Initializes the object with a file of m chunks.
//     int join(int[] ownedChunks): 
//         A new user joined the system owning some chunks of the file, 
//         the system should assign an id to the user which is the smallest positive integer not taken by any other user. 
//         Return the assigned id.
//     void leave(int userID): 
//         The user with userID will leave the system, you cannot take file chunks from them anymore.
//     int[] request(int userID, int chunkID): 
//         The user userID requested the file chunk with chunkID. 
//         Return a list of the IDs of all users that own this chunk sorted in ascending order.

// Example:
// Input:
// ["FileSharing","join","join","join","request","request","leave","request","leave","join"]
// [[4],[[1,2]],[[2,3]],[[4]],[1,3],[2,2],[1],[2,1],[2],[[]]]
// Output:
// [null,1,2,3,[2],[1,2],null,[],null,1]
// Explanation:
// FileSharing fileSharing = new FileSharing(4); // We use the system to share a file of 4 chunks.
// fileSharing.join([1, 2]);    // A user who has chunks [1,2] joined the system, assign id = 1 to them and return 1.
// fileSharing.join([2, 3]);    // A user who has chunks [2,3] joined the system, assign id = 2 to them and return 2.
// fileSharing.join([4]);       // A user who has chunk [4] joined the system, assign id = 3 to them and return 3.
// fileSharing.request(1, 3);   // The user with id = 1 requested the third file chunk, as only the user with id = 2 has the file, return [2] . Notice that user 1 now has chunks [1,2,3].
// fileSharing.request(2, 2);   // The user with id = 2 requested the second file chunk, users with ids [1,2] have this chunk, thus we return [1,2].
// fileSharing.leave(1);        // The user with id = 1 left the system, all the file chunks with them are no longer available for other users.
// fileSharing.request(2, 1);   // The user with id = 2 requested the first file chunk, no one in the system has this chunk, we return empty list [].
// fileSharing.leave(2);        // The user with id = 2 left the system.
// fileSharing.join([]);        // A user who doesn't have any chunks joined the system, assign id = 1 to them and return 1. Notice that ids 1 and 2 are free and we can reuse them.

// Constraints:
//     1 <= m <= 10^5
//     0 <= ownedChunks.length <= min(100, m)
//     1 <= ownedChunks[i] <= m
//     Values of ownedChunks are unique.
//     1 <= chunkID <= m
//     userID is guaranteed to be a user in the system if you assign the IDs correctly.
//     At most 10^4 calls will be made to join, leave and request.
//     Each call to leave will have a matching call for join.

// Follow-up:
//     1. What happens if the system identifies the user by their IP address instead of their unique ID and users disconnect and connect from the system with the same IP?
//     2. If the users in the system join and leave the system frequently without requesting any chunks, will your solution still be efficient?
//     3. If all users join the system one time, request all files, and then leave, will your solution still be efficient?
//     4. If the system will be used to share n files where the ith file consists of m[i], what are the changes you have to make?

import "fmt"
import "sort"

type FileSharing struct {
    seq int
    sp []int
    owned map[int][]int // k: seq v: 文件列表
    haved map[int][]int // k: 文件 v: 拥有者列表
}

func Constructor(m int) FileSharing {
    return FileSharing { 1, []int{}, map[int][]int{},  map[int][]int{}, }
}

func (fs *FileSharing) Join(ownedChunks []int) int {
    id := 0
    if len(fs.sp) != 0 {
        sort.Ints(fs.sp)
        id = fs.sp[0]
        fs.sp = fs.sp[1:]
    } else {
        id = fs.seq
        fs.seq++
    }
    sort.Ints(ownedChunks)
    oc := make([]int, len(ownedChunks))
    copy(oc, ownedChunks)
    fs.owned[id] = oc

    for _, c := range ownedChunks {
        idx := sort.SearchInts(fs.haved[c], id)
        fs.haved[c] = append(fs.haved[c], 0)
        copy(fs.haved[c][idx+1:], fs.haved[c][idx:])
        fs.haved[c][idx] = id
    }

    return id
}

func (fs *FileSharing) Leave(userID int)  {
    fs.sp = append(fs.sp, userID)
    for _, f := range fs.owned[userID] { // 遍历该用户所拥有的所有的文件
        idx := sort.SearchInts(fs.haved[f], userID)
        copy(fs.haved[f][idx:], fs.haved[f][idx+1:])
        fs.haved[f] = fs.haved[f][:len(fs.haved[f])-1]
    }
    delete(fs.owned, userID)
}

func (fs *FileSharing) Request(userID int, chunkID int) []int {
    if len(fs.haved[chunkID]) != 0 {
        res := make([]int, len(fs.haved[chunkID]))
        copy(res, fs.haved[chunkID])
        idx := sort.SearchInts(fs.owned[userID], chunkID)
        // 如果该用户已拥有该文件，无需添加
        if idx != len(fs.owned[userID]) && fs.owned[userID][idx] == chunkID {
            return res
        }
        // 加入 owned
        fs.owned[userID] = append(fs.owned[userID], 0)
        copy(fs.owned[userID][idx+1:], fs.owned[userID][idx:])
        fs.owned[userID][idx] = chunkID
        // 加入haved
        idx = sort.SearchInts(fs.haved[chunkID], userID)
        fs.haved[chunkID] = append(fs.haved[chunkID], 0)
        copy(fs.haved[chunkID][idx+1:], fs.haved[chunkID][idx:])
        fs.haved[chunkID][idx] = userID
        return res
    }
    return nil
}

/**
 * Your FileSharing object will be instantiated and called as such:
 * obj := Constructor(m);
 * param_1 := obj.Join(ownedChunks);
 * obj.Leave(userID);
 * param_3 := obj.Request(userID,chunkID);
 */

func main() {
    // FileSharing fileSharing = new FileSharing(4); // We use the system to share a file of 4 chunks.
    obj := Constructor(4)
    fmt.Println(obj)
    // fileSharing.join([1, 2]);    // A user who has chunks [1,2] joined the system, assign id = 1 to them and return 1.
    obj.Join([]int{1,2})
    fmt.Println(obj)
    // fileSharing.join([2, 3]);    // A user who has chunks [2,3] joined the system, assign id = 2 to them and return 2.
    obj.Join([]int{2,3})
    fmt.Println(obj)
    // fileSharing.join([4]);       // A user who has chunk [4] joined the system, assign id = 3 to them and return 3.
    obj.Join([]int{4})
    fmt.Println(obj)
    // fileSharing.request(1, 3);   // The user with id = 1 requested the third file chunk, as only the user with id = 2 has the file, return [2] . Notice that user 1 now has chunks [1,2,3].
    fmt.Println(obj.Request(1,3)) // [2]
    // fileSharing.request(2, 2);   // The user with id = 2 requested the second file chunk, users with ids [1,2] have this chunk, thus we return [1,2].
    fmt.Println(obj.Request(2,2)) // [1,2]
    // fileSharing.leave(1);        // The user with id = 1 left the system, all the file chunks with them are no longer available for other users.
    obj.Leave(1)
    fmt.Println(obj)
    // fileSharing.request(2, 1);   // The user with id = 2 requested the first file chunk, no one in the system has this chunk, we return empty list [].
    fmt.Println(obj.Request(2,1)) // []
    // fileSharing.leave(2);        // The user with id = 2 left the system.
    obj.Leave(2)
    fmt.Println(obj)
    // fileSharing.join([]);        // A user who doesn't have any chunks joined the system, assign id = 1 to them and return 1. Notice that ids 1 and 2 are free and we can reuse them.
    obj.Join([]int{})
    fmt.Println(obj)
}