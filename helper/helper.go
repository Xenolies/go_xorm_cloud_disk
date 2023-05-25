package helper

import (
	"cloud_disk/core/define"
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

/**
* @Author: Xenolies
* @Date: 2023/5/25 16:53
* @Version: 1.0
 */

// Md5 将字符串转为MD5
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func MakeToken(id int, identity string, name string) (string, error) {
	userClaim := define.UserClaim{
		Id:             id,
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
