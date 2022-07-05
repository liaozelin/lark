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

// GetDocxBlockListOfDocument 获取文档所有块的富文本内容并分页返回。
//
// 在调用此接口前, 请仔细阅读[新版文档 OpenAPI 接口校验规则](https://bytedance.feishu.cn/docx/doxcnamKaccZKqIMopnREJCZUMe#doxcn6AkCE2AUUm2WwxID7lS7Xc), 了解相关规则及约束。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block/list
func (r *DriveService) GetDocxBlockListOfDocument(ctx context.Context, request *GetDocxBlockListOfDocumentReq, options ...MethodOptionFunc) (*GetDocxBlockListOfDocumentResp, *Response, error) {
	if r.cli.mock.mockDriveGetDocxBlockListOfDocument != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDocxBlockListOfDocument mock enable")
		return r.cli.mock.mockDriveGetDocxBlockListOfDocument(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDocxBlockListOfDocument",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/docx/v1/documents/:document_id/blocks",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDocxBlockListOfDocumentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDocxBlockListOfDocument mock DriveGetDocxBlockListOfDocument method
func (r *Mock) MockDriveGetDocxBlockListOfDocument(f func(ctx context.Context, request *GetDocxBlockListOfDocumentReq, options ...MethodOptionFunc) (*GetDocxBlockListOfDocumentResp, *Response, error)) {
	r.mockDriveGetDocxBlockListOfDocument = f
}

// UnMockDriveGetDocxBlockListOfDocument un-mock DriveGetDocxBlockListOfDocument method
func (r *Mock) UnMockDriveGetDocxBlockListOfDocument() {
	r.mockDriveGetDocxBlockListOfDocument = nil
}

// GetDocxBlockListOfDocumentReq ...
type GetDocxBlockListOfDocumentReq struct {
	DocumentID         string  `path:"document_id" json:"-"`           // 文档的唯一标识, 示例值: "doxcnePuYufKa49ISjhD8Ih0ikh"
	PageSize           *int64  `query:"page_size" json:"-"`            // 分页大小, 示例值: 500, 最大值: `500`
	PageToken          *string `query:"page_token" json:"-"`           // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "aw7DoMKBFMOGwqHCrcO8w6jCmMOvw6ILeADCvsKNw57Di8O5XGV3LG4_w5HCqhFxSnDCrCzCn0BgZcOYUg85EMOYcEAcwqYOw4ojw5QFwofCu8KoIMO3K8Ktw4IuNMOBBHNYw4bCgCV3U1zDu8K-J8KSR8Kgw7Y0fsKZdsKvW3d9w53DnkHDrcO5bDkYwrvDisOEPcOtVFJ-I03CnsOILMOoAmLDknd6dsKqG1bClAjDuS3CvcOTwo7Dg8OrwovDsRdqIcKxw5HDohTDtXN9w5rCkWo"
	DocumentRevisionID *int64  `query:"document_revision_id" json:"-"` // 查询的文档版本, 1表示文档最新版本。若此时查询的版本为文档最新版本, 则需要持有文档的阅读权限；若此时查询的版本为文档的历史版本, 则需要持有文档的编辑权限, 示例值:1, 默认值: `-1`, 最小值: `-1`
	UserIDType         *IDType `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetDocxBlockListOfDocumentResp ...
type GetDocxBlockListOfDocumentResp struct {
	Items     []*DocxBlock `json:"items,omitempty"`      // 文档的 Block 信息
	PageToken string       `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool         `json:"has_more,omitempty"`   // 是否还有更多项
}

// getDocxBlockListOfDocumentResp ...
type getDocxBlockListOfDocumentResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *GetDocxBlockListOfDocumentResp `json:"data,omitempty"`
}
