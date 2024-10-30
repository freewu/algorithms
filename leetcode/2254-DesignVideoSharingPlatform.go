package main

// 2254. Design Video Sharing Platform
// You have a video sharing platform where users can upload and delete videos. 
// Each video is a string of digits, where the ith digit of the string represents the content of the video at minute i. 
// For example, the first digit represents the content at minute 0 in the video, the second digit represents the content at minute 1 in the video, and so on. 
// Viewers of videos can also like and dislike videos. 
// Internally, the platform keeps track of the number of views, likes, and dislikes on each video.

// When a video is uploaded, it is associated with the smallest available integer videoId starting from 0. 
// Once a video is deleted, the videoId associated with that video can be reused for another video.

// Implement the VideoSharingPlatform class:
//     VideoSharingPlatform() 
//         Initializes the object.
//     int upload(String video) 
//         The user uploads a video. Return the videoId associated with the video.
//     void remove(int videoId) 
//         If there is a video associated with videoId, remove the video.
//     String watch(int videoId, int startMinute, int endMinute) 
//         If there is a video associated with videoId, increase the number of views on the video by 1 
//         and return the substring of the video string starting at startMinute 
//         and ending at min(endMinute, video.length - 1) (inclusive). \
//         Otherwise, return "-1".
//     void like(int videoId) 
//         Increases the number of likes on the video associated with videoId by 1 
//         if there is a video associated with videoId.
//     void dislike(int videoId) 
//         Increases the number of dislikes on the video associated with videoId by 1 
//         if there is a video associated with videoId.
//     int[] getLikesAndDislikes(int videoId) 
//         Return a 0-indexed integer array values of length 2 where values[0] is the number of likes 
//         and values[1] is the number of dislikes on the video associated with videoId. 
//         If there is no video associated with videoId, return [-1].
//     int getViews(int videoId) 
//         Return the number of views on the video associated with videoId, 
//         if there is no video associated with videoId, return -1.

// Example 1:
// Input
// ["VideoSharingPlatform", "upload", "upload", "remove", "remove", "upload", "watch", "watch", "like", "dislike", "dislike", "getLikesAndDislikes", "getViews"]
// [[], ["123"], ["456"], [4], [0], ["789"], [1, 0, 5], [1, 0, 1], [1], [1], [1], [1], [1]]
// Output
// [null, 0, 1, null, null, 0, "456", "45", null, null, null, [1, 2], 2]
// Explanation
// VideoSharingPlatform videoSharingPlatform = new VideoSharingPlatform();
// videoSharingPlatform.upload("123");          // The smallest available videoId is 0, so return 0.
// videoSharingPlatform.upload("456");          // The smallest available videoId is 1, so return 1.
// videoSharingPlatform.remove(4);              // There is no video associated with videoId 4, so do nothing.
// videoSharingPlatform.remove(0);              // Remove the video associated with videoId 0.
// videoSharingPlatform.upload("789");          // Since the video associated with videoId 0 was deleted,
//                                              // 0 is the smallest available videoId, so return 0.
// videoSharingPlatform.watch(1, 0, 5);         // The video associated with videoId 1 is "456".
//                                              // The video from minute 0 to min(5, 3 - 1) = 2 is "456", so return "456".
// videoSharingPlatform.watch(1, 0, 1);         // The video associated with videoId 1 is "456".
//                                              // The video from minute 0 to min(1, 3 - 1) = 1 is "45", so return "45".
// videoSharingPlatform.like(1);                // Increase the number of likes on the video associated with videoId 1.
// videoSharingPlatform.dislike(1);             // Increase the number of dislikes on the video associated with videoId 1.
// videoSharingPlatform.dislike(1);             // Increase the number of dislikes on the video associated with videoId 1.
// videoSharingPlatform.getLikesAndDislikes(1); // There is 1 like and 2 dislikes on the video associated with videoId 1, so return [1, 2].
// videoSharingPlatform.getViews(1);            // The video associated with videoId 1 has 2 views, so return 2.

