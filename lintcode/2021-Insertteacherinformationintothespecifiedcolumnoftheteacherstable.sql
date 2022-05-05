-- 2021 · Insert teacher information into the specified column of the teachers table
-- # Description
-- Write an SQL statement to insert a new teacher record into the teacher table teachers,the record is as follows:

-- name	email	age	country
-- XiaoFu	XiaoFu@lintcode.com	20	CN
-- Table definition: teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher‘s name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher’s nationality

-- The id column of the teachers table teachers is the primary key, self-growing and does not require a value to be set.
-- The inserted record field type should match the table definition field type
-- Example
-- Example1

-- Table content: Teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, we will execute the following statement:

-- SELECT *
-- FROM `teachers`;
-- return to:

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- 6	XiaoFu	XiaoFu@lintcode.com	20	CN
-- Example 2

-- Table content: Teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- After running your SQL statement, we will execute the following statement:

-- SELECT *
-- FROM `teachers`;
-- return to:

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	XiaoFu	XiaoFu@lintcode.com	20	CN

INSERT INTO `teachers`(
	`name`,
	`email`,
	`age`,
	`country`
)
VALUES (
	'XiaoFu',
	'XiaoFu@lintcode.com',
	20,
	'CN'
)