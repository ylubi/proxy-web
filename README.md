# proxyWebApplication
proxyWebApplication是以[snail007/goproxy](https://github.com/snail007/goproxy/)为基础服务，完成的网页可视化应用。

---
[![stable](https://img.shields.io/badge/stable-stable-green.svg)](https://github.com/snail007/goproxy/)

### 使用前须知
 - [下载](#下载)
 - [目录位置](#目录位置)
 - [依赖包](#依赖包)
 
### 手册目录
 - [1. 运行](#运行)
 - [2. 参数介绍](#参数介绍)
     - [2.1 http参数](#http参数)
     - [2.2 tcp参数](#tcp参数)
     - [2.3 udp参数](#udp参数)
     - [2.4 socks参数](#socks参数)
     - [2.5 tclient参数](#tclient参数)
     - [2.6 tserver参数](#tserver参数)
     - [2.6 tbridge参数](#tbridge参数)
 - [3. TODO](#TODO)
 
### 下载
cd 进入GOPATH
使用git下载源码
编译go build进行编译
用命令go run main.go
然后用8080端口在网页进入，如localhost:8080

### 目录位置
下载[snail007/goproxy](https://github.com/snail007/goproxy/releases)
在与proxyWebApplication平级的目录下建proxyService目录
压缩包解压
加密文件默认在proxyService/.cert/目录里
也可在config里修改目录路径

### 依赖包
[github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)使用sqlite作为数据库
[github.com/Unknwon/goconfig](https://github.com/Unknwon/goconfig)解析配置文件

