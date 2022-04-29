-- 1992 · Check average course enrollment
-- # Description
-- Write an SQL statement to query the average of the course count student_count in the course table courses, 
-- return the field name avg_student_count, and round the result to two decimal places.
-- Table definition: courses

-- | column name | type | comment |
-- | ------------- | -------- | -------- |a
-- | id | int | primary key |
-- | name | varchar | course name |
-- | student_count | int | number of students |
-- | created_at | date | class start time |
-- | teacher_id | int | teacher id |

-- Contact me on wechat to get more FLAMG requent Interview questions . (wechat id : jiuzhang15)


-- The column names returned by the query need to be the same as the case of the column names output by the sample
-- If the number of students in the input data is NULL, the data will be skipped.
-- If the number of students in the input data is all NULL, or the input data is empty, then return NULL
-- Example
-- Sample I:

-- Table Contents : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020/6/1	4
-- 2	System Design	1350	2020/7/18	3
-- 3	Django	780	2020/2/29	3
-- 4	Web	340	2020/4/22	4
-- 5	Big Data	700	2020/9/11	1
-- 6	Artificial Intelligence	1660	2018/5/13	3
-- 7	Java P6+	780	2019/1/19	3
-- 8	Data Analysis	500	2019/7/12	1
-- 10	Object Oriented Design	300	2020/8/8	4
-- 12	Dynamic Programming	2000	2018/8/18	1
-- After running your SQL statement, the table should return.

-- avg_student_count
-- 929.00
-- Sample 2:

-- Table Contents : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	NULL	2020-6-1	4
-- 2	System Design	NULL	2020-7-18	3
-- 3	Django	NULL	2020-2-29	3
-- 4	Web	NULL	2020-4-22	4
-- 5	Big Data	NULL	2020-9-11	1
-- After running your SQL statement, the table should return.

-- avg_student_count
-- NULL
-- Because the student count in sample 2 is NULL, the average result returned is also NULL
SELECT
	TRUNCATE(SUM(student_count) / COUNT(1),2) AS avg_student_count
FROM
	courses
WHERE
	student_count IS NOT NULL