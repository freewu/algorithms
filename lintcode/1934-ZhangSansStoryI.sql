-- 1934 · Zhang San's Story I
-- # Description
-- Zhang San was admitted to a well-known university, and reporters came to Zhang San's school for an interview.
-- The students table records the student's name and class (class_id). Please write a SQL statement to find out the names of all students with the surname "zhang" from the students table (students).
-- The habit of Chinese names is the first name and the last name, this question requires to find the string beginning with "zhang"

-- Table definition: students

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	student name
-- class_id	int	student's class
-- Example
-- Example 1:

-- Table content: students

-- id	name	class_id
-- 1	zhangsan	2
-- 2	lisi	1
-- 3	wanger	4
-- 4	zhaoliu	1
-- 5	niuniu	2
-- 6	zhangfei	3
-- 7	guanyu	4
-- 8	liubei	3
-- 9	linqi	2
-- After running your SQL statement, the table should return:

-- name
-- zhangsan
-- zhangfei
-- Example 2:

-- Table content: students

-- id	name	class_id
-- 1	zhangsan	2
-- 2	lisi	1
-- 3	wanger	4
-- 4	zhaoliu	1
-- 5	niuniu	2
-- After running your SQL statement, the table should return:

-- name
-- zhangsan
SELECT 
	name
FROM 
	students
WHERE
	name LIKE "zhang%"