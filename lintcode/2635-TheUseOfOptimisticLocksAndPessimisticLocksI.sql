-- 2635 · The use of optimistic locks and pessimistic locks (I)
-- # Description
-- Databases often use optimistic locks and pessimistic locks in concurrent situations. 
-- Now we require to use optimistic locks to protect the data of the table. 
-- Our current database uses version to implement this,
-- and now we request to write SQL statement to update the information of the teacher named Western Venom in the teachers table,
-- changing the nationality to CN.

-- Current version number version = 5
-- Table Definition: teachers (Teachers table)

-- column name	type	comments
-- id	int	primary key
-- name	varchar	Instructor's name
-- email	varchar	Instructor's email
-- age	int	Tutor's age
-- country	varchar	Tutor's nationality
-- version	int	version number
-- Translated with www.DeepL.com/Translator (free version)

-- Example
-- Form content : teachers

-- id	name	email	age	country	version
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'	3
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'	2
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'	5
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'	6
-- **Return result : **

-- id	name	email	age	country	version
-- 3	'Western Venom'	'western.venom@163.com'	28	'CN'	6

-- # 悲观锁（Pessimistic Lock）
--    悲观锁的特色是先获取锁，再进行业务操做，即“悲观”的认为获取锁是很是有可能失败的，所以要先确保获取锁成功再进行业务操做。
--    一般所说的“一锁二查三更新”即指的是使用悲观锁。
--    一般来说在数据库上的悲观锁须要数据库自己提供支持，即经过经常使用的 select … for update 操做来实现悲观锁。
--    当数据库执行select for update时会获取被select中的数据行的行锁，所以其余并发执行的select for update若是试图选中同一行则会发生排斥（须要等待行锁被释放），所以达到锁的效果。
--    select for update获取的行锁会在当前事务结束时自动释放，所以必须在事务中使用
--    这里须要注意的一点是不一样的数据库对select for update的实现和支持都是有所区别的，
--    例如 oracle 支持select for update no wait，表示若是拿不到锁马上报错，而不是等待，mysql就没有no wait这个选项。
--    另外mysql还有个问题是select for update语句执行中全部扫描过的行都会被锁上，这一点很容易形成问题。
--    所以若是在mysql中用悲观锁务必要肯定走了索引，而不是全表扫描。

-- # 乐观锁（Optimistic Lock）
--    乐观锁的特色先进行业务操做，不到万不得已不去拿锁。即“乐观”的认为拿锁多半是会成功的，所以在进行完业务操做须要实际更新数据的最后一步再去拿一下锁就好。编程
--    乐观锁在数据库上的实现彻底是逻辑的，不须要数据库提供特殊的支持。通常的作法是在须要锁的数据上增长一个版本号，或者时间戳，而后按照以下方式实现
-- 
--       SELECT data AS old_data, version AS old_version FROM …;
--       -- 根据获取的数据进行业务操做，获得new_data和new_version
--       UPDATE SET data = new_data, version = new_version WHERE version = old_version
--       if (updated row > 0) {
--           // 乐观锁获取成功，操做完成
--       } else {
--          // 乐观锁获取失败，回滚并重试
--       }
-- 
--    乐观锁是否在事务中其实都是无所谓的，其底层机制是这样：
--        在数据库内部update同一行的时候是不容许并发的，即数据库每次执行一条update语句时会获取被update行的写锁，直到这一行被成功更新后才释放。
--        所以在业务操做进行前获取须要锁的数据的当前版本号，而后实际更新数据时再次对比版本号确认与以前获取的相同，并更新版本号，便可确认这之间没有发生并发的修改。
--        若是更新失败便可认为老版本的数据已经被并发修改掉而不存在了，此时认为获取锁失败，须要回滚整个业务操做并可根据须要重试整个过程。

-- # 总结
--    乐观锁在不发生取锁失败的状况下开销比悲观锁小，可是一旦发生失败回滚开销则比较大，所以适合用在取锁失败几率比较小的场景，能够提高系统并发性能性能
--    乐观锁还适用于一些比较特殊的场景，例如在业务操做过程当中没法和数据库保持链接等悲观锁没法适用的地方

-- Write your SQL here --

-- optimistic locks
UPDATE
	`teachers`
SET
	country = 'CN',
	version = version + 1
WHERE
	name = 'Western Venom' AND
	version = 5;