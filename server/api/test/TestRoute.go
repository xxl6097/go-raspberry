package test

import (
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-http/server/inter"
	"github.com/xxl6097/go-http/server/route"
	"net/http"
)

type TestRoute struct {
	controller *TestController
}

func NewRoute(ctl *TestController) inter.IRoute {
	opt := &TestRoute{
		controller: ctl,
	}
	return opt
}

func (this *TestRoute) Setup(router *mux.Router) {
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/test/post",
		Fun:    this.controller.Post,
		NoAuth: true,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodGet,
		Path:   "/test/get",
		Fun:    this.controller.Test,
		NoAuth: false,
	})
}
