package server

import (
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/server"
	"github.com/xxl6097/go-http/server/route"
	"github.com/xxl6097/go-http/server/token"
	"github.com/xxl6097/go-http/server/util"
	"go-raspberry/server/api"
	"go-raspberry/server/config"
	"strings"
)

var tokenArr []string

func init() {
	glog.Debug("init...")
	route.RouterUtil.SetApiPath(config.Get().Server.ApiPath)
	glog.SetLogFile("./log", "app.log")
	glog.SetCons(true)
	tokens := config.Get().Token
	if tokens != "" {
		tokenArr = strings.Split(tokens, ",")
	}
	token.TokenUtils.SetTokenCallBack(func(token string) (bool, map[string]interface{}) {
		if tokenArr == nil || len(tokenArr) == 0 {
			return false, nil
		}
		if util.Contains(tokenArr, token) {
			return true, nil
		}
		return false, nil
	})
}

func Do() {
	glog.Info("Do")
	api.Do()
	server.NewServerWithLogView(config.Get().Logview.Username, config.Get().Logview.Password).Start(fmt.Sprintf(":%d", config.Get().Server.Port))
}
