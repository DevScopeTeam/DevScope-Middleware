package method

import (
	"DevScope-Middleware/model"
	github_model "DevScope-Middleware/model/github"
	"DevScope-Middleware/utils"
	"encoding/json"
	"fmt"
)

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

	dev_count := 0.0
	other_count := 0.0
	for _, commit := range commits {
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

	user, err := GetGithuUserInfo(username)
	if err != nil {
		return rank, err
	}

	// 获取开发者的活动数据
	// events, err := GetUserEvents(username)
	// if err != nil {
	// 	return rank, err
	// }

	// 获取开发者的仓库信息
	repos, err := GetUserRepos(username)
	if err != nil {
		return rank, err
	}

	var detail model.DetailRank

	// 计算项目重要性
	for _, repo := range repos {
		detail.ProjectImportance.Forks += repo.ForksCount
		detail.ProjectImportance.Stars += repo.StargazersCount
		detail.ProjectImportance.Issues += repo.OpenIssuesCount
		detail.ProjectImportance.Watchers += repo.WatchersCount
		detail.ProjectImportance.Subscribers += repo.SubscribersCount
	}
	rank.ProjectImportance = detail.ProjectImportance.CalculateScore()

	// 计算代码贡献分
	if commit_count, precount, error := GetUserCommitAndPRCounts(username); error != nil {
		return rank, error
	} else {
		detail.CodeContribution.CommitCount = commit_count
		detail.CodeContribution.PrCount = precount
	}
	rank.CodeContribution = detail.CodeContribution.CalculateScore()

	// 计算社区影响力
	detail.CommunityInfluence.Followers = user.Followers
	detail.CommunityInfluence.Stars = detail.ProjectImportance.Stars
	detail.CommunityInfluence.Watchers = detail.ProjectImportance.Watchers
	rank.CommunityInfluence = detail.CommunityInfluence.CalculateScore()

	// 综合评分
	rank.CalculateOverallScore()
	rank.UpdatedAt = utils.GetXTimeNow()

	return rank, nil
}
