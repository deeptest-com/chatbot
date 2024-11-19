package router

import (
	"github.com/deeptest-com/deeptest-next/cmd/server/v1/handler"
	"github.com/kataras/iris/v12"
)

type TcbotModule struct {
	TcbotCtrl   *handler.TcbotCtrl   `inject:""`
	TcCacheCtrl *handler.TcCacheCtrl `inject:""`
}

func (m *TcbotModule) Party() func(public iris.Party) {
	return func(index iris.Party) {
		index.Post("/", m.TcbotCtrl.Index).Name = ""

		index.PartyFunc("/cache", func(party iris.Party) {
			party.Post("/get", m.TcCacheCtrl.Get).Name = ""
			party.Post("/set", m.TcCacheCtrl.Set).Name = ""
			party.Post("/clear", m.TcCacheCtrl.Clear).Name = ""
		})
	}
}
