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

清楚当前数据库 `flushdb`

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
3. redis淘汰策略
4. redis删除策略
5. redis缓存一致性
6. redis核心对象
7. redis数据类型
8. redis持久化
9. redis架构协议
10. redis实现分布式锁
11. redis实现异步队列
12. 缓存穿透、缓存雪崩
13. redis事务
14. 多线程优化







### 为什么redis是单线程还这么快

```bash
答:
	1. (核心) redis是将所有的数据全部放在内存中, 读写都是在内存中完成的;
	2. k-v型数据库, 内部构建一个哈希表(hashmap), 查找和操作的时间复杂度都是O(1);
	3. 采用单线程, 没有了多线程的上下文切换的性能消耗, 没有了访问共享资源加锁的性能消耗;
	4. 采用IO多路复用技术, 非阻塞IO;
```



### redis数据类型

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



> **String**

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



> List

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

> List小结

```
实际上是一个链表
如果key不存在, 创建新的链表
如果key存在, 新增内容
如果移除了所有值, 空链表, 也代表不存在
在两边插入或者改动值效率最高
```



> Hash 

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





> Set

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





> Zset
>
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



>Geospatial

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



> Hyperloglog
>
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



> Bitmaps
>
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





### redis事务

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

