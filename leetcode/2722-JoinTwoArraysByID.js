// 2722. Join Two Arrays by ID
// Given two arrays arr1 and arr2, return a new array joinedArray. 
// All the objects in each of the two inputs arrays will contain an id field that has an integer value. 
// joinedArray is an array formed by merging arr1 and arr2 based on their id key. 
// The length of joinedArray should be the length of unique values of id. 
// The returned array should be sorted in ascending order based on the id key.

// If a given id exists in one array but not the other, the single object with that id should be included in the result array without modification.
// If two objects share an id, their properties should be merged into a single object:
//         If a key only exists in one object, that single key-value pair should be included in the object.
//         If a key is included in both objects, the value in the object from arr2 should override the value from arr1.
 
// Example 1:
// Input: 
// arr1 = [
//     {"id": 1, "x": 1},
//     {"id": 2, "x": 9}
// ], 
// arr2 = [
//     {"id": 3, "x": 5}
// ]
// Output: 
// [
//     {"id": 1, "x": 1},
//     {"id": 2, "x": 9},
//     {"id": 3, "x": 5}
// ]
// Explanation: There are no duplicate ids so arr1 is simply concatenated with arr2.

// Example 2:
// Input: 
// arr1 = [
//     {"id": 1, "x": 2, "y": 3},
//     {"id": 2, "x": 3, "y": 6}
// ], 
// arr2 = [
//     {"id": 2, "x": 10, "y": 20},
//     {"id": 3, "x": 0, "y": 0}
// ]
// Output: 
// [
//     {"id": 1, "x": 2, "y": 3},
//     {"id": 2, "x": 10, "y": 20},
//     {"id": 3, "x": 0, "y": 0}
// ]
// Explanation: The two objects with id=1 and id=3 are included in the result array without modifiction. The two objects with id=2 are merged together. The keys from arr2 override the values in arr1.

// Example 3:
// Input: 
// arr1 = [{"id": 1, "b": {"b": 94},"v": [4, 3], "y": 48}]
// arr2 = [{"id": 1, "b": {"c": 84}, "v": [1, 3]}]
// Output: [{"id": 1, "b": {"c": 84}, "v": [1, 3], "y": 48}]
// Explanation: The two objects with id=1 are merged together. For the keys "b" and "v" the values from arr2 are used. Since the key "y" only exists in arr1, that value is taken form arr1.
 
// Constraints:
//         arr1 and arr2 are valid JSON arrays
//         Each object in arr1 and arr2 has a unique integer id key
//         2 <= JSON.stringify(arr1).length <= 10^6
//         2 <= JSON.stringify(arr2).length <= 10^6

/**
 * @param {Array} arr1
 * @param {Array} arr2
 * @return {Array}
 */
var join = function(arr1, arr2) {
    const map = new Map();
    // 循环 arr1 将 arr1 以 id 写入 map
    for (const obj of arr1) map.set(obj.id, obj);
    // 循环 arr2
    for (const obj of arr2) {
        // 如果 map 不存在 arr2.id 写入到 map
        if (!map.has(obj.id)) map.set(obj.id, obj);
        else {
            // 如果存在相同的 id 以 arr2 数级优先
            const prevObj = map.get(obj.id);
            for (const key of Object.keys(obj)) prevObj[key] = obj[key];
        }
    }
    // 循环 map 取出 value 数组
    const res = new Array();
    for (let key of map.keys()) res.push(map.get(key));
    // 按 id 排序输出
    return res.sort((a,b) => a.id-b.id); 
};

let arr1 = [
    {"id": 1, "x": 1},
    {"id": 2, "x": 9}
];
let arr2 = [
    {"id": 3, "x": 5}
];

console.log(join(arr1,arr2)); // { id: 1, x: 1 }, { id: 2, x: 9 }, { id: 3, x: 5 } ]
// [
//     {"id": 1, "x": 1},
//     {"id": 2, "x": 9},
//     {"id": 3, "x": 5}
// ]

 
arr1 = [
    {"id": 1, "x": 2, "y": 3},
    {"id": 2, "x": 3, "y": 6}
];
arr2 = [
    {"id": 2, "x": 10, "y": 20},
    {"id": 3, "x": 0, "y": 0}
]
console.log(join(arr1,arr2));
// [
//     {"id": 1, "x": 2, "y": 3},
//     {"id": 2, "x": 10, "y": 20},
//     {"id": 3, "x": 0, "y": 0}
// ]

arr1 = [{"id": 1, "b": {"b": 94},"v": [4, 3], "y": 48}]
arr2 = [{"id": 1, "b": {"c": 84}, "v": [1, 3]}]
console.log(join(arr1,arr2)); // [{"id": 1, "b": {"c": 84}, "v": [1, 3], "y": 48}]
