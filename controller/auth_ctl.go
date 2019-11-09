package controller

import (
	"flower/captcha"
	"flower/crypto"
	"flower/entity"
	"flower/handler"
	"flower/http"
	"flower/jwt"
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
	verifyResult := captcha.VerifyCaptcha(req.Id, req.VerifyValue)
	if !verifyResult {
		resp = result.CaptchaError
		return
	}
	exist, account, err := service.AccSrv.GetAccountByName(req.Name)
	if err != nil {
		resp = result.DatabaseError
	}
	if !exist {
		resp = result.AcountError
		return
	}
	ok := crypto.ValidePassword(req.Password, account.Password)
	if !ok {
		resp = result.AcountError
		return
	}
	tokenString, err := jwt.GenJwt(map[string]interface{}{http.JwtIdKey: account.Id, http.JwtRoleId: account.RoleId, http.JwtExp: time.Now().Add(100 * time.Minute).Unix()})
	if err != nil {
		// TODO
	}
	loginResp := entity.LoginResp{Token: tokenString}
	resp = result.NewSuccess(loginResp)
	return
}

func (a *Auth) ModifyPassword(ctx *fasthttp.RequestCtx, req *entity.ModifyPasswordReq) (resp *result.Result) {
	verifyResult := captcha.VerifyCaptcha(req.Id, req.VerifyValue)
	if !verifyResult {
		resp = result.CaptchaError
		return
	}
	md5Pw := crypto.GetPasswordWithMd5(req.Password)
	id, ok := http.GetJwtId(ctx)
	if !ok {
		resp = result.ForbiddenError
		return
	}
	exist, dbPw, err := service.AccSrv.GetAccountPwById(id)
	if err != nil {
		resp = result.DatabaseError
		return
	}
	if !exist {
		resp = result.CaptchaError
		return
	}
	if strings.EqualFold(md5Pw, dbPw) {
		resp = result.NewError(result.PasswordUnchangedEc, "密码和之前设置一样")
		return
	}
	err = service.AccSrv.UpdateAccountPwById(id, md5Pw)
	if err != nil {
		resp = result.DatabaseError
		return
	}
	resp = result.NewSuccess("密码修改成功")
	return
}
