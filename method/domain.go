package method

import "DevScope-Middleware/model"

func AddDomain(domain model.Domain) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(domain); err != nil {
		return err
	}

	if err := db.Create(&domain).Error; err != nil {
		return err
	}

	return nil
}

func GetDomain(uuid string) (model.Domain, error) {
	db, err := getDB()
	if err != nil {
		return model.Domain{}, err
	}
	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var domain model.Domain
	if err := db.Where("uuid = ?", uuid).First(&domain).Error; err != nil {
		return model.Domain{}, err
	}

	return domain, nil
}

func DeleteDomain(uuid string) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Where("uuid = ?", uuid).Delete(&model.Domain{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateDomain(domain model.Domain) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Model(&domain).Where("uuid = ?", domain.UUID).Updates(&domain).Error; err != nil {
		return err
	}

	return nil
}

func GetDomainList(page, pageSize int) ([]model.Domain, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var domains []model.Domain
	if err := db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&domains).Error; err != nil {
		return nil, err
	}

	return domains, nil
}

func GetUsernameListByTagUUID(tag_uuid string, page, pageSize int) ([]string, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var usernames []string
	if err := db.Model(&model.Domain{}).Where("tag_uuid = ?", tag_uuid).Limit(pageSize).Offset((page-1)*pageSize).Pluck("username", &usernames).Error; err != nil {
		return nil, err
	}

	return usernames, nil
}
