package http

import (
	. "fmt"
	"net"

	"github.com/cihub/seelog"
)

func StartTcpServer(port int) {
	var ipString = Sprintf("0.0.0.0:%d", port)
	seelog.Info("[func:StartTcpServer] starting TCP listen on ", ipString)
	listener, err := net.Listen("tcp", ipString)
	if err != nil {
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// 对每个新连接创建一个协程进行收发数据
		go HandleConnection(conn)
	}
	seelog.Info("[func:StartTcpServer] Stopping TCP listen on ", ipString)
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		// read data from connection.
		_, err := conn.Read(buf[:])
		if err != nil {
			break
		}
		// send data.
		if _, err = conn.Write([]byte("Send From Server")); err != nil {
			break
		}
	}
}
