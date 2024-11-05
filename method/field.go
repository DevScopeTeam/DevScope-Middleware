package method

import "DevScope-Middleware/model"

func AddField(field model.Field) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(field); err != nil {
		return err
	}

	if err := db.Create(&field).Error; err != nil {
		return err
	}

	return nil
}

func GetField(uuid string) (model.Field, error) {
	db, err := getDB()
	if err != nil {
		return model.Field{}, err
	}
	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var field model.Field
	if err := db.Where("uuid = ?", uuid).First(&field).Error; err != nil {
		return model.Field{}, err
	}

	return field, nil
}

func DeleteField(uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Where("uuid = ?", uuid).Delete(&model.Field{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateField(field model.Field) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Model(&field).Where("uuid = ?", field.UUID).Updates(&field).Error; err != nil {
		return err
	}

	return nil
}

func GetFieldList(page, pageSize int) ([]model.Field, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var fields []model.Field
	if err := db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&fields).Error; err != nil {
		return nil, err
	}

	return fields, nil
}

func GetUsernameListByTagUUID(tag_uuid string) ([]string, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var usernames []string
	if err := db.Model(&model.Field{}).Where("tag_uuid = ?", tag_uuid).Pluck("username", &usernames).Error; err != nil {
		return nil, err
	}

	return usernames, nil 
}
