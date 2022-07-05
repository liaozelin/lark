// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// CreateBitableView 在数据表中新增一个视图
//
// 该接口支持调用频率上限为 10 QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/create
func (r *BitableService) CreateBitableView(ctx context.Context, request *CreateBitableViewReq, options ...MethodOptionFunc) (*CreateBitableViewResp, *Response, error) {
	if r.cli.mock.mockBitableCreateBitableView != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#CreateBitableView mock enable")
		return r.cli.mock.mockBitableCreateBitableView(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "CreateBitableView",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/views",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createBitableViewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableCreateBitableView mock BitableCreateBitableView method
func (r *Mock) MockBitableCreateBitableView(f func(ctx context.Context, request *CreateBitableViewReq, options ...MethodOptionFunc) (*CreateBitableViewResp, *Response, error)) {
	r.mockBitableCreateBitableView = f
}

// UnMockBitableCreateBitableView un-mock BitableCreateBitableView method
func (r *Mock) UnMockBitableCreateBitableView() {
	r.mockBitableCreateBitableView = nil
}

// CreateBitableViewReq ...
type CreateBitableViewReq struct {
	AppToken string  `path:"app_token" json:"-"`  // bitable app token, 示例值: "appbcbWCzen6D8dezhoCH2RpMAh"
	TableID  string  `path:"table_id" json:"-"`   // table id, 示例值: "tblsRc9GRRXKqhvW"
	ViewName string  `json:"view_name,omitempty"` // 视图名字, 示例值: "表格视图1"
	ViewType *string `json:"view_type,omitempty"` // 视图类型, 示例值: "grid", 可选值有: grid: 表格视图, kanban: 看板视图, gallery: 画册视图, gantt: 甘特视图, form: 表单视图
}

// CreateBitableViewResp ...
type CreateBitableViewResp struct {
	View *CreateBitableViewRespView `json:"view,omitempty"` // 视图
}

// CreateBitableViewRespView ...
type CreateBitableViewRespView struct {
	ViewID   string `json:"view_id,omitempty"`   // 视图Id
	ViewName string `json:"view_name,omitempty"` // 视图名字
	ViewType string `json:"view_type,omitempty"` // 视图类型
}

// createBitableViewResp ...
type createBitableViewResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *CreateBitableViewResp `json:"data,omitempty"`
}
