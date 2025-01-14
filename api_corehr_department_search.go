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

// SearchCoreHRDepartment 搜索部门信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/department/search
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/organization-management/department/search
func (r *CoreHRService) SearchCoreHRDepartment(ctx context.Context, request *SearchCoreHRDepartmentReq, options ...MethodOptionFunc) (*SearchCoreHRDepartmentResp, *Response, error) {
	if r.cli.mock.mockCoreHRSearchCoreHRDepartment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHR#SearchCoreHRDepartment mock enable")
		return r.cli.mock.mockCoreHRSearchCoreHRDepartment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "SearchCoreHRDepartment",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/departments/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(searchCoreHRDepartmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRSearchCoreHRDepartment mock CoreHRSearchCoreHRDepartment method
func (r *Mock) MockCoreHRSearchCoreHRDepartment(f func(ctx context.Context, request *SearchCoreHRDepartmentReq, options ...MethodOptionFunc) (*SearchCoreHRDepartmentResp, *Response, error)) {
	r.mockCoreHRSearchCoreHRDepartment = f
}

// UnMockCoreHRSearchCoreHRDepartment un-mock CoreHRSearchCoreHRDepartment method
func (r *Mock) UnMockCoreHRSearchCoreHRDepartment() {
	r.mockCoreHRSearchCoreHRDepartment = nil
}

// SearchCoreHRDepartmentReq ...
type SearchCoreHRDepartmentReq struct {
	PageSize           int64             `query:"page_size" json:"-"`            // 分页大小, 示例值: 100, 取值范围: `1` ～ `100`
	PageToken          *string           `query:"page_token" json:"-"`           // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 6891251722631890445
	UserIDType         *IDType           `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType   *DepartmentIDType `query:"department_id_type" json:"-"`   // 此次调用中使用的部门 ID 类型, 示例值: open_department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, people_corehr_department_id: 以 people_corehr_department_id 来标识部门, 默认值: `open_department_id`
	ManagerList        []string          `json:"manager_list,omitempty"`         // manager ID 列表, 按部门直接负责人搜索, <b>字段权限要求: </b>, 按照部门负责人搜索 (corehr:department.manager.search:read), 示例值: ["7094136522860922112"], 最大长度: `100`
	DepartmentIDList   []string          `json:"department_id_list,omitempty"`   // 部门 ID 列表, 示例值: ["7094136522860922111"]
	NameList           []string          `json:"name_list,omitempty"`            // 部门名称列表, 需精确匹配, 示例值: ["后端研发部"]
	ParentDepartmentID *string           `json:"parent_department_id,omitempty"` // 上级部门 ID, 可查询直接下级部门, <b>字段权限要求: </b>, 按照上级部门搜索(corehr:department.organize.search:read), 示例值: "7094136522860922222"
	CodeList           []string          `json:"code_list,omitempty"`            // 部门 code 列表, 示例值: ["D00000123"]
	Fields             []string          `json:"fields,omitempty"`               // 返回数据的字段列表, 为空时不返回任何字段, 示例值: ["department_name"]
}

// SearchCoreHRDepartmentResp ...
type SearchCoreHRDepartmentResp struct {
	Items     []*SearchCoreHRDepartmentRespItem `json:"items,omitempty"`      // 查询的部门信息
	PageToken string                            `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                              `json:"has_more,omitempty"`   // 是否还有更多项
}

// SearchCoreHRDepartmentRespItem ...
type SearchCoreHRDepartmentRespItem struct {
	ID                 string                                          `json:"id,omitempty"`                   // 部门 ID
	VersionID          string                                          `json:"version_id,omitempty"`           // 部门记录版本 ID
	DepartmentName     []*SearchCoreHRDepartmentRespItemDepartmentName `json:"department_name,omitempty"`      // 部门名称
	SubType            *SearchCoreHRDepartmentRespItemSubType          `json:"sub_type,omitempty"`             // 部门类型, 枚举值 api_name 可通过[【获取字段详情】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/get_by_param)接口查询, 查询参数如下: object_api_name = "department", custom_api_name = "subtype"
	ParentDepartmentID string                                          `json:"parent_department_id,omitempty"` // 上级部门 ID, 字段权限要求: 获取部门组织架构信息
	Manager            string                                          `json:"manager,omitempty"`              // 部门负责人雇佣 ID, 枚举值及详细信息可通过[【搜索员工信息】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/employee/search)接口查询获得, 字段权限要求: 获取部门负责人信息
	TreeOrder          string                                          `json:"tree_order,omitempty"`           // 树形排序, 代表同层级的部门排序序号
	ListOrder          string                                          `json:"list_order,omitempty"`           // 列表排序, 代表所有部门的混排序号
	Code               string                                          `json:"code,omitempty"`                 // 编码
	IsRoot             bool                                            `json:"is_root,omitempty"`              // 是否根部门
	IsConfidential     bool                                            `json:"is_confidential,omitempty"`      // 是否保密
	EffectiveDate      string                                          `json:"effective_date,omitempty"`       // 生效日期
	ExpirationDate     string                                          `json:"expiration_date,omitempty"`      // 失效日期
	Active             bool                                            `json:"active,omitempty"`               // 是否启用
	Description        []*SearchCoreHRDepartmentRespItemDescription    `json:"description,omitempty"`          // 描述
	CustomFields       []*SearchCoreHRDepartmentRespItemCustomField    `json:"custom_fields,omitempty"`        // 自定义字段, 字段权限要求: 获取部门自定义字段
}

// SearchCoreHRDepartmentRespItemCustomField ...
type SearchCoreHRDepartmentRespItemCustomField struct {
	CustomApiName string                                         `json:"custom_api_name,omitempty"` // 自定义字段 apiname, 即自定义字段的唯一标识
	Name          *SearchCoreHRDepartmentRespItemCustomFieldName `json:"name,omitempty"`            // 自定义字段名称
	Type          int64                                          `json:"type,omitempty"`            // 自定义字段类型
	Value         string                                         `json:"value,omitempty"`           // 字段值, 是 json 转义后的字符串, 根据元数据定义不同, 字段格式不同（如 123, 123.23, "true", ["id1", "id2"], "2006-01-02 15:04:05"）
}

// SearchCoreHRDepartmentRespItemCustomFieldName ...
type SearchCoreHRDepartmentRespItemCustomFieldName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// SearchCoreHRDepartmentRespItemDepartmentName ...
type SearchCoreHRDepartmentRespItemDepartmentName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// SearchCoreHRDepartmentRespItemDescription ...
type SearchCoreHRDepartmentRespItemDescription struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// SearchCoreHRDepartmentRespItemSubType ...
type SearchCoreHRDepartmentRespItemSubType struct {
	EnumName string                                          `json:"enum_name,omitempty"` // 枚举值
	Display  []*SearchCoreHRDepartmentRespItemSubTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// SearchCoreHRDepartmentRespItemSubTypeDisplay ...
type SearchCoreHRDepartmentRespItemSubTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// searchCoreHRDepartmentResp ...
type searchCoreHRDepartmentResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *SearchCoreHRDepartmentResp `json:"data,omitempty"`
}
