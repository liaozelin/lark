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

// ApproveApprovalInstance 对于单个审批任务进行同意操作。同意后审批流程会流转到下一个审批人。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/approve
func (r *ApprovalService) ApproveApprovalInstance(ctx context.Context, request *ApproveApprovalInstanceReq, options ...MethodOptionFunc) (*ApproveApprovalInstanceResp, *Response, error) {
	if r.cli.mock.mockApprovalApproveApprovalInstance != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#ApproveApprovalInstance mock enable")
		return r.cli.mock.mockApprovalApproveApprovalInstance(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "ApproveApprovalInstance",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/tasks/approve",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(approveApprovalInstanceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalApproveApprovalInstance mock ApprovalApproveApprovalInstance method
func (r *Mock) MockApprovalApproveApprovalInstance(f func(ctx context.Context, request *ApproveApprovalInstanceReq, options ...MethodOptionFunc) (*ApproveApprovalInstanceResp, *Response, error)) {
	r.mockApprovalApproveApprovalInstance = f
}

// UnMockApprovalApproveApprovalInstance un-mock ApprovalApproveApprovalInstance method
func (r *Mock) UnMockApprovalApproveApprovalInstance() {
	r.mockApprovalApproveApprovalInstance = nil
}

// ApproveApprovalInstanceReq ...
type ApproveApprovalInstanceReq struct {
	UserIDType   *IDType `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ApprovalCode string  `json:"approval_code,omitempty"` // 审批定义 Code, 示例值: "7C468A54-8745-2245-9675-08B7C63E7A85"
	InstanceCode string  `json:"instance_code,omitempty"` // 审批实例 Code, 示例值: "81D31358-93AF-92D6-7425-01A5D67C4E71"
	UserID       string  `json:"user_id,omitempty"`       // 根据user_id_type填写操作用户id, 示例值: "f7cb567e"
	Comment      *string `json:"comment,omitempty"`       // 意见, 示例值: "OK"
	TaskID       string  `json:"task_id,omitempty"`       // 任务 ID, 审批实例详情task_list中id, 示例值: "12345"
}

// ApproveApprovalInstanceResp ...
type ApproveApprovalInstanceResp struct {
}

// approveApprovalInstanceResp ...
type approveApprovalInstanceResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *ApproveApprovalInstanceResp `json:"data,omitempty"`
}
