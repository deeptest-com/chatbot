package handler

import (
	v1 "github.com/deeptest-com/deeptest-next/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest-next/internal/server/moudules/service"
	_domain "github.com/deeptest-com/deeptest-next/pkg/domain"

	"github.com/kataras/iris/v12"
)

type TcbotCtrl struct {
	BaseCtrl
	TcbotService *service.TcbotService `inject:""`
}

func (c *TcbotCtrl) CreatePart(ctx iris.Context) {
	req := v1.TcNlpReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.TcbotService.CreatePart(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.FailErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.Success.Code, Data: data})
}
