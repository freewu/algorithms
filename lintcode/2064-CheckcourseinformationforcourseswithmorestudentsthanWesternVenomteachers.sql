-- 2064 · Check course information for courses with more students than Western Venom teachers
-- # Description
-- Write an SQL statement to obtain the number of students in each course taught by the teacher Western Venom, 
-- then query the courses with student count more than all of Western Venom taught and return the information of these courses.

-- Table Definition 1: courses

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	date	creation time
-- teacher_id	int	instructor id
-- Table Definition 2: teachers

-- column_name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- The query return column name needs to be the same case as the sample output column name.
-- If the number of students in the input data is NULL, then 0 data is returned.
-- If the query does not return any result, nothing is returned.
-- Example
-- Example I:

-- Table Contents 1: courses

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
-- Table of Contents 2: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the courses table should return.

-- id	name	student_count	created_at	teacher_id
-- 12	Dynamic Programming	2000	2018-08-18	1
-- Example 2:				
-- Table Contents 1: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- Table of Contents 2: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the courses table should return.

-- id	name	student_count	created_at	teacher_id
-- Because there is no eligible data in the input sample, only the title is shown here, no data.

SELECT
	*
FROM	
	courses
WHERE
	student_count > (
		SELECT 
			IF(COUNT(*) > 0,MAX(c.student_count),0) -- 如果没有 Western Venom 返回 0
		FROM
			courses AS c,
			teachers AS t
		WHERE
			c.teacher_id = t.id AND
			t.name = 'Western Venom'
	)