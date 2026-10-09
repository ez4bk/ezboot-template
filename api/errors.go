package api

import (
	"net/http"

	"github.com/ez4bk/ezboot/bizerr"
)

var (
	// errEchoServiceRange 服务的错误码起始范围
	errEchoServiceRange = bizerr.NewRange(10000)

	// ErrEchoBadText 表示传入的参数为空
	ErrEchoBadText = errEchoServiceRange.New(http.StatusBadRequest, "bad text") // 10001
)
