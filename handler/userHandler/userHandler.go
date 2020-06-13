package userHandler

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/liuyuexclusive/future.srv.basic/model"
	user "github.com/liuyuexclusive/future.srv.basic/proto/user"
	"github.com/liuyuexclusive/utils/db"
	"github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
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
	mySigningKey := []byte(mySigningKey)

	if id == "" {
		return "", errors.New("无效的id")
	}

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

	pwd := encrypt(key, user.Salt)

	if pwd != user.Pwd {
		logrus.Error("密码错误:" + pwd)
		return "", errors.New("密码错误")
	}

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    "test",
		Id:        id,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return token, nil
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
	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(req.Token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("无效token")
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
