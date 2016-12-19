# wego

![travis](https://travis-ci.org/repong/wego.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/repong/wego)](https://goreportcard.com/report/github.com/repong/wego)

屏蔽字过滤服务

### 用法

1. 从 https://github.com/repong/wego/releases 下载对应自己系统的版本，目前仅编译了linux和macOS版本
2. `tar zxf wego-osx-1.0.0.tar.gz`
3. `./wego` (默认8000端口，可以使用`./wego -port 1234`修改监听端口)

运行情况如下：

``` bash
$ ./wego -dict /tmp/dict1.txt,/tmp/dict2.txt
Version    : 1.1.0-1-g4345260
Git Hash   : a0bb954eeb2c277498b03da247bfe28625a5a2a9
Build Time : 2016-12-14T11:30:27Z
载入sego词典 /tmp/dict1.txt
载入sego词典 /tmp/dict2.txt
sego词典载入完毕
Listening at 8000
[GIN] 2016/12/15 - 15:45:46 | 200 |     148.311µs | 127.0.0.1 |   POST    /filter
[GIN] 2016/12/15 - 15:46:18 | 200 |      93.437µs | 127.0.0.1 |   GET    /validate
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

* ~~https://github.com/repong/hardict 封装了更新字典及检测屏蔽字的方法~~
* 使用用户自定义字典，每行一个文本

### Todo

* [x] http
* [ ] gRPC
