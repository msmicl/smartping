<p align="center">
    <a href="http://smartping.org">
        <img src="http://smartping.org/logo.png">
    </a>
    <h3 align="center">SmartPing | 开源、高效、便捷的网络质量监控神器</h3>
    <p align="center">
        一个综合性网络质量(PING)检测工具，支持正/反向PING绘图、互PING拓扑绘图与报警、全国PING延迟地图与在线检测工具等功能
        <br>
        <a href="http://smartping.org"><strong>-- Browse website --</strong></a>
        <br>
        <br>
        <a href="https://www.travis-ci.org/smartping/smartping">
            <img src="https://www.travis-ci.org/smartping/smartping.svg?branch=master" >
        </a>
        <a href="https://goreportcard.com/report/github.com/smartping/smartping">
            <img src="https://goreportcard.com/badge/github.com/smartping/smartping" >
        </a>
         <a href="https://github.com/smartping/smartping/releases">
             <img src="https://img.shields.io/github/release/smartping/smartping.svg" >
         </a>
         <a href="https://github.com/smartping/smartping/blob/master/LICENSE">
             <img src="https://img.shields.io/hexpm/l/plug.svg" >
         </a>
    </p>    
</p>

## 针对 Azure 网络的改进

Azure 网络不支持ping 的ICMP 协议，因为ICMP 协议无法通过Azure 基础网络的网关。因此，需要使用其他的方式测定Azure 网络延迟。本次改进将使用CentOS 自带的qperf 测定网络延迟，替换掉smartping 内置的ICMP 协议的方式。

qperf 工具在CentOS/RHEL 操作系统中已经存在了十几年了，稳定、可靠。qperf 的[**源代码在这里**](https://www.github.com/linux-rdma/qperf)。

改进后的smartping 强依赖qperf 工具，因此在部署改进版smartping 之前，需要在smartping 运行节点上先安装qperf：

``` bash
sudo yum install -y qperf
```

然后，请从当前repository 拉取smartping 的最新代码：

``` bash
sudo git clone https://www.github.com/msmicl/smartping.git
```

当前版本的代码，在qperf分支中，请先切换分支再进行构建：

``` bash
# 切换分支
sudo git switch -c qperf remotes/origin/qperf

# 构建smartping
cd smartping/src
sudo go build

```
qperf 工具默认使用19765 端口进行通信，请确保测试机器防火墙的19765 端口的TCP 入栈方向开放。所有添加的节点都会强制使用qperf 来进行网络延迟测试。工具的地图上那些公共DNS仍将使用ICMP 的ping 协议去测试。

其他部分，比如服务的启动、停止，smartping 服务的安装等，均按照原有说明操作。

---


## 功能 ##

- 正向PING，反向Ping绘图
- 互PING间机器的状态拓扑，自定义延迟、丢包阈值报警（声音报警与邮件报警），报警时MTR检测
- 全国PING延迟地图（各省份可分电信、联通、移动三条线路）
- 检测工具，支持使用SmartPing各节点进行网络相关检测

## 设计思路 ##

本系统的定位为轻量级工具，即使组多点成互Ping网络可以遵守无中心化原则，所有的数据均存储自身节点中，每个节点提供出方向的数据，从任意节点查询数据均会通过Ajax请求关联节点的API接口获取并组装全部数据。
## 项目截图 ##

![app-bg.jpg](http://smartping.org/assets/img/app-bg.png "")

## 技术交流

<a target="_blank" href="//shang.qq.com/wpa/qunwpa?idkey=dd689e43fd8ecfeb28bffc31d53cb058c6ea23263aa1a34fc032efaf91aae924"><img border="0" src="http://pub.idqqimg.com/wpa/images/group.png" alt="SmartPing" title="SmartPing"></a>

## 项目贡献

欢迎参与项目贡献！比如提交PR修复一个bug，或者新建 [Issue](https://github.com/smartping/smartping/issues/) 讨论新特性或者变更。

## 其他资料 ##

- 官网： http://smartping.org
- 文档： https://docs.smartping.org
- - 下载安装：https://docs.smartping.org/install/
- - API文档：https://docs.smartping.org/api/
