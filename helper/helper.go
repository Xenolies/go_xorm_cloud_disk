package helper

import (
	"cloud_disk/core/consts"
	"cloud_disk/core/define"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/http"
	"net/smtp"
	"path"
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

// GenerateToken 生成用户 Token
func GenerateToken(id int, identity string, name string, expirationTime int) (string, error) {
	userClaim := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expirationTime)).Unix(),
			//ExpiresAt: time.Now().Add(time.Second * time.Duration(expirationTime)).Unix(),
		},
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

// AliOssUpload 阿里云上传
func AliOssUpload(bucket *oss.Bucket, req *http.Request) (string, error) {
	file, fileHandler, err := req.FormFile("file")
	if err != nil {
		return "", err
	}
	UploadPath := "cloud_disk/" + GetUUID() + path.Ext(fileHandler.Filename) // path.Ext(fileHandler.Filename) 获取文件名对应的文件后缀
	fmt.Println("UPLOADPATH: ", UploadPath)
	// 上传到 OSS 上
	err = bucket.PutObject(UploadPath, file)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	// 没有错误的话返回文件路径
	return UploadPath, nil
}

// AnalyzeToken 解析 Token
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !claims.Valid {
		fmt.Println(uc.Name, uc.Identity, uc.ExpiresAt)
		fmt.Println("claims.Valid : ", claims.Valid)
		return uc, errors.New("token is invalid")
	}
	return uc, nil
}
