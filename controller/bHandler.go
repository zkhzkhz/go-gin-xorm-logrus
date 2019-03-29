package controller

import (
	"../log"
	"../util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("my_secret_key")

var users = map[string]string{
	"users1": "password1",
	"user2":  "password2",
}

type Credential struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signin(c *gin.Context) {
	var creds Credential

	err := c.Bind(&creds)
	util.HandleErr("bind credential failed", err, "return")

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "未登录",
		})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claim{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
		})
		return
	}
	//c.Header("Authorization", tokenString)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime.UTC().In(time.Local),
		HttpOnly: true, //防止XSS攻击
	})
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"token":   tokenString,
		"expires": expirationTime.Local(),
		"msg":     "登录成功",
	})
}

func Welcome(c *gin.Context) {
	//cookie, err := c.Request.Cookie("token1")
	tokenStr := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(tokenStr, "Bearer ")
	tokenStr = splitToken[1]
	//if err != nil {
	//	if err == http.ErrNoCookie {
	//		c.JSON(http.StatusUnauthorized, gin.H{
	//			"status": http.StatusUnauthorized,
	//			"msg":    "no cookie",
	//		})
	//		return
	//	}
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"status": http.StatusBadRequest,
	//		"msg":    "解析cookie错误",
	//	})
	//	return
	//}

	//tokenStr := cookie.Value
	claims := &Claim{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"msg":    "失效",
		})
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"msg":    "签名错误",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "welcome",
	})
}

func Refresh(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(tokenStr, "Bearer ")
	tokenStr = splitToken[1]

	claims := &Claim{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"msg":    "未登录",
		})
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"msg":    "key invalid",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    nil,
		})
		return
	}
	//	// We ensure that a new token is not issued until enough time has elapsed
	//	// In this case, a new token will only be issued if the old token is within
	//	// 30 seconds of expiry. Otherwise, return a bad request status
	//if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	//	c.JSON(http.StatusBadRequest, nil)
	//	return
	//}
	expirationTime := time.Now().Add(5 * time.Minute)
	log.Info("Now", time.Now())
	log.Info(expirationTime)
	claims.ExpiresAt = expirationTime.Unix()
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
		})
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime.Local(),
		HttpOnly: true, //防止XSS攻击
	})
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "刷新成功",
	})
}

//func GetDataB(c *gin.Context) {
//	//request.ParseFromRequest(req *http.Request,extractor request.Extractor,keyFunc jwt.Keyfunc)
//	var b models.StructB
//	_ = c.Bind(&b)
//	tokenStr := GetTokenStr(c)
//
//	claims := &Claim{}
//	err := ValidToken(tokenStr, claims)
//	if err == nil {
//
//		c.JSON(200, gin.H{
//			"a": b.NestedStruct,
//			"b": b.FieldB,
//		})
//		return
//	}
//	if err == util.ErrEmptyToken {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status": http.StatusBadRequest,
//			"msg":    "token值为空",
//		})
//		return
//	}
//	if err == util.ErrExpiredToken {
//		c.JSON(http.StatusUnauthorized, gin.H{
//			"status": http.StatusUnauthorized,
//			"msg":    "令牌失效",
//		})
//		return
//	}
//	if err == jwt.ErrSignatureInvalid {
//		c.JSON(http.StatusUnauthorized, gin.H{
//			"status": http.StatusUnauthorized,
//			"msg":    "key invalid",
//		})
//		return
//	} else {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status": http.StatusBadRequest,
//			"msg":    "验证出错",
//		})
//		return
//	}
//
//}
func GetTokenStr(c *gin.Context) string {
	tokenStr := c.Request.Header.Get("authorization")
	splitToken := strings.Split(tokenStr, "Bearer ")
	tokenStr = splitToken[1]
	return tokenStr
}

//func ValidToken(tokenStr string, claims *Claim) error {
//	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (i interface{}, e error) {
//		return jwtKey, nil
//	})
//	if token == nil {
//		return util.ErrEmptyToken
//	}
//	if !token.Valid {
//		return util.ErrExpiredToken
//	}
//	if err != nil {
//		return err
//	}
//	return err
//}
