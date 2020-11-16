# 字符串

## 与其它主要编程语言的差异：

1. `string` 是数据类型，不是引用或指针类型

2. `string` 是只读的 `byte slice`，`len`函数可以求出它所包含的 byte 数

3.  `string` 的 byte 数组可以存放任何数据

## Unicode 和 UTF-8

+ Unicode 是一种字符集(code point)

+ UTF-8 是 Unicode 的存储实现(转换为字节序列的规则)

### 编码与存储

|字符|`“中”`|
|----|----|
|Unicode|`0x4E2D`|
|Utf-8|`0xE4B8AD`|
|`string/[]byte`|`[0xE4, 0xB8, 0xAD]`|

## 常用字符串函数

### `strings` 包

### `strconv` 包
