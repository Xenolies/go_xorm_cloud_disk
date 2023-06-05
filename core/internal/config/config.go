package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	SMTP struct {
		SmtpCode   string
		MailAddr   string
		SmtpServer string
		Port       int
	}
	Mysql struct {
		Driver     string
		DataSource string
	}
	Redis struct {
		DataSource string
		Password   string
		DB         int
	}
	AliOSS struct {
		AccessKey    string
		SecretKey    string
		Endpoint     string
		BucketName   string
		CustomDomain string
	}
}
