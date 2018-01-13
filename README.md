# proxy-web详细介绍
proxy-web是用go语言写的，基于[snail007/goproxy](https://github.com/snail007/goproxy/)完成的可视化网页应用

---
[![stable](https://img.shields.io/badge/stable-stable-green.svg)](https://github.com/snail007/goproxy/)

### 使用前须知
 - [作用](#作用)
 - [下载](#下载)
 - [更新](#更新)
 - [配置](#配置)
 - [依赖包](#依赖包)
 
### 手册目录
- [1. 使用](#1使用)
- [2. 参数介绍](#2参数介绍)
     - [2.1 http参数](#21http参数)
     - [2.2 tcp参数](#22tcp参数)
     - [2.3 udp参数](#23udp参数)
     - [2.4 socks参数](#24socks参数)
     - [2.5 client参数](#25client参数)
     - [2.6 server参数](#26server参数)
     - [2.7 bridge参数](#27bridge参数)

### 作用
1、	用web界面的方式使用goproxy，更加方便  
2、	监控goproxy运行情况  
3、	实时显示goproxy产生的日志  
4、	能自启动goproxy  
 
### 下载
[下载地址](https://github.com/yincongcyincong/proxy-web/releases)  

### 更新
1、 可以在linux下自动生成证书和key文件  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/keygen.png?raw=true" />  
2、 支持--c参数压缩  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/compress.png?raw=true" />  
3、--always参数，使下级代理流量全部使用上级代理  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/always.png?raw=true" />   
4、支持在“代理服务器+端口”input框里面加参数
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/param.png?raw=true" />  

### 配置
配置文件为config/config.ini  
可以配置的属性有：端口（默认28080），goproxy的路径（默认[snail007/goproxy](https://github.com/snail007/goproxy/releases)路径在proxy-web目录下的proxyService目录内），登录账号和密码（都为admin）  


### 依赖包
[github.com/boltdb/bolt](https://github.com/boltdb/bolt)使用bolt扩展为数据库  
[github.com/Unknwon/goconfig](https://github.com/Unknwon/goconfig)解析配置文件  
[github.com/astaxie/beego/tree/master/session](https://github.com/astaxie/beego/tree/master/session) session模块  
这些依赖已经在源码内解决，无需go get

### 1.使用
使用28080端口进入页面（如：localhost:28080），首先到登录页面  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/login.png?raw=true" />  
账号密码都为admin，登录进入  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/empty.png?raw=true" />  
点击，添加代理，显示添加代理的弹框，可以选择代理是否开启proxy-web服务时也自动启动  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/add.png?raw=true" /> 
代理添加完成后可执行修改、删除、启用和显示日志的操作  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/preview.png?raw=true" />  

### 2.参数介绍
代理协议：需要用到的协议 如http， tcp等协议。  
本地连接类型：-t参数。  
链式代理：本地连接的类型，“顶级代理”不需要填写“上级服务器+端口”。  
代理服务器+端口：-p参数。  
上级服务器+端口：-P参数。  
父级连接类型：-T参数 ，选取后可能会有不同的加密方式，上传文件的加密方式会有默认文件，tcp形式默认不加密。  
参数具体怎样使用请查看goproxy手册

#### **2.1.http参数** 
tls形式加密：-C .crt文件 和 -K .key文件  
ssh形式加密：有密钥和密码两种方式，-u 用户名 -A 密码 -S 私钥文件 -s 私钥密码  
kcp形式加密：-B 密码  
<img src="https://github.com/yincongcyincong/proxy-web/blob/master/docs/image/http1.png?raw=true" />  
`path to proxy/proxy http -t tcp -p :8081`  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/http2.png?raw=true" />  
`path to proxy/proxy http -t tls -p :8081 -T tls -P 2.2.2.2:8081 -C path to file/proxy.crt -K path to file/proxy.key`  

#### **2.2.tcp参数** 
tls形式加密：-C .crt文件 和 -K .key文件  
kcp形式加密：-B密码  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/tcp1.png?raw=true" />  
`path to proxy/proxy tcp -t tls -p :8081 -T tls -P 2.2.2.2:8081 -C path to file/proxy.crt -K path to file/proxy.key`  

#### **2.3.udp参数** 
没有加密模式  
“本地连接类型”只有udp模式  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/udp1.png?raw=true" />  
`path to proxy/proxy udp -p :8081 -T tls -P 2.2.2.2:8081 -C path to file/proxy.crt -K path to file/proxy.key`

#### **2.4.socks参数** 
tls形式加密：-C .crt文件 和 -K .key文件  
ssh形式加密：有密钥和密码两种方式，-u 用户名 -A 密码 -S 私钥文件 -s 私钥密码  
kcp形式加密：-B 密码  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/socks.png?raw=true" />  
`path to proxy/proxy socks -t tcp -p :8081 -T kcp -P 2.2.2.2:8081 -B 1234 `

#### **2.5.client参数** 
只有tls形式的加密且必须加密  
tls形式加密：-C .crt文件 和 -K .key文件 
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/client.png?raw=true" />  
`path to proxy/proxy client -P ":8081" -C path to file/proxy.crt -K path to file/proxy.key `  
“上级服务器+端口”填写的内容无效

#### **2.6.server参数** 
只有tls形式的加密且必须加密  
tls形式加密：-C .crt文件 和 -K .key文件  
“代理服务器+端口”代表-r参数  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/server.png?raw=true" />  
`path to proxy/proxy server -r "udp://:10053@:53" -P "2.2.2.2:8081" -C path to file/proxy.crt -K path to file/proxy.key`

#### **2.7.bridge参数** 
只有tls形式的加密且必须加密  
tls形式加密：-C .crt文件 和 -K .key文件  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/bridge.png?raw=true" />  
`path to proxy/proxy bridge -P ":8081" -C path to file/proxy.crt -K path to file/proxy.key `  
“上级服务器+端口”填写的内容无效  

### 源码使用  
- 使用linux或者其他平台编译，请删除resource.syso  
- git下载源码  
   
### TODO
- -L参数进程池  
- server -r参数分解  

### License
- under GPLv3 license  

### Contact
- QQ群：189618940
