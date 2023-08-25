# 针对Azure 新增的功能

这个0.8.1 版本主要是支持了Azure 环境。

Azure 上ICMP 协议一直不支持，因此ping 无法适用于Azure 虚拟机。在Azure 上实现类似ping 的功能需要使用paping 来实现，即需要指定IP 和端口号，以TCP的方式联通测试虚拟机的连通性和时延。

详细情况请参考微软官方文档：[使用 PsPing & PaPing 进行 TCP 端口连通性测试](https://docs.azure.cn/zh-cn/articles/azure-operations-guide/virtual-network/aog-virtual-network-tcp-psping-paping-connectivity)

## 以下是为了兼容Azure 环境做出的修改


### 配置文件的修改

在config.json 配置文件中增加了一个TCPPort 配置项。该配置项用来设置用来接收TCP 连接的端口。smartping 启动后对TCP 连接的测试将通过这个端口来接收对端发起的连接请求。

该参数紧挨着Port 参数。Port 参数是本机HTTP 服务侦听端口，用来接收用户对HTTP 页面的访问。


### 增加http/tcp.go 文件

在该文件中实现TCP Server，用来接纳对面测试端发起的TCP 连接请求，并把对端发来的不大于1K 的数据返回给对端。SmartPing 启动后，会一直侦听TCP 端口的连接请求并及时做出响应。


### 在nettools/ping.go 文件中增加TCPPing 函数

该函数用于向指定IP地址和端口发起TCP 连接，并计时。计算出TCP 连接自发起到通联成功所消耗的时间。

### 在HTTP/api.go 函数中修正对IP地址的正则表达式校验

此举是为了兼容带有端口号码的IP地址可以被正确写入配置文件中。
