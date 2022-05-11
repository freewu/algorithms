-- 2704 · Zhang San's MySQL Learning Path (III)
-- # Description
-- Zhang San is playing a game in the dormitory, 
-- today's luck is not good, always matched teammates pit, so angry that Zhang San directly uninstalled the game. 
-- Zhang San, who was not in the mood to eat, opened QQ and found 999+ messages from his school teacher, 
-- who wanted him to create a SQL view v_best_teachers and asked him to write it and send it to him. 
-- Zhang San looked at the requirements and wrote SQL while reading the requirements, and found that... It's not right! 
-- The school's requirement was to find out the most popular teachers in each country (the more students in the class, the more popular they are) 
-- and sort them in ascending order by nationality. Now Zhang San in the knowledge sharing process, 
-- take this question to test you and taunt you can not do it ...... lad, 
-- the time to hit the face of Zhang San is here, what are you waiting for ? punch

-- Table Definition 1: teachers (Teachers table)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Instructor's email
-- age	int	lecturer's age
-- country	varchar	tutor's nationality

-- Table Definition 2: courses (Course List)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course_name
-- student_count	int	Total number of students attending a class
-- created_at	date	Course creation time
-- teacher_id	int unsigned	instructor v_best_teachers

-- Result view definition: v_best_teachers(view)
-- column_name	type	comment
-- teacher_id	int	teacher_name
-- student_count	int	Total number of students
-- country	varchar	Teacher's nationality
-- email	varchar	teacher's email

-- Example
-- Enter data:

-- courses table.
-- id	name	student_count	created_at	teacher_id
-- 1	'Advanced Algorithms'	880	'2020-6-1 09:03:12'	4
-- 2	'System Design'	1350	'2020-7-18 10:03:12'	8
-- 3	'Django'	780	'2020-2-29 12:03:12'	2
-- 4	'Web'	340	'2020-4-22 13:03:12'	4
-- 5	'Big Data'	700	'2020-9-11 16:03:12'	7
-- 6	'Artificial Intelligence'	1660	'2018-5-13 18:03:12'	3

-- teachers table.
-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Linghu Chong'	None	28	'CN'

-- Returned results:

-- teacher_id	student_count	country	email
-- 2	780	'CN'	'northern.beggar@qq.com'
-- 4	1220	'JP'	'southern.emperor@qq.com'
-- 3	1660	'USA'	'western.venom@163.com'

-- Write your SQL here --
-- 先生成一张老师全信息 + 学生数量的 视图
CREATE OR REPLACE VIEW
	`v_teachers`
AS 
	-- 老师全信息+学生数量
	SELECT
		t.*,
		c.student_count
	FROM
		teachers AS t,
		( -- 老师ID + 学生数量
			SELECT 
				SUM(student_count) AS student_count,
				teacher_id
			FROM
				courses
			GROUP BY	
				teacher_id
		) AS c
	WHERE
		t.id = c.teacher_id;

-- 每个国家最受欢迎的老师信息
CREATE OR REPLACE VIEW
	`v_best_teachers`
AS 
	SELECT
		t.id AS teacher_id,
		t.student_count AS student_count,
		t.country AS country,
		t.email AS email
	FROM
		v_teachers AS t,
		(
			SELECT
				MAX(student_count) AS student_count,
				country
			FROM
				`v_teachers`
			GROUP BY
				country
		) AS m
	WHERE
		m.country = t.country AND
		m.student_count = t.student_count;
