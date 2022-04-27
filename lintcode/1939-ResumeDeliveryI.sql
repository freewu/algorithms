-- 1939 · Resume Delivery I
-- # Description
-- The students table stores all student information, including student id and student name
-- The companies table stores all company information, including company id and company name
-- The recording table stores all resume delivery data, including student id (student_id) and company id (company_id)
-- Please write SQL statements to query the names of all students who have not submitted their resumes to Alibaba.

-- Table definition 1: students

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	student name
-- Table definition 2: companies

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	company name
-- address	varchar	company address
-- Table definition 3: recording

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- delivery_date	date	delivery date
-- company_id	int	company id
-- student_id	int	student id
-- Example
-- Example 1:

-- Table content 1: students

-- id	name
-- 1	Da Ming
-- 2	Amy
-- 3	Mike
-- 4	Park
-- 5	George
-- Table content 2: companies

-- id	name	address
-- 1	Alibaba	Hang Zhou
-- 2	NetEase	Guang Zhou
-- 3	Baidu	Bei Jing
-- 4	Tencent	Shen Zhen
-- Table content 3: recording

-- id	delivery_date	company_id	student_id
-- 1	2020-01-08	3	4
-- 2	2020-02-06	4	5
-- 3	2020-03-12	1	1
-- 4	2020-04-07	1	4
-- After running your SQL statement, the table should return:

-- name
-- Amy
-- Mike
-- George
-- Example 2:

-- Table content 1: students

-- id	name
-- 1	Da Ming
-- 2	Amy
-- 3	Mike
-- 4	Park
-- 5	George
-- Table content 2: companies

-- id	name	address
-- 1	Alibaba	Hang Zhou
-- 2	NetEase	Guang Zhou
-- 3	Baidu	Bei Jing
-- 4	Tencent	Shen Zhen
-- Table content 3: recording

-- id	delivery_date	company_id	student_id
-- 1	2020-01-08	1	2
-- 2	2020-02-06	1	5
-- 3	2020-03-12	1	1
-- 4	2020-04-07	1	4
-- After running your SQL statement, the table should return:

-- name
-- Mike

SELECT	
	s.name AS name
FROM
	students AS s
WHERE	
	id NOT IN (
		SELECT	
			r.student_id
		FROM
			recording AS r,
			companies AS c 
		WHERE	
			r.company_id = c.id AND
			c.name = 'Alibaba'
	)