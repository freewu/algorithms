-- 2060 · Search for the name of the teacher for the 'Big Data' course
-- # Description
-- Write an SQL statement to query the name of teacher who teaches 'Big Data' from the courses table and the teachers table.

-- Table definition: teachers

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Table Definition: courses

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	course creation time
-- teacher_id	int	teacher id

-- If there is no teacher teaches 'Big Data' in the teachers table, nothing is returned.
-- If the teacher's name is NULL, then NULL is returned.
-- Example
-- Sample 1:

-- Table content: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-6-1 09:03:12	4
-- 2	System Design	1350	2020-7-18 10:03:12	3
-- 3	Django	780	2020-2-29 12:03:12	3
-- 4	Web	340	2020-4-22 13:03:12	4
-- 5	Big Data	700	2020-9-11 16:03:12	1
-- 6	Artificial Intelligence	1660	2018-5-13 18:03:12	3
-- 7	Java P6+	780	2019-1-19 13:03:12	3
-- 8	Data Analysis	500	2019-7-12 13:03:12	1
-- 10	Object Oriented Design	300	2020-8-8 13:03:12	4
-- 12	Dynamic Programming	2000	2018-8-18 20:03:12	1
-- Table content : teachers

-- | id   | name             | email                     | age | country |
-- | ---- | ---------------- | ------------------------- | --- | ------- |
-- | 1    | Eastern Heretic  | eastern.heretic@gmail.com | 20  | UK      |
-- | 2    | Northern Beggar  | northern.beggar@qq.com    | 21  | CN      |
-- | 3    | Western Venom    | western.venom@163.com     | 28  | USA     |
-- | 4    | Southern Emperor | southern.emperor@qq.com   | 21  | JP      |
-- | 5    | Linghu Chong     | NULL                      | 18  | CN      |

-- After running your SQL statement, the table should return.

-- name
-- Eastern Heretic
-- Sample 2:

-- Table Contents: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-6-1 09:03:12	4
-- 4	Web	340	2020-4-22 13:03:12	4
-- 5	Big Data	700	2020-9-11 16:03:12	1
-- 8	Data Analysis	500	2019-7-12 13:03:12	1
-- 10	Object Oriented Design	300	2020-8-8 13:03:12	4
-- 12	Dynamic Programming	2000	2018-8-18 20:03:12	1
-- Table content : teachers

-- | id   | name             | email                     | age | country |
-- | ---- | ---------------- | ------------------------- | --- | ------- |
-- | 2    | Northern Beggar  | northern.beggar@qq.com    | 21  | CN      |
-- | 3    | Western Venom    | western.venom@163.com     | 28  | USA     |
-- | 4    | Southern Emperor | southern.emperor@qq.com   | 21  | JP      |
-- | 5    | Linghu Chong     | NULL                      | 18  | CN      |

-- After running your SQL statement, the table should return.

-- name
-- Because the teachers table does not have a teacher 'id' for the 'Big Data' course, only the title is shown here, no data.
SELECT
	name
FROM
	teachers
WHERE
	id IN (
		SELECT
			teacher_id
		FROM
			courses
		WHERE
			name = 'Big Data'
	)