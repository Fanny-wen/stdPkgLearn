### 数据库设计模式

```
1. 主扩展
2. 主从模式
3. 多对多模式
4. 名值模式
```

### 数据库设计六个设计阶段

```
1. 需求分析, 两种方式: 自顶向下, 自底向上
2. 概念设计(画E-R图, ), 四种方式: 自顶向下, 自底向上, 逐步扩张, 混合策略
3. 逻辑设计
4. 物理结构设计
5. 数据库实施
6. 数据库运行与维护
```

### 什么是事务

```
	答: 把一系列sql语句作为一个整体进行操作
```





### 数据库事务的四大特性

```
ACID
A: 原子性 Atomicity, 原子性是指事务包含的所有操作不可分割, 要么全部成功, 要么全部失败回滚
C: 一致性 Consistency, 执行的前后数据的完整性保持一致
I: 隔离性 Isolation, 一个事务执行的过程中, 不应该受到其他事务的干扰
D: 持久性 Durability, 事务一旦结束, 数据就持久到数据库
```

### 数据库事务隔离级别

```
1. Read uncommitted 读未提交
2. Read committed 读提交
	在RC级别中, 数据的读取都是不需要加锁的, 但是数据的写入, 修改, 删除都是要加锁的.
3. Repeatable read 重复读	mysql默认级别
4. Serializable 串性化
	读用读锁, 写用写锁, 读锁和写锁互斥
```

### mysql表锁&行锁&页面锁(锁的粒度)

```
1. 表级别锁
	1) 表锁
       	Lock table tablename read // 表级别的共享锁
       	Lock table tablename wirte // 表级别的排他锁
	2) 元数据锁(MDL)
		不需要显示的使用元数据锁,当我们对数据库表进行操作时,会自动给这个表加上元数据锁;元数据锁是为了在用户对表进行操作时, 防止其他线程对这个表的结构进行了修改.
   	3) 意向锁
   		意向锁的目的是为了快速判断表里是否有记录被加锁, 意向锁分为意向独占锁和意向共享锁;在想要给数据库表加[独占表锁]时, 需要遍历表里的所有记录查看是否有记录存在独占锁, 而有了意向锁后, 只需要表查看是否有意向独占锁就可以判断了.
   		
2. 行级锁
	1) 记录锁
		仅仅锁定一条记录
	2) 间隙锁
		锁定一个范围，但是不包含记录本身
	3) next-key锁
		锁定一个范围，并且锁定记录本身
		
3. 页面锁

注: myisam: 只支持表锁, 不支持行锁.  innodb: 支持表锁和行锁, 如果基于索引查询数据则是行锁, 否则就是表锁
```

### mysql排他锁&共享锁(锁之间是否兼容)

```
1. 排他锁: for update
	又称为写锁, 是修改数据时创建的锁,写锁一旦被获取,那么其他事务都无法再修改数据.
	
2. 共享锁: lock in share mode
	又称为读锁, 是读取数据时创建的锁, 可以并发读取数据, 但是任何事务都不可以修改数据.
```

### mysql乐观锁&悲观锁(资源加不加锁)

```
1: 乐观锁
	乐观锁其实是对共享资源的一种概念, 认为在操作资源的过程中不会有其他人操作资源,所以不加锁;采用的是MVCC的方式机制实现;通过给表加上额外的字段(version),在查询时会将version一起读出去,如果修改了数据,就在原来的version上+1,提交时和当前版本的数据进行对比,如果大于则予以更新, 否则认识是过期的数据.
	
2: 悲观锁
	悲观锁和乐观锁相反,认为一定会有其他人来修改数据,所以都数据全部都加上锁.
```



### 为什么mysql事务  读已提交隔离级别 可以解决 脏读脏写问题

```
答: 使用事务A进行查询的同时, 有事务B对数据进行了修改, 但是事务B还未提交, 这时事务A查询到的数据其实是事务B做了修改但还未提交的数据, 而如果此时事务B进行回滚, 这就造成了脏读; 而mysql的读提交则是规定不允许一个事务读取未提交的事务.

注: 但其实mysql在事务查询没有根据索引进行查询时, 不会真的把所有记录都加锁; 实际的过程中, mysql做了一些改进, 在msyql server 过滤条件, 发现不满足后, 会调用unlock_row方法, 将不满足条件的记录释放锁.
```



### 为什么mysql事务  可重复读隔离级别 可以解决 不可重复读问题 & 无法解决幻读问题

