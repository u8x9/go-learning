# 并发编程

## Thread vs Groutine

### 创建时默认的 stack 的大小

+ JDK5 以后，Java Thread stack 默认为1M

+ Groutine 的 stack 初始化大小为2K

### 和KSE(Kernel Space Entity) 的对应关系

+ Java Thread 是 1:1
+ Groutine 是 M:N

![image-20201118153647494](image-20201118153647494.png)

## CSP 并发机制

[`Communicating Sequential Processes`](https://en.wikipedia.org/wiki/Communicating_sequential_processes)

### CSP vs Actor

+ 和  Actor 的直接通讯不同，CSP模式是通过 Channel 进行通讯的，更加松耦合

+ Go 中的 channel 是有容量限制并且独立于处理 Groutine，而如 Erlang，Actor模式中的mailbox容量是无限的，接收进程也总是被动地处理消息

![image-20201118162345387](image-20201118162345387.png)

## channel 的关闭

+ 向关闭的 channel 发送数据，会导致panic

+ `v, ok := <-ch`，`ok` 为 bool 值，true 表示正常接收，false 表示 channel 已关闭

+ 所有的 channel 接收者都会在 channel 关闭时，立刻从阻塞等待中返回且上述 `ok` 值为 false。这个广播机制常被用来向多个订阅者同时发送信号。如：退出信号。

## Context

+ 根 Context：通过 `context.Background()` 创建

+ 子 Context: `context.WithCancel(parentContext)` 创建
	
	+ `ctx, cancel := context.WithCancel(context.Background())`
	
+ 当前 Context 被取消时，基于它的子 context 都会被取消

+ 接收取消通知 `<-ctx.Done()`
