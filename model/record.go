package model

// 记录
type Record struct {
	UUID    string `json:"uuid" validate:"required" gorm:"primaryKey"`
	Content string `json:"content" validate:"required"`
}

// RecordListResp 记录列表响应
type RecordListResp struct {
	Code       int      `json:"code"`
	RecordList []Record `json:"record_list"`
	Total      int64    `json:"total,omitempty"`
}

// RecordResp 记录响应
type RecordResp struct {
	Code   int    `json:"code"`
	Record Record `json:"record"`
}

// RecordCountResp 记录数量响应
type RecordCountResp struct {
	Code  int   `json:"code"`
	Count int64 `json:"count"`
}