// Example 2:
// Input
// ["VideoSharingPlatform", "remove", "watch", "like", "dislike", "getLikesAndDislikes", "getViews"]
// [[], [0], [0, 0, 1], [0], [0], [0], [0]]
// Output
// [null, null, "-1", null, null, [-1], -1]
// Explanation
// VideoSharingPlatform videoSharingPlatform = new VideoSharingPlatform();
// videoSharingPlatform.remove(0);              // There is no video associated with videoId 0, so do nothing.
// videoSharingPlatform.watch(0, 0, 1);         // There is no video associated with videoId 0, so return "-1".
// videoSharingPlatform.like(0);                // There is no video associated with videoId 0, so do nothing.
// videoSharingPlatform.dislike(0);             // There is no video associated with videoId 0, so do nothing.
// videoSharingPlatform.getLikesAndDislikes(0); // There is no video associated with videoId 0, so return [-1].
// videoSharingPlatform.getViews(0);            // There is no video associated with videoId 0, so return -1.

// Constraints:
//     1 <= video.length <= 10^5
//     The sum of video.length over all calls to upload does not exceed 10^5
//     video consists of digits.
//     0 <= videoId <= 10^5
//     0 <= startMinute < endMinute < 10^5
//     startMinute < video.length
//     The sum of endMinute - startMinute over all calls to watch does not exceed 10^5.
//     At most 10^5 calls in total will be made to all functions.

import "fmt"
import "container/heap"

type MinHeap struct {
    arr []int
}

func (m *MinHeap) Len() int { return len(m.arr) }
func (m *MinHeap) Less(i, j int) bool { return m.arr[i] < m.arr[j] }
func (m *MinHeap) Swap(i, j int) { m.arr[i], m.arr[j] = m.arr[j], m.arr[i] }
func (m *MinHeap) Push(x interface{}) { m.arr = append(m.arr, x.(int)) }
func (m *MinHeap) Pop() interface{} {
    top := m.arr[len(m.arr)-1]
    m.arr = m.arr[:len(m.arr)-1]
    return top
}

// 用堆保持当前可用的id，hashmap存video对应关系即可
type VideoSharingPlatform struct { 
    g *MinHeap
    v map[int] *struct {
        content              string
        like, dislike, views int
    }
}

func Constructor() VideoSharingPlatform {
    no := make([]int, 10001)
    for i := 0; i < len(no); i++ {
        no[i] = i
    }
    h := &MinHeap{ arr: no }
    heap.Init(h)
    return VideoSharingPlatform{g: h,v: map[int]*struct {
        content              string
        like, dislike, views int
    }{}}
}

func (p *VideoSharingPlatform) Upload(video string) int {
    videoId:= heap.Pop(p.g).(int)
    p.v[videoId] = &struct {
        content              string
        like, dislike, views int
    }{ content: video, like: 0, dislike:0, views: 0 }
    return videoId
}

func (p *VideoSharingPlatform) Remove(videoId int) {
    if p.v[videoId] == nil { return  }
    delete(p.v, videoId)
    heap.Push(p.g, videoId)
}

func (p *VideoSharingPlatform) Watch(videoId int, startMinute int, endMinute int) string {
    t := p.v[videoId]
    if t == nil { return "-1" }
    if endMinute > len(t.content) - 1{
        endMinute = len(t.content) - 1
    }
    p.v[videoId].views++
    return t.content[startMinute:endMinute + 1]
}

func (p *VideoSharingPlatform) Like(videoId int) {
    if p.v[videoId] == nil { return }
    p.v[videoId].like++
}

func (p *VideoSharingPlatform) Dislike(videoId int) {
    if p.v[videoId] == nil { return }
    p.v[videoId].dislike++
}

func (p *VideoSharingPlatform) GetLikesAndDislikes(videoId int) []int {
    if p.v[videoId] == nil{ return []int{ -1 } }
    return []int{p.v[videoId].like, p.v[videoId].dislike}
}

func (p *VideoSharingPlatform) GetViews(videoId int) int {
    if p.v[videoId] == nil { return -1 }
    return p.v[videoId].views
}

/**
 * Your VideoSharingPlatform object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Upload(video);
 * obj.Remove(videoId);
 * param_3 := obj.Watch(videoId,startMinute,endMinute);
 * obj.Like(videoId);
 * obj.Dislike(videoId);
 * param_6 := obj.GetLikesAndDislikes(videoId);
 * param_7 := obj.GetViews(videoId);
 */

