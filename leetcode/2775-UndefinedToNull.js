// 2775. Undefined to Null
// Write a function called undefinedToNull that takes a deeply nested object or array obj, 
// and then creates a copy of that object with any undefined values replaced by null.
// undefined values are handled differently than null values when objects are converted to a JSON string using JSON.stringify(). 
// This function helps ensure serialized data is free of unexpected errors.

// Example 1:
// Input: obj = {"a": undefined, "b": 3}
// Output: {"a": null, "b": 3}
// Explanation: The value for obj.a has been changed from undefined to null

// Example 2:
// Input: obj = {"a": undefined, "b": ["a", undefined]}
// Output: {"a": null,"b": ["a", null]}
// Explanation: The values for obj.a and obj.b[1] have been changed from undefined to null

// Constraints:
//         obj is a valid JSON object or array
//         2 <= JSON.stringify(obj).length <= 10^5

/**
 * @param {Object|Array} obj
 * @return {Object|Array}
 */
var undefinedToNull = function(obj) {
    // for(o in obj) {
    //     if(Array.isArray(obj[o])) obj[o] = undefinedToNull(obj[o]);
    //     obj[o] = (obj[o] === undefined)? null : obj[o]
    // }
    // return obj;
    if(!obj) return null
    const keys = Object.keys(obj);
    for(let key of keys) {
        if(obj[key] === undefined) obj[key] = null;
        // 如果是对象 递归处理
        if(typeof obj[key] === 'object') obj[key] = undefinedToNull(obj[key])
    }
    return obj;
};


console.log(undefinedToNull({"a": undefined, "b": 3})) // {"a": null, "b": 3}
console.log(undefinedToNull([undefined, undefined])) // [null, null] 
