package main

import (
	"gin/log"
	"io/ioutil"
	"net"
	"os"
)

func main() {

	for i := 0; i < 10000; i++ {
		log.Info(i)
		go client()
	}

}

func client() {

	server := "localhost:8080"
	addr, err := net.ResolveTCPAddr("tcp4", server)
	checkErr(err)
	conn1, err := net.DialTCP("tcp4", nil, addr)
	checkErr(err)
	_, err = conn1.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErr(err)
	response, _ := ioutil.ReadAll(conn1)
	log.Info(string(response))
	//os.Exit(0)

	var m = make(map[string]int, 0)
	log.Info(m)
	m["sb"] = 1
	conn, err := net.Dial("tcp", ":7777")
	if err != nil {
		return
	}
	for {
		buf := make([]byte, 1024)
		if length, err := conn.Read(buf); err == nil {
			if length > 0 {
				buf[length] = 0
				log.Info(string(buf[0:length]))
			}
		}
	}
}
func checkErr(e error) {
	if e != nil {
		log.Info(e)
		os.Exit(1)
	}
}
