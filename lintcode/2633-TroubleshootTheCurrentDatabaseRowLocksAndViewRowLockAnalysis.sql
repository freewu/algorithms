-- 2633 · Troubleshoot the current database row locks and view row lock analysis
-- # Description
-- Now we need to troubleshoot the current database row lock situation, please write SQL statement to view the row lock analysis

-- Example
-- After entering your SQL statement, the result should be returned:

-- Variable_name	Value
-- 'Innodb_row_lock_current_waits'	'0'
-- 'Innodb_row_lock_time'	'0'
-- 'Innodb_row_lock_time_avg'	'0'
-- 'Innodb_row_lock_time_max'	'0'
-- 'Innodb_row_lock_waits'	'0'

-- SHOW STATUS 
SHOW STATUS like '%Innodb_row_lock%';


-- # Innodb_row_lock_current_waits
-- InnoDB表上的操作当前等待的行锁的数量。

-- # Innodb_row_lock_time
-- 为InnoDB表获取行锁花费的总时间(以毫秒为单位)。

-- # Innodb_row_lock_time_avg
-- 为InnoDB表获取行锁的平均时间(以毫秒为单位)。

-- # Innodb_row_lock_time_max
-- 为InnoDB表获取行锁的最长时间(以毫秒为单位)。

-- # Innodb_row_lock_waits
-- InnoDB表上的操作必须等待行锁的次数。

-- # Innodb_rows_deleted
-- 从InnoDB表中删除的行数。

-- # Innodb_rows_inserted
-- 插入到InnoDB表中的行数。

-- # Innodb_rows_read
-- 从InnoDB表中读取的行数。

-- # Innodb_rows_updated
-- InnoDB表中更新的行数。

-- # Innodb_truncated_status_writes
-- SHOW ENGINE INNODB STATUS被截断的次数

-- 获取 InnoDB 的行锁争用情况
-- show status like 'innoDB_row_lock%'
-- Innodb_row_lock_current_waits   0  // 当前处于等待状态的锁数量
-- Innodb_row_lock_time    181621     // 启动到现在锁定的总时间长度
-- Innodb_row_lock_time_avg    321    // 平均每次锁定的时长(ms)
-- Innodb_row_lock_time_max    20824  // 最长的一次锁定时间(ms)
-- Innodb_row_lock_waits   565        // 启动到现在总计是锁定次数

-- 如果锁争用比较严重，那么字段 Innodb_row_lock_current_waits 与 Innodb_row_lock_time_avg 值都会比较高。

-- 背景知识

-- MySQL事务简单回顾

-- 3.2 行锁模式以及加锁方法
-- InnoDB 有以下两种类型的锁

-- 共享锁(S)：允许一个事务去读一行，阻止其他事务获得相同数据集的排他锁。
-- 排他锁(X)：允许获得排他锁的事务更新数据，阻止其他事务取得相同数据集的共享读锁和排他写锁。
-- 除此之外还有两种意向锁

-- 意向共享锁(IS)：事务打算给数据行加行共享锁，事务在给一个数据行加共享锁前必须先取得该表的IS锁。
-- 意向排他锁(IX)：事务打算给数据行加行排他锁，事务在给一个数据行加排他锁前必须先取得该表的IX锁。

-- 行级锁的兼容
-- 意向锁是InnoDB自动加的，不需要用户干预。UPDATE/DELETE/INSERT 语句 InnoDB 会自动给涉及的数据集加排他锁。普通的 SELECT 语句 InnoDB 不加锁。不过可以在语句中显式的给数据集加共享锁或者排他锁。

-- 在 RC（read commited）级别中，数据的读取都是不加锁的，但是数据的写入、修改和删除是需要加锁的。

-- InnoDB 行锁是通过给索引项加锁实现的，如果没有索引，InnoDB 将通过隐藏的聚簇索引来对记录加锁。 InnoDB 行锁有三种情形：

-- Record Lock：对索引项加锁
-- Gap Lock：对索引之间的“间隙”、第一条记录前的“间隙”或最后一条记录的“间隙”加锁。
-- Next-key lock：前两种的组合，对记录及其前面的间隙加锁。
-- 如果不通过索引条件检索数据，那么 InnoDB 将对表中的所有记录加锁，实际效果就和表锁一样了。

-- 如下是一些注意事项及说明

-- 1. 在不通过索引条件查询时，InnoDB 会锁定表中的所有记录

-- // session1
-- set autocommit = 0;
-- select * from test where col1 = 1 for update;
-- // session2
-- select * from test where col1 = 2 for update;
-- // waiting...
-- 这是一个串行操作，在session1给表的col1字段的某条记录加了排他锁，在理想情况下session2的操作应该不会受到影响。

-- 如果col1字段没有设置索引的话，这个阻塞操作就一定会发生。因为检索操作没有走到索引会导致 InnoDB 给所有的记录都加了行锁。这样 session2 的排他锁就无法得到从而进入阻塞状态。

-- 2. 由于 MySQL 的行锁是针对索引加的锁，不是针对记录加的锁，所以虽然访问不同行的记录，如果使用了相同的索引键，一样会出现锁冲突。

-- record1	record2
-- id=1	value = 1
-- id=1	value = 4
-- 表中两条记录，两个字段id，value中仅有id字段有索引。

-- // session1
-- set autocommit = 0;
-- select * from table_test where id = 1 and value = '1' for update;

-- // session2
-- set autocommit = 0;
-- select * from table_test where id = 1 and value = '4' for update;
-- // waiting
-- 虽然session2访问的是和session1不同的记录，但是因为使用了相同的索引，所以也还是需要等待锁。

-- 3. 当表有多个索引的时候，不同的事务可以使用不同的索引锁定不同的行，不论是使用主键索引、唯一索引还是普通索引，InnoDB 都会使用行锁来对数据加锁。

-- 4. 即便在条件中使用了索引字段，但是是否使用索引来检索数据是由 MySQL 通过判断不同执行计划的代价来决定的，如果 MySQL 认为全表扫描效率更高，比如对一些很小的表，它就不会使用索引，这种情况下 InnoDB 也会对所有记录加锁。

-- 3.3 Next-key 锁
-- 当我们使用范围条件而不是相等条件检索数据，请求获得锁，InnoDB 会给符合条件的数据的索引项加锁。

-- select * from test where id > 100 for update;
-- 这是一个范围检索，InnoDB 不仅会对符合条件的记录加锁，还会对大于100的“间隙(即不存在的记录)”加锁。这个检索执行的时候会阻塞100以后id数据的插入操作。

-- 在使用范围条件检索并锁定记录时，InnoDB 的这种加锁机制会阻塞符合条件范围内键值对的并发插入，会造成严重的锁等待。因此在实际开发中，尤其是并发插入较多的应用，我们要尽量使用相等的条件来访问和更新数据，避免使用范围检索。
