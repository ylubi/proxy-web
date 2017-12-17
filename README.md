# proxy-web
proxy-web是以[snail007/goproxy](https://github.com/snail007/goproxy/)为基础服务，完成的网页可视化应用。

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
     - [2.5 client参数](#25client参数)
     - [2.6 server参数](#26server参数)
     - [2.7 bridge参数](#27bridge参数)
 
### 下载
[下载地址](https://github.com/yincongcyincong/proxy-web/releases)  

### 目录位置
下载[snail007/goproxy](https://github.com/snail007/goproxy/releases)  
在与proxy-web平级的目录下建proxyService目录  
压缩包解压  
加密文件默认在proxyService/.cert/目录里  
也可在config里修改目录路径  

### 依赖包
[github.com/boltdb/bolt](https://github.com/boltdb/bolt)使用bolt扩展为数据库  
[github.com/Unknwon/goconfig](https://github.com/Unknwon/goconfig)解析配置文件  
[github.com/astaxie/beego/tree/master/session](https://github.com/astaxie/beego/tree/master/session) session模块 

### 1.运行
然后用28080端口在浏览器进入，如localhost:28080  
首先要登录，可在config里面配置账号密码，默认都为admin  
config里也可以修改端口  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/login.png?raw=true" />  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/preview.png?raw=true" />  
  
### 2.参数介绍
代理协议：需要用到的协议 如http， tcp等协议。  
本地连接类型：-t参数。  
链式代理：此次连接的类型，顶级代理不需要“上级服务器+端口”。  
代理服务器+端口：-p参数。  
上级服务器+端口：-P参数。  
父级连接类型：-T参数 选取后可能会有不同的加密方式，上传文件的加密方式会有默认文件，tcp形式默认不加密。 

#### **2.1.http参数** 
tls形式加密：-C .crt文件 和 -K参数 .key文件  
ssh形式加密：有密钥和密码两种方式，-u 用户名 -A 密码 -S .key文件 -s 密钥密码 
kcp形式加密：-B 密码  
<img src="https://github.com/yincongcyincong/proxy-web/blob/master/docs/image/http1.png?raw=true" />  
`path to proxy/proxy http -t tcp -p :8081`  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/http2.png?raw=true" />  
`path to proxy/proxy http -t tls -p :8081 -T tls -P 2.2.2.2:8081 -C path to file/proxy.crt -K path to file/proxy.key`  

#### **2.2.tcp参数** 
tls形式加密：-C .crt文件 和 -K参数 .key文件  
kcp形式加密：-B密码  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/tcp1.png?raw=true" />  
`path to proxy/proxy tcp -t tls -p :8081 -T tls -P 2.2.2.2:8081 -C path to file/proxy.crt -K path to file/proxy.key`  

#### **2.3.udp参数** 
没有加密模式  
“本次连接类型”只有udp模式  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/udp1.png?raw=true" />  
`path to proxy/proxy udp -p :8081 -T tls -P 2.2.2.2:8081 -C path to file/proxy.crt -K path to file/proxy.key`

#### **2.4.socks参数** 
tls形式加密：-C .crt文件 和 -K参数 .key文件  
ssh形式加密：有密钥和密码两种方式，-u 用户名 -A 密码 -S .key文件 -s 密钥密码   
kcp形式加密：-B 密码  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/socks.png?raw=true" />  
`path to proxy/proxy socks -t tcp -p :8081 -T kcp -P 2.2.2.2:8081 -B 1234 `

#### **2.5.client参数** 
只有tls形式的加密且必须加密  
tls形式加密：-C .crt文件 和 -K参数 .key文件 
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/client.png?raw=true" />  
`path to proxy/proxy client -P ":8081" -C path to file/proxy.crt -K path to file/proxy.key `  
“上级服务器+端口”填写的内容无效

#### **2.6.server参数** 
只有tls形式的加密且必须加密  
tls形式加密：-C .crt文件 和 -K参数 .key文件  
“代理服务器+端口”填写-r参数  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/server.png?raw=true" />  
`path to proxy/proxy server -r "udp://:10053@:53" -P "2.2.2.2:8081" -C path to file/proxy.crt -K path to file/proxy.key`

#### **2.7.bridge参数** 
只有tls形式的加密且必须加密  
tls形式加密：-C .crt文件 和 -K参数 .key文件  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/bridge.png?raw=true" />  
`path to proxy/proxy bridge -P ":8081" -C path to file/proxy.crt -K path to file/proxy.key `  
“上级服务器+端口”填写的内容无效  

### 源码使用  
- git下载源码  
- 配置文件在config,可以修改路径、端口和登录账号密码   
   
### TODO
- -L参数进程池  
- server -r参数分解  

### License
- under GPLv3 license  

### Contact
- QQ群：189618940
