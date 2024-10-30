package model

import "DevScope-Middleware/model/github"

type DevScoreResp struct {
	Code     int             `json:"code"`
	Score    float64         `json:"score"`
	UserInfo github.UserInfo `json:"user_info"`
}
