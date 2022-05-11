-- 2685 · Zhang San's MySQL Learning Path (II)
-- #  Description
-- Zhang San once had a dream to listen to every teacher's course in the school, 
-- so he wrote a view v_mydream to show all the courses with the information of the teachers, 
-- and go in ascending order by teacher's name to achieve, but Zhang San found that there are too many courses, 
-- the dream needs to be changed, 
-- so he changed his dream to listen to each teacher's most popular course (the more students listening to the course If the teacher does not teach a course, 
-- the number of courses and students will be displayed as NULL) 
-- The school was very supportive and heard that he was recently learning MySQL and gave him 2 tables of data to find out for himself, 
-- the teachers table contains information about teachers and the courses table contains information about courses.
-- Zhang San looks at it and says, "Wow, this is too easy, thanks to the school for providing the data",
--  but he is about to rewrite the original dream view when he finds Li Si passing by. 
-- Li Si hungry croak, see Zhang San asked to go to dinner together, Zhang San a flash of light, "
-- so coincidentally ~ I also intend to go to dinner, but now I have a very simple problem here, you help me solve I will buy you dinner ",
-- Li Si heard that there is such a good thing, immediately agreed to down, 
-- and found that it will not. In order not to lose face and to be able to eat a meal, 
-- he secretly contacted you and asked you to help him solve it.

-- Table Definition 1: teachers (Teachers table)
-- column name	type	comments
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Instructor's email
-- age	int	lecturer's age
-- country	varchar	tutor's nationality

-- Table Definition 2: courses (Course List)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	Total number of students attending a class
-- created_at	date	Course creation time
-- teacher_id	int unsigned	instructor id

-- Result View Definition: v_mydream(view)
-- column_name	type	comment
-- teacher_name	int	instructor_name
-- course_name	varchar	Course name
-- student_count	int	Total number of students attending the course

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
-- 7	'Java P6+'	780	'2019-1-19 13:03:12'	3
-- 8	'Data Analysis'	500	'2019-7-12 13:03:12'	6
-- 10	'Object Oriented Design'	300	'2020-8-8 13:03:12'	4
-- 12	'Dynamic Programming'	2000	'2018-8-18 20:03:12'	1
-- teachers table.

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'JP'
-- 3	'Western Venom'	'western.venom@163.com'	28	'JP'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Linghu Chong'	None	18	'CN'
-- Returned results:

-- teacher_name	course_name	student_count
-- 'Eastern heretic'	None	None
-- 'Linghu Chong'	None	None
-- 'Northern Beggar'	'Django'	780
-- 'Southern Emperor'	'Advanced Algorithms'	880
-- 'Western Venom'	'Artificial Intelligence'	1660

-- Write your SQL here --
CREATE OR REPLACE VIEW
	`v_mydream`
AS
	SELECT
		t.name AS teacher_name,
		c.name AS course_name,
		c.student_count AS student_count
	FROM
		teachers AS t
	LEFT JOIN 
		(
			SELECT
				cc.teacher_id,
				cc.name,
				cc.student_count
			FROM
				courses AS cc,
				(
					SELECT
						teacher_id,
						MAX(student_count) AS num
					FROM
						courses 
					GROUP BY
						teacher_id
				) AS p
			WHERE
				cc.teacher_id = p.teacher_id AND
				cc.student_count = p.num
		) AS c
	ON 
		t.id = c.teacher_id
	ORDER BY 
		t.name DESC