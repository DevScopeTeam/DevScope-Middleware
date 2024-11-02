package model

// ProjectImportance 表示项目的重要性
type ProjectImportance struct {
	Stars        int     `json:"stars"`        // 项目星标数
	Forks        int     `json:"forks"`        // 项目Fork数
	Dependencies int     `json:"dependencies"` // 项目依赖数
	WorkWeight   float64 `json:"workWeight"`   // 开发者的工作占比
}

// CodeContribution 表示代码贡献量
type CodeContribution struct {
	CommitCount      int     `json:"commitCount"`      // 提交代码的数量
	CodeLinesAdded   int     `json:"codeLinesAdded"`   // 添加的代码行数
	CodeQualityScore float64 `json:"codeQualityScore"` // 代码质量评分
	IssueResolutions int     `json:"issueResolutions"` // 解决的issue数量
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
