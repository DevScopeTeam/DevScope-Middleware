package method

import (
	"DevScope-Middleware/config"
	github_model "DevScope-Middleware/model/github"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
		log.Fatalf("Error fetching user info: %v", err)
	}

	// 解析并展示数据
	var user github_model.UserInfo
	if err := json.Unmarshal(data, &user); err != nil {
		log.Fatalf("Error parsing user info: %v", err)
	}

	return user, nil
}

// 获取开发者活动数据
func GetUserEvents(username string) ([]github_model.Event, error) {
	endpoint := fmt.Sprintf("/users/%s/events", username)
	data, err := makeRequest(endpoint)
	if err != nil {
		log.Fatalf("Error fetching user events: %v", err)
	}

	// 解析并展示数据
	var events []github_model.Event
	if err := json.Unmarshal(data, &events); err != nil {
		log.Fatalf("Error parsing user events: %v", err)
	}

	return events, nil
}

// 获取仓库和语言偏好
func GetUserRepos(username string) ([]github_model.Repo, error) {
	endpoint := fmt.Sprintf("/users/%s/repos", username)
	data, err := makeRequest(endpoint)
	if err != nil {
		log.Fatalf("Error fetching user repos: %v", err)
	}

	// 解析并展示数据
	var repos []github_model.Repo
	if err := json.Unmarshal(data, &repos); err != nil {
		log.Fatalf("Error parsing user repos: %v", err)
	}

	return repos, nil
}

// 获取开发者拉取请求
func GetUserPullRequests(owner, repo, username string) ([]github_model.PullRequest, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/pulls?state=all&author=%s", owner, repo, username)
	data, err := makeRequest(endpoint)
	if err != nil {
		log.Fatalf("Error fetching pull requests: %v", err)
	}

	// 解析并展示数据
	var pulls []github_model.PullRequest
	if err := json.Unmarshal(data, &pulls); err != nil {
		log.Fatalf("Error parsing pull requests: %v", err)
	}

	return pulls, nil
}

// 获取开发者提交的 Issues
func GetUserIssues(username string) ([]github_model.Issue, error) {
	endpoint := fmt.Sprintf("/search/issues?q=author:%s", username)
	data, err := makeRequest(endpoint)
	if err != nil {
		log.Fatalf("Error fetching user issues: %v", err)
	}

	// 解析并展示数据
	var issues []github_model.Issue
	if err := json.Unmarshal(data, &issues); err != nil {
		log.Fatalf("Error parsing user issues: %v", err)
	}

	return issues, nil
}

// 获取开源项目
func GetRepo(owner, repo string) (github_model.Repo, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s", owner, repo)
	data, err := makeRequest(endpoint)
	if err != nil {
		log.Fatalf("Error fetching repo info: %v", err)
	}

	// 解析并展示数据
	var repoInfo github_model.Repo
	if err := json.Unmarshal(data, &repoInfo); err != nil {
		log.Fatalf("Error parsing repo info: %v", err)
	}

	return repoInfo, nil
}

// 获取项目贡献者
func GetRepoContributors(owner, repo string) ([]github_model.UserInfo, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/contributors", owner, repo)
	data, err := makeRequest(endpoint)
	if err != nil {
		log.Fatalf("Error fetching contributors: %v", err)
	}

	// 解析并展示数据
	var contributors []github_model.UserInfo
	if err := json.Unmarshal(data, &contributors); err != nil {
		log.Fatalf("Error parsing contributors: %v", err)
	}
	fmt.Println("Contributors:", contributors)

	return contributors, nil
}
