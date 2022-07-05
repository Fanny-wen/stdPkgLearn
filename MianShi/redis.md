# Redis

## 概述

>redis是什么?

Redis ( Remote Dictionary Server ), 即远程字典服务



> 特性

1. 多样的数据类型
2. 持久化
3. 集群
4. 事务
5. ...



> 基础知识

```
默认端口 6379
有16个数据库
```

清除当前数据库 `flushdb`

```sql
> flushdb
```

清楚所有数据库 `flushall`

```sql
> flushall
```



> redis是单线程的!

redis是基于内存操作,  CPU不是redis性能的瓶颈,  是根据机器的内存和网络带宽, 既然可以使用单线程来实现, 就使用单线程了

### 基本命令

```bash
# 判断key是否存在
EXISTS [key]

# 设置key过期时间
EXPIRE [key] seconds
# 查看过期剩余时间
ttl [key]

# 查看当前key的类型
type [key]
```





## 问题

1. redis为什么这么快?
2. redis中的多线程
3. redis删除策略
4. redis淘汰策略
5. redis缓存一致性
6. redis核心对象
7. redis数据类型
8. redis持久化
9. redis如何进行缓存预热
10. 缓存穿透、缓存击穿、缓存雪崩
11. redis主从复制
12. 哨兵机制
13. redis分片cluster
14. redis实现异步队列
15. redis实现分布式锁
16. redis事务
17. 多线程优化



<img src="http://s3.51cto.com/oss/202009/04/6039cd0b01a4b1b0de8397fdbfcca076.png" style="zoom: 150%;" />



### 1.为什么redis是单线程还这么快

```bash
答:
	1. (核心) redis是将所有的数据全部放在内存中, 读写都是在内存中完成的;
	2. k-v型数据库, 内部构建一个哈希表(hashmap), 查找和操作的时间复杂度都是O(1);
	3. 采用单线程, 没有了多线程的上下文切换的性能消耗, 没有了访问共享资源加锁的性能消耗;
	4. 采用IO多路复用技术, 非阻塞IO;
```





### 2.redis中的多线程

>  https://www.modb.pro/db/231149

```bash
	redis的单线程指的是网络IO和键值对读写只有一个线程
	但是对于其他功能如: 持久化,异步删除, 集群数据同步等, 其实是由额外的线程执行的
```



### 3.redis删除策略

> 博客链接: https://www.cnblogs.com/ysocean/p/12422635.html
>
> 博客链接: https://developer.aliyun.com/article/666405



#### 定时删除

```bash
创建一个定时器，当key设置有过期时间，且过期时间到达时，由定时器任务立即执行对键的删除操作

优点：节省内存，到时就删除，快速释放掉不必要的内存空间

缺点：CPU压力大，无论此时CPU过载有多高，都会占用CPU，会影响Redis服务器的响应时间和吞吐量

总结：用处理器性能换取内存空间(时间换空间)
```

