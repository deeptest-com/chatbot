package service

import (
	v1 "github.com/deeptest-com/deeptest-next/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest-next/internal/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type TcCacheService struct {
	InstructionDef *domain.InstructionDef
}

func (s *TcCacheService) Get(req v1.TcCacheReq, ctx iris.Context) (ret interface{}, err error) {
	sess := sessions.Get(ctx)

	ret = sess.Get(req.Key)

	return
}

func (s *TcCacheService) Set(req v1.TcCacheReq, ctx iris.Context) (err error) {
	sess := sessions.Get(ctx)

	sess.Set(req.Key, req.Data)

	return
}

func (s *TcCacheService) Clear(req v1.TcCacheReq, ctx iris.Context) (err error) {
	sess := sessions.Get(ctx)

	sess.Clear()

	return
}
