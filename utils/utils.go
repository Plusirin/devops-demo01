package utils

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gin-gonic/gin"

	"devops/model"
	"devops/myerr"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

const (
	Limit = 20
)

func GetRequestID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

// 将一个字符串进行MD5加密后返回加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// 给文本加密
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// 比较密码
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// 验证字段是否有效.
func Validate(m model.Account) error {
	validate := validator.New()
	return validate.Struct(m)
}

func CheckParam(accountName, password string) myerr.Err {
	if accountName == "" {
		return myerr.New(*myerr.ErrValidation, nil).Add("用户名为空.")
	}

	if password == "" {
		return myerr.New(*myerr.ErrValidation, nil).Add("密码为空.")
	}
	return myerr.Err{ErrNum: *myerr.PassParamCheck, Err: nil}
}
