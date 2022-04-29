-- 1994 · Update all course information
-- # Description
-- Write an SQL statement to change the number of students to 0 for all courses in the course table.

-- Table definition: courses

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	datetime	course start time
-- teacher_id	int	teacher id
-- Contact me on wechat to get more FLAMG requent Interview questions . (wechat id : jiuzhang15)


-- Example
-- Table content: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- 12	Dynamic Programming	2000	2018-08-18	1
-- After running your SQL statement, we will execute the following statement:				
-- SELECT *
-- FROM `courses`;
-- return to:

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	0	2020-06-01	4
-- 2	System Design	0	2020-07-18	3
-- 3	Django	0	2020-02-29	3
-- 4	Web	0	2020-04-22	4
-- 5	Big Data	0	2020-09-11	1
-- 6	Artificial Intelligence	0	2018-05-13	3
-- 7	Java P6+	0	2019-01-19	3
-- 8	Data Analysis	0	2019-07-12	1
-- 10	Object Oriented Design	0	2020-08-08	4
-- 12	Dynamic Programming	0	2018-08-18	1
UPDATE 
	courses
SET
	student_count = 0