package model

type Field struct {
	UUID     string `json:"uuid" gorm:"primaryKey" validate:"required"`
	Username string `json:"username" gorm:"type:varchar(191); not null" validate:"required"`
	TagUUID  string `json:"tag_uuid" gorm:"type:varchar(191); not null" validate:"required"`
}

type FieldResp struct {
	Code int   `json:"code"`
	Tag  Field `json:"field"`
}

type FieldListResp struct {
	Code int     `json:"code"`
	List []Field `json:"list"`
}
