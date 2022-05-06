-- 2056 · Copy the data in the teachers table that are older than 20 to another table
-- # Description
-- Write an SQL statement to copy the data of teachers older than 20 to an empty table teachers_bkp which has the same structure with the teachers table.
-- Table definition: teachers

-- Column Name	Type	Comments
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


-- If the eligible data in teachers is null, the returned teachers_bkp is also null.
-- Example
-- example I

-- Table content: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	CN
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, we will execute the following statement.

-- SELECT * FROM `teachers_bkp`;
-- Returns.

-- id	name	email	age	country
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	CN
-- example II

-- Table Contents: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	19	UK
-- 2	Northern Beggar	northern.beggar@qq.com	18	CN
-- 3	Western Venom	western.venom@163.com	19	USA
-- 4	Southern Emperor	southern.emperor@qq.com	18	CN
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, we will execute the following statement.

-- SELECT * FROM `teachers_bkp`;
-- Returns.

-- id	name	email	age	country
-- The teachers whose age is greater than 20 are empty, so the returned teachers_bkp table is also empty, so only the title is shown here, no data.
INSERT INTO
	teachers_bkp
SELECT
	*
FROM
	teachers
WHERE
	age > 20