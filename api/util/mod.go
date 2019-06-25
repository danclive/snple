package util

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/danclive/mqtt-console/log"
	"github.com/gin-gonic/gin"
)

func Page(c *gin.Context) (int64, int64, error) {
	var list_parms struct {
		Page  int64 `form:"page"`
		Limit int64 `form:"limit"`
	}

	if err := c.Bind(&list_parms); err != nil {
		return 0, 0, err
	}

	if list_parms.Page == 0 {
		list_parms.Page = 1
	}

	if list_parms.Limit == 0 {
		list_parms.Limit = 10
	}

	page := list_parms.Page
	limit := list_parms.Limit

	offset := (page - 1) * limit

	return limit, offset, nil
}

func Success(data interface{}) (int, gin.H) {
	return http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	}
}

func Error(code int, info string) (int, gin.H) {
	log.Error(info)
	return http.StatusOK, gin.H{
		"success": false,
		"message": gin.H{
			"code": code,
			"info": info,
		},
	}
}

func PassEncode(value string) (str string) {
	sha_ctx := sha256.New()
	sha_ctx.Write([]byte(value))

	cipher_str := sha_ctx.Sum(nil)

	str = hex.EncodeToString(cipher_str)
	return
}
