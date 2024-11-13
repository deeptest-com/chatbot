package router

import (
	"github.com/deeptest-com/deeptest-next/cmd/server/v1/handler"
	"github.com/kataras/iris/v12"
)

type TcbotModule struct {
	TcbotCtrl *handler.TcbotCtrl `inject:""`
}

func (m *TcbotModule) Party() func(public iris.Party) {
	return func(party iris.Party) {
		party.Post("/create_part", m.TcbotCtrl.CreatePart).Name = ""
	}
}
