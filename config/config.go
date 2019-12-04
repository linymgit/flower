package config

const News_Type_Id = 37

type Config struct {
	ServerPort    int           `json:"server_port"`
	NewsTypeId    int           `json:"news_type_id"`
	Md5Salt       string        `json:"md5_salt"`
	CaptchaConfig CaptchaConfig `json:"captcha_config"`
	JwtConfig     JwtConfig     `json:"jwt_config"`
	QNconfig      QNconfig      `json:"qn_config"`
	CorsConfig    CorsConfig    `json:"cors_config,omitempty"`
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
}

type CorsConfig struct {
	CorsAllowHeaders     string `json:"cors_allow_headers"`
	CorsAllowMethods     string `json:"cors_allow_methods"`
	CorsAllowOrigin      string `json:"cors_allow_origin"`
	CorsAllowCredentials string `json:"cors_allow_credentials"`
}

type MysqlConfig struct {
	ConnUrl string `json:"conn_url"`
}
