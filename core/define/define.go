package define

import (
	"github.com/dgrijalva/jwt-go"
)

/**
* @Author: Xenolies
* @Date: 2023/5/25 17:01
* @Version: 1.0
 */

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud_disk"

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpiration 验证码过期时间 单位: 分钟
var CodeExpiration = 30

// ExpirationTime Token 过期时间 单位: 分钟
var ExpirationTime int = 30

// RefreshTokenExpire 刷新 Token 时间
var RefreshTokenExpire int = 60

// PageSize 分页的默认参数
var PageSize int = 20

// DataFormat 时间格式化格式
var DataFormat = "2006-01-02 15:04:05"
