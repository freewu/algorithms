-- 2581 · Automatic backup when deleting teacher information
-- # Description
-- Please design a trigger for the teachers table teachers to copy the deleted data in teachers to the backup table teachers_bkp which has the same structure

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
-- Table Content 1: TEACHERS

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	CN
-- 5	Linghu Chong	NULL	18	CN
-- Table of contents 2: teachers_bkp

-- id	name	email	age	country
-- After running your SQL statement, we will execute the following statement.

-- DELETE * FROM `teachers`;
-- SELECT * FROM `teachers_bkp`;
-- Return results.

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	CN
-- 5	Linghu Chong	NULL	18	CN

-- Write your SQL here --
DROP TRIGGER IF EXISTS `after_teachers_delete`;
CREATE TRIGGER `after_teachers_delete`
AFTER DELETE ON `teachers`
FOR EACH ROW
BEGIN
    -- 同步到 teachers_bkp 表上
    INSERT INTO 
        `teachers_bkp` (
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
END