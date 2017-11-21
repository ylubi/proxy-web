# proxyWebApplication
proxyWebApplication是以[snail007/goproxy](https://github.com/snail007/goproxy/)为基础服务，完成的网页可视化应用。

---
[![stable](https://img.shields.io/badge/stable-stable-green.svg)](https://github.com/snail007/goproxy/)

### 使用前须知
 - [下载](#下载)
 - [目录位置](#目录位置)
 - [依赖包](#依赖包)
 
### 手册目录
 - [1. 运行](#1运行)
 - [2. 参数介绍](#2参数介绍)
     - [2.1 http参数](#21http参数)
     - [2.2 tcp参数](#22tcp参数)
     - [2.3 udp参数](#23udp参数)
     - [2.4 socks参数](#24socks参数)
     - [2.5 tclient参数](#25tclient参数)
     - [2.6 tserver参数](#26tserver参数)
     - [2.7 tbridge参数](#27tbridge参数)
 - [3. TODO](#3TODO)
 - [4. License](#4License)
 
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

### 1.运行
<img src="https://github.com/yincongcyincong/proxyWebApplication/raw/master/docs-images/preview.png?raw=true" width="200"/> 

### 2.参数介绍  
#### **2.1.http参数** 
#### **2.2.tcp参数** 
#### **2.3.udp参数**  
#### **2.4.socks参数**  
#### **2.5.tclient参数**  
#### **2.6.tserver参数**  
#### **2.7.tbridge参数**  
  
### 3.TODO
- -L参数进程池  
- tserver -r参数分解  

### 4.License
- under GPLv3 license
   
