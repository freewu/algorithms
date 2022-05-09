-- 2573 · Backup New Data Trigger
-- # Description
-- Please design a trigger for the teachers table teachers to copy the new data in teachers to the backup table teachers_bkp which has the same structure

-- Table definition: teachers (teachers table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	instructor's name
-- email	varchar	Instructor's email
-- age	int	instructor's age
-- country	varchar	tutor's nationality
-- Table definition: teachers_bkp (backup table)

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	Instructor's name
-- email	varchar	Instructor's email
-- age	int	instructor's age
-- country	varchar	Tutor's nationality

-- Example
-- Table content 1: teachers

-- id	name	email	age	country
-- 1	'Linghu Chong'	None	18	'CN'
-- Table of contents 2: teachers_bkp

-- id	name	email	age	country
-- After running your SQL statement, we will execute the following statement.
--      SELECT *  FROM `teachers_bkp`;
-- Returns.

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	33	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'JP'
-- 3	'Western Venom'	'western.venom@163.com'	28	'CN'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	31	'USA'

-- Write your SQL here --
DROP TRIGGER IF EXISTS `after_teachers_insert`;
CREATE TRIGGER `after_teachers_insert`
AFTER INSERT ON `teachers`
FOR EACH ROW
BEGIN
	-- 插入数据到 teachers_bkp 
	INSERT INTO
		`teachers_bkp`(
			`name`,
			`email`,
			`age`,
			`country`
		) 
	VALUES (
		new.name,
		new.email,
		new.age,
		new.country
	);
END