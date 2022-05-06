-- 2074 · Search for the names of teachers whose average number of students in classes taught by all teachers exceeds the average number of students in all courses
-- # Description
-- Write an SQL statement that query the namename of teachers whose average number of students of all courses 
-- they taught more than the average of all courses from the course table courses.

-- Table Definition 1: teachers

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Table Definition 2: courses (Course Table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	course creation time
-- teacher_id	int unsigned	instructor id

-- The column name returned by the query needs to be the same as the case of the column name output by the sample.
-- If the input data is null, NULL is returned.
-- Example
-- Example I

-- Table Contents 1: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- Table Contents 2: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Java P6+	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- 12	Dynamic Programming	2000	2018-08-18	1
-- After running your SQL statement, the table should return.

-- name
-- Eastern Heretic
-- Western Venom
-- Example II

-- Table Contents 1: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	linghu.chong@linghuchong.com	18	CN
-- 6	Jia He	jia.he@163.com	22	CN
-- 7	Men Tu	men.tu@guju.com	26	CN
-- Table Contents 2: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Web	800	2019-8-9	1
-- 2	Database	1440	2018-10-8	1
-- 3	cloud computing	850	2020-4-5	1
-- 4	C++	970	2020-5-28	1
-- 5	virtual reality	350	2020-11-21	1
-- After running your SQL statement, the table should return.

-- name
-- There is no eligible data in Example 2, so the output contains only table headers and no data.
SELECT
	name
FROM
	teachers
WHERE
	id IN (
		SELECT 
			p.teacher_id
		FROM
			(
				SELECT
					c.teacher_id AS teacher_id,
					(SUM(c.student_count) / COUNT(*)) AS avg_num
				FROM
					courses  AS c
				GROUP BY 
					c.teacher_id
			) AS p
		WHERE
			p.avg_num > ( -- 取平均数
				SELECT
					IF(COUNT(*) > 0,SUM(student_count) / COUNT(*),0)
				FROM
					courses
			)
		
	)