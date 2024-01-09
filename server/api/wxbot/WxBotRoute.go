package wxbot

import (
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-http/server/inter"
	"github.com/xxl6097/go-http/server/route"
	"net/http"
)

type TestRoute struct {
	controller *WxBotController
}

func NewRoute(ctl *WxBotController) inter.IRoute {
	opt := &TestRoute{
		controller: ctl,
	}
	return opt
}

func (this *TestRoute) Setup(router *mux.Router) {
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/wxbot/sendMsg",
		Fun:    this.controller.SendMsg,
		NoAuth: false,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/wxbot/webhook",
		Fun:    this.controller.WebHook_ddnsgo,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodGet,
		Path:   "/wxbot/test",
		Fun:    this.controller.TestSendMsg,
		NoAuth: true,
	})
}
