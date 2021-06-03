// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetApprovalInstance
//
// 通过审批实例 Instance Code  获取审批实例详情。Instance Code 由 [批量获取审批实例](https://open.feishu.cn/document/ukTMukTMukTM/uQDOyUjL0gjM14CN4ITN) 接口获取。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uEDNyUjLxQjM14SM0ITN
func (r *ApprovalService) GetApprovalInstance(ctx context.Context, request *GetApprovalInstanceReq, options ...MethodOptionFunc) (*GetApprovalInstanceResp, *Response, error) {
	if r.cli.mock.mockApprovalGetApprovalInstance != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#GetApprovalInstance mock enable")
		return r.cli.mock.mockApprovalGetApprovalInstance(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "GetApprovalInstance",
		Method:                "POST",
		URL:                   "https://www.feishu.cn/approval/openapi/v2/instance/get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApprovalInstanceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockApprovalGetApprovalInstance(f func(ctx context.Context, request *GetApprovalInstanceReq, options ...MethodOptionFunc) (*GetApprovalInstanceResp, *Response, error)) {
	r.mockApprovalGetApprovalInstance = f
}

func (r *Mock) UnMockApprovalGetApprovalInstance() {
	r.mockApprovalGetApprovalInstance = nil
}

type GetApprovalInstanceReq struct {
	InstanceCode string  `json:"instance_code,omitempty"` // 审批实例 Code
	Locale       *string `json:"locale,omitempty"`        // zh-CN - 中文<br>en-US - 英文
}

type getApprovalInstanceResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 返回码的描述
	Data *GetApprovalInstanceResp `json:"data,omitempty"` // 返回业务信息
}

type GetApprovalInstanceResp struct {
	ApprovalCode string                             `json:"approval_code,omitempty"` // 审批定义 Code
	ApprovalName string                             `json:"approval_name,omitempty"` // 审批名称
	StartTime    int64                              `json:"start_time,omitempty"`    // 审批创建时间
	EndTime      int64                              `json:"end_time,omitempty"`      // 审批完成时间，未完成为 0
	UserID       string                             `json:"user_id,omitempty"`       // 发起审批用户
	OpenID       string                             `json:"open_id,omitempty"`       // 发起审批用户 open id
	SerialNumber string                             `json:"serial_number,omitempty"` // 审批单编号
	DepartmentID string                             `json:"department_id,omitempty"` // 发起审批用户所在部门
	Status       string                             `json:"status,omitempty"`        // 审批实例状态<br>PENDING    - 审批中<br>APPROVED - 通过<br>REJECTED  - 拒绝<br>CANCELED -  撤回<br>DELETED    -  删除
	Form         ApprovalWidgetList                 `json:"form,omitempty"`          // json数组，**控件值**
	TaskList     []*GetApprovalInstanceRespTask     `json:"task_list,omitempty"`     // 审批任务列表
	CommentList  []*GetApprovalInstanceRespComment  `json:"comment_list,omitempty"`  // 评论列表
	Timeline     []*GetApprovalInstanceRespTimeline `json:"timeline,omitempty"`      // 审批动态
}

type GetApprovalInstanceRespTask struct {
	ID           string  `json:"id,omitempty"`             // task id
	UserID       *string `json:"user_id,omitempty"`        // 审批人<br>自动通过、自动拒绝 task user_id 为空
	OpenID       *string `json:"open_id,omitempty"`        // 审批人 open id
	Status       string  `json:"status,omitempty"`         // 任务状态<br>PENDING - 审批中<br>APPROVED - 同意<br>REJECTED  - 拒绝<br>TRANSFERRED - 已转交<br>DONE -  完成
	NodeID       *string `json:"node_id,omitempty"`        // task 所属节点 id
	CustomNodeID *string `json:"custom_node_id,omitempty"` // task 所属节点自定义 id, 如果没设置自定义 id, 则不返回该字段
	Type         string  `json:"type,omitempty"`           // 审批方式<br>AND -会签<br>OR - 或签<br>AUTO_PASS -自动通过<br>AUTO_REJECT - 自动拒绝<br>SEQUENTIAL - 按顺序
	StartTime    int64   `json:"start_time,omitempty"`     // task 开始时间
	EndTime      int64   `json:"end_time,omitempty"`       // task 完成事件, 未完成为 0
}

type GetApprovalInstanceRespComment struct {
	ID         string `json:"id,omitempty"`          // comment id
	UserID     string `json:"user_id,omitempty"`     // 发表评论用户
	OpenID     string `json:"open_id,omitempty"`     // 发表评论用户 open id
	Comment    string `json:"comment,omitempty"`     // 评论内容
	CreateTime int64  `json:"create_time,omitempty"` // 评论时间
}

type GetApprovalInstanceRespTimeline struct {
	Type       string                              `json:"type,omitempty"`         // 动态类型，不同类型 ext 内的 user_id_list 含义不一样<br>START - 审批开始<br>PASS - 通过<br>REJECT  - 拒绝<br>AUTO_PASS -  自动通过<br>AUTO_REJECT - 自动拒绝<br>REMOVE_REPEAT - 去重<br>TRANSFER - 转交 <br>ADD_APPROVER_BEFORE  - 前加签<br>ADD_APPROVER -  并加签<br>ADD_APPROVER_AFTER -  后加签 <br>DELETE_APPROVER  - 减签<br>ROLLBACK_SELECTED -  指定回退<br>ROLLBACK - 全部回退<br>CANCEL -  撤回<br>DELETE - 删除<br>CC - 抄送
	CreateTime int64                               `json:"create_time,omitempty"`  // 发生时间
	UserID     *string                             `json:"user_id,omitempty"`      // 动态产生用户
	OpenID     *string                             `json:"open_id,omitempty"`      // 动态产生用户 open id
	UserIDList []string                            `json:"user_id_list,omitempty"` // 被抄送人列表
	OpenIDList []string                            `json:"open_id_list,omitempty"` // 被抄送人列表
	TaskID     *string                             `json:"task_id,omitempty"`      // 产生动态关联的task_id
	Comment    *string                             `json:"comment,omitempty"`      // 理由
	Ext        *GetApprovalInstanceRespTimelineExt `json:"ext,omitempty"`          // 动态其他信息，目前包括 user_id_list, user_id
}

type GetApprovalInstanceRespTimelineExt struct {
	UserIDList []string `json:"user_id_list,omitempty"` // **type类型** - **user_id_list 含义**<br>TRANSFER - 被转交人 <br>ADD_APPROVER_BEFORE  -  被加签人<br>ADD_APPROVER -   被加签人<br>ADD_APPROVER_AFTER -   被加签人 <br>DELETE_APPROVER  - 被减签人
	OpenIDList []string `json:"open_id_list,omitempty"` // user_id_list 对应的 open id
	UserID     *string  `json:"user_id,omitempty"`      // **type类型** - **user_id 含义**<br>CC - 抄送人
	OpenID     *string  `json:"open_id,omitempty"`      // user_id 对应的 open_id
}
