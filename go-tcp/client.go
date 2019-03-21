package main

import (
	"gin/log"
	"net"
)

func main() {
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
