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

// EventV2ContactEmployeeTypeEnumActivedV3 启用人员类型会发出对应事件。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=contact&version=v3&resource=employee_type_enum&event=actived)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/events/actived
func (r *EventCallbackService) HandlerEventV2ContactEmployeeTypeEnumActivedV3(f EventV2ContactEmployeeTypeEnumActivedV3Handler) {
	r.cli.eventHandler.eventV2ContactEmployeeTypeEnumActivedV3Handler = f
}

// EventV2ContactEmployeeTypeEnumActivedV3Handler event EventV2ContactEmployeeTypeEnumActivedV3 handler
type EventV2ContactEmployeeTypeEnumActivedV3Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2ContactEmployeeTypeEnumActivedV3) (string, error)

// EventV2ContactEmployeeTypeEnumActivedV3 ...
type EventV2ContactEmployeeTypeEnumActivedV3 struct {
	OldEnum *EventV2ContactEmployeeTypeEnumActivedV3OldEnum `json:"old_enum,omitempty"` // 旧枚举类型
	NewEnum *EventV2ContactEmployeeTypeEnumActivedV3NewEnum `json:"new_enum,omitempty"` // 新枚举类型
}

// EventV2ContactEmployeeTypeEnumActivedV3NewEnum ...
type EventV2ContactEmployeeTypeEnumActivedV3NewEnum struct {
	EnumID      string                                                       `json:"enum_id,omitempty"`      // 枚举值id
	EnumValue   string                                                       `json:"enum_value,omitempty"`   // 枚举的编号值, 创建新的人员类型后, 系统生成对应编号。对应[创建用户接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/create)中用户信息的employee_type字段值
	Content     string                                                       `json:"content,omitempty"`      // 枚举内容, 长度范围: `1` ～ `100` 字符
	EnumType    int64                                                        `json:"enum_type,omitempty"`    // 类型, 可选值有: `1`: 内置类型, `2`: 自定义
	EnumStatus  int64                                                        `json:"enum_status,omitempty"`  // 使用状态, 可选值有: `1`: 激活, `2`: 未激活
	I18nContent []*EventV2ContactEmployeeTypeEnumActivedV3NewEnumI18nContent `json:"i18n_content,omitempty"` // i18n定义
}

// EventV2ContactEmployeeTypeEnumActivedV3NewEnumI18nContent ...
type EventV2ContactEmployeeTypeEnumActivedV3NewEnumI18nContent struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// EventV2ContactEmployeeTypeEnumActivedV3OldEnum ...
type EventV2ContactEmployeeTypeEnumActivedV3OldEnum struct {
	EnumID      string                                                       `json:"enum_id,omitempty"`      // 枚举值id
	EnumValue   string                                                       `json:"enum_value,omitempty"`   // 枚举的编号值, 创建新的人员类型后, 系统生成对应编号。对应[创建用户接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/create)中用户信息的employee_type字段值
	Content     string                                                       `json:"content,omitempty"`      // 枚举内容, 长度范围: `1` ～ `100` 字符
	EnumType    int64                                                        `json:"enum_type,omitempty"`    // 类型, 可选值有: `1`: 内置类型, `2`: 自定义
	EnumStatus  int64                                                        `json:"enum_status,omitempty"`  // 使用状态, 可选值有: `1`: 激活, `2`: 未激活
	I18nContent []*EventV2ContactEmployeeTypeEnumActivedV3OldEnumI18nContent `json:"i18n_content,omitempty"` // i18n定义
}

// EventV2ContactEmployeeTypeEnumActivedV3OldEnumI18nContent ...
type EventV2ContactEmployeeTypeEnumActivedV3OldEnumI18nContent struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}