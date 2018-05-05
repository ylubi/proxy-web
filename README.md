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

### 作用
1、	用web界面的方式使用goproxy，更加方便  
2、	监控goproxy运行情况  
3、	实时显示goproxy产生的日志  
4、	能自启动goproxy  
 
### 下载
[下载地址](https://github.com/yincongcyincong/proxy-web/releases)  

### 更新
v 2.0 全面更新，可以自由配置参数

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
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/empty.jpg?raw=true" />  
点击，添加代理，显示添加代理的弹框，可以选择代理是否开启proxy-web服务时也自动启动  
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/add.jpg?raw=true" />
修改操作
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/update.jpg?raw=true" />
启动操作
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/start.jpg?raw=true" />
查看日志操作
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/log.jpg?raw=true" />
删除操作
<img src="https://github.com/yincongcyincong/proxy-web/raw/master/docs/image/delete.jpg?raw=true" />

### 2.参数介绍
名称：代理的名称。  
参数：指[snail007/goproxy](https://github.com/snail007/goproxy/)中的各种参数。  
开启proxy-web自动启动goproxy功能。  
加密文件上传。  
参数具体怎样使用请查看goproxy手册。  

### 源码使用  
- 使用非windows编译，请删除resource.syso  
- git下载源码  
   
### TODO
- -查找bug

### License
- under GPLv3 license  

### Contact
- QQ群：189618940
