package method

import "DevScope-Middleware/model"

func AddTag(tag model.Tag) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(tag); err != nil {
		return err
	}

	if err := db.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

func GetTagList() ([]model.Tag, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var tags []model.Tag

	if err := db.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func GetTag(uuid string) (model.Tag, error) {
	db, err := getDB()
	if err != nil {
		return model.Tag{}, err
	}
	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var tag model.Tag

	if err := db.Where("uuid = ?", uuid).First(&tag).Error; err != nil {
		return model.Tag{}, err
	}

	return tag, nil
}

func DeleteTag(uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Where("uuid = ?", uuid).Delete(&model.Tag{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateTag(tag model.Tag) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Model(&model.Tag{}).Where("uuid = ?", tag.UUID).Updates(tag).Error; err != nil {
		return err
	}

	return nil
}