```
答: 使用事务进行查询记录时, mysql会给所查询的记录上锁, 通过这样一个方式解决不可重复读问题. 但是, 在可重复读中, 事务A第一次读取记录时, 给这些记录加锁, 其他事务无法修改删除记录, 可以实现可重复读,却无法避免幻读; 因为这种方法无法锁住insert的数据, 所以当事务A先读取了记录或者修改了全部记录, 事务B还是可以insert记录提交, 这时事务A就会发现莫名其妙多出了一条记录, 这就是换读, 不能通过行锁来解决.
```



### 不可重复读和幻读的区别

```
不可重复读的重点在于 update和delete
幻读的重点在于 insert

所以说不可重复读和幻读最大的区别，就在于如何通过锁机制来解决他们产生的问题.
```



### 悲观锁和乐观锁

```
悲观锁:
	指的是对数据被外界修改持保守态度, 在悲观锁情况下, 为了保证事务的隔离性, 读取记录时需要加锁, 无法同时被其他事务修改; 修改时也需要加锁, 无法被其他事务读取.
	
乐观锁:
	乐观锁大多是基于数据版本(Version)来实现的. 通过给数据库表增加一个"version"字段, 读取数据时, 会把版本号一同读出, 更新时会将版本号加1, 此时将提交的数据版本和数据库内对应记录的当前版本进行对比, 如果提交的版本号大于数据库当前版本则予以更新, 否则认为是过期数据.
```



### MVCC(Multi-Verison Concurrency Control)

```
多版本并发控制解决了哪些问题?
1. 读写之间阻塞的问题(通过 MVCC 可以让读写互相不阻塞，即读不阻塞写，写不阻塞读，这样就可以提升事务并发处理能力)
2. 降低了死锁的概率(mysql采用乐观锁的方式, 读取的时候不需要锁, 对于写操作, 也只锁定必要的行)
3. 解决一致性读的问题(一致性读也被称作 快照读, 当我们查询在数据库在某个时间点的快照时, 只能看见这个时间点之前事务提交更新的结果, 而不能看到这个时间点之后事务提交的更新结果)
```



### 快照读 & 当前读

```
快照读:
	事务读取到的数据，要么是事务开始前就已经存在的数据，要么是事务自身插入或者修改过的数据。
	
	实现方式: 一般的SELECT, 通过undo log + MVCC 实现, 例:
	
	
	更新后: DB_Row_ID  DB_TRX_ID  DB_ROLL_Ptr   Pri_id(主键)     name
					   2			Ox012445    10          python
									|
								    ﹀
	更新前: DB_Row_ID  DB_TRX_ID  DB_ROLL_Ptr   Pri_id(主键)     name
					   1			Null		10			golang		 [undo log]

注1: insert undo log 只在事务回滚时需要, 事务提交就可以删掉了。update undo log 包括 update 和 delete , 回滚和快照读都需要。

注2: DB_TRA_ID:事务id, 6 字节 DB_TRX_ID 字段，表示最后更新的事务 id ( update , delete , insert ) . 此外，删除在内部被视为更新，其中行中的特殊位被设置为将其标记为已软删除.
	
注3: 	DB_ROLL_Ptr:回滚指针, 指向前一个版本的 undo log 记录, 7 字节回滚指针，指向前一个版本的 undo log 记录，组成 undo 链表。如果更新了行，则撤消日志记录包含在更新行之前重建行内容所需的信息。
	
注4: 	DB_ROW_ID：6-byte，隐藏的行 ID，用来生成默认聚簇索引。如果我们创建数据表的时候没有指定聚簇索引，这时 InnoDB 就会用这个隐藏 ID 来创建聚集索引。采用聚簇索引的方式可以提升数据的查找效率。

当前读:
	事务读取到的是最新的数据, 并且当前读返回的记录都会上锁, 保证其他并发事务不会修改这些记录
	
	实现方式: Next-key锁(record lock 行锁 + Gap lock 间隙锁)
	
注5: Gap lock 是 Innodb 为了解决幻读问题时引入的锁机制，所以只有在 Read Repeatable 、Serializable 隔离级别才有
```



### Innodb 和 MyIsum的区别

