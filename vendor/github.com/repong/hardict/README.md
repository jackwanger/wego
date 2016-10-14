hardict
===

基于sego的屏蔽字检索工具，自带同步和谐字典功能

### 字典

* 来源

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

### 用法

* 检查是否包含屏蔽字

  ``` go
  hardict.ExistInvalidWord(text)
  ```

* 屏蔽字替换为*号

  ``` go
  hardict.ReplaceInvalidWords(text)
  ```
