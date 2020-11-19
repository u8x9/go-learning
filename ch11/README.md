# 测试

## 单元测试框架

+ `Fail`, `Error`：该测试失败，该测试继续，其他测试继续执行

+ `FailNow`, `Fatal`：该测试失败，该测试终止，其他测试继续执行

## 内置单元测试框架

### 代码覆盖率

```bash
go test -v -cover
```

### 断言

[`github.com/stretchr/testify`](https://github.com/stretchr/testify)