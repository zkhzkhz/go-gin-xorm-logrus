package controller

import (
	"gin/log"
	"gin/models"
	"gin/util"
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
	//http.SetCookie(c.Writer, &http.Cookie{
	//	Name:     "token1",
	//	Value:    tokenString,
	//	Expires:  expirationTime,
	//	HttpOnly: true, //防止XSS攻击
	//})
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"token":tokenString,
		"expires":expirationTime,
		"msg":    "登录成功",
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

func GetDataB(c *gin.Context) {
	//request.ParseFromRequest(req *http.Request,extractor request.Extractor,keyFunc jwt.Keyfunc)
	var b models.StructB
	_ = c.Bind(&b)
	tokenString := c.Request.Header.Get("token")
	if tokenString == "" {
		log.Info(tokenString)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "未登录",
		})
		tokenString = util.CreateToken("welcome to XXY's code world", "xiaoming")
		log.Info(tokenString)
		return
	}
	info, bool := util.ParseToken(tokenString, "welcome to XXY's code world")
	log.Info(info)
	log.Info(bool)
	log.Info(b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
	return

}
