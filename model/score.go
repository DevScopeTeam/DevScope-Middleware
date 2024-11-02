package model

type DevScoreResp struct {
	Code  int           `json:"code"`
	Score DeveloperRank `json:"score"`
	// UserInfo github.UserInfo `json:"user_info"`
}
