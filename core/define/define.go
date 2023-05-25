package define

import "github.com/dgrijalva/jwt-go"

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
