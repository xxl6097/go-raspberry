package wxbot

import (
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/server/util"
	"net/http"
)

type WxBotController struct {
	srv *WxBotService
}

func NewController(service *WxBotService) *WxBotController {
	return &WxBotController{
		srv: service,
	}
}

func (this *WxBotController) SendMsg(w http.ResponseWriter, r *http.Request) {
	req := util.GetReqMapData(w, r)
	if req != nil {
		glog.Warnf("resp---->%+v", req)
	}
	Respond(w, Ignore(false))
}

func (this *WxBotController) TestSendMsg(w http.ResponseWriter, r *http.Request) {
	text := util.GetRequestParam(r, "text")
	if text != "" {
		glog.Warn("resp---->", text)
		this.srv.Test(text)
	}
	Respond(w, Ok())
}

func (this *WxBotController) WebHook_ddnsgo(w http.ResponseWriter, r *http.Request) {
	req := util.GetReqData[DDNSGOEntity](w, r)
	if req != nil {
		glog.Warn("resp---->", req)
		this.srv.Webhook(req)
	}
	Respond(w, Ok())
}
