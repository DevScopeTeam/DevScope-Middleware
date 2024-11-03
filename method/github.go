package method

import (
	"DevScope-Middleware/config"
	"DevScope-Middleware/model"
	github_model "DevScope-Middleware/model/github"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// GitHub API 基础 URL
const apiURL = "https://api.github.com"

// 生成 HTTP 请求
func makeRequest(endpoint string) ([]byte, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", apiURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	// 设置 Token 认证 (如果有)
	if config.GithubToken != "" {
		req.Header.Set("Authorization", "Bearer "+config.GithubToken)
	}

	// 设置 X-GitHub-Api-Version
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

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

	return user.Location, nil
}

// 获取开发者活动数据
func GetUserEvents(username string) ([]github_model.Event, error) {
	endpoint := fmt.Sprintf("/users/%s/events/public", username)
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

// 获取用户最近提交过commit记录的仓库（暂未指定时间范围）
func GetUserRecentRepos(username string) ([]github_model.Repo, error) {
	events, err := GetUserEvents(username)
	if err != nil {
		return nil, err
	}

	repos := []github_model.Repo{}
	id_list := make(map[int64]bool)
	for _, event := range events {
		if event.Type == "PushEvent" && event.Actor.Login == username {
			if _, ok := id_list[event.Repo.ID]; !ok {
				repo_name := strings.Split(event.Repo.Name, "/")
				fmt.Println(repo_name)
				if repo, err := GetRepo(repo_name[0], repo_name[1]); err != nil {
					return nil, err
				} else {
					repos = append(repos, repo)
				}
				id_list[event.Repo.ID] = true
			}
		}
	}
	fmt.Println(id_list)

	return repos, nil
}

// 获取用户自己的仓库和语言偏好
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
	// fmt.Println("Contributors:", contributors)

	return contributors, nil
}

// 计算开发者在仓库中的工作占比
func CalculateWorkWeightInRepo(username, owner, repo_name string) (float64, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/commits", owner, repo_name)
	data, err := makeRequest(endpoint)
	if err != nil {
		return 0, err
	}

	// 解析并展示数据
	var commits []github_model.Commit
	if err := json.Unmarshal(data, &commits); err != nil {
		return 0, err
	}
	// total := float64(len(commits))
	dev_count := 0.0
	other_count := 0.0
	for _, commit := range commits {
		// fmt.Println(commit.Author.Login, commit.Committer.Login)
		// fmt.Println("---------------")
		if commit.Author.Login == username {
			dev_count++
		} else if commit.Author.Login != "" {
			other_count++
		}
	}

	return float64(dev_count / (other_count + dev_count)), nil
}

// 计算开发者评分
func CalculateDeveloperScore(username string) (model.DeveloperRank, error) {
	var rank model.DeveloperRank
	rank.Username = username

	// user, err := GetGithuUserInfo(username)
	// if err != nil {
	// 	return rank, err
	// }

	// 获取开发者的活动数据
	events, err := GetUserEvents(username)
	if err != nil {
		return rank, err
	}

	// 获取开发者的仓库信息
	repos, err := GetUserRepos(username)
	if err != nil {
		return rank, err
	}

	// 计算项目重要性
	for _, repo := range repos {
		// fmt.Println(repo.Owner.Login)
		rank.Score.ProjectImportance += float64(repo.ForksCount) * float64(repo.StargazersCount)
	}

	// 计算代码贡献分
	rank.Score.CodeContribution = float64(len(events))

	// 计算社区影响力
	for _, repo := range repos {
		rank.Score.CommunityInfluence += float64(repo.StargazersCount + repo.ForksCount)
	}

	// 综合评分
	rank.Score.CalculateOverallScore()

	return rank, nil
}
