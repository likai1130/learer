package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	DELIMITER      = '\t'
)
var wg sync.WaitGroup

func serverGo() {
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
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
		log.Printf("Established a connection with a client application. (remote address: %s)\n", conn.RemoteAddr())
		go handleconn(conn)
	}
}

/**
	处理连接
 */
func handleconn(conn net.Conn) {
	defer func() {
		conn.Close()
		wg.Done()
	}()
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
func converToInt32(str string)(int32, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("Parse Error: %s\n",err.Error())
		return 0, errors.New(fmt.Sprintf("'%s' is not integer!", str))
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		log.Printf(fmt.Sprintf("Convert Error: The integer %s is too large/small.\n", num))
		return 0, errors.New(fmt.Sprintf("'%s' is not 32-bit integer!", num))
	}
	return int32(num), nil
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

/**
	客户端
 */
func clientGo(id int) {
	defer wg.Done()
	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2*time.Second)
	if err != nil {
		log.Printf("Dial Error: %s (Client[%d])\n",err, id)
		return
	}
	defer conn.Close()
	log.Printf("Connected to server. (remote address: %s, local address: %s) (Client[%d])\n",conn.RemoteAddr(), conn.LocalAddr(), id)
	time.Sleep(200 * time.Millisecond)
	requestNum := 5
	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))
	for i := 0; i < requestNum; i++ {
		i32Req := rand.Int31()
		n ,err := write(conn,fmt.Sprintf("%d", i32Req))
		if err != nil {
			log.Printf("Write Error: %s (Client[%d])\n",err, id)
			continue
		}
		log.Printf("Sent request (written %d bytes): %d (Client[%d])\n",n,i32Req,id)
	}
	for i := 0; i < requestNum; i++ {
		strResp, err := read(conn)
		if err != nil {
			if err == io.EOF {
				log.Printf("The connection is closed by another side. (Client[%d])\n", id)
			}else {
				log.Printf("Read Error: %s (Client[%d])\n", err, id)
			}
			break
		}
		log.Printf("Received response: %s (Client[%d])\n", strResp, id)
	}

}

func main() {
	wg.Add(2)
	go serverGo()
	time.Sleep(500 * time.Millisecond)
	go clientGo(1)
	wg.Wait()
}
