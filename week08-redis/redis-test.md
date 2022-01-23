# week08作业：redis测试

### **作业内容：**

1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

   

   

### 第一部分：

10字节：

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220123233740294.png)

20字节：

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220123233808713.png)

50字节：

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220123233840317.png)

100字节：

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220123233912902.png)



200字节：

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220123233943950.png)



1k字节：

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220123234014847.png)

5k字节：

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220123234048568.png)





### 第二部分：

测试前memory大小 5.89M：

[root@hserver1 redis]# redis-cli -p 6379 info memory

used_memory:6194120
used_memory_human:5.91M
used_memory_rss:14024704
used_memory_rss_human:13.38M
used_memory_peak:155196112
used_memory_peak_human:148.01M
total_system_memory:8254799872
total_system_memory_human:7.69G
used_memory_lua:37888
used_memory_lua_human:37.00K
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
mem_fragmentation_ratio:2.26
mem_allocator:jemalloc-4.0.3

执行10字节30万kv写入，单个key平均76字节：

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220124001146259.png)

执行20字节30万k v写入，单个ke y平均65字节：

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220124001357260.png)

执行50字节30万k v写入，单个key平均84字节：

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220124001543046.png)

执行100字节30万k v写入，单个ke y平均100字节：

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220124001726741.png)

执行200字节30万k v写入，单个key平均171字节：

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220124001828236.png)

执行1k字节30万k v写入，单个key平均818字节：

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220124002020511.png)

执行5k字节30万k v写入，单个key平均4,144字节：

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220124002134572.png)

![image](https://github.com/xiehengkang/golang-training-camp/blob/xiehengkang/week08-redis/img/image-20220124002149284.png)





### 总结：

在key的字节长度递增的情况下，key字节越长，平均单个key占用的字节越多，在redis缓存key设计时，使用的key名称越短，越能节省内存资源。