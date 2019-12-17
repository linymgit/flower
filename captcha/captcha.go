package captcha

import (
	"flower/config"
	"github.com/mojocn/base64Captcha"
	"github.com/mojocn/base64Captcha/store"
	"time"
)

var configC base64Captcha.ConfigCharacter

func Init() {
	configC = base64Captcha.ConfigCharacter{
		Height: config.Conf.CaptchaConfig.CaptchaHeight,
		Width:  config.Conf.CaptchaConfig.CaptchaWidth,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeArithmetic,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	// 默认的十分钟太长了 自定义一个吧
	base64Captcha.SetCustomStore(store.NewMemoryStore(10240, time.Duration(config.Conf.CaptchaConfig.CaptchaExpiredMin*int(time.Minute))))
}

func GetCaptcha() (id, base64Png string) {
	id, digitCap := base64Captcha.GenerateCaptcha("", configC)
	base64Png = base64Captcha.CaptchaWriteToBase64Encoding(digitCap)
	return
}

func VerifyCaptcha(id, verifyValue string) (verifyResult bool) {
	verifyResult = base64Captcha.VerifyCaptcha(id, verifyValue)
	return
}
