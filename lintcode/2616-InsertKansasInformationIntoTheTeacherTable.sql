-- 2616 · Insert Kansas information into the teacher table
-- # Description
-- We want to insert a Kansas whose age is 41 and nationality is UK into the teachers table, 
-- please add SQL statements to implement the insertion of Kansas.

-- Table definition: teachers (Teachers table)

-- column name	type	comment
-- id	int	primary key
-- name	varchar	Instructor's name
-- email	varchar	Instructor's email
-- age	int	Tutor's age
-- country	varchar	Tutor's nationality

-- Example
-- Table content : teachers

-- 不要删除预置代码 --
-- 开启一个事务 -- 
BEGIN;

-- 插入 Kansas 的信息 --
-- Write your SQL Query here --
INSERT INTO 
	`teachers` (
		`name`,
		`age`,
		`country`
	)
VALUES (
	'Kansas',
	41,
	'UK'
);

COMMIT;