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

// UpdateCoreHrJobData 更新任职信息
//
// - 任职信息是一个按版本更新的对象, 更新任职信息实际上是增加一个新的任职版本
// - 「任职原因」不允许填写为「onboarding」, 当上一个任职版本的「任职原因」为「onboarding」时, 「任职原因」必填
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_data/patch
func (r *CoreHrService) UpdateCoreHrJobData(ctx context.Context, request *UpdateCoreHrJobDataReq, options ...MethodOptionFunc) (*UpdateCoreHrJobDataResp, *Response, error) {
	if r.cli.mock.mockCoreHrUpdateCoreHrJobData != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#UpdateCoreHrJobData mock enable")
		return r.cli.mock.mockCoreHrUpdateCoreHrJobData(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "UpdateCoreHrJobData",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/job_datas/:job_data_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateCoreHrJobDataResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrUpdateCoreHrJobData mock CoreHrUpdateCoreHrJobData method
func (r *Mock) MockCoreHrUpdateCoreHrJobData(f func(ctx context.Context, request *UpdateCoreHrJobDataReq, options ...MethodOptionFunc) (*UpdateCoreHrJobDataResp, *Response, error)) {
	r.mockCoreHrUpdateCoreHrJobData = f
}

// UnMockCoreHrUpdateCoreHrJobData un-mock CoreHrUpdateCoreHrJobData method
func (r *Mock) UnMockCoreHrUpdateCoreHrJobData() {
	r.mockCoreHrUpdateCoreHrJobData = nil
}

// UpdateCoreHrJobDataReq ...
type UpdateCoreHrJobDataReq struct {
	JobDataID                string                                       `path:"job_data_id" json:"-"`                  // 任职信息ID, 示例值: "151515"
	ClientToken              *string                                      `query:"client_token" json:"-"`                // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	UserIDType               *IDType                                      `query:"user_id_type" json:"-"`                // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType         *DepartmentIDType                            `query:"department_id_type" json:"-"`          // 此次调用中使用的部门 ID 类型, 示例值: open_department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, people_corehr_department_id: 以 people_corehr_department_id 来标识部门, 默认值: `people_corehr_department_id`
	JobLevelID               *string                                      `json:"job_level_id,omitempty"`                // 职务级别 ID, 枚举值及详细信息可通过【查询单个职务级别】接口查询获得, 示例值: "6890452208593372679"
	EmployeeTypeID           *string                                      `json:"employee_type_id,omitempty"`            // 人员类型 ID, 枚举值及详细信息可通过【查询单个人员类型】接口查询获得, 示例值: "6890452208593372679"
	WorkingHoursTypeID       *string                                      `json:"working_hours_type_id,omitempty"`       // 工时制度 ID, 枚举值及详细信息可通过【查询单个工时制度】接口查询获得, 示例值: "6890452208593372679"
	WorkLocationID           *string                                      `json:"work_location_id,omitempty"`            // 工作地点 ID, 枚举值及详细信息可通过【查询单个地点】接口查询获得, 示例值: "6890452208593372679"
	DepartmentID             *string                                      `json:"department_id,omitempty"`               // 部门 ID, 枚举值及详细信息可通过【查询单个部门】接口查询获得, 示例值: "6890452208593372679"
	JobID                    *string                                      `json:"job_id,omitempty"`                      // 职务 ID, 枚举值及详细信息可通过【查询单个职务】接口查询获得, 示例值: "6890452208593372679"
	ProbationStartDate       *string                                      `json:"probation_start_date,omitempty"`        // 试用期开始日期, 示例值: "2018-03-16"
	ProbationEndDate         *string                                      `json:"probation_end_date,omitempty"`          // 试用期结束日期（实际结束日期）, 示例值: "2019-05-24"
	PrimaryJobData           *bool                                        `json:"primary_job_data,omitempty"`            // 是否为主任职, 示例值: true
	EmploymentID             *string                                      `json:"employment_id,omitempty"`               // 雇佣 ID, 示例值: "6893014062142064135"
	EffectiveTime            *string                                      `json:"effective_time,omitempty"`              // 生效时间, 示例值: "2020-05-01 00:00:00"
	ExpirationTime           *string                                      `json:"expiration_time,omitempty"`             // 失效时间, 示例值: "2020-05-02 00:00:00"
	JobFamilyID              *string                                      `json:"job_family_id,omitempty"`               // 职务序列 ID, 枚举值及详细信息可通过【查询单个职务序列】接口查询获得, 示例值: "1245678"
	AssignmentStartReason    *UpdateCoreHrJobDataReqAssignmentStartReason `json:"assignment_start_reason,omitempty"`     // 任职原因, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)任职原因（assignment_start_reason）枚举定义部分获得, 示例值: onboarding
	ProbationExpectedEndDate *string                                      `json:"probation_expected_end_date,omitempty"` // 预计试用期结束日期, 示例值: "2006-01-02"
	DirectManagerID          *string                                      `json:"direct_manager_id,omitempty"`           // 实线主管的任职记录ID, 示例值: "6890452208593372679"
	DottedLineManagerIDList  []string                                     `json:"dotted_line_manager_id_list,omitempty"` // 虚线主管的任职记录ID, 示例值: ["6971723901730686501"]
	SecondDirectManagerID    *string                                      `json:"second_direct_manager_id,omitempty"`    // 第二实线主管的任职记录ID, 示例值: "6890452208593372679"
	CostCenterRate           []*UpdateCoreHrJobDataReqCostCenterRate      `json:"cost_center_rate,omitempty"`            // 成本中心分摊信息
	CustomFields             []*UpdateCoreHrJobDataReqCustomField         `json:"custom_fields,omitempty"`               // 自定义字段
}

// UpdateCoreHrJobDataReqAssignmentStartReason ...
type UpdateCoreHrJobDataReqAssignmentStartReason struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// UpdateCoreHrJobDataReqCostCenterRate ...
type UpdateCoreHrJobDataReqCostCenterRate struct {
	CostCenterID *string `json:"cost_center_id,omitempty"` // 支持的成本中心id, 示例值: "6950635856373745165"
	Rate         *int64  `json:"rate,omitempty"`           // 分摊比例, 示例值: 100
}

// UpdateCoreHrJobDataReqCustomField ...
type UpdateCoreHrJobDataReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "\"Sandy\""
}

// UpdateCoreHrJobDataResp ...
type UpdateCoreHrJobDataResp struct {
	JobData *UpdateCoreHrJobDataRespJobData `json:"job_data,omitempty"` // 任职信息
}

// UpdateCoreHrJobDataRespJobData ...
type UpdateCoreHrJobDataRespJobData struct {
	ID                       string                                               `json:"id,omitempty"`                          // 任职信息 ID
	JobLevelID               string                                               `json:"job_level_id,omitempty"`                // 职务级别 ID, 枚举值及详细信息可通过【查询单个职务级别】接口查询获得
	EmployeeTypeID           string                                               `json:"employee_type_id,omitempty"`            // 人员类型 ID, 枚举值及详细信息可通过【查询单个人员类型】接口查询获得
	WorkingHoursTypeID       string                                               `json:"working_hours_type_id,omitempty"`       // 工时制度 ID, 枚举值及详细信息可通过【查询单个工时制度】接口查询获得
	WorkLocationID           string                                               `json:"work_location_id,omitempty"`            // 工作地点 ID, 枚举值及详细信息可通过【查询单个地点】接口查询获得
	DepartmentID             string                                               `json:"department_id,omitempty"`               // 部门 ID, 枚举值及详细信息可通过【查询单个部门】接口查询获得
	JobID                    string                                               `json:"job_id,omitempty"`                      // 职务 ID, 枚举值及详细信息可通过【查询单个职务】接口查询获得
	ProbationStartDate       string                                               `json:"probation_start_date,omitempty"`        // 试用期开始日期
	ProbationEndDate         string                                               `json:"probation_end_date,omitempty"`          // 试用期结束日期（实际结束日期）
	PrimaryJobData           bool                                                 `json:"primary_job_data,omitempty"`            // 是否为主任职
	EmploymentID             string                                               `json:"employment_id,omitempty"`               // 雇佣 ID
	EffectiveTime            string                                               `json:"effective_time,omitempty"`              // 生效时间
	ExpirationTime           string                                               `json:"expiration_time,omitempty"`             // 失效时间
	JobFamilyID              string                                               `json:"job_family_id,omitempty"`               // 职务序列 ID, 枚举值及详细信息可通过【查询单个职务序列】接口查询获得
	AssignmentStartReason    *UpdateCoreHrJobDataRespJobDataAssignmentStartReason `json:"assignment_start_reason,omitempty"`     // 任职原因, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)任职原因（assignment_start_reason）枚举定义部分获得
	ProbationExpectedEndDate string                                               `json:"probation_expected_end_date,omitempty"` // 预计试用期结束日期
	WeeklyWorkingHours       int64                                                `json:"weekly_working_hours,omitempty"`        // 周工作时长
	DirectManagerID          string                                               `json:"direct_manager_id,omitempty"`           // 实线主管的任职记录ID
	DottedLineManagerIDList  []string                                             `json:"dotted_line_manager_id_list,omitempty"` // 虚线主管的任职记录ID
	SecondDirectManagerID    string                                               `json:"second_direct_manager_id,omitempty"`    // 第二实线主管的任职记录ID
	CostCenterRate           []*UpdateCoreHrJobDataRespJobDataCostCenterRate      `json:"cost_center_rate,omitempty"`            // 成本中心分摊信息
	CustomFields             []*UpdateCoreHrJobDataRespJobDataCustomField         `json:"custom_fields,omitempty"`               // 自定义字段
}

// UpdateCoreHrJobDataRespJobDataAssignmentStartReason ...
type UpdateCoreHrJobDataRespJobDataAssignmentStartReason struct {
	EnumName string                                                        `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHrJobDataRespJobDataAssignmentStartReasonDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHrJobDataRespJobDataAssignmentStartReasonDisplay ...
type UpdateCoreHrJobDataRespJobDataAssignmentStartReasonDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// UpdateCoreHrJobDataRespJobDataCostCenterRate ...
type UpdateCoreHrJobDataRespJobDataCostCenterRate struct {
	CostCenterID string `json:"cost_center_id,omitempty"` // 支持的成本中心id
	Rate         int64  `json:"rate,omitempty"`           // 分摊比例
}

// UpdateCoreHrJobDataRespJobDataCustomField ...
type UpdateCoreHrJobDataRespJobDataCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// updateCoreHrJobDataResp ...
type updateCoreHrJobDataResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *UpdateCoreHrJobDataResp `json:"data,omitempty"`
}