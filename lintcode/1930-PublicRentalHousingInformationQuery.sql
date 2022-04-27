-- 1930 · Public Rental Housing Information Query
-- # Description
-- The rooms table records tenant information (tenant_id) and rent of public rental housing
-- The tenants table records the name of the tenant (name)
-- Please write a SQL statement to query the id, rent and name of the tenant of all rooms. If there is no tenant yet, it will be null.

-- Table definition 1: rooms

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- tenant_id	int	tenant's id
-- rent	int	rent
-- Table definition 2: tenants

-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	tenant's name
-- Example
-- Example 1:

-- Table content 1: rooms

-- id	tenant_id	rent
-- 1	2	300
-- 2	3	400
-- 3	null	300
-- 4	1	500
-- Table content 2: tenants

-- id	name
-- 1	zhangsan
-- 2	lisi
-- 3	wanger
-- After running your SQL statement, the table should return:

-- id	rent	name
-- 1	300	lisi
-- 2	400	wanger
-- 3	300	null
-- 4	500	zhangsan
-- Example 2:

-- Table content 1: rooms

-- id	tenant_id	rent
-- 1	2	300
-- 2	3	400
-- 3	null	300
-- Table content 2: tenants

-- id	name
-- 1	zhangsan
-- 2	lisi
-- 3	wanger
-- After running your SQL statement, the table should return:

-- id	rent	name
-- 1	300	lisi
-- 2	400	wanger
-- 3	300	null

SELECT
	r.id AS id,
	r.rent AS rent,
	t.name AS name
FROM 
	rooms AS r
LEFT JOIN	
	tenants AS t 
ON
	r.tenant_id = t.id
