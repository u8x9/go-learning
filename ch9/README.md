# 并发编程

## Thread vs Groutine

### 创建时默认的 stack 的大小

+ JDK5 以后，Java Thread stack 默认为1M

+ Groutine 的 stack 初始化大小为2K

### 和KSE(Kernel Space Entity) 的对应关系

+ Java Thread 是 1:1
+ Groutine 是 M:N

![image-20201118153647494](image-20201118153647494.png)