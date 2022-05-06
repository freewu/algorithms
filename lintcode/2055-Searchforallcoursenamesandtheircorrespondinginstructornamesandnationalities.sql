-- 2055 · Search for all course names and their corresponding instructor names and nationalities
-- # Description
-- Write an SQL statement to full outer join the courses table courses and the teachers table teachers,
-- then query all course names and their corresponding teacher's names and nationalities, 
-- with the result columns named course name course_name, teacher name teacher_name and teacher nationality teacher_country respectively.

-- Table Definition 1: courses (course table)

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id
-- Table Definition 2 : teachers (teachers table)

-- column_name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- This query is performed with a full outer join. Since MySQL does not support FULL OUTER JOIN, you can use the keyword UNION to merge the left and right joins to achieve a full join.
-- The column name returned by the query should be the same as the case of the column name output in the sample.
-- If the query of both tables are empty, nothing will be returned.
-- Example
-- Example I:

-- Table Contents 1 : courses (Courses table)

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithms	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- 12	Dynamic Programming	2000	2018-08-18	2
-- Table Contents 2 : teachers (Teachers table)

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- course_name	teacher_name	teacher_country
-- Advanced Algorithms	Southern Emperor	JP
-- System Design	Western Venom	USA
-- Django	Western Venom	USA
-- Web	Southern Emperor	JP
-- Big Data	Eastern Heretic	UK
-- Artificial Intelligence	Western Venom	USA
-- Java P6+	Western Venom	USA
-- Data Analysis	Eastern Heretic	UK
-- Object Oriented Design	Southern Emperor	JP
-- Dynamic Programming	Northern Beggar	CN
-- NULL	Linghu Chong	CN
-- Example 2
-- Table Contents 1 : courses (Courses table)

-- id	name	student_count	created_at	teacher_id
-- 1	NULL	880	2020-06-01	4
-- 2	NULL	1350	2020-07-18	3
-- 3	NULL	780	2020-02-29	3
-- 4	NULL	340	2020-04-22	4
-- Table Contents 2 : teachers (Teachers table)

-- id	name	email	age	country
-- 1	NULL	eastern.heretic@gmail.com	20	NULL
-- 3	NULL	western.venom@163.com	28	NULL
-- 4	NULL	southern.emperor@qq.com	21	NULL
-- After running your SQL statement, the table should return.

-- course_name	teacher_name	teacher_country
-- Because there is no data for the query in the sample, only the title is shown here, no data.
SELECT 
	p.course_name AS course_name,
	p.teacher_name AS teacher_name,
	p.teacher_country AS teacher_country
FROM 
	(
		(
			SELECT
				c.name AS course_name,
				t.name AS teacher_name,
				t.country AS teacher_country
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
				t.name AS teacher_name,
				t.country AS teacher_country
			FROM
				teachers AS t
			RIGHT JOIN 
				courses  AS c
			ON
				t.id = c.teacher_id
		)
	) AS p
GROUP BY 
	p.course_name,p.teacher_name,p.teacher_country -- query result de-duplication