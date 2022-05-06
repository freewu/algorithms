-- 2048 · Search for information on all courses taught by a given teacher
-- # Description
-- Write an SQL statement to query the information of all courses taught by 'Western Venom' from the teachers table teachers and the courses table courses.
-- Table definition: teachers

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Table Definition: courses

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	date	course creation time
-- teacher_id	int	teacher's id

-- If there are no courses in the courses table taught by 'Western Venom', nothing will be returned.

-- Example
-- Example 1:

-- Table content : teachers

-- | id   | name             | email                     | age  | country |
-- | ---- | ---------------- | ------------------------- | ---- | ------- |
-- | 1    | Eastern Heretic  | eastern.heretic@gmail.com | 20   | UK      |
-- | 2    | Northern Beggar  | northern.beggar@qq.com    | 21   | CN      |
-- | 3    | Western Venom    | western.venom@163.com     | 28   | USA     |
-- | 4    | Southern Emperor | southern.emperor@qq.com   | 21   | JP      |
-- | 5    | Linghu Chong     | NULL                      | 18   | CN      |

-- Table Contents: courses

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
-- After running your SQL statement, the table should return.

-- id	name	student_count	created_at	teacher_id
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- Example 2:

-- Table Contents : teachers

-- | id   | name             | email                     | age  | country |
-- | ---- | ---------------- | ------------------------- | ---- | ------- |
-- | 1    | Eastern Heretic  | eastern.heretic@gmail.com | 20   | UK      |
-- | 2    | Northern Beggar  | northern.beggar@qq.com    | 21   | CN      |
-- | 3    | Western Venom    | western.venom@163.com     | 28   | USA     |
-- | 4    | Southern Emperor | southern.emperor@qq.com   | 21   | JP      |
-- | 5    | Linghu Chong     | NULL                      | 18   | CN      |

-- Table Contents: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-6-1 09:03:12	4
-- 4	Web	340	2020-4-22 13:03:12	4
-- 5	Big Data	700	2020-9-11 16:03:12	1
-- 8	Data Analysis	500	2019-7-12 13:03:12	1
-- 10	Object Oriented Design	300	2020-8-8 13:03:12	4
-- 12	Dynamic Programming	2000	2018-8-18 20:03:12	1
-- After running your SQL statement, the table should return.

-- id	name	student_count	created_at	teacher_id
-- Because there are no courses taught by Western Venom teachers in the courses table, only the titles are shown here, not the data.

SELECT
	*
FROM
	courses
WHERE
	teacher_id IN (
		SELECT
			id
		FROM
			teachers
		WHERE
			name = 'Western Venom'
	)