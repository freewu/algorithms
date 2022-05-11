-- 2676 · Creating a simple SQL View (II)
-- # Description
-- Now, your supervisor asks you to write a simple view v_courses. 
-- You are asked to query the courses taught by teachers whose country is not USA and JP 
-- and the corresponding number of students student_count from the teachers table (teachers) and the courses table (courses).

-- Table definition: teachers (teachers table)
-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	instructor's name
-- email	varchar	Instructor's email
-- age	int	instructor's age
-- country	varchar	tutor's nationality

-- Table definition: courses
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total_students
-- created_at	datetime	Course creation time
-- teacher_id	int unsigned	instructor id

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
-- 2	System Design	1350	2020-7-18	2
-- 3	Django	780	2020-2-29	5
-- 4	Web	340	2020-4-22	4
-- 5	Big Data	700	2020-9-11	1
-- 6	Artificial Intelligence	1660	2018-5-13	3
-- 7	Java P6+	780	2019-1-19	3
-- 8	Data Analysis	500	2019-7-12	1
-- 10	Object Oriented Design	300	2020-8-8	4
-- 12	Dynamic Programming	2000	2018-8-18	2
-- After running your SQL statement, the table should return.

-- id	name	student_count
-- 2	'System Design'	1350
-- 3	'Django'	780
-- 5	'Big Data'	700
-- 8	'Data Analysis'	500
-- 12	'Dynamic Programming'	2000
-- Sample 2:

-- Table Contents : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	19	JP
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

-- id	name	student_count
-- The query is empty because the teacher_id in the sample course table is NULL, 
-- i.e., there is no teacher teaching, and the course id and student count are empty, so only the title is shown here and there is no data.

-- Write your SQL here --
CREATE VIEW 
	`v_courses`
AS 
	SELECT 
		c.id AS id,
		c.name AS name,
		c.student_count AS student_count
	FROM 
		courses  AS c
	LEFT JOIN
		teachers AS t
	ON
		c.teacher_id = t.id
	WHERE
		t.country NOT IN ('USA','JP') OR 
		t.country IS NULL;