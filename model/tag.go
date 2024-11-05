package model

type Tag struct {
	UUID string `json:"uuid" gorm:"primaryKey" validate:"required"`
	Name string `json:"name" gorm:"unique" validate:"required"`
}

type TagResp struct {
	Code int `json:"code"`
	Tag  Tag `json:"tag"`
}

type TagListResp struct {
	Code int   `json:"code"`
	List []Tag `json:"list"`
}
