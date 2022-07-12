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

// EventV1ApprovalInstance ## 什么是审批事件监听
//
// 开发者通过审批开放接口创建审批实例后, 如要在审批完成后进行额外处理, 需不断轮询审批详情接口获取最新状态, 既增加开发复杂度, 又造成了不必要的接口查询消耗。
// 审批事件监听使用开放平台事件订阅机制, 可将指定消息推送到开发者在应用后台配置的回调 URL 上。审批事件包括审批实例状态变更、审批任务状态变更、审批抄送事件。
// 消息推送可能会有重复, 需要开发者做幂等。
// 审批实例状态变更事件, 会在实例状态变更后, 向开发者推送审批实例状态消息。
// 1. 用户创建审批后, 推送【PENDING】状态
// 2. 任一审批人拒绝后, 推送【REJECTED】状态
// 3. 流程中所有人同意后, 推送【APPROVED】状态
// 4. 发起人撤回审批后, 推送【CANCELED】状态
// 5. 审批定义被管理员删除后, 推送【DELETED】状态
// 6. 发起人撤销已通过的审批, 推送【REVERTED】状态
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ugDNyUjL4QjM14CO0ITN
func (r *EventCallbackService) HandlerEventV1ApprovalInstance(f EventV1ApprovalInstanceHandler) {
	r.cli.eventHandler.eventV1ApprovalInstanceHandler = f
}

// EventV1ApprovalInstanceHandler event EventV1ApprovalInstance handler
type EventV1ApprovalInstanceHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1ApprovalInstance) (string, error)

// EventV1ApprovalInstance ...
type EventV1ApprovalInstance struct {
	AppID        string `json:"app_id,omitempty"` // 如: cli_xxx
	TenantKey    string `json:"tenant_key,omitempty"`
	Type         string `json:"type,omitempty"`          // approval_instance 固定字段
	ApprovalCode string `json:"approval_code,omitempty"` // 审批定义 Code
	InstanceCode string `json:"instance_code,omitempty"` // 审批实例 Code
	Status       string `json:"status,omitempty"`        // 实例状态 PENDING - 进行中 APPROVED - 已通过 REJECTED - 已拒绝 CANCELED -  已撤回 DELETED - 已删除 REVERTED - 已撤销
	OperateTime  string `json:"operate_time,omitempty"`  // 事件发生时间
}