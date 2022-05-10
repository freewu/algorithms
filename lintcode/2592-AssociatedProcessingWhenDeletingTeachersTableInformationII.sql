-- 2592 · Associated processing when deleting teachers table information (II)
-- # Description
-- A new group of researchers has been recruited to the LintCode department 
-- and their information needs to be deleted from the teachers table in order to protect their personal safety. 
-- Write a trigger to record their information in teachers_bkp each time they are deleted from the teachers table. 
-- Some highly respected researchers are keen to join the Chinese (CN) nationality, please write a trigger to achieve this.

-- We need to add two new triggers to the teachers table when deleting a teacher's data.

-- trigger bkp_teachers_delete: backs up the deleted data to a backup table teachers_bkp with the same structure.
-- Trigger before_teachers_delete: changes the nationality of the oldest teacher in the teachers_bkp table to 'CN'.
-- The trigger bkp_teachers_delete needs to be executed before the trigger before_teachers_delete.
-- Table Definition 1: teachers (Teachers table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Instructor's email
-- age	int	lecturer's age
-- country	varchar	tutor's nationality
-- Table Definition 2: courses (Course List)

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total_students
-- created_at	date	Course creation time
-- teacher_id	int unsigned	instructor id
-- Table Definition 3: teachers_bkp (backup table)

-- column_name	type	comment
-- id	int unsigned	primary key
-- name	varchar	Instructor's name
-- email	varchar	Instructor's email
-- age	int	instructor's age
-- country	varchar	Tutor's nationality
-- Contact me on wechat to get Amazon、Google requent Interview questions . (wechat id : jiuzhang15)


-- Example
-- Enter data:
-- teachers table.

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Linghu Chong'	None	18	'CN'
-- teachers_bkp Table.

-- id	name	email	age	country
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'CN'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'CN'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Linghu Chong'	None	18	'CN'

-- Write your SQL here --
-- 触发器 bkp_teachers_delete：将删除的数据备份到结构相同的备份表 teachers_bkp 中
DROP TRIGGER IF EXISTS `bkp_teachers_delete`;
CREATE TRIGGER `bkp_teachers_delete`
AFTER DELETE ON `teachers`
FOR EACH ROW
BEGIN
	-- 将删除的数据备份到结构相同的备份表 teachers_bkp 中；
	INSERT INTO
		`teachers_bkp`(
			`id`,
			`name`,
			`email`,
			`age`,
			`country`
		) 
	VALUES (
		old.id,
		old.name,
		old.email,
		old.age,
		old.country
	);
END;

-- 触发器 before_teachers_delete：将 teachers_bkp 表中最年长的教师国籍改为 'CN'
DROP TRIGGER IF EXISTS `before_teachers_delete`;
CREATE TRIGGER `before_teachers_delete`
AFTER DELETE ON `teachers`
FOR EACH ROW
BEGIN
	-- 获取最大年龄
	DECLARE max_age int;
	SET max_age = (SELECT MAX(age) FROM `teachers_bkp`);

	-- 将 teachers_bkp 表中最年长的教师国籍改为 'CN'
	UPDATE 
		`teachers_bkp`
	SET
		country = 'CN'
	WHERE
		age = max_age;
END;