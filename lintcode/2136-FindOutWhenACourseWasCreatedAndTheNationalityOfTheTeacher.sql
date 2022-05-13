-- 2136 · Find out when a course was created and the nationality of the teacher
-- # Description
-- Write a SQL statement to cross join the courses table courses and the teachers table teachers 
-- and combine it with a WHERE clause to query the course creation date 
-- and the teacher's nationality in the courses table teacher_id
--  and the teachers table id, with the result column names displayed as course_date and teacher_country respectively.

-- Table Definition 1: courses
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	course creation time
-- teacher_id	int	teacher id

-- Table Definition 2 : teachers
-- column_name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- This question is a cross join of two tables
-- If the query of both tables are empty, nothing is returned

-- Example
-- Example 1：

-- Table Contents 1 : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	5
-- 4	Web	340	2020-04-22	4
-- Table Contents 2 : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- course_date	teacher_country
-- 2020-06-01	JP
-- 2020-07-18	USA
-- 2020-02-29	CN
-- 2020-04-22	JP
-- Example 2:

-- Table Contents 1 : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	5
-- 4	Web	340	2020-04-22	4
-- Table Contents 2 : teachers

-- id	name	email	age	country
-- 11	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 12	Northern Beggar	northern.beggar@qq.com	21	CN
-- 13	Western Venom	western.venom@163.com	28	USA
-- 14	Southern Emperor	southern.emperor@qq.com	21	JP
-- 15	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the table should return.

-- course_date	teacher_country
-- There is no data that matches the join condition in the WHERE clause in Sample 2, so the output contains only the table header and no data.

-- Write your SQL Query here --
SELECT
    c.created_at AS course_date,
    t.country AS teacher_country
FROM
    teachers AS t,
    courses AS c
WHERE
    t.id = c.teacher_id