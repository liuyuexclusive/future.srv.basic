package userHandler

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/liuyuexclusive/future.srv.basic/model"
	user "github.com/liuyuexclusive/future.srv.basic/proto/user"
	"github.com/liuyuexclusive/utils/db"
	"github.com/liuyuexclusive/utils/jwt"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	user.UserHandler
}

const (
	mySigningKey = "sadhasldjkko126312jljdkhfasu0"
)

//Md5 生成32位md5字串
func encrypt(s string, salt string) string {
	h := sha256.New()
	h.Write([]byte(s + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func auth(id, key string) (string, error) {
	var user model.User

	err := db.Open(func(db *db.DB) error {
		db.Where("name=?", id).First(&user)
		return nil
	})

	if err != nil {
		return "", err
	}

	if user.ID == 0 {
		return "", errors.New("无效的登陆名")
	}

	if key == "" {
		return "", errors.New("请输入密码")
	}

	pwd := jwt.Sha256(key, user.Salt)

	if pwd != user.Pwd {
		logrus.Error("密码错误:" + pwd)
		return "", errors.New("密码错误")
	}

	return jwt.GetToken(id)
}

func (e *Handler) Auth(ctx context.Context, req *user.AuthRequest, rsp *user.AuthResponse) error {
	token, err := auth(req.Id, req.Key)
	if err != nil {
		return err
	}
	rsp.Token = token
	return nil
}

func (e *Handler) Validate(ctx context.Context, req *user.ValidateRequest, rsp *user.ValidateResponse) error {
	claims, err := jwt.GetClaims(req.Token)

	if err != nil {
		return err
	}

	rsp.Name = claims["jti"].(string)
	return nil
}

func (e *Handler) Get(ctx context.Context, req *user.GetRequest, rsp *user.GetResponse) error {
	return db.Open(func(db *db.DB) error {
		var user model.User
		db.Where("name=?", req.Name).First(&user)
		if user.ID == 0 {
			return errors.New("找不到用户 " + req.Name)
		}
		rsp.Name = user.Name
		rsp.Access = strings.Split(user.Access, ",")
		rsp.Avatar = user.Avatar
		return nil
	})
}
