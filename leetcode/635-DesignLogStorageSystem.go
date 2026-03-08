package main

// 635. Design Log Storage System
// You are given several logs, where each log contains a unique ID and timestamp. 
// Timestamp is a string that has the following format: Year:Month:Day:Hour:Minute:Second, 
// for example, 2017:01:01:23:59:59. All domains are zero-padded decimal numbers.

// Implement the LogSystem class:
//     LogSystem() Initializes the LogSystem object.
//     void put(int id, string timestamp) Stores the given log (id, timestamp) in your storage system.
//     int[] retrieve(string start, string end, string granularity) Returns the IDs of the logs whose timestamps are within the range from start to end inclusive. start and end all have the same format as timestamp, and granularity means how precise the range should be (i.e. to the exact Day, Minute, etc.). For example, start = "2017:01:01:23:59:59", end = "2017:01:02:23:59:59", and granularity = "Day" means that we need to find the logs within the inclusive range from Jan. 1st 2017 to Jan. 2nd 2017, and the Hour, Minute, and Second for each log entry can be ignored.

// Example 1:
// Input
// ["LogSystem", "put", "put", "put", "retrieve", "retrieve"]
// [[], [1, "2017:01:01:23:59:59"], [2, "2017:01:01:22:59:59"], [3, "2016:01:01:00:00:00"], ["2016:01:01:01:01:01", "2017:01:01:23:00:00", "Year"], ["2016:01:01:01:01:01", "2017:01:01:23:00:00", "Hour"]]
// Output
// [null, null, null, null, [3, 2, 1], [2, 1]]
// Explanation
// LogSystem logSystem = new LogSystem();
// logSystem.put(1, "2017:01:01:23:59:59");
// logSystem.put(2, "2017:01:01:22:59:59");
// logSystem.put(3, "2016:01:01:00:00:00");
// // return [3,2,1], because you need to return all logs between 2016 and 2017.
// logSystem.retrieve("2016:01:01:01:01:01", "2017:01:01:23:00:00", "Year");
// // return [2,1], because you need to return all logs between Jan. 1, 2016 01:XX:XX and Jan. 1, 2017 23:XX:XX.
// // Log 3 is not returned because Jan. 1, 2016 00:00:00 comes before the start of the range.
// logSystem.retrieve("2016:01:01:01:01:01", "2017:01:01:23:00:00", "Hour");

// Constraints:
//     1 <= id <= 500
//     2000 <= Year <= 2017
//     1 <= Month <= 12
//     1 <= Day <= 31
//     0 <= Hour <= 23
//     0 <= Minute, Second <= 59
//     granularity is one of the values ["Year", "Month", "Day", "Hour", "Minute", "Second"].
//     At most 500 calls will be made to put and retrieve.

import "fmt"
import "strings"
import "strconv"

type LogSystem struct {
	list [][2]int64
}

func Constructor() LogSystem {
	return LogSystem{}
}

func (this *LogSystem) Put(id int, timestamp string) {
	st := this.split(timestamp)
	this.list = append(this.list, [2]int64{this.convert(st), int64(id)})
}

func (this *LogSystem) convert(st []int) int64 {
	if st[1] > 0 {
		st[1]--
	}
	if st[2] > 0 {
		st[2]--
	}
	return int64(st[0]-1999)*(31*12)*24*60*60 + 
	       int64(st[1])*31*24*60*60 + 
	       int64(st[2])*24*60*60 + 
	       int64(st[3])*60*60 + 
	       int64(st[4])*60 + 
	       int64(st[5])
}

func (this *LogSystem) split(s string) []int {
	parts := strings.Split(s, ":")
	result := make([]int, len(parts))
	for i, part := range parts {
		result[i], _ = strconv.Atoi(part)
	}
	return result
}

func (this *LogSystem) Retrieve(s string, e string, gra string) []int {
	res := []int{}
	start := this.Granularity(s, gra, false)
	end := this.Granularity(e, gra, true)
	for _, item := range this.list {
		if item[0] >= start && item[0] < end {
			res = append(res, int(item[1]))
		}
	}
	return res
}

func (this *LogSystem) Granularity(s string, gra string, end bool) int64 {
	h := map[string]int{
		"Year": 0, "Month": 1, "Day": 2,
		"Hour": 3, "Minute": 4, "Second": 5,
	}
	
	res := []string{"1999", "00", "00", "00", "00", "00"}
	st := this.split(s)
	
	for i := 0; i <= h[gra]; i++ {
		res[i] = strconv.Itoa(st[i])
	}
	
	t := make([]int, len(res))
	for i, str := range res {
		t[i], _ = strconv.Atoi(str)
	}
	
	if end {
		t[h[gra]]++
	}
	return this.convert(t)
}

// import "sort"

// type Time struct {
//     val string
//     id  int
// }

// type LogSystem struct {
//     times []*Time
//     id2t map[int]*Time
// }

