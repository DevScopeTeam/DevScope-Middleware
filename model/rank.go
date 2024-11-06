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
	CommitCount int `json:"commitCount"` // 提交Commit的数量
	PrCount     int `json:"prCount"`     // 提交Pull Request的数量
}

// CommunityInfluence 表示社区影响力
type CommunityInfluence struct {
	Followers int `json:"followers"` // 用户的Follower数量
	Stars     int `json:"stars"`     // 项目的Star数量
	Watchers  int `json:"watchers"`  // 项目的Watcher数量
}

// DetailRank 表示详细评分
type DetailRank struct {
	ProjectImportance  ProjectImportance  `json:"projectImportance"`  // 项目重要性评分
	CodeContribution   CodeContribution   `json:"codeContribution"`   // 代码贡献量评分
	CommunityInfluence CommunityInfluence `json:"communityInfluence"` // 社区影响力评分
}

// DeveloperRank 表示开发者的评分
type DeveloperRank struct {
	Username           string  `json:"username" gorm:"primaryKey" validate:"required"`                 // 用户名
	ProjectImportance  float64 `json:"project" validate:"required"`                                    // 项目重要性评分
	CodeContribution   float64 `json:"code" validate:"required"`                                       // 代码贡献量评分
	CommunityInfluence float64 `json:"influence" validate:"required"`                                  // 社区影响力评分
	Overall            float64 `json:"overall" validate:"required"`                                    // 综合评分
	Nation             string  `json:"nation,omitempty"  validate:"required" gorm:"type:varchar(191)"` // 国家
	UpdatedAt          XTime   `json:"updated_at,omitempty" gorm:"autoUpdateTime"`                     // 更新时间
}

// DevScoreResp 表示开发者分数的响应
type DevScoreResp struct {
	Code  int           `json:"code"`
	Score DeveloperRank `json:"score"`
}

// RankListResp 表示开发者排名列表的响应
type RankListResp struct {
	Code int             `json:"code"`
	List []DeveloperRank `json:"list"`
}

// CalculateOverallScore 计算综合评分
func (s *DeveloperRank) CalculateOverallScore() {
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

	// 将分数缩放到0-100的范围
	return sigmoid(sum) * 100
}

// 计算开发者的代码贡献量评分
func (c *CodeContribution) CalculateScore() float64 {
	// 权重可以根据实际情况进行调整
	const (
		commitWeight = 0.01
		prWeight     = 0.01
	)

	// 计算加权和
	sum := commitWeight*float64(c.CommitCount) + prWeight*float64(c.PrCount)

	// 将分数缩放到0-100的范围
	return sigmoid(sum) * 100
}

func (c *CommunityInfluence) CalculateScore() float64 {
	// 权重可以根据实际情况进行调整
	const (
		followersWeight = 0.01
		starsWeight     = 0.005
		watchersWeight  = 0.005
	)

	// 计算加权和
	sum := followersWeight*float64(c.Followers) + starsWeight*float64(c.Stars) + watchersWeight*float64(c.Watchers)

	// 将分数缩放到0-100的范围
	return sigmoid(sum) * 100
}
