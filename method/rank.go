package method

import (
	"DevScope-Middleware/model"
)

func AddRank(rank model.DeveloperRank) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(rank); err != nil {
		return err
	}

	if err := db.Create(&rank).Error; err != nil {
		return err
	}

	return nil
}

func UpdateRank(rank model.DeveloperRank) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 校验数据
	if err := validate.Struct(rank); err != nil {
		return err
	}

	if err := db.Save(&rank).Error; err != nil {
		return err
	}

	return nil
}

func GetRank(username string) (model.DeveloperRank, error) {
	db, err := getDB()
	if err != nil {
		return model.DeveloperRank{}, err
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var rank model.DeveloperRank
	if err := db.Where("username = ?", username).First(&rank).Error; err != nil {
		return model.DeveloperRank{}, err
	}

	return rank, nil
}

// 获取排行榜，按overall排序
func GetRankList(page, pageSize int, nation string) ([]model.DeveloperRank, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var ranks []model.DeveloperRank

	if nation != "" {
		if err := db.Where("nation = ?", nation).Order("overall desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&ranks).Error; err != nil {
			return nil, err
		}
		return ranks, nil
	}

	if err := db.Order("overall desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&ranks).Error; err != nil {
		return nil, err
	}

	return ranks, nil
}

func GetRankCount() (int64, error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var count int64
	if err := db.Model(&model.DeveloperRank{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
