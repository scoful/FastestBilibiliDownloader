# FastestBibiliDownloader

#### 项目地址：**[ FastestBilibiliDownloader](https://github.com/sodaling/FastestBilibiliDownloader)**

**PS：** 原项目作者@sodaling 似乎不怎么上github，又急用，所以直接fork修改了，如@sodaling看到，请联系，然后提pr到原项目

## 修改清单
1. 增加了bvid下载方式
2. 修改了用aid下载后，文件夹和文件目录的命名方式，更直观看到视频名字，其实用aid下载是可以下载多p的，比如那种教程类的
3. 修改了报错信息，更直观知道是哪个视频报错了   
4. 提供了打包exe，在Windows环境直接运行就可

## 使用方法
1. 下载[最新版本](https://github.com/scoful/FastestBilibiliDownloader/releases)
2. 选择一种方式来下载，手动输入，aid 或 upid 或 bvid
-   aid是老版的b站视频id，是一串数字，现在基本是找不到类似的，eg: 801638963
-   upid是指up主id，是一串数字，点up主头像进去，地址上可以看到，eg：110930331
-   bvid是现在的b站视频id，是一串字符串，基本是BV开头的，eg：BV1ey4y1E7k4
3. 在地址上获取对应id类型的具体内容，复制粘贴，回车，静等下载完毕
![img.png](img.png)
   
**PS1：** 如有问题，请提issue

最后感谢原作者的代码思路！！！

===========以下是原文===========


**东半球第二快的Bilibili.com（B站）视频下载器！**

如果你想下载b站某个up主的所有视频，而且要飞快的那种，那么你可以试试这个项目-.-

目前提供两个视频下载方案:

1. 通过视频的aid,下载单个视频.
2. 通过up主的upid(b站叫mid),下载这个up主所投稿的所有视频.



> 特性:
>
> Github上下载b站视频代码已经有很多了.那么本下载器的特点是啥呢?
>
> 因为这是用Golang写的,当然了,也就利用了Golang的特性:goroutine.
>
> 简单来说,特点就是:
>
> **快!贼快!下载的视频越多越快！**
>
> * 当单个aid视频分了若干个part时候,或者当你选了下载up主下所有视频时候.多个视频将会同时并行下载,跑满你的网速绝对不是问题.
> * 下载与合并视频并行处理.如果视频分了多个part,下载完成的同时就会立即合并.该视频合并处理和其他与其他下载和合并同时进行且互不影响.

## 运行

下载的视频会存放在运行路径下的**download**文件夹下，每个视频（aid）一个文件夹。

```shell
Please enter your id type(`aid` or `upid`) #aid是视频id，upid是up主id，可以在up主主页地址栏获得。
Please enter your id #输入id
```



### 使用Golang编译环境
1. 安装Golang编译环境
* Ubuntu
```shell
sudo apt install golang
```

1.1 如果你在中国大陆，那么你大概率可能或许maybe需要配置代理才能顺利进行下一步。
```shell
go env -w GO111MODULE=on #启用Go Moledules
go env -w  GOPROXY=https://goproxy.io #使用官方代理
```

2. 一次性运行FastestBibiliDownloader
程序入口在**cmd/start-concurrent-engine.go**，只需要
```shell
go run cmd/start-concurrent-engine.go
```
首次运行会花时间下一大堆东西，然后按提示操作即可。
注意，合并视频需要FFmpeg的支持。不然只会下载并不会自动合并。FFmpeg的安装教程请咨询搜索引擎。

3. 编译FastestBibiliDownloader
```shell
go build cmd/start-concurrent-engine.go
```
之后直接运行./start-concurrent-engine即可。

### 如果你没有Golang编译环境，或者没有FFmeg环境。那么推荐用docker方式运行。已经写好了dockefile和makefile。你只需要：

   ```shell
   $ cd FastestBilibiliDownloader
   $ make build #下载镜像
   $ make run #运行镜像
   ```

   

### 后续有空会打包bin文件到release的。

## 感谢

1. engine部分的框架参考**ccmouse**的思路，后面自己调整了整体架构部分，非常感谢。
2. [bilibili-downloader](https://github.com/stevenjoezhang/bilibili-downloader)：b站请求视频的API等等都是从这位的代码获得，本身的py代码注释也非常清晰，非常感谢。
3. @sshwy帮忙抓虫纠错
4. @justin201802不厌其烦的帮忙修改

>欢迎各位提pr或者fork或者什么都行，能帮助到你的话欢迎star！疫情无聊在家磨时间的产物，粗糙了一点，欢迎各位完善～

