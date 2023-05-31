package helper

import (
	"cloud_disk/core/consts"
	"cloud_disk/core/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"time"
)

/**
* @Author: Xenolies
* @Date: 2023/5/25 16:53
* @Version: 1.0
 */

// Md5 将字符串转为 MD5
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// MakeToken 生成用户 Token
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

// MailSendCode 验证码发送
func MailSendCode(mail string, code string) error {
	e := email.NewEmail()
	e.From = "Jordan Wright <" + consts.MailAddr + ">"
	e.To = []string{mail}
	e.Subject = " 验证码 "
	e.HTML = []byte(" 验证码: <h1>" + code + "</h1>")
	err := e.SendWithTLS(
		"smtp.qq.com:465",
		smtp.PlainAuth(
			"", consts.MailAddr,
			consts.SmtpCodeQQMail,
			"smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		return err
	} else {
		return nil
	}
}

// MakeRandCode 生成随机数
func MakeRandCode() string {
	s := "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

// GetUUID 生成 UUID 作为判断用户的唯一标识
func GetUUID() string {
	uuid := uuid.New()
	return uuid.String()
}
