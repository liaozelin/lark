// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateSheetProtectedDimension 该接口用于根据 spreadsheetToken 和维度信息增加多个保护范围；单次操作不超过5000行或列。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ugDNzUjL4QzM14CO0MTN
func (r *DriveService) CreateSheetProtectedDimension(ctx context.Context, request *CreateSheetProtectedDimensionReq, options ...MethodOptionFunc) (*CreateSheetProtectedDimensionResp, *Response, error) {
	if r.cli.mock.mockDriveCreateSheetProtectedDimension != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateSheetProtectedDimension mock enable")
		return r.cli.mock.mockDriveCreateSheetProtectedDimension(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Drive#CreateSheetProtectedDimension call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateSheetProtectedDimension request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_dimension",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedUserAccessToken: true,
	}
	resp := new(createSheetProtectedDimensionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Drive#CreateSheetProtectedDimension POST https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_dimension failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Drive#CreateSheetProtectedDimension POST https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_dimension failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "CreateSheetProtectedDimension", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateSheetProtectedDimension success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveCreateSheetProtectedDimension(f func(ctx context.Context, request *CreateSheetProtectedDimensionReq, options ...MethodOptionFunc) (*CreateSheetProtectedDimensionResp, *Response, error)) {
	r.mockDriveCreateSheetProtectedDimension = f
}

func (r *Mock) UnMockDriveCreateSheetProtectedDimension() {
	r.mockDriveCreateSheetProtectedDimension = nil
}

type CreateSheetProtectedDimensionReq struct {
	SpreadSheetToken      string                                                 `path:"spreadsheetToken" json:"-"`       // spreadsheet 的 token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN) 的第 4 项
	AddProtectedDimension *CreateSheetProtectedDimensionReqAddProtectedDimension `json:"addProtectedDimension,omitempty"` // 需要增加保护范围的维度信息，可多个范围
}

type CreateSheetProtectedDimensionReqAddProtectedDimension struct {
	Dimension *CreateSheetProtectedDimensionReqAddProtectedDimensionDimension `json:"dimension,omitempty"` // 需要保护行列的维度信息
	Editors   []int64                                                         `json:"editors,omitempty"`   // 允许编辑保护范围的用户的 userID
	LockInfo  *string                                                         `json:"lockInfo,omitempty"`  // 保护范围的信息
}

type CreateSheetProtectedDimensionReqAddProtectedDimensionDimension struct {
	SheetID        string  `json:"sheetId,omitempty"`        // sheetId
	MajorDimension *string `json:"majorDimension,omitempty"` // 默认 ROWS ，可选 ROWS、COLUMNS
	StartIndex     int64   `json:"startIndex,omitempty"`     // 开始的位置
	EndIndex       int64   `json:"endIndex,omitempty"`       // 结束的位置
}

type createSheetProtectedDimensionResp struct {
	Code int64                              `json:"code,omitempty"`
	Msg  string                             `json:"msg,omitempty"`
	Data *CreateSheetProtectedDimensionResp `json:"data,omitempty"`
}

type CreateSheetProtectedDimensionResp struct {
	AddProtectedDimension *CreateSheetProtectedDimensionRespAddProtectedDimension `json:"addProtectedDimension,omitempty"` // 需要增加保护范围的维度信息，可多个范围
}

type CreateSheetProtectedDimensionRespAddProtectedDimension struct {
	Dimension *CreateSheetProtectedDimensionRespAddProtectedDimensionDimension `json:"dimension,omitempty"` // 需要保护行列的维度信息
	Editors   []int64                                                          `json:"editors,omitempty"`   // 允许编辑保护范围的用户的 userID
	LockInfo  string                                                           `json:"lockInfo,omitempty"`  // 保护范围的信息
	ProtectID string                                                           `json:"protectId,omitempty"` // 保护区域的唯一 uid ，可用做后续解除保护
}

type CreateSheetProtectedDimensionRespAddProtectedDimensionDimension struct {
	SheetID        string `json:"sheetId,omitempty"`        // sheetId
	MajorDimension string `json:"majorDimension,omitempty"` // 默认 ROWS ，可选 ROWS、COLUMNS
	StartIndex     int64  `json:"startIndex,omitempty"`     // 开始的位置
	EndIndex       int64  `json:"endIndex,omitempty"`       // 结束的位置
}
