-- 2577 · Backup Update Data Trigger
-- # Description
-- Please design a trigger for the teachers table teachers to copy the old data to a backup table teachers_bkp 
-- with the same structure when teachers updates its data

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
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Linghu Chong'	None	18	'CN'
-- Table of contents 2: teachers_bkp

-- id	name	email	age	country
-- Return result.

-- id	name	email	age	country
-- 1	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 2	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 3	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 4	'Linghu Chong'	None	18	'CN'

DROP TRIGGER IF EXISTS `after_teachers_update`;
CREATE TRIGGER `after_teachers_update`
AFTER UPDATE ON `teachers`
FOR EACH ROW
BEGIN
    -- 同步到 teachers_bkp 表上
    INSERT INTO 
        `teachers_bkp` (
            `name`,
            `email`,
            `age`,
            `country`
        )
    VALUES (
        old.name,
        old.email,
        old.age,
        old.country
    );
END