-- 2047 · Copy all the data in the teacher table to another table
-- # Description
-- Write an SQL statement that copies all the data of the teachers table teachers to the empty table teachers_bkp with same structure.
-- Table definition: teachers

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Table definition: teachers_bkp

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- If the data of teachers is null, then the returned teachers_bkp is also null
-- Example
-- Sample I

-- Table content: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	CN
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, we will execute the following statement.

-- SELECT *
-- FROM `teachers_bkp`;
-- Returns.

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	CN
-- 5	Linghu Chong	NULL	18	CN
-- Sample 2

-- Table Contents: teachers

-- id	name	email	age	country
-- After running your SQL statement, we will execute the following statement.

-- SELECT *
-- FROM `teachers_bkp`;
-- Returns.

-- id	name	email	age	country
-- The teachers table is empty, so the returned teachers_bkp table is also empty

-- insert into
INSERT INTO
	teachers_bkp 
SELECT 
	id,
	name,
	email,
	age,
	country
FROM
	teachers

--  SELECT INTO 要求目标表不存在，因为在插入时会自动创建；INSERT INTO SELECT要求目标表存在；
-- select into
SELECT 
	*
INTO 
	`teachers_bkp`
FROM
	`teachers`