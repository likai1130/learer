package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

/**
	TCP通讯，需求：
	1. 接收客户端程序的请求
	2. 计算数据立方根
	3. 把结果返回给客户端程序

	分析：

	1. 需要根据事先约定好的数据边界把接收到的请求数据分成数据块
	2. 仅接受可以由int32类型表示的请求数据数据块，对于不符合要求的数据块，要生成错误信息，返回给客户端程序，发送给客户端程序的每块数据都约定好数据边界
	3. 对于每个符合要求的数据块，计算其立方根，生成结果描述返回给客户端
	4. 需要鉴别闲置的通讯连接并且主动关闭。闲置连接的依据：在过去10s中没有任何数据经过该连接到达服务端。这就有效减少相关资源的消耗。
 */
func main() {
	// listener 监听器。网络必须是“tcp”、“tcp4”、“tcp6”、“unix”或“unixpacket”。
	listener, e := net.Listen("tcp", "127.0.0.1:8085")
	if e != nil {
		panic(e)
	}
	//等待连接
	conn, e := listener.Accept()
	if e != nil {
		panic(e)
	}

	var dataBuffer bytes.Buffer
	b := make([]byte, 10)
	for  {
		n, e := conn.Read(b)
		if e != nil {
			if e == io.EOF {
				fmt.Println("The connection is closed")
				conn.Close()
			}else {
				fmt.Println("Read Error: %s \n",e)
			}
			break
		}
		dataBuffer.Write(b[:n])
	}
	fmt.Println(dataBuffer.String())
	fmt.Println(conn.RemoteAddr().String())


}
