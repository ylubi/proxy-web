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
  代理协议：需要用到的协议 如http， tcp等协议。  
  本次链接类型：-t参数。  
  链式代理：区分此次链接类型，顶级代理不需要“上级服务器+端口”。  
  代理服务器+端口：-p参数。  
  上级服务器+端口：-P参数。  
  父级连接类型：-T参数 选取后可能会有不同的加密方式，上传文件的加密方式会有默认文件，tcp形式默认不加密。  
#### **2.1.http参数** 
tls形式加密：-C .crt文件 和 -K参数 .key文件   
ssh形式加密：有密钥和密码两种方式，-u用户名 -A密码 -S .key文件
kcp形式加密：-B密码  
#### **2.2.tcp参数** 
tls形式加密：-C .crt文件 和 -K参数 .key文件    
kcp形式加密：-B密码 
#### **2.3.udp参数**  
没有加密模式  
本次链接类型只有udp模式  
#### **2.4.socks参数**  
tls形式加密：-C .crt文件 和 -K参数 .key文件   
ssh形式加密：有密钥和密码两种方式，-u用户名 -A密码 -S .key文件
kcp形式加密：-B密码   
#### **2.5.tclient参数**  
只有tls形式的机密且必须加密  
tls形式加密：-C .crt文件 和 -K参数 .key文件   
#### **2.6.tserver参数**  
只有tls形式的机密且必须加密  
tls形式加密：-C .crt文件 和 -K参数 .key文件   
#### **2.7.tbridge参数**  
只有tls形式的机密且必须加密  
tls形式加密：-C .crt文件 和 -K参数 .key文件   
  
### TODO
- -L参数进程池  
- tserver -r参数分解  

### License
- under GPLv3 license
   
