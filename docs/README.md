# Go 入门指南

---

## 关键特性

* 开发快
* 运行快
* 高并发
* 易学习

---

## 开发快

* 强大的库支持
* 编译快（模块机制）
* 易分析、易调试（静态类型）
* 易布署（静态链接，最小化依赖）


## 运行快

* 直接编译为本地代码，类似C/C++
	* （ELF on Unix/PE on Windows）

## 高并发

* 语法级别的 协程(goroutine)/管道(channel) 支持

---

## 易学习

* 自带GC（无内存泄漏）
* 没有class
	* 没有继承、多态(OOP, Object-Orentied Programming)
	* 有struct
		* 支持数据抽象
		* 支持封装(encapsulation)
			* 隐藏部分字段
	* 有method (Object-Based)
* 没有template，没有泛型编程（GP, Generic Programming）
	* 有interface

---

## 内容概要

* 变量
* 基本类型
* 复合类型
* 控制结构
* 函数 func
* 结构 struct
* 接口 interface
* 协程 goroutine
* 管道 channel

---

## 变量 = 名称 + 类型 [+ 初始值]

```go
var i int     // declare but not assigned, 0 by default
var j int = 1 // declare and assigned

// infer type by value
var isprime = false // bool

// declare more
var a, b, c int = 3, 4, 5

// declare and assigned
k := 2
x, y := 320, 240
```

---

## 关键字

（不能用于变量名）

![keywords](./images/keywords.png)

---

## 预定义的名称

内置常量、类型、函数
（不能用于变量名）

![predeclared names](./images/prenames.png)

---


## 基本类型

* 数值类型
* 布尔类型
* 字符串
* 常量
* 指针

---

## 基本类型

* 数值类型
	* 整型
		* `int`, `int8`, `int16`, `int32`, `int64`
		* `uint`, `uint8`, `uint16`, `uint32`, `uint64`
	* 浮点型
		* `float32`
		* `float64`
	* 复数 = 实部（实数）+ 虚部（虚数）
		* 虚数单位 `i`, `j`
		* `complex64` (两个`float32`)
		* `complex128` (两个`float64`)
		* `complex` (同上)

---

* 布尔类型
	* `bool`, `true`, `false`
* 字符串
	* `string`
	* UTF-8
* 常量
	* `const`
* 指针
	* 没有解引用操作符（`->`）
	* 也用成员操作符（`.`），自动解引用

---

## 基本类型

```go
	var number uint32 = 0x12345678
	var number uint64 = 0x1234567812345678

	var r float32 = 5.0
	var pi float64 = 3.141592653525
	c1 := 3 + 4i // complex128

	var name string = "Brian W. Kernighan"
```

---

## 复合类型


