package main

import (
	"encoding/binary"
	"flag"
	"gin/log"
	"gin/util"
	"net"
	"os"
	"time"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "37", "port")

func main() {
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", "localhost:8083")
	util.HandleErr("Can't resolve address", err, "exit1")

	conn, err := net.DialUDP("udp", nil, addr)
	util.HandleErr("Can't dial", err, "exit1")
	defer conn.Close()
	_, err = conn.Write([]byte(""))
	util.HandleErr("failed:", err, "exit1")

	data := make([]byte, 4)
	_, err = conn.Read(data)
	util.HandleErr("failed to Read UDP msg because of", err, "exit1")
	t := binary.BigEndian.Uint32(data)
	log.Info(time.Unix(int64(t), 0).String())

	os.Exit(0)

}
