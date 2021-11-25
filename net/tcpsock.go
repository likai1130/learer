package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	DELIMITER      = '\t'
)

func serverGo() {
	var listener net.Listener
	listener, err := net.Listen(SERVER_ADDRESS, SERVER_ADDRESS)
	if err != nil {
		log.Printf("Listen Error :%s\n", err.Error())
		return
	}
	defer listener.Close()
	log.Printf("Got listener for the server.(local address: %s)\n",listener.Addr())

	for  {

		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept Error: %s \n",err.Error())
		}
		log.Printf("Established a connection with a client application. (remote address: %s)\n")
		go handleconn(conn)
	}
}

/**
	处理连接
 */
func handleconn(conn net.Conn) {
	for  {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq,err := read(conn)
		if err != nil {
			if err == io.EOF {
				log.Printf("The connertion is closed by another side. (Server) \n")
			}else {
				log.Printf("Read Error: %s (Server)\n",err)
			}
			break
		}
		log.Printf("Received reauest: %s (Server)\n",strReq)

		i32Req,err := converToInt32(strReq)
		if err != nil {
			n,err := write(conn,err.Error())
			if err != nil {
				log.Printf("Write Error (written %d bytes): %s (Server)\n", err)
			}
			log.Printf("Sent response (written %d bytes): %s (Server)\n", n, err)
			continue
		}

		f64Resp := cbrt(i32Req)
		respMsg := fmt.Sprintf("The cube root of %d is %f.", i32Req, f64Resp)
		n,err := write(conn,respMsg)
		if err != nil {
			log.Printf("Write Error: %s (Server)\n", err)
		}
		log.Printf("Sent response (written %d bytes): %s (Server)\n", n, respMsg)
	}
}

/**
	求根
 */
func cbrt(param int32) float64 {
	return math.Cbrt(float64(param))
}

/**
	服务端把数据写给客户端
 */
func write(conn net.Conn, content string) (int,  error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}

/**
	校验是不是int32
 */
func converToInt32(s string)(int32, error) {
	return 0,nil
}

/**
	服务端读客户端数据
 */
func read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for  {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

func clientGo() {

}
func main() {

}
