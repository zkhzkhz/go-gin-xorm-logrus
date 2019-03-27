package main

import (
	"encoding/binary"
	"flag"
	"gin/log"
	"gin/util"
	"net"
	"time"
)

var host = flag.String("host","","host")
var port = flag.String("port","37","port")

func main() {
	flag.Parse()

	addr,err:=net.ResolveUDPAddr("udp","localhost:8082")
	util.HandleErr("Can't resolve address",err,"exit1")
	log.Info("udp server running at ",addr.IP,addr.Port)
	conn,err:=net.ListenUDP("udp",addr)
	util.HandleErr("Error listening",err,"exit1")

	defer conn.Close()
	for   {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	data:=make([]byte,1024)
	n,remoteAddr,err:=conn.ReadFromUDP(data)
	util.HandleErr("failed to read UDP msg because of",err,"return")
	daytime:=time.Now().Unix()
	fMap:=make(map[string]interface{},0)
	fMap["n"]=n
	fMap["dayTime"]=daytime
	log.InfoWithFields("",fMap)
	b:=make([]byte,4)
	binary.BigEndian.PutUint32(b,uint32(daytime))
	_, _ = conn.WriteToUDP(b, remoteAddr)
}

