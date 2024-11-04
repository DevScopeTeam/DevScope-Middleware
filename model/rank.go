package model

// ProjectImportance 表示项目的重要性
type ProjectImportance struct {
	Stars       int     `json:"stars"`       // 项目星标数
	Forks       int     `json:"forks"`       // 项目Fork数
	Issues      int     `json:"issues"`      // 项目issue数量
	Watchers    int     `json:"watchers"`    // 项目的watcher数量
	Subscribers int     `json:"subscribers"` // 项目的订阅者数量
	WorkWeight  float64 `json:"workWeight"`  // 开发者的工作占比
}

// CodeContribution 表示代码贡献量
type CodeContribution struct {
	CommitCount      int `json:"commitCount"`      // 提交Commit的数量
	IssueResolutions int `json:"issueResolutions"` // 解决的issue数量
}

// CommunityInfluence 表示社区影响力
type CommunityInfluence struct {
	EventParticipation int     `json:"eventParticipation"` // 参与社区活动的次数
	SocialMediaImpact  float64 `json:"socialMediaImpact"`  // 社交媒体影响力评分
	ConflictResolution int     `json:"conflictResolution"` // 调解社区冲突的次数
}

// Score 表示综合评分
type Score struct {
	ProjectImportance  float64 `json:"project"`   // 项目重要性评分
	CodeContribution   float64 `json:"code"`      // 代码贡献量评分
	CommunityInfluence float64 `json:"influence"` // 社区影响力评分
	Overall            float64 `json:"overall"`   // 综合评分
}

// DetailRank 表示详细评分
type DetailRank struct {
	ProjectImportance  ProjectImportance  `json:"projectImportance"`  // 项目重要性评分
	CodeContribution   CodeContribution   `json:"codeContribution"`   // 代码贡献量评分
	CommunityInfluence CommunityInfluence `json:"communityInfluence"` // 社区影响力评分
}

// DeveloperRank 表示开发者的评分
type DeveloperRank struct {
	Username string `json:"username"` // 开发者用户名
	Score    Score  `json:"score"`
}

// CalculateOverallScore 计算综合评分
func (s *Score) CalculateOverallScore() {
	// 假设权重分别为：项目重要性0.3，代码贡献0.4，社区影响力0.3
	weights := [3]float64{0.3, 0.4, 0.3}
	s.Overall = s.ProjectImportance*weights[0] + s.CodeContribution*weights[1] + s.CommunityInfluence*weights[2]
}

// 计算开发者的项目重要性评分
func (p *ProjectImportance) CalculateScore() float64 {
	// 权重可以根据实际情况进行调整
	const (
		forksWeight       = 0.01
		starsWeight       = 0.01
		issuesWeight      = 0.01
		watchersWeight    = 0.01
		subscribersWeight = 0.01
	)

	// 计算加权和
	sum := forksWeight*float64(p.Forks) + starsWeight*float64(p.Stars) + issuesWeight*float64(p.Issues) + watchersWeight*float64(p.Watchers) + subscribersWeight*float64(p.Subscribers)

	// 应用Sigmoid函数
	score := sigmoid(sum)

	// 将分数缩放到0-100的范围
	return score * 100
}