![](http://s4.51cto.com/oss/202108/02/886233cccac300380d3baab354674328.png)



#### 惰性删除

```bash
数据到达过期时间后，不做处理。等下次访问时，

如果未过期，返回数据
如果已过期，删除并返回不存在
优点：节约CPU性能，发现必须删除时才删除

缺点：内存压力大，出现长期占用内存空间的数据

总结：用内存空间换取CPU处理性能(空间换时间)
```

![](http://s6.51cto.com/oss/202108/02/96509571892e959247873b956d449468.png)



#### 定期删除

```bash
每隔默认的 100 ms 随机抽取一些设置了过期时间的 key，检查是否过期，如果过期就删除

特点： 
	CPU占用设置有峰值，检测频度可以自定义
	内存压力不是很大，长期占用内存的冷数据会被持续清理

注: 为什么是100ms, Redis服务器启动初始化时，读取配置server.hz的值，默认为10, 意思是每秒运行10次
```



#### redis过期删除策略

```bash
Redis的过期删除策略就是：惰性删除和定期删除两种策略配合使用。

　　惰性删除：Redis的惰性删除策略由 db.c/expireIfNeeded 函数实现，所有键读写命令执行之前都会调用 expireIfNeeded 函数对其进行检查，如果过期，则删除该键，然后执行键不存在的操作；未过期则不作操作，继续执行原有的命令。

　　定期删除：由redis.c/activeExpireCycle 函数实现，函数以一定的频率运行，每次运行时，都从一定数量的数据库中取出一定数量的随机键进行检查，并删除其中的过期键。

　　注意：并不是一次运行就检查所有的库，所有的键，而是随机检查一定数量的键。

　　定期删除函数的运行频率，在Redis2.6版本中，规定每秒运行10次，大概100ms运行一次。在Redis2.8版本后，可以通过修改配置文件redis.conf 的 hz 选项来调整这个次数。
```





### 4.redis淘汰策略

> 博客链接: https://www.cnblogs.com/ysocean/p/12422635.html



#### 设置Redis最大内存

```bash

在配置文件redis.conf 中，可以通过参数 maxmemory <bytes> 来设定最大内存：
不设定该参数默认是无限制的，但是通常会设定其为物理内存的四分之三

注: maxmermory：占用物理内存的比例。默认值是0，标识不限制。生产上根据需要设置，一般在50%以上 
```



#### 设置内存淘汰方式

```bash
当现有内存大于 maxmemory 时，便会触发redis主动淘汰内存方式，通过设置 maxmemory-policy ，有如下几种淘汰方式：

　　1）volatile-lru   利用LRU算法移除设置过过期时间的key (LRU:最近使用 Least Recently Used ) 。

　　2）allkeys-lru   利用LRU算法移除任何key （和上一个相比，删除的key包括设置过期时间和不设置过期时间的）。通常使用该方式。

　　3）volatile-random 移除设置过过期时间的随机key 。

　　4）allkeys-random  无差别的随机移除。

　　5）volatile-ttl   移除即将过期的key(minor TTL) 

　　6）noeviction 不移除任何key，只是返回一个写错误 ，默认选项，一般不会选用。
　　
注: maxmermroy-samples: 选取待删除的数据时，如果扫描全库，会严重消耗性能，降低读写性能。因为采用随机获取数据的方式作为待检测删除数据

注: maxmermory-policy ：达到最大内存后，对被挑选出来的数据进行删除的策略 
```



> redis LRU实现
>
> 博客链接: https://segmentfault.com/a/1190000017555834



### 5.redis缓存一致性

>博客链接: https://coolshell.cn/articles/17416.html
>
>博客链接: https://blog.51cto.com/u_14299052/2988627
>
>博客链接: https://blog.51cto.com/u_15257216/3024323

```
什么是缓存一致性问题?
	不管是先写MySQL数据库，再删除Redis缓存；还是先删除缓存，再写库，都有可能出现数据不一致的情况。
	举一个例子： 1.如果删除了缓存Redis，还没有来得及写库MySQL，另一个线程就来读取，发现缓存为空，则去数据库中读取数据写入缓存，此时缓存中为脏数据。 2.如果先写了库，在删除缓存前，写库的线程宕机了，没有删除掉缓存，则也会出现数据不一致情况。 因为写和读是并发的，没法保证顺序,就会出现缓存和数据库的数据不一致的问题
	
更新缓存有4中模式: Cache aside, Read through, Write through, Write behind caching

最好的办法是给缓存设置过期时间

Cache aside:
	失效：应用程序先从cache取数据，没有得到，则从数据库中取数据，成功后，放到缓存中
	命中：应用程序从cache中取数据，取到后返回
	更新：先把数据存到数据库中，成功后，再让缓存失效


Read-Through/Write-Through:
	Read-Through: 缓存配置一个读模块，它知道如何将数据库中的数据写入缓存。在数据被请求的时候，如果未命中，则将数据从数据库载入缓存。
	Write-Through: 缓存配置一个写模块，它知道如何将数据写入数据库。当应用要写入数据时，缓存会先存储数据，并调用写模块将数据写入数据库。
	Read Through/Write Through适用于写入之后经常被读取的应用

Write-behind:
	和page cache一样;
	在更新数据的时候，只更新缓存，不更新数据库，而缓存会异步地批量更新数据库
	适合频繁写的场景, MySQL的InnoDB Buffer Pool机制就使用到这种模式
	但有数据丢失的风险，如果缓存挂掉而数据没有及时写到数据库中，那么缓存中的有些数据将永久的丢失
```





### 6.redis核心对象





### 7.redis数据类型

```bash
1. String
2. List
3. Set
4. Hash
5. Zset(sorted set)

6. Bit arrays
7. HyperLogLogs
8. Stream
```



#### String

```bash
# 设置key
> SET [key]

# 获取key
> GET [key]

# 获取字符串长度
> STRLEN [key]

# 自增1
> INCR [key]

# 设置步长 指定增量
> INCRBY [key]

# 设置步长, 指定float类型增量
> INCRBYFLOAT [key] [increment]

# 自减1
> DECR [key]

# 设置步长 指定减量
> DECRBy [key]

# 字符串范围 range
> GETRANGE [key] start end

# 替换指定位置的字符串
> SETRANGE [key] offset value

# set with expire设置值并设置过期时间
> SETEX [key] seconds value

# set with not exist
# 如果key不存在则创建
# 如果存在, 则无效
# (在分布式锁回常常使用)
> SETNX [key] value

# 设置多个值
> MSET [key value ...]

# 获取多个值
> MGET [key ...]

# 是一个原子性的操作
> MSETNX

# 先获取再设置, 如果不存在则返回nil, 存在则返回原来的值
> GETSET [key]

# 追加value到key上
> APPEND [key] [value]
```



#### List

```bash
# 将一个值或者多个值, 插入到list的左边
> LPUSH [key] [value ...]

# 将一个值或者多个值, 插入到list的右边
> RPUSH [key] [value ...]

# 通过区间获取具体的值
> LRANGE [key] start end

# 移除list第一个元素
> LPOP [key]

# 移除list最后一个元素
> RPOP [key]

# 通过下标获取list中的某一个值
> LINDEX [key] [index]

# 获取list的长度
> LLEN [key]

# 移除指定的值
> LREM [key] [count] [value]

# 通过下标截取指定的长度, list被改变, 只剩下截取的元素
> LTRIM [key] [start] [stop]

# 移除list最后一个元素, 将元素移动到另一个的list
> RPOPLPUSH [source] [destination]

# 将list中指定下标的值替换为另一个值, 类似与更新操作
# 如果下标值不存在则报错
> LSET [key] [index] [value]

# 将某个值插入到list中某个值的前面或者后面
> LINSERT [key] [defore]|[after] [pivot] [value]
```

##### List小结

```
实际上是一个链表
如果key不存在, 创建新的链表
如果key存在, 新增内容
如果移除了所有值, 空链表, 也代表不存在
在两边插入或者改动值效率最高
```



#### Hash 

```bash
# 设置hash里一个字段的值
> HSET	[key] [field] [value]

# 设置hash里多个字段的值
> HMSET [key] [key] [field value ...]

# 获取hash指定field的值
> HGET [key] [field]

# 获取hash多个字段和值
> HMGET [key] [field ...]

# 获取hash所有字段和值
> HGETALL

# 获取hash所有字段
> HKEYs [key] 

# 获取hash所有值
> HVALS

# 删除hash指定field
> HDEL [key] [field]

# 获取hash字段数量
> HLEN [key]

# 判断hash是否有该字段
> HEXISTS [key] [field]

# 将hash指定field增加指定增量
> HINCRBY [key] [field] [increment]

# 将hash指定field增加指定浮点型增量
> HINCRBYFLOAT [key] [field] [increment]

# 如果不存在则设置, 存在则无效
> HSETNX [key] [field] [value]
```





#### Set

```bash
# 向set添加元素
> SADD [key] [value]

# 查看set所有值
> SMEMBERS [key]

# 判断某个value是否在set中
> SISMEMBER [key] [value]

# 获取set中的元素个数
> SCARD [key]

# 移除set中的指定元素
> SREM [key] [value ...]

# 随机抽选出指定个数的元素
> SRANDMEMBER [key] [count]

# 随机删除元素
> SPOP [key] [count]

# 将指定的值移动到另一个set
> SMOVE [source] [destination] [member]

# 并集
> SUNION [key] [key ...]

# 交集
> SINTER [key] [key ...]

# 差集
> SDIFF [key] [key]
```



#### Zset

> 在set的基础上, 增加了一个值

```bash
# 将一个或多个score/member对 添加到zset中
# XX: 仅仅更新存在的成员，不添加新成员
# NX: 不更新存在的成员。只添加新成员
# CH: 修改返回值为发生变化的成员总数，原始是返回新添加成员的总数 (CH 是 changed 的意思)
# INCR: 当ZADD指定这个选项时，成员的操作就等同ZINCRBY命令，对成员的分数进行递增操作。
> ZADD [key] [nx | xx] [ch] [incr] [score] [member] [score member ...]

# 获取zset元素个数, key存在时返回个数, 否则返回0
> ZCARD [key]

# 获取zset min<= score <= max 的个数, 可以使用 -inf 和 +inf, 这样就是获取所有
> ZCOUNT [key] [min] [max]

# 获取zset min<= member <= max 的个数
# 成员名称前需要加 [ 符号作为开头, [ 符号与成员之间不能有空格
# 可以使用 - 和 +, 这样就是获取所有
# 计算成员之间的成员数量时,参数 min 和 max 的位置也计算在内。
> ZLEXCOUNT [key] [min] [max]

# 给zset的 member的score增加increment 
# 如果member不存在, 就添加一个member, score为increment
# 如果key不存在, 就创建一个zset, 这个zset只含有一个元素member, score为increment
> ZINCRBY [key] [increment] [member]

# 删除并返回有序集合key中的最多count个具有最高score的成员
# 如未指定，count的默认值为1。指定一个大于有序集合的基数的count不会产生错误。 
# 当返回多个元素时候，得分最高的元素将是第一个元素，然后是分数较低的元素
> ZPOPMAX [key] [count]

# 删除并返回有序集合key中的最多count个具有最低score的成员
# 如未指定，count的默认值为1。指定一个大于有序集合的基数的count不会产生错误。 
# 当返回多个元素时候，得分最低的元素将是第一个元素，然后是分数较高的元素。
> ZPOPMIN [key] [count]

# 移除zset中的一个或多个成员，不存在的成员将被忽略
> ZREM [key] [member ...]

# 取 numkeys个zset的交集并把结果放到destination中
# numkeys 表示多少个zset
# weights 表示为每个给定的有序集指定一个乘法因子 
# AGGREGATE 表示聚合方式, 有sum, min, max
# 例: ZINTERSTORE out 2 zset1 zset2 WEIGHTS 2 3
> ZINTERSTORE [destination] [numkeys] [key ...] [weights] [aggregate sum|min|max]

# 取 numkeys个zset的并集并把结果放到destination中
> ZUNIONSTORE [destination] [numkeys] [key ...] [weights] [aggregate sum|min|max]

# 返回 start <= score <= stop区间的成员, 按分数值递增(从小到大)来排序, 具有相同分数值的成员按字典序(lexicographical order)来排列。
# WITHSCORES 将元素的分数与元素一起返回
> ZRANGE [key] [start] [stop] [WITHSCORES]

# 返回zset中 min<= member <= max，按成员字典正序排序, score必须相同!!!
# “max” 和 “min” 参数前必须加 “[” 符号作为开头, “max” 和 “min” 参数前可以加 “(“ 符号作为开头表示小于
# 例: ZRANGEBYLEX zset - + LIMIT 0 3
> ZRANGEBYLEX [key] [min] [max] [LIMIT offset count]

# 返回 start <= score <= stop区间的成员, 按分数值递增(从小到大)来排序, 具有相同分数值的成员按字典序(lexicographical order)来排列。
# LIMIT参数指定返回结果的数量及区间
# 可以使用 -inf +inf
> ZRANGEBYSCORE [key] [min] [max] [WITHSCORES] [LIMIT offset count]

# 获得member按score值递减(从大到小)排列的排名
> ZRANK [key] [member]

# 获得member按score值递减(从小到大)排列的排名
> ZREVRANK [key] [member]

# 删除Zset中指定排序中 start <= index <= stop 的member
# 可以使用 -1 -2
# 0处 是小的元素
> ZREMRANGEBYRANK [key] [start] [stop]

# 返回Zset中 指定member的score值
> ZSCORE [key] [member]
```



#### Geospatial

```bash
# 向GEO中添加多个元素
# 规则: 有效的经度从-180度到180度; 有效的纬度从-85.05112878度到85.05112878度
> GEOADD [key] [longitude latitude member ...]

# 返回members的11个字符的Geohash字符串
# 将二维的经纬度转换为一维的字符串
> GEOHASH [key] [member ...]

# 获取 指定members的经度和纬度
> GEOPOS [key] [member ...]

# 获取两个member之间的距离
# unit是单位, m|km|ft|mi
> GEODIST [key] [member1] [member2] [unit]

# 以给定的经度、纬度为中心, 以指定的半径查找周围的member
> GEORADUIS [key] [longitude] [latitude] [radius m|km|ft|mi] [withcoord] [withdist] [count] [ASC|DESC]

# 以给定的member为中心, 以指定的半径查找周围的member
> GEORADUISBYMEMBER [key] [member] [radius m|km|ft|mi] [withcoord] [withdist] [withhash] [count] [ASC|DESC]
```

>GEO底层实现原理其实是Zset, 所以可以使用Zset命令操作GEO



```bash
# 查看GEO全部元素
> ZRANGE [key] [start] [stop] [withscores]

# 删除GEO 指定元素
> ZREM [key] [member ...]
```



#### Hyperloglog

> 基数统计的算法
>
> 优点: 占用的内存是固定的, 2^64 不同的元素的基数, 只需要12KB内存

```bash
# 创建Hyperloglog
> PFADD [key] [element ...]

# 统计Hyperloglog的基数数量
> PFCOUNT [key ...]

# 合并多组 Hyperloglog 到 destkey
> PFMERGE [destkey] [sourcekey ...]
```



#### Bitmaps

> 位存储, 只有0和1两个状态

```bash
# 设置或清空bitmaps在offset处的value
# 0<= offset < 232
# value 只能是 0或1
> SETBIT [key] [offset] [value]

# 获取bitmap在offset处的value
> GETBIT [key] [offset]

# 统计bitmaps为1的数量
> BITCOUNT [key] [start end]

# 返回bitmaps第一个被设为0或1的bit位
# bit 只能填 0 或 1
> BITPOS [key] [bit] [start] [end]

# 一个或多个保存二进制位的字符串 key 进行位元操作，并将结果保存到 destkey 上
# operation 支持 AND(求逻辑并) 、 OR(求逻辑或) 、 NOT (求逻辑非)、 XOR(求逻辑异或) 
> BITOP [operation] [destkey] [key ...]

# 
> BITFIELD

# 
> BITFIELD_RO
```



### 8.redis持久化

> 博客链接: https://whetherlove.github.io/2018/10/05/Redis%E5%85%A5%E9%97%A8-%E6%8C%81%E4%B9%85%E5%8C%96/

```bash
为了防止数据丢失以及服务重启时能够恢复数据，Redis支持数据的持久化，主要分为两种方式，分别是RDB和AOF; 当然实际场景下还会使用这两种的混合模式
```

#### RDB持久化

```
RDB 就是 Redis DataBase 的缩写，中文名为快照/内存快照，RDB持久化是把当前进程数据生成快照保存到磁盘上的过程，由于是某一时刻的快照，那么快照中的值要早于或者等于内存中的值。

触发方式:
	手动触发:
		save命令: 阻塞当前Redis服务器，直到RDB过程完成为止，对于内存 比较大的实例会造成长时间阻塞，线上环境不建议使用
		bgsave命令:Redis进程执行fork操作创建子进程，RDB持久化过程由子进程负责，完成后自动结束。阻塞只发生在fork阶段，一般时间很短
		bgsave流程如下:
		1) 执行bgsave命令，Redis父进程判断当前是否存在正在执行的子进程，如只RDB/AOF子进程，如果存在bgsave命令直接返回。

		2) 父进程执行fork操作创建子进程，fork操作过程中父进程会阻塞，通过info stats命令查看latest_fork_usec选项，可以获取最近一个fork以操作的耗时，单位为微秒。

		3) 父进程仍fork完成后，bgsave命令返回“Background saving started”信息并不再阻塞父进程，可以继续响应其他命令。

		4) 子进程创建RDB文件，根据父进程内存生成临时快照文件，完成后对原有文件进行原子替换。执行lastsave命令可以获取最后一次生成尺RDB的时间，对应info统计的rdb_last_save_time选项。

		5) 进程发送信号给父进程衣示完成，父进程更新统计信息，具体见info Persistence下的rdb_*相关选项。
		
	自动触发:
		有如下四种情况时自动触发:
		1) redis.conf中配置 save m n, 即在m秒内有n次修改时, 自动触发bgsave生成rdb文件
		2) 主从复制时, 从节点要从主节点进行全量复制时也会触发bgsave操作,生成当时的快照发送到从节点
		3) 执行debug reload命令重新加载redis时也会触发bgsave操作
		4) 默认情况下执行shutdown命令时, 如果没有开启aof持久化, 也会触发bgsave操作
		
	优点:
		1) RDB文件是某个时间节点的快照，默认使用LZF算法进行压缩，压缩后的文件体积远远小于内存大小，适用于备份、全量复制等场景；
		2) Redis加载RDB文件恢复数据要远远快于AOF方式；
	缺点:
		1) RDB方式实时性不够，无法做到秒级的持久化； 
		2) 每次调用bgsave都需要fork子进程，fork子进程属于重量级操作，频繁执行成本较高； 
		3) RDB文件是二进制的，没有可读性，AOF文件在了解其结构的情况下可以手动修改或者补全； 
		4) 版本兼容RDB文件问题；
```

![](http://whetherlove.github.io/images/rdb.png)

#### AOF持久化 

```bash
AOF日志采用写后日志，即先写内存，后写日志。

	为什么采用写后日志？ 
		Redis要求高性能，采用写日志有两方面好处：1避免额外的检查开销：Redis 在向 AOF 里面记录日志的时候，并不会先去对这些命令进行语法检查。所以，如果先记日志再执行命令的话，日志中就有可能记录了错误的命令，Redis 在使用日志恢复数据时，就可能会出错; 2不会阻塞当前的写操作 
		但这种方式存在潜在风险： 如果命令执行完成，写日志之前宕机了，会丢失数据。 主线程写磁盘压力大，导致写盘慢，阻塞后续操作。
```



##### 如何实现AOF

```bash
AOF日志记录Redis的每个写命令，步骤分为：命令追加（append）、文件写入（write）和文件同步（sync）

	1.命令追加 当AOF持久化功能打开了，服务器在执行完一个写命令之后，会以协议格式将被执行的写命令追加到服务器的 aof_buf 缓冲区。
	2.文件写入和同步 关于何时将 aof_buf 缓冲区的内容写入AOF文件中，Redis提供了三种写回策略:
		appendfsync always：将aof_buf缓冲区的所有内容写入并同步到AOF文件，每个写命令同步写入磁盘
		appendfsync everysec：将aof_buf缓存区的内容写入AOF文件，每秒同步一次，该操作由一个线程专门负责
		appendfsync no：将aof_buf缓存区的内容写入AOF文件，什么时候同步由操作系统来决定

```

![](http://whetherlove.github.io/images/aof.png)



##### 深入理解AOF重写

```bash
	AOF会记录每个写命令到AOF文件，随着时间越来越长，AOF文件会变得越来越大。如果不加以控制，会对Redis服务器，甚至对操作系统造成影响，而且AOF文件越大，数据恢复也越慢。为了解决AOF文件体积膨胀的问题，Redis提供AOF文件重写机制来对AOF文件进行“瘦身”
```

![](https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/7/30/16c4345eb3719dac~tplv-t2oaga2asx-zoom-in-crop-mark:1304:0:0:0.awebp)

```bash
AOF重写会阻塞吗?

	AOF重写过程是由后台进程bgrewriteaof来完成的。主线程fork出后台的bgrewriteaof子进程，fork会把主线程的内存拷贝一份给bgrewriteaof子进程，这里面就包含了数据库的最新数据。然后，bgrewriteaof子进程就可以在不影响主线程的情况下，逐一把拷贝的数据写成操作，记入重写日志。 所以aof在重写时，在fork进程时是会阻塞住主线程的
```



```bash'
AOF日志何时会重写？

	有两个配置项控制AOF重写的触发： 
	auto-aof-rewrite-min-size:表示运行AOF重写时文件的最小大小，默认为64MB。 
	auto-aof-rewrite-percentage:这个值的计算方式是，当前aof文件大小和上一次重写后aof文件大小的差值，再除以上一次重写后aof文件大小。也就是当前aof文件比上一次重写后aof文件的增量大小，和上一次重写后aof文件大小的比值。
```



```bash
重写日志时，有新数据写入咋整？

	重写过程总结为：“一个拷贝，两处日志”。在fork出子进程时的拷贝，以及在重写时，如果有新数据写入，主线程就会将命令记录到两个aof日志内存缓冲区中。如果AOF写回策略配置的是always，则直接将命令写回旧的日志文件，并且保存一份命令至AOF重写缓冲区，这些操作对新的日志文件是不存在影响的。（旧的日志文件：主线程使用的日志文件，新的日志文件：bgrewriteaof进程使用的日志文件）
    而在bgrewriteaof子进程完成会日志文件的重写操作后，会提示主线程已经完成重写操作，主线程会将AOF重写缓冲中的命令追加到新的日志文件后面。这时候在高并发的情况下，AOF重写缓冲区积累可能会很大，这样就会造成阻塞，Redis后来通过Linux管道技术让aof重写期间就能同时进行回放，这样aof重写结束后只需回放少量剩余的数据即可。
    最后通过修改文件名的方式，保证文件切换的原子性。 
    在AOF重写日志期间发生宕机的话，因为日志文件还没切换，所以恢复数据时，用的还是旧的日志文件
```

![](http://whetherlove.github.io/images/rewrite.png)

```
主线程fork出子进程的是如何复制内存数据的？

	fork采用操作系统提供的写时复制（copy on write）机制，就是为了避免一次性拷贝大量内存数据给子进程造成阻塞。fork子进程时，子进程时会拷贝父进程的页表，即虚实映射关系（虚拟内存和物理内存的映射索引表），而不会拷贝物理内存。这个拷贝会消耗大量cpu资源，并且拷贝完成前会阻塞主线程，阻塞时间取决于内存中的数据量，数据量越大，则内存页表越大。拷贝完成后，父子进程使用相同的内存地址空间
	
	但主进程是可以有数据写入的，这时候就会拷贝物理内存中的数据。如下图（进程1看做是主进程，进程2看做是子进程）
	在主进程有数据写入时，而这个数据刚好在页c中，操作系统会创建这个页面的副本（页c的副本），即拷贝当前页的物理数据，将其映射到主进程中，而子进程还是使用原来的的页c
```

![](.\Static\redis-x-aof-3.png)



### 9.redis如何进行缓存预热

```
缓存预热就是系统上线后，将相关的缓存数据直接加载到缓存系统。这样就可以避免在用户请求的时候，先查询数据库，然后再将数据缓存的问题！用户直接查询事先被预热的缓存数据！
解决思路：
1.提前把数据塞入redis, 哪些是热点数据?
2.开发逻辑上也要规避差集(没缓存的数据), 会造成击穿,穿透,雪崩(加锁)

1、直接写个缓存刷新页面，上线时手工操作下；
2、数据量不大，可以在项目启动的时候自动进行加载；
3、定时刷新缓存
```

![](http://tva1.sinaimg.cn/large/006y8mN6ly1g8mg8yjuyqj30ug0u0guk.jpg)



### 10. 缓存穿透、缓存击穿(失效)、缓存雪崩



#### 缓存穿透

```
意味着有特殊请求在查询一个不存在的数据，即数据不存在 Redis 也不存在于数据库。

导致每次请求都会穿透到数据库，缓存成了摆设，对数据库产生很大压力从而影响正常服务。
```

![缓存穿透](http://s5.51cto.com/oss/202203/08/a3515296228971d7425711b9bd9fa20bf2e737.png)

##### 缓存穿透解决方案

```
1.缓存空值：当请求的数据不存在 Redis 也不存在数据库的时候，设置一个缺省值(比如：None)。当后续再次进行查询则直接返回空值或者缺省值, 缓存有效时间可以设置短点，如30秒（设置太长会导致正常情况也没法使用）。这样可以防止攻击用户反复用同一个id暴力攻击
2.布隆过滤器：在数据写入数据库的同时将这个 ID 同步到到布隆过滤器中，当请求的 id 不存在布隆过滤器中则说明该请求查询的数据一定没有在数据库中保存，就不要去数据库查询了
```

##### 布隆过滤器

```
BloomFilter 的算法是，首先分配一块内存空间做 bit 数组，数组的 bit 位初始值全部设为 0。
加入元素时，采用 k 个相互独立的 Hash 函数计算，然后将元素 Hash 映射的 K 个位置全部设置为 1。
检测 key 是否存在，仍然用这 k 个 Hash 函数计算出 k 个位置，如果位置全部为 1，则表明 key 存在，否则不存在。
哈希函数会出现碰撞，所以布隆过滤器会存在误判。
这里的误判率是指，BloomFilter 判断某个 key 存在，但它实际不存在的概率，因为它存的是 key 的 Hash 值，而非 key 的值。
所以有概率存在这样的 key，它们内容不同，但多次 Hash 后的 Hash 值都相同。
对于 BloomFilter 判断不存在的 key ，则是 100% 不存在的，反证法，如果这个 key 存在，那它每次 Hash 后对应的 Hash 值位置肯定是 1，而不会是 0。布隆过滤器判断存在不一定真的存在。
```

![布隆过滤器](http://s9.51cto.com/oss/202203/08/e1961550674b29bf300579b216ce3753332379.png)

#### 缓存击穿(失效)

```
高并发流量，访问的这个数据是热点数据，请求的数据在 DB 中存在，但是 Redis 存的那一份已经过期，后端需要从 DB 从加载数据并写到 Redis。

关键字：单一热点数据、高并发、数据失效。

但是由于高并发，同时读缓存没读到数据，又同时去数据库去取数据，引起数据库压力瞬间增大，造成过大压力, 可能会把 DB 压垮，导致服务不可用
```

![缓存击穿](http://s7.51cto.com/oss/202203/08/859c5d912b9f799ae5e451c489017358041ce5.png)

##### 缓存击穿解决方案:

```
1.使用锁
当发现缓存失效的时候，不是立即从数据库加载数据。

而是先获取分布式锁，获取锁成功才执行数据库查询和写数据到缓存的操作，获取锁失败，则说明当前有线程在执行数据库查询操作，当前线程睡眠一段时间再重试。

这样只让一个请求去数据库读取数据。

2.过期时间 + 随机值
对于热点数据，我们不设置过期时间，这样就可以把请求都放在缓存中处理，充分把 Redis 高吞吐量性能利用起来。

或者过期时间再加一个随机值。

设计缓存的过期时间时，使用公式：过期时间=baes 时间+随机时间。

即相同业务数据写缓存时，在基础过期时间之上，再加一个随机的过期时间，让数据在未来一段时间内慢慢过期，避免瞬时全部过期，对 DB 造成过大压力。

3.预热
预先把热门数据提前存入 Redis 中，并设热门数据的过期时间超大值
```



#### 缓存雪崩

```
缓存雪崩指的是大量的请求无法在 Redis 缓存系统中处理，请求全部打到数据库，导致数据库压力激增，甚至宕机。

出现该原因主要有两种：

1.大量热点数据同时过期，导致大量请求需要查询数据库并写到缓存；
2.Redis 故障宕机，缓存系统异常

缓存大量数据同时过期
数据保存在缓存系统并设置了过期时间，但是由于在同时一刻，大量数据同时过期。

系统就把请求全部打到数据库获取数据，并发量大的话就会导致数据库压力激增。

缓存雪崩是发生在大量数据同时失效的场景，而缓存击穿(失效)是在某个热点数据失效的场景，这是他们最大的区别
```

![缓存雪崩](http://s2.51cto.com/oss/202203/08/047043171e0ef79ce2f228416eeb288ba9999d.png)

##### 缓存雪崩解决方案--缓存大量数据同时过期

```
1.过期时间添加随机值
要避免给大量的数据设置一样的过期时间，过期时间 = baes 时间+ 随机时间(较小的随机数，比如随机增加 1~5 分钟)。

这样一来，就不会导致同一时刻热点数据全部失效，同时过期时间差别也不会太大，既保证了相近时间失效，又能满足业务需求。

2.接口限流
当访问的不是核心数据的时候，在查询的方法上加上接口限流保护。比如设置 10000 req/s。

如果访问的是核心数据接口，缓存不存在允许从数据库中查询并设置到缓存中。

这样的话，只有部分请求会发送到数据库，减少了压力。

限流，就是指，我们在业务系统的请求入口前端控制每秒进入系统的请求数，避免过多的请求被发送到数据库。
```

![限流](http://s3.51cto.com/oss/202203/08/436b91510f59355d8938383343887068130bd3.png)

##### 缓存雪崩解决方案--Redis 故障宕机

```
对于缓存系统故障导致的缓存雪崩的解决方案有两种：

1.服务熔断和限流
在业务系统中，针对高并发的使用服务熔断来有损提供服务从而保证系统的可用性。
服务熔断就是当从缓存获取数据发现异常，则直接返回错误数据给前端，防止所有流量打到数据库导致宕机。
服务熔断和限流属于在发生了缓存雪崩，如何降低雪崩对数据库造成的影响的方案。

2.构建高可用的缓存集群
所以，缓存系统一定要构建一套 Redis 高可用集群，如果 Redis 的主节点故障宕机了，从节点还可以切换成为主节点，继续提供缓存服务，避免了由于缓存实例宕机而导致的缓存雪崩问题。
```



### 11.redis主从复制

> 博客链接 https://segmentfault.com/a/1190000039242024
>
> https://www.51cto.com/article/700563.html
>
> https://www.cnblogs.com/kismetv/p/9236731.html



```

```









### 12.redis哨兵机制

>博客链接 https://www.cnblogs.com/kismetv/p/9609938.html



```

```



### 13.redis分片cluster

> 博客链接 https://www.cnblogs.com/kismetv/p/9853040.html



#### redis cluster工作原理



##### 数据分区方案

###### 哈希取余分区

```
	哈希取余分区思路非常简单：计算key的hash值，然后对节点数量进行取余，从而决定数据映射到哪个节点上。该方案最大的问题是，当新增或删减节点时，节点数量发生变化，系统中所有的数据都需要重新计算映射关系，引发大规模数据迁移。
```

###### 一致性哈希分区

```
	一致性哈希算法将整个哈希值空间组织成一个虚拟的圆环，如下图所示，范围为0-2^32-1；对于每个数据，根据key计算hash值，确定数据在环上的位置，然后从此位置沿环顺时针行走，找到的第一台服务器就是其应该映射到的服务器
	与哈希取余分区相比，一致性哈希分区将增减节点的影响限制在相邻节点。以下图为例，如果在node1和node2之间增加node5，则只有node2中的一部分数据会迁移到node5；如果去掉node2，则原node2中的数据只会迁移到node4中，只有node4会受影响。

一致性哈希分区的主要问题在于，当节点数量较少时，增加或删减节点，对单个节点的影响可能很大，造成数据的严重不平衡。还是以上图为例，如果去掉node2，node4中的数据由总数据的1/4左右变为1/2左右，与其他节点相比负载过高
```

> 一致性hash算法
>
> https://www.cnblogs.com/lpfuture/p/5796398.html

![](https://img2018.cnblogs.com/blog/1174710/201810/1174710-20181025213424713-1246878063.png)

###### 带虚拟节点的一致性哈希分区

```
	该方案在一致性哈希分区的基础上，引入了虚拟节点的概念。Redis集群使用的便是该方案，其中的虚拟节点称为槽（slot）。槽是介于数据和实际节点之间的虚拟概念；每个实际节点包含一定数量的槽，每个槽包含哈希值在一定范围内的数据。引入槽以后，数据的映射关系由数据hash->实际节点，变成了数据hash->槽->实际节点。

	在使用了槽的一致性哈希分区中，槽是数据管理和迁移的基本单位。槽解耦了数据和实际节点之间的关系，增加或删除节点对系统的影响很小。仍以上图为例，系统中有4个实际节点，假设为其分配16个槽(0-15)； 槽0-3位于node1，4-7位于node2，以此类推。如果此时删除node2，只需要将槽4-7重新分配即可，例如槽4-5分配给node1，槽6分配给node3，槽7分配给node4；可以看出删除node2后，数据在其他节点的分布仍然较为均衡。

	槽的数量一般远小于2^32，远大于实际节点的数量；在Redis集群中，槽的数量为16384(2^14)。

	下面这张图很好的总结了Redis集群将数据映射到实际节点的过程
	
	Redis对数据的特征值（一般是key）计算哈希值，使用的算法是CRC16, slot = CRC16(key) & 16384

	根据哈希值，计算数据属于哪个槽。

	根据槽与节点的映射关系，计算数据属于哪个节点。
```



![](https://img2018.cnblogs.com/blog/1174710/201908/1174710-20190802191100758-1110624103.png)



##### 节点通信机制

###### 两个端口

```
在哨兵系统中，节点分为数据节点和哨兵节点：前者存储数据，后者实现额外的控制功能。在集群中，没有数据节点与非数据节点之分：所有的节点都存储数据，也都参与集群状态的维护。为此，集群中的每个节点，都提供了两个TCP端口：

	1)普通端口：即我们在前面指定的端口(7000等)。普通端口主要用于为客户端提供服务（与单机节点类似）；但在节点间数据迁移时也会使用。
	2)集群端口：端口号是普通端口+10000（10000是固定值，无法改变），如7000节点的集群端口为17000。集群端口只用于节点之间的通信，如搭建集群、增减节点、故障转移等操作时节点间的通信；不要使用客户端连接集群接口。为了保证集群可以正常工作，在配置防火墙时，要同时开启普通端口和集群端口。
```

###### Gossip协议

```
节点间通信，按照通信协议可以分为几种类型：单对单、广播、Gossip协议等。重点是广播和Gossip的对比。

	广播是指向集群内所有节点发送消息；优点是集群的收敛速度快(集群收敛是指集群内所有节点获得的集群信息是一致的)，缺点是每条消息都要发送给所有节点，CPU、带宽等消耗较大。

	Gossip协议的特点是：在节点数量有限的网络中，每个节点都“随机”的与部分节点通信（并不是真正的随机，而是根据特定的规则选择通信的节点），经过一番杂乱无章的通信，每个节点的状态很快会达到一致。Gossip协议的优点有负载(比广播)低、去中心化、容错性高(因为通信有冗余)等；缺点主要是集群的收敛速度慢
```

###### Gossip协议消息类型

| 消息类型 | **消息内容**                                                 |
| :------: | ------------------------------------------------------------ |
|   MEET   | 在节点握手阶段，当节点收到客户端的CLUSTER MEET命令时，会向新加入的节点发送MEET消息，请求新节点加入到当前集群；新节点收到MEET消息后会回复一个PONG消息 |
|   PING   | 每隔一秒钟，选择5个最久没有通信的节点，发送PING消息，检测对应的节点是否在线；同时还有一种策略是，如果某个节点的通信延迟大于了`cluster-node-time`的值的一半，就会立即给该节点发送PING消息，避免数据交换延迟过久 |
|   PONG   | 当节点接收到MEET或者PING消息之后，会回一个PONG消息给发送方，代表自己收到了MEET或者PING消息。同时，节点也可以主动的通过PONG消息向集群中广播自己的信息，让其他节点获取到自己最新的属性，就像完成了故障转移之后新的master向集群发送PONG消息一样 |
|   FAIL   | 用于广播自己的对某个节点的宕机判断，假设当前节点对A节点判断为宕机，就会立即向Redis Cluster广播自己对于A节点的判断，所有收到消息的节点就会对A节点做标记 |
| PUBLISH  | 用于向指定的Channel发送消息，某个节点收到PUBLISH消息之后会直接在集群内广播，这样一来，客户端无论连接到任何节点都能够订阅这个Channel |

###### 使用Gossip的优劣

| 优点       | 描述                                                         |
| ---------- | ------------------------------------------------------------ |
| 扩展性     | 网络可以允许节点的任意增加和减少，新增加的节点的状态最终会与其他节点一致 |
| 容错性     | 由于每个节点都持有一份完整元数据，所以任何节点宕机都不会影响gossip的运行 |
| 健壮性     | 与容错性类似，由于所有节点都持有数据，地位平台，是一个去中心化的设计，任何节点都不会影响到服务的运行 |
| 最终一致性 | 当有新的信息需要传递时，消息可以快速的发送到所有的节点，让所有的节点都拥有最新的数据 |

| 缺点                                                         |
| :----------------------------------------------------------- |
| gossip仍然存在一些缺点。例如消息可能最终会经过很多轮才能到达目标节点，而这可能会带来较大的延迟。同时由 |



##### 数据结构

```
节点需要专门的数据结构来存储集群的状态。所谓集群的状态，是一个比较大的概念，包括：集群是否处于上线状态、集群中有哪些节点、节点是否可达、节点的主从状态、槽的分布……

节点为了存储集群状态而提供的数据结构中，最关键的是clusterNode和clusterState结构：前者记录了一个节点的状态，后者记录了集群作为一个整体的状态。
```



###### clusterNode

```
clusterNode结构保存了一个节点的当前状态，包括创建时间、节点id、ip和端口号等。每个节点都会用一个clusterNode结构记录自己的状态，并为集群内所有其他节点都创建一个clusterNode结构来记录节点状态。
```

```C
typedef struct clusterNode {
    mstime_t ctime; /* 创建节点的时间 */
    char name[CLUSTER_NAMELEN]; /* 节点的名字 */
    int flags;      /* 节点标识，标记节点角色或者状态，比如主节点从节点或者在线和下线 */
    uint64_t configEpoch; /* 当前节点已知的集群统一epoch */
    unsigned char slots[CLUSTER_SLOTS/8]; /* slots handled by this node */
    int numslots;   /* Number of slots handled by this node */
    int numslaves;  /* Number of slave nodes, if this is a master */
    struct clusterNode **slaves; /* pointers to slave nodes */
    struct clusterNode *slaveof; /* pointer to the master node. Note that it
                                    may be NULL even if the node is a slave
                                    if we don't have the master node in our
                                    tables. */
    mstime_t ping_sent;      /* 当前节点最后一次向该节点发送 PING 消息的时间 */
    mstime_t pong_received;  /* 当前节点最后一次收到该节点 PONG 消息的时间 */
    mstime_t fail_time;      /* FAIL 标志位被设置的时间 */
    mstime_t voted_time;     /* Last time we voted for a slave of this master */
    mstime_t repl_offset_time;  /* Unix time we received offset for this node */
    mstime_t orphaned_time;     /* Starting time of orphaned master condition */
    long long repl_offset;      /* 当前节点的repl便宜 */
    char ip[NET_IP_STR_LEN];  /* 节点的IP 地址 */
    int port;                   /* 端口 */
    int cport;                  /* 通信端口，一般是端口+1000 */
    clusterLink *link;          /* 和该节点的 tcp 连接 */
    list *fail_reports;         /* 下线记录列表 */
} clusterNode;
```



###### clusterState

```C
typedef struct clusterState {
 
    //自身节点
    clusterNode *myself;
 
    //配置纪元
    uint64_t currentEpoch;
 
    //集群状态：在线还是下线
    int state;
 
    //集群中至少包含一个槽的节点数量
    int size;
 
    //哈希表，节点名称->clusterNode节点指针
    dict *nodes;
  
    //槽分布信息：数组的每个元素都是一个指向clusterNode结构的指针；如果槽还没有分配给任何节点，则为NULL
    clusterNode *slots[16384];
 
    …………
     
} clusterState;
```



#####  集群命令的实现

这一部分将以cluster meet(节点握手)、cluster addslots(槽分配)为例，说明节点是如何利用上述数据结构和通信机制实现集群命令的

###### cluster meet

```
假设要向A节点发送cluster meet命令，将B节点加入到A所在的集群，则A节点收到命令后，执行的操作如下：

1)  A为B创建一个clusterNode结构，并将其添加到clusterState的nodes字典中

2)  A向B发送MEET消息

3)  B收到MEET消息后，会为A创建一个clusterNode结构，并将其添加到clusterState的nodes字典中

4)  B回复A一个PONG消息

5)  A收到B的PONG消息后，便知道B已经成功接收自己的MEET消息

6)  然后，A向B返回一个PING消息

7)  B收到A的PING消息后，便知道A已经成功接收自己的PONG消息，握手完成

8)  之后，A通过Gossip协议将B的信息广播给集群内其他节点，其他节点也会与B握手；一段时间后，集群收敛，B成为集群内的一个普通节点

通过上述过程可以发现，集群中两个节点的握手过程与TCP类似，都是三次握手：A向B发送MEET；B向A发送PONG；A向B发送PING
```



###### cluster addslots

```
集群中槽的分配信息，存储在clusterNode的slots数组和clusterState的slots数组中，两个数组的结构前面已做介绍；二者的区别在于：前者存储的是该节点中分配了哪些槽，后者存储的是集群中所有槽分别分布在哪个节点。

cluster addslots命令接收一个槽或多个槽作为参数，例如在A节点上执行cluster addslots {0..10}命令，是将编号为0-10的槽分配给A节点，具体执行过程如下：

1)  遍历输入槽，检查它们是否都没有分配，如果有一个槽已分配，命令执行失败；方法是检查输入槽在clusterState.slots[]中对应的值是否为NULL。

2)  遍历输入槽，将其分配给节点A；方法是修改clusterNode.slots[]中对应的比特为1，以及clusterState.slots[]中对应的指针指向A节点

3)  A节点执行完成后，通过节点通信机制通知其他节点，所有节点都会知道0-10的槽分配给了A节点
```











### 16.redis事务

> redis单条命令保持原子性, 但是事务不保证原子性
>
> redis事务没有隔离级别的概念, 所有命令在事务中, 并没有直接被执行, 只有发起执行命令的时候才会执行
>
> * 开启事务( multi )
>
> * 命令入队( ... )
>
> * 执行事务( exec ) & 取消事务( discard)
>
>   ```
>   编译型异常: 命令错误, 事务中所有的命令都不会执行
>   运行时异常: 如果事务队列中存在语法错误, 那么执行命令的时候, 其他命令是可以执行的
>   ```
>
> * 乐观锁( watch ) & 取消锁( unwatch )
>
>   ```bash
>   开始事务前, 使用watch可以作为乐观锁
>   如果watch 的key在exec之前被修改了, 那么这次事务就会失败
>   当exec被调用时, 不管事务是否成功执行, 对所有key的监视都会被取消
>   ```
>
>   



```bash

```

