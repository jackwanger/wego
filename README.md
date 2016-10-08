# hope_word

![travis](https://travis-ci.org/repong/hope_word.svg?branch=master)

基于sego的屏蔽字过滤器

### 用法

``` bash
go get -u github.com/repong/hope_word
hope_word -port ":3000" -files "dict1.txt,dict2.txt"
```

### 例子

``` bash
curl -XPOST http://localhost:8000/validate -d "message=测试看看"                                                                                                                                                                     19:17:30
{"result":"false"}

curl -XPOST http://localhost:8000/filter -d "message=测试看看"                                                                                                                                                                       19:17:35
{"result":"**看看"}
```

### 敏感词来源

* 和谐宝典 https://github.com/wear/harmonious_dictionary