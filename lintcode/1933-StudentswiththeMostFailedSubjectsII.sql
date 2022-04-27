-- 1933 · Students with the Most Failed Subjects II
-- # Description
-- The exams records of students are stored in the exam table
-- Please use SQL statement to find the student_id corresponding to the student with the largest number of failed subjects.

-- Table definition: exams

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- student_id	int	student's id
-- date	date	date of exam
-- is_pass	int	grade status (0 means fail, others means pass)
-- Data guarantee that there are many students with the largest number of failed subjects

-- Example
-- Example 1:

-- Table content: exams

-- id	student_id	date	is_pass
-- 1	1	2020-11-15	0
-- 2	2	2020-11-17	1
-- 3	3	2020-11-24	0
-- 4	3	2020-11-28	0
-- 5	1	2020-11-29	0
-- After running your SQL statement, the table should return:

-- student_id
-- 1
-- 3
-- Example 2:

-- Table content: exams

-- id	student_id	date	is_pass
-- 1	1	2020-11-15	0
-- 2	2	2020-11-17	0
-- 3	3	2020-11-24	0
-- 4	3	2020-11-28	0
-- 5	1	2020-11-29	0
-- 6	2	2020-11-13	0
-- After running your SQL statement, the table should return:

-- student_id
-- 1
-- 2
-- 3

SELECT
	student_id
FROM
	(
		SELECT
			student_id,
			COUNT(*) AS num
		FROM 
			exams
		WHERE 
			is_pass = 0 
		GROUP BY student_id
	) AS p 
WHERE 
	p.num = (
		SELECT 
			MAX(w.num)
		FROM
			(
				SELECT
					student_id,
					COUNT(*) AS num
				FROM 
					exams
				WHERE 
					is_pass = 0 
				GROUP BY student_id
			) AS w
	)