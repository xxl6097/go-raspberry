package wxbot

import (
	"fmt"
	"go-raspberry/server/config"
	"go-raspberry/server/httpclient"
	"sync"
)

/*
	{
	    "to": "uuxia",
	    "type": "text",
	    "content": "# 侨城豪苑公网IP变了\n## IPV4更新了\n - 公网地址：#{ipv4Addr} \n - 域名地址：#{ipv4Domains} \n - 域名更新结果：#{ipv4Result} \n## IPV6更新了\n - 公网地址：#{ipv6Addr} \n - 域名地址：#{ipv6Domains} \n - 域名更新结果：#{ipv6Result} \n"
	}

	{
	    "msgtype": "markdown",
	    "markdown": {
	        "title": "侨城豪苑公网IP变了",
	        "text": "# 侨城豪苑公网IP变了\n## IPV4更新了\n - 公网地址：#{ipv4Addr} \n - 域名地址：#{ipv4Domains} \n - 域名更新结果：#{ipv4Result} \n## IPV6更新了\n - 公网地址：#{ipv6Addr} \n - 域名地址：#{ipv6Domains} \n - 域名更新结果：#{ipv6Result} \n"
	    }
	}
*/
type WxBotService struct {
	mutex sync.Mutex
	//http://x.uuxia.cn:3001/webhook/msg
	wxhost string
	//https://oapi.dingtalk.com/robot/send?access_token=e3fd8d1759709fdb214a778f8bccc411acb533c698d8949cf34cf2ac73aedba9
	ddhost string
}

func NewService() *WxBotService {
	return &WxBotService{
		wxhost: config.Get().Webhook.WX.Host,
		ddhost: config.Get().Webhook.DD.Host,
	}
}

func (this *WxBotService) pushwx(user, text string) {
	header := map[string]string{"Content-Type": "application/json"}
	httpclient.PostStruct(this.wxhost, header, WXRespondv2(user, "text", text))
}

func (this *WxBotService) pushdd(text string) {
	httpclient.PostStruct(this.ddhost, nil, DDRespond("markdown", "侨城豪苑公网IP变了", text))
}

func (this *WxBotService) Webhook(entity *DDNSGOEntity) {
	text := fmt.Sprintf("# 侨城豪苑公网IP变了\n## IPV4更新了\n - 公网地址：%s \n - 域名地址：%s \n - 域名更新结果：%s \n## IPV6更新了\n - 公网地址：%s \n - 域名地址：%s \n - 域名更新结果：%s", entity.Ipv4Addr, entity.Ipv4Domains, entity.Ipv4Result, entity.Ipv6Addr, entity.Ipv6Domains, entity.Ipv6Result)
	this.pushwx(config.Get().Webhook.WX.User, text)
	this.pushdd(text)
}

func (this *WxBotService) Test(text string) {
	this.pushwx(config.Get().Webhook.WX.User, text)
	this.pushdd(text)
}
