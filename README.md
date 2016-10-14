# wego

![travis](https://travis-ci.org/repong/wego.svg?branch=master)

基于sego的屏蔽字过滤器

### 用法

1. 从 https://github.com/repong/wego/releases 下载对应自己系统的版本，目前仅编译了linux和macOS版本
2. `tar zxf wego-osx-1.0.0.tar.gz`
3. `./wego` (默认8000端口，可以使用`./wego -port 1234`修改监听端口)

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

* 字典来源

  https://github.com/wear/harmonious_dictionary

* 如何更新

  ``` bash
  make update_dict
  ```

* 字典格式

  一般只需要如下配置，其中，词频是一个数字，数字大的会优先匹配，必须大于1

  ```
  屏蔽字1 词频
  屏蔽字2 词频
  ```
