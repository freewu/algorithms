-- 2082 · Statistics on the number of courses taught by each teacher
-- Description
-- Write a SQL statement to query the teacher table teachers and the course table courses, 
-- count the number of courses taught by each teacher, and sort the results in descending order by the number of courses, 
-- and in ascending order by the teacher's name if the number of courses is the same, 
-- and returns the teacher's name and the number of courses, and named teacher_name and course_count.

-- Table definition: teachers

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Table definition: courses

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	datetime	course creation time
-- teacher_id	int unsigned	teacher id

-- The column name returned by the query must match the column name of the sample output.
-- If the number of courses is empty,then 0 will be returned.
-- Example
-- Sample 1:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- Table Contents: courses

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

-- teacher_name	course_count
-- Western Venom	4
-- Eastern Heretic	3
-- Southern Emperor	3
-- Linghu Chong	0
-- Northern Beggar	0
-- Sample 2:

-- Table Contents : Teachers

-- id	name	email	age	country
-- 5	Linghu Chong	NULL	18	CN
-- 6	Buqun Yue	NULL	18	CN
-- Table Contents: courses

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

-- teacher_name	course_count
-- Buqun Yue	0
-- Linghu Chong	0

SELECT 
	t.name AS teacher_name,
	IFNULL(c.num,0) AS course_count
FROM
	teachers AS t
LEFT JOIN 
	(
		SELECT 
			teacher_id,
			COUNT(*) AS num
		FROM
			courses
		GROUP BY
			teacher_id
	) AS c
ON 
	t.id = c.teacher_id
ORDER BY 
	course_count DESC, teacher_name ASC