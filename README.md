# hope_word

![travis](https://travis-ci.org/repong/hope_word.svg?branch=master)

基于sego的屏蔽字过滤器

### 用法

``` bash
go get -u github.com/repong/hope_word
hope_word -port 8000
```

### 例子

``` bash
curl -XPOST http://localhost:8000/validate -d "message=测试看看"
{"result":"false"}

curl -XPOST http://localhost:8000/filter -d "message=测试看看"
{"result":"**看看"}
```

### 更新字典

``` bash
make update_dict
```