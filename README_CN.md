# shenyu-client-golang

English | [简体中文](README_CN.md)

[![Build and Test](https://github.com/apache/shenyu-client-golang/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/apache/shenyu-client-golang/actions)
[![codecov.io](https://codecov.io/gh/apache/shenyu-client-golang/coverage.svg?branch=main)](https://app.codecov.io/gh/apache/shenyu-client-golang?branch=main)
[![GoDoc](https://godoc.org/github.com/apache/shenyu-client-golang?status.svg)](https://godoc.org/github.com/apache/shenyu-client-golang)

---

## Shenyu-client-golang
Shenyu-client-golang是提供了Go语言访问ShenYu网关的功能，并支持服务注册到ShenYu网关。

---
## 已支持注册到ShenYu网关的方式
* **以Http方式注册**
* **以Nacos方式注册**
* **以Zookeeper方式注册**
* **以Consul方式注册**

---

## 要求

要求Go语言版本 **1.13**

SDK支持ShenYu的版本 **2.4.3及以上**

## 安装方法

使用 `go get命令` 安装 SDK：

```sh
$ go get -u github.com/apache/shenyu-client-golang
```

## 代码列子路径

* shenyu-client-golang/example/**_client/main.go
---

## 开始

* Http 示例  [简体中文](doc/HTTP_CN.md) | [English](doc/HTTP_EN.md)  
* Nacos 示例 [简体中文](doc/NACOS_CN.md) | [English](doc/NACOS_EN.md)
* Zookeeper 示例 [简体中文](doc/ZK_CN.md) | [English](doc/ZK_EN.md)
* Consul 示例  [简体中文](doc/CONSUL_CN.md) | [English](doc/CONSUL_EN.md)

---

