-- 2569 · Normative courses table data insert
-- # Description
-- We need to add a new trigger to the courses table named before_courses_insert 
-- to set the instructor id to 0 and the creation time to NULL if the instructor id does not exist in the instructor table when the course data is added.

-- Table Definition 1: teachers (teachers table)

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Table Definition 2: courses (Course List)

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	course creation time
-- teacher_id	int unsigned	teacher id
-- Contact me on wechat to get Amazon、Google requent Interview questions . (wechat id : jiuzhang15)


-- Example
-- Input data:
-- courses table :

-- id	name	student_count	created_at	teacher_id
-- teachers table :

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Linghu Chong'	None	18	'CN'
-- Returned results:

-- id	name	student_count	created_at	teacher_id
-- 1	'Sing Java'	50	'2021-01-28'	2
-- 2	'Math For Core'	45	None	0
-- 3	'Core Writer'	0	None	0
-- 4	'Think Python'	300	'2021-01-26'	1

-- Write your SQL here --
DROP TRIGGER IF EXISTS `before_courses_insert`;
CREATE TRIGGER `before_courses_insert`
BEFORE INSERT ON `courses`
FOR EACH ROW
BEGIN
	DECLARE c int;
	-- 判断 teacher_id 是否存在
	SET c = (SELECT COUNT(*) FROM `teachers` WHERE id = new.teacher_id);
	-- teacher_id 不存在
	IF c = 0 THEN
		-- 则将该条课程数据的教师 id 置为 0
		SET new.teacher_id = 0;
		-- 且创建时间置为 NULL 
		SET new.created_at = NULL;
	END IF;
END