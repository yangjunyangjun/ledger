package user

import (
	"github.com/gin-gonic/gin"
	"ledger/api/user"
	"ledger/api/user/db"
	"ledger/lib"
	"ledger/pkg/server/middlware"
	"ledger/pkg/server/webbase"
	"ledger/sdk"
	"time"
)

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

type Register struct {
	UserName string `json:"user_name"  binding:"required"`
	PassWord string `json:"password" binding:"required"`
	Icon     string `json:"icon"`
	Code     int64  `json:"code" binding:"required"`
	Email    string `json:"email" binding:"email"`
}

type Activation struct {
	Email string `json:"email" form:"email" binding:"email"`
}

type userList struct {
	Id       int64  `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
	Email    string `json:"email" form:"email"`
	webbase.RequestPage
}

// @Title 用户登录
// @Author y18175612315@163.com
// @Description 用户登录
// @Tags 用户相关接口
// @Param body body	LoginRequest true "JSON数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/user/login [post]
func (w *UserWebServer) login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		sdk.Log.Error("invalid request")
		w.Error(c, webbase.InvalidRequest, "invalid request")
		return
	}
	u, err := user.GetByUser(req.UserName, req.UserName)
	if err != nil {
		sdk.Log.Errorf("db err %v", err)
		w.Error(c, 5000, "db error")
		return
	}
	if u == nil {
		sdk.Log.Error("user does not exist")
		w.Error(c, 5000, "invalid user_name password")
		return
	}
	if lib.Md5(req.PassWord) != u.Password {
		sdk.Log.Error("invalid user_name password")
		w.Error(c, 5000, "invalid user_name password")
		return
	}
	m := middlware.Claims{
		UserId:   int64(u.ID),
		Username: u.UserName,
		Role:     u.Role,
		Email:    u.Email,
	}
	token, err := m.GenerateToken()
	if err != nil {
		sdk.Log.Errorf("Issue token error %s", err)
		w.Error(c, webbase.InvalidToken, "invalid request")
		return
	}
	c.Header("Authorization", token)
	w.Success(c, nil)
}

// @Title 用户注册
// @Author y18175612315@163.com
// @Description 用户注册
// @Tags 用户相关接口
// @Param body body	Register true "JSON数据"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/user/register [post]
func (w *UserWebServer) register(c *gin.Context) {
	var req Register
	if err := c.ShouldBind(&req); err != nil {
		sdk.Log.Error("invalid request")
		w.Error(c, webbase.InvalidRequest, "invalid request")
		return
	}
	// 从redis取出验证码 校验是否挣钱
	if code, err := sdk.RedisCli.Get(req.Email).Int64(); err != nil || code != req.Code {
		sdk.Log.Errorf("get activate code err  %v or code error", err)
		w.Error(c, 5000, "code error or expired")
		return
	}
	registerUser := db.User{
		UserName: req.UserName,
		Password: lib.Md5(req.PassWord),
		Email:    req.Email,
		Icon:     req.Icon,
	}
	var (
		u   *db.User
		err error
	)
	u, err = user.GetByUser(req.Email, req.UserName)
	if err != nil {
		sdk.Log.Errorf("db err %v", err)
		w.Error(c, 5000, "db error")
		return
	}
	if u.UserName == req.UserName {
		w.Error(c, 5000, "用户名已存在")
		return
	}
	if u.Email == req.Email {
		w.Error(c, 5000, "该邮箱已注册")
		return
	}
	if err = user.AddUser(registerUser); err != nil {
		sdk.Log.Errorf("db err %v", err)
		w.Error(c, 5000, "db error")
		return
	}
	w.Success(c, nil)
}

// @Title 发送激活码
// @Author y18175612315@163.com
// @Description 发送激活码
// @Tags 用户相关接口
// @Param query query Activation true "param参数"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/user/send_activate [get]
func (w *UserWebServer) sendAct(c *gin.Context) {
	var req Activation
	if err := c.ShouldBind(&req); err != nil {
		sdk.Log.Error("invalid request")
		w.Error(c, webbase.InvalidRequest, "invalid request")
		return
	}
	code := lib.RandStr()
	email := lib.Email{
		To:      req.Email,
		Subject: "activate",
		Msg:     code,
	}
	if err := sdk.RedisCli.Set(req.Email, code, time.Duration(180)*time.Second).Err(); err != nil {
		sdk.Log.Errorf("redis set error %v ", err)
		w.Error(c, 5000, "redis set error")
		return
	}
	lib.RegisterMail(&email)
	w.Success(c, nil)
}

// @Title 用户列表
// @Author y18175612315@163.com
// @Description 用户列表
// @Tags 用户相关接口
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param query query userList true "param参数"
// @Success 200 {object} webbase.Response
// @Router	/ledger/v1/user/user_list [get]
func (w *UserWebServer) userList(c *gin.Context) {
	var req userList
	if err := c.ShouldBind(&req); err != nil {
		sdk.Log.Error("invalid request")
		w.Error(c, webbase.InvalidRequest, "invalid request")
		return
	}
	if u := w.GetUser(c); u == nil && u.Role != 2 {
		sdk.Log.Error("user no permission")
		w.Error(c, 5000, "user no permission")
		return
	}
	userList, err := user.QueryUserList(req.Id, req.UserName, req.Email, req.Limit, req.Offset)
	if err != nil {
		sdk.Log.Errorf("query user list err %v", err)
		w.Error(c, 5000, "query user list err")
		return
	}
	w.Success(c, userList)
}
