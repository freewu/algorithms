-- 2070 · Search for the name of a course created later than the creation time of any of the specified teacher's courses
-- # Description
-- Please write an SQL statement. For all courses of teacher Southern Emperor, 
-- query the course name (excluding Southern Emperor's course) that creation time is later than any course of Southern Emperor's from the courses table and the teachers table.

-- Table definition: teachers

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Table Definition: courses

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	course creation time
-- teacher_id	int	teacher id

-- If the query returns no results, nothing is returned .
-- Later than the specified time does not include the specified time.
-- Any one course does not mean [all courses], as long as it is later than one of the courses, it is considered to meet the requirements.
-- Example
-- Sample 1:

-- Table content: courses

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
-- Table Contents : Teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- name
-- System Design
-- Big Data
-- Course System Design creation time is 2020-7-18 whitch is later than Web or Senior Algorithm, so it should be selected.
-- Course Big Data creation time is 2020-9-11 whitch is later than Web or Advanced Algorithms or Object Oriented Design, so it should be selected.

-- Sample 2:

-- Table Contents: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	400	2020-6-1	4
-- 4	Web	340	2020-4-22	4
-- 5	Big Data	700	2019-9-11	1
-- 8	Data Analysis	500	2019-7-12	1
-- 10	Object Oriented Design	300	2020-8-8	4
-- 12	Dynamic Programming	2000	2018-8-18	1
-- Table Contents : Teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- name
-- Because there are no courses in the courses table with any course creation time greater than "Southern Emperor", only the title is shown here, no data.
SELECT
	c.name
FROM
	courses AS c,
	teachers AS t
WHERE
	c.teacher_id = t.id AND
	t.name != 'Southern Emperor' AND
	c.created_at > (
		SELECT 
			MIN(c1.created_at)
		FROM
			courses AS c1,
			teachers AS t1
		WHERE
			c1.teacher_id = t1.id AND
			t1.name = 'Southern Emperor'
	)