package router

import (
	"github.com/kataras/iris/v12"
)

type IndexModule struct {
	AccountModule *AccountModule `inject:""`
	OptLogModule  *OptLogModule  `inject:""`
	PermModule    *PermModule    `inject:""`
	RoleModule    *RoleModule    `inject:""`
	UserModule    *UserModule    `inject:""`

	AibotModule *AibotModule `inject:""`
	TcbotModule *TcbotModule `inject:""`
}

func NewIndexModule() *IndexModule {
	return &IndexModule{}
}

func (m *IndexModule) ApiParty() func(rbac iris.Party) {
	return func(rbac iris.Party) {
		rbac.PartyFunc("/accounts", m.AccountModule.Party())
		rbac.PartyFunc("/optlogs", m.OptLogModule.Party())
		rbac.PartyFunc("/roles", m.RoleModule.Party())
		rbac.PartyFunc("/perms", m.PermModule.Party())
		rbac.PartyFunc("/users", m.UserModule.Party())

		rbac.PartyFunc("/aichat", m.AibotModule.Party())
		rbac.PartyFunc("/tcchat", m.TcbotModule.Party())
	}
}
