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

// EventV2CorehrJobDataChangedV1 员工在飞书人事异动完成后将触发该事件。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=corehr&version=v1&resource=job_data&event=changed)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_data/events/changed
func (r *EventCallbackService) HandlerEventV2CorehrJobDataChangedV1(f EventV2CorehrJobDataChangedV1Handler) {
	r.cli.eventHandler.eventV2CorehrJobDataChangedV1Handler = f
}

// EventV2CorehrJobDataChangedV1Handler event EventV2CorehrJobDataChangedV1 handler
type EventV2CorehrJobDataChangedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2CorehrJobDataChangedV1) (string, error)

// EventV2CorehrJobDataChangedV1 ...
type EventV2CorehrJobDataChangedV1 struct {
	JobDataID string `json:"job_data_id,omitempty"` // 主对象ID
}