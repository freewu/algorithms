-- 2575 · Normative courses table data update
-- # Description
-- We need to add a new trigger to the courses table named before_courses_update and when updating the course data,
-- if the teacher id does not exist in the teachers table, then the teacher id for that data will not be updated. 
-- write SQL statements to implement this trigger.

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
-- student_count	int	total_students
-- created_at	date	Course creation time
-- teacher_id	int unsigned	instructor id
-- Contact me on wechat to get Amazon、Google requent Interview questions . (wechat id : jiuzhang15)


-- Example
-- Data entry:

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
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Linghu Chong'	None	18	'CN'
-- Returned results:

-- id	name	student_count	created_at	teacher_id
-- 1	'Advanced Algorithms'	880	'2020-06-01'	4
-- 2	'System Design'	1350	'2020-07-18'	4
-- 3	'Sing Java'	50	'2020-02-29'	2
-- 4	'Web'	340	'2020-04-22'	5
-- 5	'Big Data'	66	'2020-09-11'	7
-- 6	'Artificial Intelligence'	1660	'2018-05-13'	3
-- 7	'Java P6+'	780	'2019-01-19'	3
-- 8	'Data Analysis'	500	'2019-07-12'	6
-- 10	'Object Oriented Design'	300	'2020-08-08'	4
-- 12	'Dynamic Programming'	2000	'2018-08-18'	2

-- Write your SQL here --
DROP TRIGGER IF EXISTS `before_courses_update`;
CREATE TRIGGER `before_courses_update`
BEFORE UPDATE ON `courses`
FOR EACH ROW
BEGIN
	-- DECLARE c int;
	-- -- 判断 教师 id 在教师表中是否存在
	-- SET c = (SELECT COUNT(*) FROM `teachers` WHERE id = new.teacher_id);
	-- -- 如果教师 id 在教师表中不存在
	-- IF c = 0 THEN
	-- 	-- 则不更新该条数据的教师 id
	-- 	SET new.teacher_id = old.teacher_id;
	-- END IF;

	IF new.teacher_id NOT IN (SELECT id FROM teachers) THEN 
		SET new.teacher_id = old.teacher_id;
	END IF;
END