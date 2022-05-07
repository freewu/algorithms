-- 2079 · Count the total number of students in each teacher's course
-- # Description
-- Write an SQL statement to count the total number of students in each teacher's course.
-- Return 0 if the teacher taught nothing or the number of students in the course were 0. 
-- And alias the returned teacher's name and student count as teacher_name and student_count.
-- Table Definition: teachers (teachers table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	instructor's name
-- email	varchar	instructor's email
-- age	int	instructor's age
-- country	varchar	instructor's nationality
-- Table definition: courses

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	datetime	course creation time
-- teacher_id	int unsigned	instructor id

-- If the data in the teachers table is null, the returned data is also null.
-- Example
-- Sample I:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- Table Contents: Courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-6-1	4
-- 2	System Design	1350	2020-7-18	3
-- 3	Django	780	2020-2-29	3
-- 4	Web	340	2020-4-22	4
-- 5	Big Data	700	2020-9-11	1
-- 6	Artificial Intelligence	1660	2018-5-13	3
-- 7	Java P6+	780	2019-1-19	3
-- 8	Data Analysis	500	2019-7-12	1
-- 10	Object Oriented Design	300	2020-8-8	4
-- 12	Dynamic Programming	2000	2018-8-18	1
-- After running your SQL statement, the table should return.

-- teacher_name	student_count
-- Eastern Heretic	3200
-- Northern Beggar	0
-- Western Venom	4570
-- Southern Emperor	1520
-- Linghu Chong	0
-- **Example 2

-- Table Contents : Teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- Table Contents: Courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-6-1	NULL
-- 2	System Design	1350	2020-7-18	NULL
-- 3	Django	780	2020-2-29	NULL
-- 4	Web	340	2020-4-22	NULL
-- 5	Big Data	700	2020-9-11	NULL
-- 6	Artificial Intelligence	1660	2018-5-13	NULL
-- 7	Java P6+	780	2019-1-19	NULL
-- 8	Data Analysis	500	2019-7-12	NULL
-- 10	Object Oriented Design	300	2020-8-8	NULL
-- 12	Dynamic Programming	2000	2018-8-18	NULL
-- After running your SQL statement, the table should return.

-- teacher_name	student_count
-- Eastern Heretic	0
-- Northern Beggar	0
-- Western Venom	0
-- Southern Emperor 0
-- Linghu Chong     0
-- If the teacher_id column in the courses table is all null, it means there is no teacher teaching, 
-- so it returns the teacher name and the total number of students in the teachers table as 0.

SELECT
	t.name AS teacher_name,
	IFNULL(SUM(c.student_count),0) AS student_count
FROM
	teachers AS t
LEFT JOIN
	courses AS c
ON 
	t.id = c.teacher_id
GROUP BY
	t.id