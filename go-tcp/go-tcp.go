package main

import (
	"gin/log"
	"io/ioutil"
	"net"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Info(http.ListenAndServe(":8081", nil))
	}()
	tcpserver, _ := net.ResolveTCPAddr("tcp4", ":8080")
	server2, _ := net.ListenTCP("tcp", tcpserver)
	go func() {
		for {
			conn, err := server2.Accept()
			if err != nil {
				log.Info(err)
				continue
			}
			go handle(conn)
		}
	}()
	server, err := net.Listen("tcp", ":7777")
	if err != nil {
		return
	}
	for {
		conn, err := server.Accept()
		log.Info("a new client connection")
		if err != nil {
			log.Error(err)
			//return
		}
		_, _ = conn.Write([]byte("hello world!"))
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	//读取客户端传送的消息
	go func() {
		response, _ := ioutil.ReadAll(conn)
		log.Info(string(response))
	}()

	//send message to client
	time.Sleep(1 * time.Second)
	now := time.Now().String()
	_, _ = conn.Write([]byte(now))
}
