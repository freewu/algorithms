package main

// 2102. Sequentially Ordinal Rank Tracker
// A scenic location is represented by its name and attractiveness score, where name is a unique string among all locations and score is an integer. 
// Locations can be ranked from the best to the worst. 
// The higher the score, the better the location. 
// If the scores of two locations are equal, then the location with the lexicographically smaller name is better.

// You are building a system that tracks the ranking of locations with the system initially starting with no locations. 
// It supports:
//     1. Adding scenic locations, one at a time.
//     2. Querying the ith best location of all locations already added, 
//        where i is the number of times the system has been queried (including the current query).
//             For example, when the system is queried for the 4th time, it returns the 4th best location of all locations already added.

// Note that the test data are generated so that at any time, the number of queries does not exceed the number of locations added to the system.

// Implement the SORTracker class:
//     SORTracker() 
//         Initializes the tracker system.
//     void add(string name, int score) 
//         Adds a scenic location with name and score to the system.
//     string get() 
//         Queries and returns the ith best location, 
//         where i is the number of times this method has been invoked (including this invocation).

// Example 1:
// Input
// ["SORTracker", "add", "add", "get", "add", "get", "add", "get", "add", "get", "add", "get", "get"]
// [[], ["bradford", 2], ["branford", 3], [], ["alps", 2], [], ["orland", 2], [], ["orlando", 3], [], ["alpine", 2], [], []]
// Output
// [null, null, null, "branford", null, "alps", null, "bradford", null, "bradford", null, "bradford", "orland"]
// Explanation
// SORTracker tracker = new SORTracker(); // Initialize the tracker system.
// tracker.add("bradford", 2); // Add location with name="bradford" and score=2 to the system.
// tracker.add("branford", 3); // Add location with name="branford" and score=3 to the system.
// tracker.get();              // The sorted locations, from best to worst, are: branford, bradford.
//                             // Note that branford precedes bradford due to its higher score (3 > 2).
//                             // This is the 1st time get() is called, so return the best location: "branford".
// tracker.add("alps", 2);     // Add location with name="alps" and score=2 to the system.
// tracker.get();              // Sorted locations: branford, alps, bradford.
//                             // Note that alps precedes bradford even though they have the same score (2).
//                             // This is because "alps" is lexicographically smaller than "bradford".
//                             // Return the 2nd best location "alps", as it is the 2nd time get() is called.
// tracker.add("orland", 2);   // Add location with name="orland" and score=2 to the system.
// tracker.get();              // Sorted locations: branford, alps, bradford, orland.
//                             // Return "bradford", as it is the 3rd time get() is called.
// tracker.add("orlando", 3);  // Add location with name="orlando" and score=3 to the system.
// tracker.get();              // Sorted locations: branford, orlando, alps, bradford, orland.
//                             // Return "bradford".
// tracker.add("alpine", 2);   // Add location with name="alpine" and score=2 to the system.
// tracker.get();              // Sorted locations: branford, orlando, alpine, alps, bradford, orland.
//                             // Return "bradford".
// tracker.get();              // Sorted locations: branford, orlando, alpine, alps, bradford, orland.
//                             // Return "orland".

// Constraints:
//     name consists of lowercase English letters, and is unique among all locations.
//     1 <= name.length <= 10
//     1 <= score <= 10^5
//     At any time, the number of calls to get does not exceed the number of calls to add.
//     At most 4 * 10^4 calls in total will be made to add and get.

import "fmt"
import "container/heap"

type Location struct {
    Name string
    Score int
}
  
// Checks whether self is greater than other location
func (this *Location) Greater(other *Location) bool {
    if this.Score == other.Score { return this.Name < other.Name }
    return this.Score > other.Score
}
  
type MinHeap []*Location
func (mnh MinHeap) Len() int           { return len(mnh) }
func (mnh MinHeap) Top() interface{}   { return mnh[0] }
func (mnh MinHeap) Swap(i, j int)      { mnh[i], mnh[j] = mnh[j], mnh[i] }
func (mnh MinHeap) Less(i, j int) bool {
    if mnh[i].Score == mnh[j].Score {  return mnh[i].Name > mnh[j].Name }
    return mnh[i].Score < mnh[j].Score
}
func (mnh *MinHeap) Push(v interface{}) { *mnh = append(*mnh, v.(*Location)) }
func (mnh *MinHeap) Pop() interface{}   {
    n := len(*mnh)
    v := (*mnh)[n-1]
    *mnh = (*mnh)[:n-1]
    return v
}

