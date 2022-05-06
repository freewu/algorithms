-- 2054 · Check the course name and the age of the corresponding instructor
-- # Description
-- Please write a SQL statement to make a full join between the course table courses and the teacher table teachers, 
-- querying the course name and the age of the corresponding teacher,
-- with the result columns named course name course_name and teacher’s age teacher_age respectively.

-- Course table: courses
-- Column Name	Type	Comments
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	total number of students
-- created_at	date	course start time
-- teacher_id	int	teacher id

-- Teacher table: teachers
-- column_name	type	comment
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- MySQL database does not support full joins. To achieve full joins you need to combine the results of left and right joins together (query result de-duplication).
-- If the query of both table queries are empty, nothing is returned.
-- Example
-- Sample I

-- Table Contents 1 : Courses (courses table)

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 9	Object Oriented Design	300	2020-08-08	4
-- 10	Dynamic Programming	2000	2018-08-18	1
-- Table Contents 2 : teachers (Teachers table)

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- course_name	teacher_age
-- Advanced Algorithms	21
-- System Design	28
-- Django	28
-- Web Southern	21
-- Big Data	20
-- Artificial Intelligence	28
-- Java P6+	28
-- Data Analysis Eastern	20
-- Object Oriented Design	21
-- Dynamic Programming	20
-- Linghu Chong	18
-- NULL	21
-- NULL	18
-- Example 2:	
-- Table Contents 1: courses (courses table)	
-- id	name	student_count	created_at	teacher_id
-- 1	NULL	880	2020-06-01	4
-- 2	NULL	1350	2020-07-18	3
-- 3	NULL	780	2020-02-29	3
-- 4	NULL	340	2020-04-22	1
-- 5	NULL	700	2020-09-11	1
-- Table Contents 2 : teachers (Teachers table)

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	NULL	UK
-- 2	Northern Beggar	northern.beggar@qq.com	NULL	CN
-- 3	Western Venom	western.venom@163.com	NULL	USA
-- 4	Southern Emperor	southern.emperor@qq.com	NULL	JP
-- 5	Linghu Chong	NULL	NULL	CN
-- After running your SQL statement, the table should return.

-- course_name	teacher_age
-- Because there is no data for the query in the sample, only the title is shown here, no data.
SELECT 
	p.course_name AS course_name,
	p.teacher_age AS teacher_age
FROM 
	(
		(
			SELECT
				c.name AS course_name,
				t.age AS teacher_age
			FROM
				teachers AS t
			LEFT JOIN 
				courses  AS c
			ON
				t.id = c.teacher_id
		) 
		UNION
		(
			SELECT
				c.name AS course_name,
				t.age AS teacher_age
			FROM
				teachers AS t
			RIGHT JOIN 
				courses  AS c
			ON
				t.id = c.teacher_id
		)
	) AS p
GROUP BY 
	p.course_name,p.teacher_age -- query result de-duplication