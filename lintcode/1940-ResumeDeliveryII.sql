-- 1940 · Resume Delivery II
-- # Description
-- The students table stores information about all students, including student id and student name name.
-- The companies table stores all company information, including company id, company name name and company address address.
-- The records table stores all CV submissions, including student id (student_id) and company id (company_id)
-- Write an SQL statement to query the name and address of the company that receives the most resumes.

-- Table definition 1: students

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	student name
-- Table definition 2: companies

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	company name
-- address	varchar	company address
-- Table definition 3: records

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
-- Table content 3: records

-- id	delivery_date	company_id	student_id
-- 1	2020-01-08	3	4
-- 2	2020-02-06	4	5
-- 3	2020-03-12	4	1
-- 4	2020-04-07	1	4
-- 5	2020-04-13	3	2
-- After running your SQL statement, the table should return:

-- name	address
-- Baidu	Bei Jing
-- Tencent	Shen Zhen
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
-- Table content 3: records

-- id	delivery_date	company_id	student_id
-- 1	2020-01-08	3	4
-- 2	2020-02-06	4	5
-- 3	2020-03-12	1	1
-- 4	2020-04-07	2	3
-- After running your SQL statement, the table should return:

-- name	address
-- Alibaba	Hang Zhou
-- NetEase	Guang Zhou
-- Baidu	Bei Jing
-- Tencent	Shen Zhen

-- sultion 1 one sql
SELECT
	c.name,
	c.address
FROM 
	companies AS c
WHERE 
	c.id IN (
		SELECT
			p.company_id
		FROM	
			(
				SELECT
					COUNT(1) AS num, company_id
				FROM records
				GROUP BY  company_id
			) AS p
		WHERE
			p.num = (
				SELECT
					MAX(w.num)
				FROM
					(
						SELECT
							COUNT(1) AS num, company_id
						FROM records
						GROUP BY company_id
					) AS w
			)
	)

-- solution 2 with create view
CREATE VIEW rv AS SELECT COUNT(1) AS num, company_id FROM records GROUP BY company_id;
SELECT 
    c.name,
	c.address
FROM 
    companies AS c
WHERE 
	c.id IN (
        SELECT
            company_id
        FROM rv 
        WHERE num = ( SELECT Max(num) FROM rv)
    )