func main() {
    // Example 1:
    // Input
    // ["VideoSharingPlatform", "upload", "upload", "remove", "remove", "upload", "watch", "watch", "like", "dislike", "dislike", "getLikesAndDislikes", "getViews"]
    // [[], ["123"], ["456"], [4], [0], ["789"], [1, 0, 5], [1, 0, 1], [1], [1], [1], [1], [1]]
    // Output
    // [null, 0, 1, null, null, 0, "456", "45", null, null, null, [1, 2], 2]
    // Explanation
    // VideoSharingPlatform videoSharingPlatform = new VideoSharingPlatform();
    obj := Constructor()
    fmt.Println(obj)
    // videoSharingPlatform.upload("123");          // The smallest available videoId is 0, so return 0.
    fmt.Println(obj.Upload("123")) // 0
    fmt.Println(obj)
    // videoSharingPlatform.upload("456");          // The smallest available videoId is 1, so return 1.
    fmt.Println(obj.Upload("456")) // 1
    fmt.Println(obj)
    // videoSharingPlatform.remove(4);              // There is no video associated with videoId 4, so do nothing.
    obj.Remove(4)
    fmt.Println(obj)
    // videoSharingPlatform.remove(0);              // Remove the video associated with videoId 0.
    obj.Remove(0)
    fmt.Println(obj)
    // videoSharingPlatform.upload("789");          // Since the video associated with videoId 0 was deleted,
    //                                              // 0 is the smallest available videoId, so return 0.
    fmt.Println(obj.Upload("789")) // 0
    fmt.Println(obj)
    // videoSharingPlatform.watch(1, 0, 5);         // The video associated with videoId 1 is "456".
    //                                              // The video from minute 0 to min(5, 3 - 1) = 2 is "456", so return "456".
    fmt.Println(obj.Watch(1, 0, 5)) // "456"
    fmt.Println(obj)
    // videoSharingPlatform.watch(1, 0, 1);         // The video associated with videoId 1 is "456".
    //                                              // The video from minute 0 to min(1, 3 - 1) = 1 is "45", so return "45".
    fmt.Println(obj.Watch(1, 0, 1)) // "456"
    fmt.Println(obj)
    // videoSharingPlatform.like(1);                // Increase the number of likes on the video associated with videoId 1.
    obj.Like(1)
    fmt.Println(obj)
    // videoSharingPlatform.dislike(1);             // Increase the number of dislikes on the video associated with videoId 1.
    obj.Dislike(1)
    fmt.Println(obj)
    // videoSharingPlatform.dislike(1);             // Increase the number of dislikes on the video associated with videoId 1.
    obj.Dislike(1)
    fmt.Println(obj)
    // videoSharingPlatform.getLikesAndDislikes(1); // There is 1 like and 2 dislikes on the video associated with videoId 1, so return [1, 2].
    fmt.Println(obj.GetLikesAndDislikes(1)) // [1, 2].
    // videoSharingPlatform.getViews(1);            // The video associated with videoId 1 has 2 views, so return 2.
    fmt.Println(obj.GetViews(1)) // 2

    // Example 2:
    // Input
    // ["VideoSharingPlatform", "remove", "watch", "like", "dislike", "getLikesAndDislikes", "getViews"]
    // [[], [0], [0, 0, 1], [0], [0], [0], [0]]
    // Output
    // [null, null, "-1", null, null, [-1], -1]
    // Explanation
    // VideoSharingPlatform videoSharingPlatform = new VideoSharingPlatform();
    obj1 := Constructor()
    fmt.Println(obj1)
    // videoSharingPlatform.remove(0);              // There is no video associated with videoId 0, so do nothing.
    obj1.Remove(0)
    fmt.Println(obj1)
    // videoSharingPlatform.watch(0, 0, 1);         // There is no video associated with videoId 0, so return "-1".
    fmt.Println(obj1.Watch(0, 0, 1)) // "-1"
    fmt.Println(obj1)
    // videoSharingPlatform.like(0);                // There is no video associated with videoId 0, so do nothing.
    obj1.Like(0)
    fmt.Println(obj1)
    // videoSharingPlatform.dislike(0);             // There is no video associated with videoId 0, so do nothing.
    obj1.Dislike(0)
    fmt.Println(obj1)
    // videoSharingPlatform.getLikesAndDislikes(0); // There is no video associated with videoId 0, so return [-1].
    fmt.Println(obj.GetLikesAndDislikes(0)) // [-1]
    // videoSharingPlatform.getViews(0);            // There is no video associated with videoId 0, so return -1.
    fmt.Println(obj.GetViews(0)) // -1
}