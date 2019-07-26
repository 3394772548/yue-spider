# yue-spider 爬虫程序

**包含了淘宝app 闲鱼app的爬虫程序 仅供测试使用**

当前可用API，其他API慢慢增加中

| api     | 功能 |
| --------- | -----:|
| mtop.taobao.detail.getdetail  | 淘宝商品基本信息 |
| mtop.taobao.detail.getdesc     |   淘宝商品商品详情 |
| mtop.taobao.rate.detaillist.get      |  淘宝商品评价   |
|mtop.taobao.wsearch.appsearch| 淘宝商品搜索 |
|mtop.taobao.wsearch.appsearch.shop| 淘宝店铺商品搜索 |
|mtop.taobao.ugc.tql.facade| 淘宝商品买家秀 |
|mtop.taobao.idle.main.item.search| 闲鱼商品搜索|
|mtop.taobao.idle.item.detail| 闲鱼宝贝详情|
|mtop.taobao.sharepassword.querypassword|淘口令解析|

[在线测试地址](http://212.64.118.243:8080/swagger/index.html)
（有请求限流最好大家下载在本地测试）



下载程序后
针对所属平台直接运行启动程序

| 平台      | 启动命令 |
| --------- | -----:|
| windows  | yue-win.exe |
| mac     |   yue-mac |
| linux      |    yue-linux |


## 命令行启动

![](https://user-images.githubusercontent.com/53135265/61588113-808c5480-abc8-11e9-9df4-b802a729cd58.jpg)

## 服务器启动后，浏览器打开接口文档地址

启动参数

指定启动端口

-port=9090 

设置代理

-proxy=http://127.0.0.1:8888

**接口文档基于swagger-go实现，可以直接在线模拟测试调用 
http://localhost:8080/swagger/index.html**


----

![](https://user-images.githubusercontent.com/53135265/61588114-808c5480-abc8-11e9-9fbe-9fbd4b2128ec.jpg)

----

## 点击其中一个接口展开接口介绍 点击 try it out 进行接口调用

----
![](https://user-images.githubusercontent.com/53135265/61588115-8124eb00-abc8-11e9-9930-b16408e781ef.jpg)

----

## 填写接口参数点击excute 运行

----
![](https://user-images.githubusercontent.com/53135265/61588116-81bd8180-abc8-11e9-81ec-9f0929905bb7.jpg)

----

## 查看对应请求命令和返回内容

----
![](https://user-images.githubusercontent.com/53135265/61588117-81bd8180-abc8-11e9-9d5c-14619f323c73.jpg)