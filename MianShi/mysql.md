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



### 数据库范式

##### 第一范式

```
数据库表的每一列都是不可分割的
```

#### 第二范式

```
表中的属性必须完全依赖于主键, 而不是部分主键(可以设置多个属性合为主键)
```

#### 第三范式

```
任何非主属性不依赖于其它非主属性
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
4. Serializable 串行化
	读用读锁, 写用写锁, 读锁和写锁互斥
```



### mysql全局锁&表锁&行锁&页面锁(锁的粒度)

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

4. 全局锁
	1): 全局读锁: Flush tables with read lock (FTWRL), 全局锁的典型使用场景是, 做全库逻辑备份.

注1: myisam: 只支持表锁, 不支持行锁.  innodb: 支持表锁和行锁, 如果基于索引查询数据则是行锁, 否则就是表锁


注2: 行锁的两阶段锁概念: 在innodb事务中, 行锁是在需要的时候才加上, 但并不是不需要了就立即释放, 而是需要等事务被提交了才释放.  有什么帮助: 事务中需要锁多个行, 把最可能造成锁冲突、最有可能影响并发度的锁尽量往后放
```



### 死锁的概念

```
	答: 当并发系统中不同线程出现循环资源依赖, 涉及的线程都在等待别的线程释放资源时, 就会导致线程都进入无限等待的状态, 称为死锁
	
注1: 死锁的两种处理策略:
		1). 直接进入等待, 直到超时, 这个超时时间可以通过参数innodb_lock_wait_timeout来设置
		2). 发起死锁检测, 发现死锁后, 主动回滚死锁链条中的某一个事务, 让其他事务得以继续执行, 将参数innodb_deadlock_detect设置为on，表示开启这个逻辑
		
注2: 等待超时处理死锁的机制是什么? 有什么局限?
		机制就是等待超时自动退出, 但是超时时间设置过长会导致其他事务等待时间过长, 超时时间设置太短又容易造成误伤, 多会有损业务

注3: 死锁检测处理所说的机制是什么? 有什么局限?
		每当一个事务被锁的时候, 就要看看它所依赖的线程有没有被别人锁住, 如此循环, 最后判断是否出现循环等待, 也就是死锁; 判断死锁是一个时间复杂度O(n)的操作

注4: 有哪些思路可以解决 热点行更新导致 的并发问题?
		1). 关闭死锁检测
		2). 控制并发度, 并发控制要做在数据库服务端, 基本思路就是，对于相同行的更新，在进入引擎之前排队。
		3). 将一行改成逻辑上的多行来减少锁冲突
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
	悲观锁和乐观锁相反,认为一定会有其他人来修改数据,所以数据全部都加上锁.
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
2. Innodb支持表级锁和行级锁, MyiSum只支持表级锁
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



### binlog 写入机制

```
	事务执行过程中, 先把日志写入binlog cache, 等事务提交时, 再将binlog cache写入binlog文件中
	
	系统给binlog cache分配了一片内存，每个线程一个，参数 binlog_cache_size用于控制单个线程内binlog cache所占内存的大小。如果超过了这个参数规定的大小，就要暂存到磁盘

	事务提交的时候，执行器把binlog cache里的完整事务写入到binlog中，并清空binlog cache
	
	每个线程有自己binlog cache，但是共用同一份binlog文件。
	
    图中的write，指的就是指把日志写入到文件系统的page cache，并没有把数据持久化到磁盘，所以速度比较快。
    图中的fsync，才是将数据持久化到磁盘的操作。一般情况下，我们认为fsync才占磁盘的IOPS。
    
    write 和fsync的时机，是由参数sync_binlog控制的：
    sync_binlog=0的时候，表示每次提交事务都只write，不fsync；
    sync_binlog=1的时候，表示每次提交事务都会执行fsync；
    sync_binlog=N(N>1)的时候，表示每次提交事务都write，但累积N个事务后才fsync。

    因此，在出现IO瓶颈的场景里，将sync_binlog设置成一个比较大的值，可以提升性能。在实际的业务场景中，考虑到丢失日志量的可控性，一般不建议将这个参数设成0，比较常见的是将其设置为100~1000中的某个数值。

    但是，将sync_binlog设置为N，对应的风险是：如果主机发生异常重启，会丢失最近N个事务的binlog日志
```

