package controller

import (
	"flower/captcha"
	"flower/config"
	"flower/crypto"
	"flower/entity"
	"flower/handler"
	"flower/http"
	"flower/jwt"
	"flower/log"
	"flower/result"
	"flower/router"
	"flower/service"
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

type Auth struct {
}

func init() {
	auth := &Auth{}
	router.AddRoute(
		"/admin/login",
		http.POST,
		auth.Login,
	)
	router.AddRoute(
		"/admin/modify/password",
		http.POST,
		auth.ModifyPassword,
		handler.CheckAdmin,
	)
}

func (a *Auth) Login(ctx *fasthttp.RequestCtx, req *entity.LoginReq) (resp *result.Result) {
	captchaCheck := false
	if captchaCheck {
		verifyResult := captcha.VerifyCaptcha(req.Id, req.VerifyValue)
		if !verifyResult {
			resp = result.CaptchaError
			log.WarnF("Login->[验证码校验失败]->[%v]", req)
			return
		}
	}
	exist, account, err := service.AccSrv.GetAccountByName(req.Name)
	if err != nil {
		resp = result.DatabaseError
		log.ErrorF("Login-> service.AccSrv.GetAccountByName(req.Name)->[%v]", err)
		return
	}
	if !exist {
		resp = result.AcountError
		log.WarnF("Login-> [%s]", resp.Msg)
		return
	}
	ok := crypto.ValidePassword(req.Password, account.Password)
	if !ok {
		resp = result.AcountError
		log.WarnF("Login-> [%s]", resp.Msg)
		return
	}
	//tokenString, err := jwt.GenJwt(map[string]interface{}{http.JwtIdKey: account.Id, http.JwtRoleId: account.RoleId, http.JwtExp: time.Now().Add(24*60 * time.Minute).Unix()})
	duration := int(time.Minute) * config.Conf.JwtConfig.JwtExpiredMin
	tokenString, err := jwt.GenJwt(map[string]interface{}{http.JwtIdKey: account.Id, http.JwtRoleId: account.RoleId, http.JwtExp: time.Now().Add(time.Duration(duration)).Unix()})
	if err != nil {
		log.ErrorF("Login-> jwt.GenJwt->[%v]", err)
	}
	loginResp := entity.LoginResp{Token: tokenString}
	resp = result.NewSuccess(loginResp)
	return
}

func (a *Auth) ModifyPassword(ctx *fasthttp.RequestCtx, req *entity.ModifyPasswordReq) (resp *result.Result) {
	captchaCheck := false
	if captchaCheck {
		verifyResult := captcha.VerifyCaptcha(req.Id, req.VerifyValue)
		if !verifyResult {
			resp = result.CaptchaError
			log.WarnF("Login->[验证码校验失败]->[%v]", req)
			return
		}
	}
	md5Pw := crypto.GetPasswordWithMd5(req.Password)
	id, ok := http.GetJwtId(ctx)
	if !ok {
		resp = result.ForbiddenError
		log.WarnF("Login->[%s]", resp.Msg)
		return
	}
	exist, dbPw, err := service.AccSrv.GetAccountPwById(id)
	if err != nil {
		resp = result.DatabaseError
		log.ErrorF("Login->service.AccSrv.GetAccountPwById(id)->[%v]", err)
		return
	}
	if !exist {
		resp = result.CaptchaError
		log.WarnF("Login->[%s]", resp.Msg)
		return
	}
	if strings.EqualFold(md5Pw, dbPw) {
		resp = result.NewError(result.PasswordUnchangedEc, "密码和之前设置一样")
		log.WarnF("Login->[密码和之前设置一样]->[%v]", req)
		return
	}
	err = service.AccSrv.UpdateAccountPwById(id, md5Pw)
	if err != nil {
		resp = result.DatabaseError
		log.ErrorF("Login->service.AccSrv.UpdateAccountPwById(id, md5Pw)->[%v]", err)
		return
	}
	resp = result.NewSuccess("密码修改成功")
	return
}
