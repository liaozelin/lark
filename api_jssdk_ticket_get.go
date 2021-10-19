// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetJssdkTicket
//
// 通过在你的网页中引入**飞书网页组件**，可在你的网页中直接使用飞书的功能。
// - 网页组件只适用于自建应用，暂不支持商店应用。
// - 网页组件适用于普通web网页，不建议在小程序webview中调用此组件
// ## 准备工作
// - 在开发者后台[创建一个企业自建应用](https://open.feishu.cn/document/home/introduction-to-custom-app-development/self-built-application-development-process)。
// - 引入组件库。在网页 html 中引入如下代码：
// ```html
// <script src="https://lf1-cdn-tos.bytescm.com/goofy/locl/lark/external_js_sdk/h5-js-sdk-1.0.8.js"></script>
// ```
// ::: warning
// 若要使用成员卡片组件，SDK需要在`<body>`加载后引入。
// :::
// ## 鉴权流程
// ### 1. 获取 access_token
// - 不同的 token 代表了组件使用者的身份。user_access_token代表以用户身份鉴权，app_access_token代表以应用身份授权。
// - 成员名片组件仅支持以用户身份(user_access_token)鉴权
// - 云文档组件可以同时支持以用户身份(user_access_token)和应用身份(app_access_token)授权。但是以应用身份授权云文档时，访问量受 80 次/分钟限制，且组件不支持 “编辑”、“评论”、“点赞” 等功能。
// - 开发者需要通过以下两种方式之一获取 token，再通过接口生成 ticket。
// -  方法一：获取用户身份。通过 [第三方网站免登](https://open.feishu.cn/document/ukTMukTMukTM/uETOwYjLxkDM24SM5AjN)获得 `user_access_token`
// - 方法二：获取应用身份。通过[服务端API](https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/app_access_token_internal)获得 `app_access_token`。
// ### 2. 获取 jsapi_ticket
// 为了降低泄漏风险，这一步应当在你的服务端进行。
//
// doc: https://open.feishu.cn/document/uYjL24iN/uUDO3YjL1gzN24SN4cjN
func (r *JssdkService) GetJssdkTicket(ctx context.Context, request *GetJssdkTicketReq, options ...MethodOptionFunc) (*GetJssdkTicketResp, *Response, error) {
	if r.cli.mock.mockJssdkGetJssdkTicket != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Jssdk#GetJssdkTicket mock enable")
		return r.cli.mock.mockJssdkGetJssdkTicket(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Jssdk",
		API:                 "GetJssdkTicket",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/jssdk/ticket/get",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedAppAccessToken:  true,
		NeedUserAccessToken: true,
	}
	resp := new(getJssdkTicketResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockJssdkGetJssdkTicket(f func(ctx context.Context, request *GetJssdkTicketReq, options ...MethodOptionFunc) (*GetJssdkTicketResp, *Response, error)) {
	r.mockJssdkGetJssdkTicket = f
}

func (r *Mock) UnMockJssdkGetJssdkTicket() {
	r.mockJssdkGetJssdkTicket = nil
}

type GetJssdkTicketReq struct{}

type getJssdkTicketResp struct {
	Code int64               `json:"code,omitempty"` // `返回码，非 0 表示失败`
	Msg  string              `json:"msg,omitempty"`  // `对返回码的文本描述`
	Data *GetJssdkTicketResp `json:"data,omitempty"` // `返回内容`
}

type GetJssdkTicketResp struct {
	ExpireIn int64  `json:"expire_in,omitempty"` // `jsapi_ticket 的有效时间`
	Ticket   string `json:"ticket,omitempty"`    // `jsapi_ticket`
}