type MaxHeap []*Location
func (mxh MaxHeap) Len() int           { return len(mxh) }
func (mxh MaxHeap) Top() interface{}   { return mxh[0] }
func (mxh MaxHeap) Swap(i, j int)      { mxh[i], mxh[j] = mxh[j], mxh[i] }
func (mxh MaxHeap) Less(i, j int) bool {
    if mxh[i].Score == mxh[j].Score { return mxh[i].Name < mxh[j].Name }
    return mxh[i].Score > mxh[j].Score
}
func (mxh *MaxHeap) Push(v interface{}) { *mxh = append(*mxh, v.(*Location)) }
func (mxh *MaxHeap) Pop() interface{}   {
    n := len(*mxh)
    v := (*mxh)[n-1]
    *mxh = (*mxh)[:n-1]
    return v
}

type SORTracker struct {
    Low     MinHeap
    High    MaxHeap
}

func Constructor() SORTracker {
    return SORTracker{ MinHeap{}, MaxHeap{} }
}

func (this *SORTracker) Add(name string, score int)  {
    l := &Location{ Name: name, Score: score }
    if this.Low.Len() > 0 && l.Greater(this.Low.Top().(*Location)) {
        heap.Push(&this.Low, l)
        heap.Push(&this.High, heap.Pop(&this.Low).(*Location))
    } else {
        heap.Push(&this.High, l)
    }
}

func (this *SORTracker) Get() string {
    heap.Push(&this.Low, heap.Pop(&this.High).(*Location))
    return this.Low.Top().(*Location).Name
}

/**
 * Your SORTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(name,score);
 * param_2 := obj.Get();
 */

func main() {
    // SORTracker tracker = new SORTracker(); // Initialize the tracker system.
    obj := Constructor()
    fmt.Println(obj)
    // tracker.add("bradford", 2); // Add location with name="bradford" and score=2 to the system.
    obj.Add("bradford", 2)
    fmt.Println(obj)
    // tracker.add("branford", 3); // Add location with name="branford" and score=3 to the system.
    obj.Add("branford", 3)
    fmt.Println(obj)
    // tracker.get();              // The sorted locations, from best to worst, are: branford, bradford.
    //                             // Note that branford precedes bradford due to its higher score (3 > 2).
    //                             // This is the 1st time get() is called, so return the best location: "branford".
    fmt.Println(obj.Get()) // branford
    // tracker.add("alps", 2);     // Add location with name="alps" and score=2 to the system.
    obj.Add("alps", 2)
    fmt.Println(obj)
    // tracker.get();              // Sorted locations: branford, alps, bradford.
    //                             // Note that alps precedes bradford even though they have the same score (2).
    //                             // This is because "alps" is lexicographically smaller than "bradford".
    //                             // Return the 2nd best location "alps", as it is the 2nd time get() is called.
    fmt.Println(obj.Get()) // alps
    // tracker.add("orland", 2);   // Add location with name="orland" and score=2 to the system.
    obj.Add("orland", 2)
    fmt.Println(obj)
    // tracker.get();              // Sorted locations: branford, alps, bradford, orland.
    //                             // Return "bradford", as it is the 3rd time get() is called.
    fmt.Println(obj.Get()) // bradford
    // tracker.add("orlando", 3);  // Add location with name="orlando" and score=3 to the system.
    obj.Add("orlando", 3)
    fmt.Println(obj)
    // tracker.get();              // Sorted locations: branford, orlando, alps, bradford, orland.
    //                             // Return "bradford".
    fmt.Println(obj.Get()) // bradford
    // tracker.add("alpine", 2);   // Add location with name="alpine" and score=2 to the system.
    obj.Add("alpine", 2)
    fmt.Println(obj)
    // tracker.get();              // Sorted locations: branford, orlando, alpine, alps, bradford, orland.
    //                             // Return "bradford".
    fmt.Println(obj.Get()) // bradford
    // tracker.get();              // Sorted locations: branford, orlando, alpine, alps, bradford, orland.
    //                             // Return "orland".
    fmt.Println(obj.Get()) // orland
}