```
1. Innodb支持事务, MyiSum不支持事务
2. Innodb支持表级锁和行级锁, MyiSum只支持行级锁
3. Innodb支持外键, MyiSum不支持外键
4. MyiSum允许没有主键和索引的表存在, 索引都是保存行的地址; Innodb如果没有设置主键或非空唯一索引, 会自动生成6字节的主键
5. 存储结构方面: 
	Innodb: 聚簇索引, 将主键组织到一颗B+树上, 行数据存储在叶子节点上; 如果是主键索引, 按照B+树的检索算法直接就可以找到数据; 如果是辅助索引, 则需要两个步骤, 第一步在辅助索引B+树中检索, 到达叶子节点获取主键, 接着是根据主键在主键索引中找到数据.
	MyiSum: 非聚簇索引, 将主键组织到一颗B+树上, 表数据存在在独立的地方; 主键索引和辅助索引都只需要一步, 因为数据是独立存储的.
	
	
聚簇索引的好处:
	1. 因为行数据和叶子节点存储在一起, 在同一页中有多行数据, 当访问同一数据页不同行记录时, 已经把页加载到了buffer中, 再次访问时直接在内存中完成访问, 不需要访问磁盘.
	2. 辅助索引叶子节点存储的是主键, 而不是地址, 这样当数据页分裂或行数据移动时的维护工作
```



### 什么是数据页分裂

```
答: [1, 2, 3, 4, 5, 6, 7, 8] # page 5

    [10, 11, 12, 13, 14, 15, 16, 17] # page 6

	insert 9
	两页都是满的, 但是又不能不按顺序插入, 这个时候就会发生页分裂
	新建 page 7
	[1, 2, 3, 4       ] # page 5
	[5, 6, 7, 8, 9] # page 7
	[10, 11, 12, 13, 14, 15, 16, 17] # page 6
	
	重新定义页之间的关系
```



### 什么是数据页合并

```
答: 当删除记录时, 实际上记录并没有被物理删除, 而是被标记为删除, 并且记录的空间可以被其他记录声明使用, 当页中的记录达到阈值的时候, innodb会开始寻找最靠近的页看看能否将两页的记录合并以优化空间使用
```



### 什么是存储过程

```
	答: 存储过程是为了完成特定功能的sql语句集, 编译后存在数据库中
	
	优点: 
		1. 重复使用, 只需要一次创建, 后续就可以直接调用
		2. 执行更快, 因为已经编译过了
		3. 更安全, 可以防止sql注入攻击
	
```





### SQL查询语句是怎么执行的

```bash
# 例
mysql> select * from T where ID=10；
```



<img src="https://static001.geekbang.org/resource/image/0d/d9/0d2070e8f84c4801adbfa03bda1f98d9.png" style="zoom: 33%;" />





### SQL更新语句是如何执行的

```bash
# 例
mysql> update T set c=c+1 where ID=2;
```

<img src="https://static001.geekbang.org/resource/image/2e/be/2e5bff4910ec189fe1ee6e2ecc7b4bbe.png" style="zoom:50%;" />



```
	更新流程会涉及到 redo log(重做日志) 和 binlog(归档日志) 
	
#### 区别 ####

	1. redo log: 是固定大小的, 循环写 ;binlog: 是追加写入
	2. redo log 属于Innodb独有; binlog是Mysql的Server层实现的
	3. redo log 是物理日志, 记录的是某个数据页上做了什么修改;
	binlog是逻辑日志, 记录的是这个语句的原始逻辑, 有两种模式, statement和raw, statement记得是sql语   	 句, row记得是行的内容, 更新前和更新后都会记录

#### redo log两阶段提交 ####
	两份日志之间的逻辑一致
	由于redo log和binlog是两个独立的逻辑，如果不用两阶段提交，要么就是先写完redo log再写binlog，或者采用反过来的顺序。我们看看这两种方式会有什么问题。

	仍然用前面的update语句来做例子。假设当前ID=2的行，字段c的值是0，再假设执行update语句过程中在写完第一个日志后，第二个日志还没有写完期间发生了crash，会出现什么情况呢？

	先写redo log后写binlog。假设在redo log写完，binlog还没有写完的时候，MySQL进程异常重启。由于我们前面说过的，redo log写完之后，系统即使崩溃，仍然能够把数据恢复回来，所以恢复后这一行c的值是1。
但是由于binlog没写完就crash了，这时候binlog里面就没有记录这个语句。因此，之后备份日志的时候，存起来的binlog里面就没有这条语句。
然后你会发现，如果需要用这个binlog来恢复临时库的话，由于这个语句的binlog丢失，这个临时库就会少了这一次更新，恢复出来的这一行c的值就是0，与原库的值不同。

	先写binlog后写redo log。如果在binlog写完之后crash，由于redo log还没写，崩溃恢复以后这个事务无效，所以这一行c的值是0。但是binlog里面已经记录了“把c从0改成1”这个日志。所以，在之后用binlog来恢复的时候就多了一个事务出来，恢复出来的这一行c的值就是1，与原库的值不同。

	可以看到，如果不使用“两阶段提交”，那么数据库的状态就有可能和用它的日志恢复出来的库的状态不一致。
```


