package method

import (
	github_model "DevScope-Middleware/model/github"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

var system_prompt = `根据用户提供的信息推测归属国家，并直接返回json。

返回结果样例1：
{
    "code": 200,
    "nation": "XXX"
}

返回结果样例2：
{
    "code": 400,
    "nation": "Unknown"
}`

// 获取开发者国籍
func GetUserNationality(username string) (string, error) {
	endpoint := fmt.Sprintf("/users/%s", username)
	data, err := makeRequest(endpoint)
	if err != nil {
		return "", err
	}

	// 解析并展示数据
	var user github_model.UserInfo
	if err := json.Unmarshal(data, &user); err != nil {
		return "", err
	}

	nation_resp, err := RequestQwen(system_prompt, user.Location)
	if err != nil {
		return "", err
	}

	nation_resp = strings.TrimSpace(nation_resp)
	nation := gjson.Get(nation_resp, "nation").String()

	return nation, nil
}
