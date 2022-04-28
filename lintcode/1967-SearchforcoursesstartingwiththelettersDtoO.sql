-- 1967 · Search for courses starting with the letters 'D' to 'O'
-- # Description
-- Write a SQL statement to query the course names which are start with the capital letters 'D' to 'O' in the course table courses (including the single characters 'D' and 'O').

-- Table definition : courses

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	courses' name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id
-- Please use BETWEEN ... AND ... keyword to solve the question
-- If there is no query results, nothing will be returned
-- Example
-- Example 1:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1	4
-- 2	System Design	1350	2020-7-18	3
-- 3	Django	780	2020-2-29	3
-- 4	Web	340	2020-4-22	4
-- 5	Big Data	700	2020-9-11	1
-- 6	Artificial Intelligence	1660	2018-5-13	3
-- 7	Java P6+	780	2019-1-19	3
-- 8	Data Analysis	500	2019-7-12	1
-- 10	Object Oriented Design	300	2020-8-8	4
-- 12	Dynamic Programming	2000	2018-8-18	1
-- After running your SQL statement, the table should return:

-- name
-- Django
-- Java P6+
-- Data Analysis
-- Object Oriented Design
-- Dynamic Programming
-- Example 2:

-- Table content : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-6-1	4
-- 2	Big Data	700	2020-9-11	1
-- 3	Artificial Intelligence	1660	2018-5-13	3
-- 4	Zookeeper	999	2019-8-19	2
-- 5	P	90	2019-8-19	3
-- After running your SQL statement, the table should return:

-- name
-- Since none of the input data starts with a course name from 'D' to 'O', the output is only the title, not the data.

-- Challenge
-- Can you use regular expressions to solve this problem?

-- use SUBSTR
SELECT
	name
FROM
	courses
WHERE
	SUBSTR(name,1,1) BETWEEN 'D' AND 'O'

-- use REGEXP
SELECT
	name
FROM
	courses
WHERE
	name REGEXP '^[D-O]'

-- MySQL 正则表达式
-- 
-- ^	匹配输入字符串的开始位置。如果设置了 RegExp 对象的 Multiline 属性，^ 也匹配 '\n' 或 '\r' 之后的位置。
-- $	匹配输入字符串的结束位置。如果设置了RegExp 对象的 Multiline 属性，$ 也匹配 '\n' 或 '\r' 之前的位置。
-- .	匹配除 "\n" 之外的任何单个字符。要匹配包括 '\n' 在内的任何字符，请使用像 '[.\n]' 的模式。
-- [...]	字符集合。匹配所包含的任意一个字符。例如， '[abc]' 可以匹配 "plain" 中的 'a'。
-- [^...]	负值字符集合。匹配未包含的任意字符。例如， '[^abc]' 可以匹配 "plain" 中的'p'。
-- p1|p2|p3	匹配 p1 或 p2 或 p3。例如，'z|food' 能匹配 "z" 或 "food"。'(z|f)ood' 则匹配 "zood" 或 "food"。
-- *	匹配前面的子表达式零次或多次。例如，zo* 能匹配 "z" 以及 "zoo"。* 等价于{0,}。
-- +	匹配前面的子表达式一次或多次。例如，'zo+' 能匹配 "zo" 以及 "zoo"，但不能匹配 "z"。+ 等价于 {1,}。
-- {n}	n 是一个非负整数。匹配确定的 n 次。例如，'o{2}' 不能匹配 "Bob" 中的 'o'，但是能匹配 "food" 中的两个 o。
-- {n,m}	m 和 n 均为非负整数，其中n <= m。最少匹配 n 次且最多匹配 m 次。

-- 查找name字段中以'st'为开头的所有数据：
-- mysql> SELECT name FROM person_tbl WHERE name REGEXP '^st';

-- 查找name字段中以'ok'为结尾的所有数据：
-- mysql> SELECT name FROM person_tbl WHERE name REGEXP 'ok$';

-- 查找name字段中包含'mar'字符串的所有数据：
-- mysql> SELECT name FROM person_tbl WHERE name REGEXP 'mar';

-- 查找name字段中以元音字符开头或以'ok'字符串结尾的所有数据：
-- mysql> SELECT name FROM person_tbl WHERE name REGEXP '^[aeiou]|ok$';