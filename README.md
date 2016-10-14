# wego

![travis](https://travis-ci.org/repong/wego.svg?branch=master)

屏蔽字过滤服务

### 用法

1. 从 https://github.com/repong/wego/releases 下载对应自己系统的版本，目前仅编译了linux和macOS版本
2. `tar zxf wego-osx-1.0.0.tar.gz`
3. `./wego` (默认8000端口，可以使用`./wego -port 1234`修改监听端口)

运行情况如下：

``` bash
Loading dict...
载入sego词典 /var/folders/yy/ksdhbgf551bgg_5p0dkyd5kr0000gn/T/chinese_dictionary.txt844575244
载入sego词典 /var/folders/yy/ksdhbgf551bgg_5p0dkyd5kr0000gn/T/english_dictionary.txt811498491
sego词典载入完毕
Version    : 1.0.0-5-g8c1e0b9
Git Hash   : 8c1e0b943ea21f72e4eca8adc7a931ebb287da8f
Build Time : 2016-10-14T07:45:22Z
[GIN] 2016/10/14 - 15:45:46 | 200 |     148.311µs | 127.0.0.1 |   POST    /filter
[GIN] 2016/10/14 - 15:46:18 | 200 |      93.437µs | 127.0.0.1 |   POST    /validate
```

### 客户端如何调用？

1. 验证是否包含屏蔽字

  ``` bash
  curl -XPOST http://localhost:8000/validate -d "message=测试封杀"
  {"result":"false"}
  ```

2. 过滤掉屏蔽字，以*号代替

  ``` bash
  curl -XPOST http://localhost:8000/filter -d "message=测试封杀"
  {"result":"测试**"}
  ```

### 字典

* https://github.com/repong/hardict 封装了更新字典及检测屏蔽字的方法

### Todo

* [x] http
* [ ] gRPC
