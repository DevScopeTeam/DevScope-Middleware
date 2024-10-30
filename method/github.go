package method

import (
	"DevScope-Middleware/config"
	github_model "DevScope-Middleware/model/github"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GitHub API 基础 URL
const apiURL = "https://api.github.com"

// 生成 HTTP 请求
func makeRequest(endpoint string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	// 设置 Token 认证 (如果有)
	if config.GithubToken != "" {
		req.Header.Set("Authorization", "token "+config.GithubToken)
	}

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查 HTTP 状态码
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GitHub API returned status code %d", resp.StatusCode)
	}

	// 读取响应数据
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// 获取开发者基本信息
func GetGithuUserInfo(username string) (github_model.UserInfo, error) {
	endpoint := fmt.Sprintf("/users/%s", username)
	data, err := makeRequest(endpoint)
	if err != nil {
		return github_model.UserInfo{}, err
	}

	// 解析并展示数据
	var user github_model.UserInfo
	if err := json.Unmarshal(data, &user); err != nil {
		return github_model.UserInfo{}, err
	}

	return user, nil
}

// 获取开发者活动数据
func GetUserEvents(username string) ([]github_model.Event, error) {
	endpoint := fmt.Sprintf("/users/%s/events", username)
	data, err := makeRequest(endpoint)
	if err != nil {
		return nil, err
	}

	// 解析并展示数据
	var events []github_model.Event
	if err := json.Unmarshal(data, &events); err != nil {
		return nil, err
	}

	return events, nil
}

// 获取仓库和语言偏好
func GetUserRepos(username string) ([]github_model.Repo, error) {
	endpoint := fmt.Sprintf("/users/%s/repos", username)
	data, err := makeRequest(endpoint)
	if err != nil {
		return nil, err
	}

	// 解析并展示数据
	var repos []github_model.Repo
	if err := json.Unmarshal(data, &repos); err != nil {
		return nil, err
	}

	return repos, nil
}

// 获取开发者拉取请求
func GetUserPullRequests(owner, repo, username string) ([]github_model.PullRequest, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/pulls?state=all&author=%s", owner, repo, username)
	data, err := makeRequest(endpoint)
	if err != nil {
		return nil, err
	}

	// 解析并展示数据
	var pulls []github_model.PullRequest
	if err := json.Unmarshal(data, &pulls); err != nil {
		return nil, err
	}

	return pulls, nil
}

// 获取开发者提交的 Issues
func GetUserIssues(username string) ([]github_model.Issue, error) {
	endpoint := fmt.Sprintf("/search/issues?q=author:%s", username)
	data, err := makeRequest(endpoint)
	if err != nil {
		return nil, err
	}

	// 解析并展示数据
	var issues []github_model.Issue
	if err := json.Unmarshal(data, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

// 获取开源项目
func GetRepo(owner, repo string) (github_model.Repo, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s", owner, repo)
	data, err := makeRequest(endpoint)
	if err != nil {
		return github_model.Repo{}, err
	}

	// 解析并展示数据
	var repoInfo github_model.Repo
	if err := json.Unmarshal(data, &repoInfo); err != nil {
		return github_model.Repo{}, err
	}

	return repoInfo, nil
}

// 获取项目贡献者
func GetRepoContributors(owner, repo string) ([]github_model.UserInfo, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/contributors", owner, repo)
	data, err := makeRequest(endpoint)
	if err != nil {
		return nil, err
	}

	// 解析并展示数据
	var contributors []github_model.UserInfo
	if err := json.Unmarshal(data, &contributors); err != nil {
		return nil, err
	}
	fmt.Println("Contributors:", contributors)

	return contributors, nil
}
