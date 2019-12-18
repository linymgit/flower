package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

//const News_Type_Id = 37

var configFilePathFlag = flag.String("cfg", "config.json", "配置文件")

type Configuration struct {
	ServerPort    int           `json:"server_port"`
	NewsTypeId    int           `json:"news_type_id"`
	Md5Salt       string        `json:"md5_salt"`
	CaptchaConfig CaptchaConfig `json:"captcha_config"`
	JwtConfig     JwtConfig     `json:"jwt_config"`
	QNconfig      QNconfig      `json:"qn_config"`
	CorsConfig    CorsConfig    `json:"cors_config,omitempty"`
	MysqlConfig   MysqlConfig   `json:"mysql_config"`
	LogConfig     LogConfig     `json:"log_config"`
}

type CaptchaConfig struct {
	CaptchaHeight     int `json:"captcha_height"`
	CaptchaWidth      int `json:"captcha_weight"`
	CaptchaExpiredMin int `json:"captcha_expired_min"`
}

type JwtConfig struct {
	JwtSecretKey  string `json:"jwt_secret_key"`
	JwtExpiredMin int    `json:"jwt_expired_min"`
}

type QNconfig struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
}

type CorsConfig struct {
	CorsAllowHeaders     string `json:"cors_allow_headers"`
	CorsAllowMethods     string `json:"cors_allow_methods"`
	CorsAllowOrigin      string `json:"cors_allow_origin"`
	CorsAllowCredentials string `json:"cors_allow_credentials"`
}

type MysqlConfig struct {
	ConnUrl         string `json:"conn_url"`
	ShowSQL         bool   `json:"show_sql"`
	ConnMaxLifetime int    `json:"conn_max_lifetime"`
	MaxIdleConns    int    `json:"max_idle_conns"`
	MaxOpenConns    int    `json:"max_open_conns"`
}

type LogConfig struct {
	LogPath   string `json:"log_path"`
	LogPrefix string `json:"log_prefix"`
}

var Conf = new(Configuration)

func getConfigFilePath() string {
	return *configFilePathFlag
}

func LoadConfig() {
	path := getConfigFilePath()

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		//todo
		log.Fatalf("load config error [%v]", err)
		return
	}

	err = json.Unmarshal(buf, Conf)
	if err != nil {
		//todo
		log.Fatalf("load config error [%v]", err)
		return
	}

	//validate(config)
}

//func validate(cfg *Config) {
//	v := validator.New()
//	err := v.Struct(cfg)
//	if err != nil {
//		//todo
//	}
//}
