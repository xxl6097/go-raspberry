package test

import (
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/server/util"
	"go-raspberry/server/httpclient"
	"net/http"
	"strings"
)

type TestController struct {
}

// NewController http://openai.clife.net:9010/v1/api/user/signin
func NewController() *TestController {
	return &TestController{}
}

func (this *TestController) Test(w http.ResponseWriter, r *http.Request) {
	glog.Warn("Test---->", r)
	body, _ := httpclient.Get1("https://www.baidu.com", nil)
	glog.Info("->", string(body))
	Respond(w, Ignore(false))
	//params := url.Values{}    // Values是一个map类型
	//params.Set("name", "tyy") // Set的时候会把string转成[]string
	//params.Set("hobby", "足球")
}

func (this *TestController) Post(w http.ResponseWriter, r *http.Request) {
	req := util.GetReqMapData(w, r)
	if req != nil {
		glog.Warn("resp---->", req)
	}
	Respond(w, Ignore(false))
}

func (this *TestController) Auth(w http.ResponseWriter, r *http.Request) {
	req := util.GetReqMapData(w, r)
	glog.Warn(req)
	username := req["username"]
	if username == nil || username.(string) == "" {
		Respond(w, Deny(false))
		return
	}
	if strings.Compare("admin", username.(string)) == 0 {
		Respond(w, Allow(true))
		return
	}
	if strings.Compare("uuxia", username.(string)) == 0 {
		Respond(w, Allow(false))
		return
	}
	Respond(w, Ignore(false))
}