![](https://static001.geekbang.org/resource/image/9e/3e/9ed86644d5f39efb0efec595abb92e3e.png)

​									**binlog写盘状态**



### redo log 写入机制

```
	redo log buffer 是一块内存，用来先存 redo 日志的。 真正把日志写到 redo log 文件（文件名是 ib_logfile+ 数字），是在执行 commit 语句的时候做的。

	事务还没提交的时候，redo log buffer 中的部分日志有没有可能被持久化到磁盘呢？
这个问题，要从 redo log 可能存在的三种状态说起。这三种状态，对应的就是图中的三个颜色块

	这三种状态分别是：
    存在 redo log buffer 中，物理上是在 MySQL 进程内存中，就是图中的红色部分；
    写到磁盘 (write)，但是没有持久化（fsync)，物理上是在文件系统的 page cache 里面，也就是图中的黄色部分；
    持久化到磁盘，对应的是 hard disk，也就是图中的绿色部分。
    日志写到 redo log buffer 是很快的，wirte 到 page cache 也差不多，但是持久化到磁盘的速度就慢多了。为了控制 redo log 的写入策略，InnoDB 提供了 innodb_flush_log_at_trx_commit 参数，它有三种可能取值：

    设置为 0 的时候，表示每次事务提交时都只是把 redo log 留在 redo log buffer 中 ;
    设置为 1 的时候，表示每次事务提交时都将 redo log 直接持久化到磁盘；
    设置为 2 的时候，表示每次事务提交时都只是把 redo log 写到 page cache。
    InnoDB 有一个后台线程，每隔 1 秒，就会把 redo log buffer 中的日志，调用 write 写到文件系统的 page cache，然后调用 fsync 持久化到磁盘。

    注意，事务执行中间过程的 redo log 也是直接写在 redo log buffer 中的，这些 redo log 也会被后台线程一起持久化到磁盘。也就是说，一个没有提交的事务的 redo log，也是可能已经持久化到磁盘的。

    实际上，除了后台线程每秒一次的轮询操作外，还有两种场景会让一个没有提交的事务的 redo log 写入到磁盘中。

    一种是，redo log buffer 占用的空间即将达到 innodb_log_buffer_size 一半的时候，后台线程会主动写盘。注意，由于这个事务并没有提交，所以这个写盘动作只是 write，而没有调用 fsync，也就是只留在了文件系统的 page cache。
    另一种是，并行的事务提交的时候，顺带将这个事务的 redo log buffer 持久化到磁盘。假设一个事务 A 执行到一半，已经写了一些 redo log 到 buffer 中，这时候有另外一个线程的事务 B 提交，如果 innodb_flush_log_at_trx_commit 设置的是 1，那么按照这个参数的逻辑，事务 B 要把 redo log buffer 里的日志全部持久化到磁盘。这时候，就会带上事务 A 在 redo log buffer 里的日志一起持久化到磁盘。
```

![](https://static001.geekbang.org/resource/image/9d/d4/9d057f61d3962407f413deebc80526d4.png)

​									**redo log存储状态**



### 索引

```bash
	答: 关系数据库中, 对某一列或多个列的值进行排序的一种存储结构
	
1. 唯一索引和普通索引区别?
	1):	普通索引可以重复, 唯一索引和主键一样不可以写入重复的值
	2): 对于数据的修改, 普通索引可以使用change buffer, 而唯一索引不行

2. 全文索引(FULLTEXT)
	只有char, varchar, text列上可以创建, 为了解决WHERE name LIKE “%word%"这类针对文本的模糊查询效率较低的问题
```



### buffer pool & change buffer

> **buffer pool**

```
	innodb内存结构:	buffer pool 缓冲池是主内存中的一个区域, 用于innodb访问表和索引数据时进行缓存. 缓冲池允许直接从内存中访问经常使用的数据, 从而加快处理速度. 一种常见的降低磁盘访问的机制, 通常以页为单位缓存数据, 常见管理算法是LRU(least recently used).
```



> **change buffer** -- 适用于非唯一普通索引

```m
	对于读请求, 缓冲池能够减少磁盘IO, 提升性能. 那写请求呢?
	情况一: 要修改的数据刚好在缓冲池中,直接修改缓冲池中的数据,一次内存操作,再写入redo log, 一次io操作.
	情况二: 要修改的数据不在缓冲池, 需要从磁盘中加载到内存中, 一次io操作, 修改缓冲池的页, 一次内存操作, 写入redo log, 一次io操作.

	情况二没有命中缓冲池的时候，至少产生一次磁盘IO，对于写多读少的业务场景，是否还有优化的空间呢？
	
	答: change buffer.
	加入 change buffer后, 情况二变为: 在change buffer 中记录这个操作, 一次内存操作, 写入redo log, 一次io操作

注1: 描述数据
	当数据页被加载到缓冲池中后，Buffer Pool 中也有叫缓存页的概念与其一一对应，大小同样是 16KB，但是 MySQL还为每个缓存也开辟额外的一些空间，用来描述对应的缓存页的一些信息，例如：数据页所属的表空间，数据页号，这些描述数据块的大小大概是缓存页的5%左右（约800B）;【每个描述信息中有 free_pre、free_next 两个指针, Free链表 就是由这两个指针连接起来形成的一个双向链表。然后 Free链表 有一个基础节点，这个基础节点存放了链表的头节点地址、尾节点地址，以及当前链表中节点数的信息】

注2: 缓存页是什么时候被创建的
	当 MySql 启动的时候，就会初始化 Buffer Pool，这个时候 MySQL 会根据系统中设置的 innodb_buffer_pool_size 大小去内存中申请一块连续的内存空间，实际上在这个内存区域比配置的值稍微大一些，因为【描述数据】也是占用一定的内存空间的，当在内存区域申请完毕之后， MySql 会根据默认的缓存页的大小（16KB）和对应`缓存页*5%`大小(800B左右)的数据描述的大小，将内存区域划分为一个个的缓存页和对应的描述数据
```

![](http://117.78.7.18/admin/uploads/d8aa0c2a-da0f-4f01-8dec-a0d6d20258ea.png)

![](https://image-static.segmentfault.com/213/835/213835386-3ee266923478eaad_fix732)

​									**buffer pool**



### buffer pool 的并发性能 & 动态调整buffer pool大小

> **并发性能**

```
	Buffer Pool 一次只能允许一个线程来操作，一次只有一个线程来执行这一系列的操作，因为MySQL 为了保证数据的一致性，操作的时候必须缓存池加锁，一次只能有一个线程获取到锁, 所以 Buffer Pool 是可以有多个的, 可以通过配置mysql配置文件来配置:
    #  Buffer Pool  的总大小
    innodb_buffer_pool_size: 8589934592
    #  Buffer Pool  的实例数（个数）
    innodb_buffer_pool_instance: 4
	
	多个Buffer Pool带来的问题: 不同的Buffer Pool缓存中会去缓存相同的数据页吗? 答: 【数据页缓存哈希表】
```



> **如何动态调整大小? 【chunk机制】**

```
	chunk是 MySQL 设计的一种机制，这种机制的原理是将 Buffer Pool 拆分一个一个大小相等的 chunk 块，每个 chunk 默认大小为 128M（可以通过参数innodb_buffer_pool_chunk_size 来调整大小），也就是说 Buffer Pool 是由一个个的chunk组成的, 假设 Buffer Pool 大小是2GB，而一个chunk大小默认是128M，也就是说一个2GB大小的 Buffer Pool 里面由16个 chunk 组成，每个chunk中有自己的缓存页和描述数据，而 free 链表、flush 链表和 lru 链表是共享的
```

![](https://i2.wp.com/img-blog.csdnimg.cn/752c0667191f47c38b2f3bbdc6e04f91.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAT2NlYW4mJlN0YXI=,size_20,color_FFFFFF,t_70,g_se,x_16)



### LRU链表

```
	LRU, mysql的LRU和传统的LRU有所区别, 主要解决了 预读失效和缓冲池污染问题.

	传统LRU: 把加入缓冲池的页放到LRU的头部，作为最近访问的元素，从而最晚被淘汰。
	这里又分两种情况： 
	1）页已经在缓冲池里，那就只做“移至”LRU头部的动作，而没有页被淘汰；
	2）页不在缓冲池里，除了做“放入”LRU头部的动作，还要做“淘汰”LRU尾部页的动作；
	
	预读失效: 由于预读(Read-Ahead)，提前把页放入了缓冲池，但最终MySQL并没有从页中读取数据，称为预读失效.
	如何解决: 1) 让预读失败的页，停留在缓冲池LRU里的时间尽可能短
			2) 让真正被读取的页，才挪到缓冲池LRU的头部
	
	mysql LRU: 
		1) 将LRU分为两个部分(新生代new sublist, 老生代old sublist);
		2) 新老生代首尾相连，即：新生代的尾(tail)连接着老生代的头(head);
		3) 新页（例如被预读的页）加入缓冲池时，只加入到老生代头部; 如果数据真正被读取（预读成功），才会加入到新生代的头部; 如果数据没有被读取，则会比新生代里的“热数据页”更早被淘汰出缓冲池
		
		
	缓冲池污染: 当某一个SQL语句，要批量扫描大量数据时，可能导致把缓冲池的所有页都替换出去，导致大量热数据被换出，MySQL性能急剧下降，这种情况叫缓冲池污染
	
	mysql LRU: 
		1) 假设T=老生代停留时间窗口
		2) 插入老生代头部的页，即使立刻被访问，并不会立刻放入新生代头部
		3) 只有满足“被访问”并且“在老生代停留时间”大于T，才会被放入新生代头部
		
	innodb_buffer_pool_size: 配置缓冲池的大小，在内存允许的情况下，建议调大这个参数，越多数据和索引放到内存里，数据库的性能会越好
	innodb_old_blocks_pct: 老生代占整个LRU链长度的比例，默认是37，即整个LRU中新生代与老生代长度比例是63:37
	innodb_old_blocks_time: 老生代停留时间窗口，单位是毫秒，默认是1000，即同时满足“被访问”与“在老生代停留时间超过1秒”两个条件，才会被插入到新生代头部
```

![](https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/6/25/16b8cf6be0371901~tplv-t2oaga2asx-watermark.awebp)

![](https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/6/25/16b8cf6be0b40dec~tplv-t2oaga2asx-watermark.awebp)

![](https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/6/25/16b8cf6be0249553~tplv-t2oaga2asx-watermark.awebp)

​										 **传统LRU**



![](https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/6/25/16b8cf6be0470ed5~tplv-t2oaga2asx-watermark.awebp)

![](https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/6/25/16b8cf6be06a8fc1~tplv-t2oaga2asx-watermark.awebp)

![](https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/6/25/16b8cf6bf49c19c4~tplv-t2oaga2asx-watermark.awebp)

​								**Mysql LRU解决预读失效问题**



![](https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/6/25/16b8cf6bf4aa4dbc~tplv-t2oaga2asx-watermark.awebp)

![](https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/6/25/16b8cf6bf7ea6308~tplv-t2oaga2asx-watermark.awebp)

![](https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/6/25/16b8cf6bf92f95ed~tplv-t2oaga2asx-watermark.awebp)

![](https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/6/25/16b8cf6bfe072707~tplv-t2oaga2asx-watermark.awebp)

​								 	**Mysql LRU解决缓冲池污染问题**



### free 链表

```free链表
    用来存放空闲的缓存页的描述数据，如果某个缓存页被使用了，那么该缓存页对应的描述数据就会被从free链表中移除; 一个双向链表数据结构，这个链表的每个节点就是一个空闲缓存页的描述信息
    
注: 链表的基础节点占用的内存空间并不包含在 Buffer Pool 之内，而是单独申请的一块内存空间，每个基节点只占用40字节大小
```

![](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/e36431beec1745269b7673c84d2b56d5~tplv-k3u1fbpfcp-watermark.awebp)

​																					**freee 链表**

### Flush 链表

```
	被修改的脏数据都记录在 Flush 中，同时会有一个后台线程会不定时的将 Flush 中记录的描述数据对应的缓存页刷新到磁盘中，如果某个缓存页被刷新到磁盘中了，那么该缓存页对应的描述数据会从 Flush 中移除，同时也会从LRU链表中移除（因为该数据已经不在 Buffer Pool 中了，已经被刷入到磁盘，所以就也没必要记录在 LRU 链表中了），同时还会将该缓存页的描述数据添加到free链表中，因为该缓存页变得空闲了
	
	注: Flush链表也有一个基础节点，如果一个缓存页被修改了，就会加入到 Flush链表 中。但是不像 LRU链表 是从 Free链表 中来的，描述信息块中还有两个指针 flush_pre、flush_next用来连接形成 flush 链表，所以 Flush链表 中的缓存页一定是在 LRU 链表中的，而 LRU 链表中不在 Flush链表 中的缓存页就是未修改过的页, 脏页既存在于 LRU链表 中，也存在于 Flush链表 中。LRU链表 用来管理 Buffer Pool 中页的可用性，Flush链表 用来管理将页刷新回磁盘，二者互不影响。
```

![](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/53efe751907244d78f9588de4e6a6d37~tplv-k3u1fbpfcp-watermark.awebp)

​																			**Flush 链表**



### 缓存页哈希表

```
	有些数据页被加载到 Buffer Pool 的缓存页中了，那怎么知道一个数据页有没有被缓存呢？ 所以InnoDB还会有一个哈希表数据结构，它用 表空间号+数据页号 作key，value 就是缓存页的地址。当使用一个数据页的时候，会先通过表空间号+数据页号作为key去这个哈希表里查一下，如果没有就从磁盘读取数据页，如果已经有了，就直接使用该缓存页。
```

![](https://www.teqng.com/wp-content/uploads/2021/09/wxsync-2021-09-11654488abd8b04a5c4a7f7ac10772c7.png)



### char和varchar的区别

```
	char: 定长, 不足长度的字符串在其后补空字符, 范围是0~255字节
	varchar: 不定长, 范围是64k(64k是整行的长度, 需要考虑其他column)
```



### blob & text

```
blob:
	binary large object, 用于存储二进制大对象, 例如图片, 音视频.
	
	1) TINYBLOB 	0 - 255字节 	短文本二进制字符串
 	2) BLOB 		0 - 65KB 	 二进制字符串
 	3) MEDIUMBLOB 	0 - 16MB 	 二进制形式的长文本数据
 	4) LONGBLOB 	0 - 4GB 	 二进制形式的极大文本数据
	
text:
	text 类型同 char、varchar 类似，都可用于存储字符串，一般情况下，遇到存储长文本字符串的需求时可以考虑使用 text 类型
	1) TINYTEXT 	0 - 255字节 			一般文本字符串
 	2) TEXT 		0 - 65535字节 		长文本字符串
	3) MEDIUMTEXT 	0 - 16772150字节 		较大文本数据
	4) LONGTEXT 	0 - 4294967295字节 	极大文本数据

    对比 varchar ，text 类型有以下特点：
    1) text 类型无须指定长度。
    2) 若数据库未启用严格的 sqlmode ，当插入的值超过 text 列的最大长度时，则该值会被截断插入并生成警告。
    3) text 类型字段不能有默认值。
    4) varchar 可直接创建索引，text 字段创建索引要指定前多少个字符。
    5) text 类型检索效率比 varchar 要低
```





### MYSQL主备的基本原理

![](https://static001.geekbang.org/resource/image/fd/10/fd75a2b37ae6ca709b7f16fe060c2c10.png)

​									**MySQL主备切换流程**



```
	在状态1中，客户端的读写都直接访问节点A，而节点B是A的备库，只是将A的更新都同步过来，到本地执行。这样可以保持节点B和A的数据是相同的。 当需要切换的时候，就切成状态2。这时候客户端读写访问的都是节点B，而节点A是B的备库。
	在状态1中，虽然节点B没有被直接访问，但是依然建议把节点B（也就是备库）设置成只读（readonly）模式。这样做，有以下几个考虑：

	1. 有时候一些运营类的查询语句会被放到备库上去查，设置为只读可以防止误操作；

	2. 防止切换逻辑有bug，比如切换过程中出现双写，造成主备不一致；

	3. 可以用readonly状态，来判断节点的角色。

注: 把备库设置成只读了，还怎么跟主库保持同步更新呢？
	readonly设置对超级(super)权限用户是无效的，而用于同步更新的线程，就拥有超级权限
```

```
	节点A到B这条线的内部流程如下图就是一个update语句在节点A执行，然后同步到节点B的完整流程图
```

![](https://static001.geekbang.org/resource/image/a6/a3/a66c154c1bc51e071dd2cc8c1d6ca6a3.png)

​											**主备流程图**



```
	备库B跟主库A之间维持了一个长连接。主库A内部有一个线程，专门用于服务备库B的这个长连接。一个事务日志同步的完整过程是这样的：

	1.在备库B上通过change master命令，设置主库A的IP、端口、用户名、密码，以及要从哪个位置开始请求binlog，这个位置包含文件名和日志偏移量。

	2.在备库B上执行start slave命令，这时候备库会启动两个线程，就是图中的io_thread和sql_thread。其中io_thread负责与主库建立连接。

	3.主库A校验完用户名、密码后，开始按照备库B传过来的位置，从本地读取binlog，发给B。

	4.备库B拿到binlog后，写到本地文件，称为中转日志（relay log）。

	5.sql_thread读取中转日志，解析出日志里的命令，并执行
```



![](https://static001.geekbang.org/resource/image/20/56/20ad4e163115198dc6cf372d5116c956.png)

​									**MySQL主备切换流程--双M结构**



```
	binlog的特性确保了在备库执行相同的binlog，可以得到与主库相同的状态。

	因此，我们可以认为正常情况下主备的数据是一致的。也就是说，A、B两个节点的内容是一致的。第一张画的是M-S结构，但实际生产上使用比较多的是双M结构。
	
	但是，双M结构还有一个问题需要解决。

	业务逻辑在节点A上更新了一条语句，然后再把生成的binlog 发给节点B，节点B执行完这条更新语句后也会生成binlog。（建议把参数log_slave_updates设置为on，表示备库执行relay log后生成binlog）。 那么，如果节点A同时是节点B的备库，相当于又把节点B新生成的binlog拿过来执行了一次，然后节点A和B间，会不断地循环执行这个更新语句，也就是循环复制了。这个要怎么解决呢？

	MySQL在binlog中记录了这个命令第一次执行时所在实例的server id。因此，我们可以用下面的逻辑，来解决两个节点间的循环复制的问题：

	1.规定两个库的server id必须不同，如果相同，则它们之间不能设定为主备关系；

	2.一个备库接到binlog并在重放的过程中，生成与原binlog的server id相同的新的binlog；

	3.每个库在收到从自己的主库发过来的日志后，先判断server id，如果跟自己的相同，表示这个日志是自己生成的，就直接丢弃这个日志。

按照这个逻辑，如果我们设置了双M结构，日志的执行流就会变成这样：

	1.从节点A更新的事务，binlog里面记的都是A的server id；

	2.传到节点B执行一次以后，节点B生成的binlog 的server id也是A的server id；

	3.再传回给节点A，A判断到这个server id与自己的相同，就不会再处理这个日志。所以，死循环在这里就断掉了。
```



### MYSQL主备高可用性

> 主备延迟

```
	主备切换可能是一个主动运维动作，比如软件升级、主库所在机器按计划下线等，也可能是被动操作，比如主库所在机器掉电。

	在介绍主动切换流程的详细步骤之前，我要先跟你说明一个概念，即“同步延迟”。与数据同步有关的时间点主要包括以下三个：

	1.主库A执行完成一个事务，写入binlog，我们把这个时刻记为T1;

	2.之后传给备库B，我们把备库B接收完这个binlog的时刻记为T2;

	3.备库B执行完成这个事务，我们把这个时刻记为T3。

所谓主备延迟，就是同一个事务，在备库执行完成的时间和主库执行完成的时间之间的差值，也就是T3-T1。

可以在备库上执行show slave status命令，它的返回结果里面会显示seconds_behind_master，用于表示当前备库延迟了多少秒。

seconds_behind_master的计算方法是这样的：

每个事务的binlog 里面都有一个时间字段，用于记录主库上写入的时间；

备库取出当前正在执行的事务的时间字段的值，计算它与当前系统时间的差值，得到seconds_behind_master。

可以看到，其实seconds_behind_master这个参数计算的就是T3-T1。所以，可以用seconds_behind_master来作为主备延迟的值，这个值的时间精度是秒
```



> **主备延迟的来源**

```
1.备库所在机器的性能要比主库所在的机器性能差
	处理方式: 对称部署

2.备库的压力大
	处理方式:
		1) 一主多从。除了备库外，可以多接几个从库，让这些从库来分担读的压力。
		2) 通过binlog输出到外部系统，比如Hadoop这类系统，让外部系统提供统计类查询的能力。

3.大事务
	情况:
		1) delete语句删除太多数据
		2) 大表DDL
```



> **可靠性优先策略**

```
在双M结构下，从状态1到状态2切换的详细过程是这样的：

	1.判断备库B现在的seconds_behind_master，如果小于某个值（比如5秒）继续下一步，否则持续重试这一步；

	2.把主库A改成只读状态，即把readonly设置为true；

	3.判断备库B的seconds_behind_master的值，直到这个值变成0为止；

    4.把备库B改成可读写状态，也就是把readonly 设置为false；

	5.把业务请求切到备库B。

这个切换流程，一般是由专门的HA系统来完成的，我们暂时称之为可靠性优先流程。

可以看到，这个切换流程中是有不可用时间的。因为在步骤2之后，主库A和备库B都处于readonly状态，也就是说这时系统处于不可写状态，直到步骤5完成后才能恢复。

在这个不可用状态中，比较耗费时间的是步骤3，可能需要耗费好几秒的时间。这也是为什么需要在步骤1先做判断，确保seconds_behind_master的值足够小。

试想如果一开始主备延迟就长达30分钟，而不先做判断直接切换的话，系统的不可用时间就会长达30分钟，这种情况一般业务都是不可接受的。

当然，系统的不可用时间，是由这个数据可靠性优先的策略决定的。你也可以选择可用性优先的策略，来把这个不可用时间几乎降为0。
```



> **可用性优先策略**

```
如果强行把步骤4、5调整到最开始执行，也就是说不等主备数据同步，直接把连接切到备库B，并且让备库B可以读写，那么系统几乎就没有不可用时间了。

把这个切换流程，暂时称作可用性优先流程。这个切换流程的代价，就是可能出现数据不一致的情况。

注: 
	1.使用row格式的binlog时，数据不一致的问题更容易被发现。而使用mixed或者statement格式的binlog时，数据很可能悄悄地就不一致了。如果你过了很久才发现数据不一致的问题，很可能这时的数据不一致已经不可查，或者连带造成了更多的数据逻辑不一致。

	2.主备切换的可用性优先策略会导致数据不一致。因此，大多数情况下，都建议你使用可靠性优先策略。毕竟对数据服务来说的话，数据的可靠性一般还是要优于可用性的。
```



### 数据库优化的几个阶段

```bash
第一阶段: 优化sql和索引
	1. 选择合适的字段属性, 例如身份证为18位, 那么使用char(18)比varchar(255)要好
	2. 尽量把字段设置为 NOT NULL, 这样在查询时可以不用比较null值
	3. 对于某些文本字段来说, 例如"性别", "省份", 可以用枚举(ENUM)类型, 因为在mysql中, ENUM类型被当作数值型数据来处理
	4. 使用连接(join)来代替子查询(Sub-Queries)
	5. 避免使用函数索引

第二阶段: 引入缓存数据库
	1. 将复杂的、耗时的、不常变的执行结果缓存起来
	
第三阶段: 读写分离
	1. 一主多从, 主从的好处: 实现数据库备份，实现数据库负载均衡，提高数据库可用性
	2. 主从的原理: 主库和从库有长连接, 主库有一个线程把bin log发送给从库, 从库有一个线程把接收bin log写入relay log, 从库还有一个线程负责把relay log读取内容写入从库数据库
	3. 如何解决主从一致性:
		1). 半同步复制: 在主库写完binlog后需要从库返回一个已接收才返回给客户端(主库写请求延时增加, 减低吞吐量)
		2). 数据库中间件: 所有读写都走中间件, 通常情况下, 写请求路由到主库, 读请求路由到从库记录所有路由到写库的key, 在主从同步时间内, 如果有读请求访问中间件, 此时从库可能还是旧数据, 就把这个key上的读请求路由到主库, 在主从同步时间过完后, 对应key的读请求继续路由到从库(绝对的一致性, 但成本高)
		3). 缓存记录写key法: 写流程, 如果key要发生写操作, 记录在cache里, 并设置"经验主从同步时间"的cache的超时时间, 然后修改主库; 读流程, 先到缓存里查看, 对应key有没有相关数据, 有就说明命中缓存, 这个key刚发生了写操作, 此时需要路由到主库读数据, 如果没有命中缓存, 说明没有发生过写操作, 此时路由到从库, 继续读写分离(系统变得更为复杂)

第四阶段: 分库分表
	1. 垂直拆分
	2. 水平拆分

```



### inner join

|  id  |   name   |
| :--: | :------: |
|  1   |  Google  |
|  2   |   淘宝   |
|  3   |   微博   |
|  4   | Facebook |

|  id  | name |
| :--: | :--: |
|  1   | 美国 |
|  5   | 中国 |
|  3   | 中国 |
|  6   | 美国 |

```mysql
# 语法  INNER JOIN产生的结果集中，是1和2的交集。
# select column_name(s)
# from table 1
# INNER JOIN table 2
# ON
# table 1.column_name=table 2.column_name

select * 
from Table A 
inner join Table B
on 
Table A.id=Table B.id
```

|  id  |  name  | address |
| :--: | :----: | :-----: |
|  1   | Google |  美国   |
|  3   |  微博  |  中国   |



### left join

```mysql
# 语法  LEFT JOIN产生表1的完全集，而2表中匹配的则有值，没有匹配的则以null值取代。
# select column_name(s)
# from table 1
# LEFT JOIN table 2
# ON table 1.column_name=table 2.column_name

select * 
from Table A 
left join Table B
on Table A.id=Table B.id
```

|  id  |   name   | address |
| :--: | :------: | :-----: |
|  1   |  Google  |  美国   |
|  2   |   淘宝   |  null   |
|  3   |   微博   |  中国   |
|  4   | Facebook |  null   |



### right join

```mysql
# 语法  RIGHT JOIN产生表2的完全集，而1表中匹配的则有值，没有匹配的则以null值取代。
# select column_name(s)
# from table 1
# RIGHT JOIN table 2
# ON table 1.column_name=table 2.column_name

select * 
from Table A 
right join Table B
on Table A.id=Table B.id
```

|  id  |  name  | address |
| :--: | :----: | :-----: |
|  1   | Google |  美国   |
|  5   |  null  |  中国   |
|  3   |  微博  |  中国   |
|  6   |  null  |  美国   |



### full outer join

```mysql
# 语法  FULL OUTER JOIN产生1和2的并集。但是需要注意的是，对于没有匹配的记录，则会以null做为值。
# select column_name(s)
# from table 1
# FULL OUTER JOIN table 2
# ON table 1.column_name=table 2.column_name

select * 
from Table A 
full outer join Table B
on Table A.id=Table B.id
```

|  id  |   name   | address |
| :--: | :------: | :-----: |
|  1   |  Google  |  美国   |
|  2   |   淘宝   |  null   |
|  3   |   微博   |  中国   |
|  4   | Facebook |  null   |
|  5   |   null   |  美国   |
|  6   |   null   |  美国   |





















