package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"gin/log"
	"net"
	"os"
)

func sendFile(path string, conn net.Conn) {
	defer conn.Close()
	fs, err := os.Open(path)
	defer fs.Close()
	if err != nil {
		fmt.Println("os.Open err = ", err)
		return
	}
	buf := make([]byte, 1024*10)
	for {
		//  打开之后读取文件
		n, err1 := fs.Read(buf)
		log.Info(buf)
		if err1 != nil {
			fmt.Println("fs.Open err = ", err1)
			return
		}
		log.Info(Base64Encode(string(buf)))

		//  发送文件
		_, _ = conn.Write(buf[:n])
	}

}

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var gBase64Coder = base64.NewEncoding(base64Table)

//MD5 string md5
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//Base64Decode base64 decode
func Base64Decode(src string) ([]byte, error) {
	return gBase64Coder.DecodeString(src)
}

//Base64Encode base64 encode
func Base64Encode(src string) string {
	return gBase64Coder.EncodeToString([]byte(src))
}

func main() {
	for {

		fmt.Println("请输入一个全路径的文件,比如,D:\\a.jpg")
		//  获取命令行参数
		var path string
		fmt.Scan(&path)
		// 获取文件名,
		info, err := os.Stat(path)
		if err != nil {
			fmt.Println("os.Stat err = ", err)
			return
		}
		// 发送文件名
		conn, err1 := net.Dial("tcp", ":8000")
		defer conn.Close()
		if err1 != nil {
			fmt.Println("net.Dial err = ", err1)
			return
		}
		conn.Write([]byte(info.Name()))
		// 接受到是不是ok
		buf := make([]byte, 1024)
		n, err2 := conn.Read(buf)
		if err2 != nil {
			fmt.Println("conn.Read err = ", err2)
			return
		}
		if "ok" == string(buf[:n]) {
			fmt.Println("成功")
			sendFile(path, conn)
		}
		// 如果是ok,那么开启一个连接,发送文件
	}
}
