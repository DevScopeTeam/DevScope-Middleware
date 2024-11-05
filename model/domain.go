package model

type Domain struct {
	UUID     string `json:"uuid" gorm:"primaryKey" validate:"required"`
	Username string `json:"username" gorm:"type:varchar(191); not null" validate:"required"`
	TagUUID  string `json:"tag_uuid" gorm:"type:varchar(191); not null" validate:"required"`
}

type DomainResp struct {
	Code int    `json:"code"`
	Tag  Domain `json:"domain"`
}

type DomainListResp struct {
	Code int      `json:"code"`
	List []Domain `json:"list"`
}

type UsernameListResp struct {
	Code int      `json:"code"`
	List []string `json:"list"`
}
