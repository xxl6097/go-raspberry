package api

import (
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/api"
	"go-raspberry/server/api/test"
	"go-raspberry/server/api/upload"
	"go-raspberry/server/api/wxbot"
)

func Do() {
	glog.Info("Do")
	api.GetApi().Add(test.NewRoute(test.NewController()))
	api.GetApi().Add(upload.NewRoute(upload.NewController()))
	api.GetApi().Add(wxbot.NewRoute(wxbot.NewController(wxbot.NewService())))
}