// func Constructor() LogSystem {
//     times := []*Time{}
//     id2t := make(map[int]*Time)
//     return LogSystem{times, id2t}
// }

// func (this *LogSystem) Put(id int, timestamp string) {
//     if _, ok := this.id2t[id]; ok {
//         return
//     }
//     nt := &Time{timestamp, id}
//     this.times = append(this.times, nt)
//     this.id2t[id] = nt
//     sort.Slice(this.times, func(i, j int) bool {
//         return this.times[i].val < this.times[j].val
//     })
// }

// func (this *LogSystem) Retrieve(start string, end string, granularity string) []int {
//     res := []int{}
//     if granularity == "Second" {
//         startIndex := sort.Search(len(this.times), func(i int) bool {
//             return this.times[i].val >= start
//         })
//         endIndex := sort.Search(len(this.times), func(i int) bool {
//             return this.times[i].val > end
//         })
//         for i := startIndex; i < endIndex; i++ {
//             res = append(res, this.times[i].id)
//         }
//         return res
//     }
//     timeLen := len(start)
//     if granularity == "Minute" {
//         newStart := start[:timeLen-2] + "00"
//         newEnd := end[:timeLen-2] + "59"
//         return this.Retrieve(newStart, newEnd, "Second")
//     }
//     if granularity == "Hour" {
//         newStart := start[:timeLen-5] + "00:00"
//         newEnd := end[:timeLen-5] + "59:59"
//         return this.Retrieve(newStart, newEnd, "Second")
//     }
//     if granularity == "Day" {
//         newStart := start[:timeLen-8] + "00:00:00"
//         newEnd := end[:timeLen-8] + "23:59:59"
//         return this.Retrieve(newStart, newEnd, "Second")
//     }
//     if granularity == "Month" {
//         newStart := start[:timeLen-11] + "01:00:00:00"
//         newEnd := end[:timeLen-11] + "31:23:59:59" // 不用纠结，一个月是几天，肯定小于31天
//         return this.Retrieve(newStart, newEnd, "Second")
//     }

//     if granularity == "Year" {
//         newStart := start[:timeLen-14] + "01:00:00:00:00"
//         newEnd := end[:timeLen-14] + "12:31:23:59:59" // 不用纠结，一个月是几天，肯定小于31天
//         return this.Retrieve(newStart, newEnd, "Second")
//     }
//     return res
// }

// type pair struct {
//     id int
//     ts string
// }

// type LogSystem struct {
//     idsAndTs   []pair
//     tsWithGran map[string]int
// }

// func Constructor() LogSystem {
//     tsWithGran := map[string]int{
//         "Year":   5,
//         "Month":  8,
//         "Day":    11,
//         "Hour":   14,
//         "Minute": 17,
//         "Second": 19,
//     }
//     return LogSystem{
//         idsAndTs:   []pair{},
//         tsWithGran: tsWithGran,
//     }
// }

// func (this *LogSystem) Put(id int, timestamp string) {
//     this.idsAndTs = append(this.idsAndTs, pair{id, timestamp})
// }

// func (this *LogSystem) getTm(t string, gran string) string {
//     l := this.tsWithGran[gran]
//     return t[:l]
// }

// func (this *LogSystem) Retrieve(start string, end string, granularity string) []int {
//     s := this.getTm(start, granularity)
//     e := this.getTm(end, granularity)
//     res := []int{}
//     for _, p := range this.idsAndTs {
//         ts := this.getTm(p.ts, granularity)
//         if ts >= s && ts <= e {
//             res = append(res, p.id)
//         }
//     }
//     return res
// }

func main() {
    // Explanation
    // LogSystem logSystem = new LogSystem();
    obj := Constructor()
    fmt.Println(obj)
    // logSystem.put(1, "2017:01:01:23:59:59");
    obj.Put(1,"2017:01:01:23:59:59")
    fmt.Println(obj)
    // logSystem.put(2, "2017:01:01:22:59:59");
    obj.Put(2, "2017:01:01:22:59:59")
    fmt.Println(obj)
    // logSystem.put(3, "2016:01:01:00:00:00");
    obj.Put(3, "2016:01:01:00:00:00")
    fmt.Println(obj)
    // // return [3,2,1], because you need to return all logs between 2016 and 2017.
    // logSystem.retrieve("2016:01:01:01:01:01", "2017:01:01:23:00:00", "Year");
    fmt.Println(obj.Retrieve("2016:01:01:01:01:01", "2017:01:01:23:00:00", "Year")) // [3,2,1]
    // // return [2,1], because you need to return all logs between Jan. 1, 2016 01:XX:XX and Jan. 1, 2017 23:XX:XX.
    // // Log 3 is not returned because Jan. 1, 2016 00:00:00 comes before the start of the range.
    // logSystem.retrieve("2016:01:01:01:01:01", "2017:01:01:23:00:00", "Hour");
    fmt.Println(obj.Retrieve("2016:01:01:01:01:01", "2017:01:01:23:00:00", "Hour")) // [2,1]
}