package proxy

import (
	"github.com/elazarl/goproxy"
	"github.com/vaulty/proxy/model"
)

type userData struct {
	vault *model.Vault
	route *model.Route
}

func ctxUserData(ctx *goproxy.ProxyCtx) *userData {
	if ctx.UserData == nil {
		ctx.UserData = &userData{}
	}

	return ctx.UserData.(*userData)
}