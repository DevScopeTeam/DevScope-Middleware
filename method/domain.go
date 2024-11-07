package method

import (
	"DevScope-Middleware/model"
	"strings"

	"github.com/tidwall/gjson"
)

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

func GetDomainUUID(name string) (string, error) {
	db, err := getDB()
	if err != nil {
		return "", err
	}
	//结束后关闭 DB
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var tag model.Tag
	if err := db.Where("name = ?", name).First(&tag).Error; err != nil {
		return "", err
	}

	return tag.UUID, nil
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

func GetUserDomains(username string) (string, error) {
	db, err := getDB()
	if err != nil {
		return "", err
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var domain_tag_list []string
	if err := db.Model(&model.Domain{}).Where("username = ?", username).Pluck("tag_uuid", &domain_tag_list).Error; err != nil {
		return "", err
	}

	var tag_name_list []string
	for _, tag_uuid := range domain_tag_list {
		tag, err := GetTag(tag_uuid)
		if err != nil {
			return "", err
		}
		tag_name_list = append(tag_name_list, tag.Name)
	}

	return strings.Join(tag_name_list, ", "), nil
}

func GetUserRankListByTagUUID(tag_uuid string, page, pageSize int) ([]model.DeveloperRank, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var rank_list []model.DeveloperRank
	uname_list, err := GetUsernameListByTagUUID(tag_uuid, page, 1000000)
	if err != nil {
		return nil, err
	}

	for _, uname := range uname_list {
		rank, err := GetRank(uname)
		if err != nil {
			return nil, err
		}
		rank_list = append(rank_list, rank)
	}

	// 排序
	for i := 0; i < len(rank_list)-1; i++ {
		for j := 0; j < len(rank_list)-i-1; j++ {
			if rank_list[j].Overall < rank_list[j+1].Overall {
				rank_list[j], rank_list[j+1] = rank_list[j+1], rank_list[j]
			}
		}
	}

	// 分页
	if len(rank_list) > pageSize {
		rank_list = rank_list[(page-1)*pageSize : page*pageSize]
	}

	return rank_list, nil
}

var domain_prompt = `根据用户提供的所有项目信息推测其开发领域，并直接返回json。

领域集合: ["artificial-intelligence", "machine-learning", "data-science", "software-development", "web-development", "mobile-development", "game-development", "blockchain", "cybersecurity", "cloud-computing", "devops", "database", "internet-of-things", "embedded-systems", "robotics", "quantum-computing"]

返回结果样例1:
{
  "code": 200,
  "list": [
    "domain1",
    "domain2",
	"domain3"
  ]
}

返回结果样例2:
{
  "code": 200,
  "list": []
}`

// 分析用户领域
func AnalyzeUserDomainList(username string) ([]string, error) {
	descriptions, err := GetUserRepoDescriptions(username)
	if err != nil {
		return nil, err
	}

	prompt := ``
	for repo_name, description := range descriptions {
		prompt += repo_name + ": " + description + "\n"
	}

	resp, err := RequestQwen(domain_prompt, prompt)
	if err != nil {
		return nil, err
	}

	resp = strings.TrimSpace(resp)
	list := gjson.Get(resp, "list").Array()

	// 转换为 string 数组
	var result []string
	for _, item := range list {
		result = append(result, item.Str)
	}

	return result, nil
}
