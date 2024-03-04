// 2700. Differences Between Two Objects
// Write a function that accepts two deeply nested objects or arrays obj1 and obj2 and returns a new object representing their differences.
// The function should compare the properties of the two objects and identify any changes. 
// The returned object should only contains keys where the value is different from obj1 to obj2.
// For each changed key, the value should be represented as an array [obj1 value, obj2 value]. 
// Keys that exist in one object but not in the other should not be included in the returned object.
// When comparing two arrays, the indices of the arrays are considered to be their keys. 
// The end result should be a deeply nested object where each leaf value is a difference array.
// You may assume that both objects are the output of JSON.parse.

// Example 1:
// Input: 
// obj1 = {}
// obj2 = {
//   "a": 1, 
//   "b": 2
// }
// Output: {}
// Explanation: There were no modifications made to obj1. New keys "a" and "b" appear in obj2, but keys that are added or removed should be ignored.

// Example 2:
// Input: 
// obj1 = {
//   "a": 1,
//   "v": 3,
//   "x": [],
//   "z": {
//     "a": null
//   }
// }
// obj2 = {
//   "a": 2,
//   "v": 4,
//   "x": [],
//   "z": {
//     "a": 2
//   }
// }
// Output: 
// {
//   "a": [1, 2],
//   "v": [3, 4],
//   "z": {
//     "a": [null, 2]
//   }
// }
// Explanation: The keys "a", "v", and "z" all had changes applied. "a" was changed from 1 to 2. "v" was changed from 3 to 4. "z" had a change applied to a child object. "z.a" was changed from null to 2.

// Example 3:
// Input: 
// obj1 = {
//   "a": 5, 
//   "v": 6, 
//   "z": [1, 2, 4, [2, 5, 7]]
// }
// obj2 = {
//   "a": 5, 
//   "v": 7, 
//   "z": [1, 2, 3, [1]]
// }
// Output: 
// {
//   "v": [6, 7],
//   "z": {
//     "2": [4, 3],
//     "3": {
//       "0": [2, 1]
//     }
//   }
// }
// Explanation: In obj1 and obj2, the keys "v" and "z" have different assigned values. "a" is ignored because the value is unchanged. In the key "z", there is a nested array. Arrays are treated like objects where the indices are keys. There were two alterations to the the array: z[2] and z[3][0]. z[0] and z[1] were unchanged and thus not included. z[3][1] and z[3][2] were removed and thus not included.

// Example 4:
// Input: 
// obj1 = {
//   "a": {"b": 1}, 
// }
// obj2 = {
//   "a": [5],
// }
// Output: 
// {
//   "a": [{"b": 1}, [5]]
// }
// Explanation: The key "a" exists in both objects. Since the two associated values have different types, they are placed in the difference array.

// Example 5:
// Input: 
// obj1 = {
//   "a": [1, 2, {}], 
//   "b": false
// }
// obj2 = {   
//   "b": false,
//   "a": [1, 2, {}]
// }
// Output: 
// {}
// Explanation: Apart from a different ordering of keys, the two objects are identical so an empty object is returned.
 
// Constraints:
//         obj1 and obj2 are valid JSON objects or arrays
//         2 <= JSON.stringify(obj1).length <= 10^4
//         2 <= JSON.stringify(obj2).length <= 10^4


/**
 * @param {Object|Array} obj1
 * @param {Object|Array} obj2
 * @return {Object|Array}
 */
function objDiff(obj1, obj2) {
    // 递归时 相同直接返回 {} 不加入到 res   有 判断 Object.keys(subDiff).length > 0
    if (obj1 === obj2) return {};
    if (obj1 === null || obj2 === null) return [obj1, obj2];
    // 存在一个为对象 一个不为空的时的处理)
    if (typeof obj1 !== 'object' || typeof obj2 !== 'object') return [obj1, obj2];
    if (Array.isArray(obj1) !== Array.isArray(obj2)) return [obj1, obj2];
 
    const res = {};
    for (const key in obj1) { // 循环 obj1 
        if (key in obj2) { // 如果 obj2 存在相关 key 需要比较相关值
            // 递归到值
            const subDiff = objDiff(obj1[key], obj2[key]);
            // 如果返回的不是空对象 说明存在差值
            if (Object.keys(subDiff).length > 0) res[key] = subDiff;
        }
    }
    return res;
};

// best solution
function objDiff1(obj1, obj2) {
    if((Array.isArray(obj1) && !Array.isArray(obj2)) || 
        (!Array.isArray(obj1) && Array.isArray(obj2))) {
            return [obj1, obj2]
    }
    const res = {};
    Object.keys(obj1).forEach(key => {
        if(obj2[key] !== undefined) {
            if(obj1[key] instanceof Object && obj2[key] instanceof Object) {
                let diffRes = objDiff(obj1[key], obj2[key]);
                if(Object.keys(diffRes).length > 0) {
                    res[key] = diffRes;
                }
            } else {
                if(obj1[key] !== obj2[key]) {
                    res[key] = [obj1[key], obj2[key]];
                }
            }
        }
    })
    return res;
};

// Example 1:
console.log(objDiff(
    {},
    {"a": 1,"b": 2}
)); // {}

// Example 2:
console.log(objDiff(
    { "a": 1, "v": 3, "x": [], "z": { "a": null} },
    { "a": 2, "v": 4, "x": [], "z": { "a": 2 } }
)); // {"a": [1, 2], "v": [3, 4],"z": { "a": [null, 2] } }

// Example 3:
console.log(objDiff(
    { "a": 5, "v": 6, "z": [1, 2, 4, [2, 5, 7]]},
    { "a": 5, "v": 7, "z": [1, 2, 3, [1]]}
)); // { "v": [6, 7], "z": {"2": [4, 3], "3": { "0": [2, 1] } } }

// Example 4:
console.log(objDiff(
    { "a": {"b": 1} },
    { "a": [5] }
)); // { "a": [{"b": 1}, [5]] }

// Example 5:
console.log(objDiff(
    { "a": [1, 2, {}], "b": false },
    { "b": false, "a": [1, 2, {}] }
)); // {}


// Example 1:
console.log(objDiff1(
    {},
    {"a": 1,"b": 2}
)); // {}

// Example 2:
console.log(objDiff1(
    { "a": 1, "v": 3, "x": [], "z": { "a": null} },
    { "a": 2, "v": 4, "x": [], "z": { "a": 2 } }
)); // {"a": [1, 2], "v": [3, 4],"z": { "a": [null, 2] } }

// Example 3:
console.log(objDiff1(
    { "a": 5, "v": 6, "z": [1, 2, 4, [2, 5, 7]]},
    { "a": 5, "v": 7, "z": [1, 2, 3, [1]]}
)); // { "v": [6, 7], "z": {"2": [4, 3], "3": { "0": [2, 1] } } }

// Example 4:
console.log(objDiff1(
    { "a": {"b": 1} },
    { "a": [5] }
)); // { "a": [{"b": 1}, [5]] }

// Example 5:
console.log(objDiff1(
    { "a": [1, 2, {}], "b": false },
    { "b": false, "a": [1, 2, {}] }
)); // {}