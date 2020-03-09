a little store

##### 1. 思路
主要是用Patricia Trie对所有的 k/v进行索引，Patricia Trie有很好的压缩效果，设想把它
完全放到内存里面, 索引里有val对应的offset。

##### 2. 使用

```bash
go build -i -o bin/gendb cmd/GEN/main.go
go build -i -o bin/cli cmd/CLI/main.go
go build -i -o bin/server cmd/SERVER/main.go
```
```bash
./bin/gendb >> gen.log # 可以生成测试数据
./bin/server db1.data # 启动数据库
./bin/cli 127.0.0.1:9527 KeyString
```

##### 3. 测试及总结
######预处理时间
  - 本地,100M数据,生成索引需要8秒；
  - 本地,1G数据,生成索引需要90秒
  - 腾讯云,100G数据,好久,已经不想测了
  
生成数据后面改成用bufio的方式，应该可以加速
######占用的内存
当key比较小的时候，< 512B，这个索引占用的内存竟然跟原来磁盘文件差不多大小
应该是由于我用了一个uint64数字的当作磁盘指针，索引占用的空间比原来的更大，
后续可以改成不同大小的指针使用不用不同类型；还有就是用指针的方式实现Patricia，
指针大小64位操作系统也占用64位，后面这改成bitmap+array的方式；当key>100K,
压缩效果就好多了。

要索引1T，小key数据的话，应该还是要将部分索引放到磁盘上，内存中是稀疏索引，
将数据分块储存，块内有序，内存只要存块的指针。

一周的期限到了，刚开始还是把这个想简单了，最后的测试结果完全达不到索引1T的目的了。