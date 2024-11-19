package handler

import (
	"fmt"
	v1 "github.com/deeptest-com/deeptest-next/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest-next/internal/server/moudules/service"
	_domain "github.com/deeptest-com/deeptest-next/pkg/domain"
	"github.com/kataras/iris/v12"
)

type TcCacheCtrl struct {
	BaseCtrl
	TcCacheService *service.TcCacheService `inject:""`
}

func (c *TcCacheCtrl) Get(ctx iris.Context) {
	req := v1.TcCacheReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.TcCacheService.Get(req, ctx)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.FailErr.Code, Msg: err.Error()})
		return
	}

	msg := ""
	if data == nil {
		msg = fmt.Sprintf("NO DATA for KEY '%s'", req.Key)
	}

	ctx.JSON(_domain.Response{Code: _domain.Success.Code, Data: data, Msg: msg})
}

func (c *TcCacheCtrl) Set(ctx iris.Context) {
	req := v1.TcCacheReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.TcCacheService.Set(req, ctx)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.FailErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.Success.Code})
}

func (c *TcCacheCtrl) Clear(ctx iris.Context) {
	req := v1.TcCacheReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.TcCacheService.Clear(req, ctx)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.FailErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.Success.Code})
}
