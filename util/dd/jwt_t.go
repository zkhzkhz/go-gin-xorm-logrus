package main

import (
	"fmt"
	"gin/util"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

func main() {
	type UserInfo map[string]interface{}

	t := time.Now()
	key := "welcome to XXY's code world"
	userInfo := make(UserInfo)
	var expTime int64 = 1000 * 1000 * 1000 * 60 * 10
	var tokenState string

	userInfo["exp"] = strconv.FormatInt(t.UTC().UnixNano(), 10) //  strconv.FormatInt(t.UTC().UnixNano(), 10)
	userInfo["local"] = strconv.FormatInt(t.Local().UnixNano(), 10)
	userInfo["loca"] = strconv.FormatInt(t.UnixNano(), 10)
	userInfo["iat"] = "0"

	tokenString := util.CreateToken(key, userInfo)
	claims, ok := util.ParseToken(tokenString, key)
	if ok {
		oldT, _ := strconv.ParseInt(claims.(jwt.MapClaims)["exp"].(string), 10, 64)
		ct := time.Now().UTC().UnixNano()
		c := ct - oldT
		if c > expTime {
			ok = false
			tokenState = "Token 已过期"
		} else {
			tokenState = "Token 正常"
		}
	} else {
		tokenState = "token无效"
	}

	fmt.Println(tokenState)
	fmt.Println(claims)
}
