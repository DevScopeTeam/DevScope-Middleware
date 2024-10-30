package github

import "time"

// PullRequest represents a pull request on GitHub.
type PullRequest struct {
	URL                string     `json:"url,omitempty"`
	ID                 int64      `json:"id,omitempty"`
	NodeID             string     `json:"node_id,omitempty"`
	HTMLURL            string     `json:"html_url,omitempty"`
	DiffURL            string     `json:"diff_url,omitempty"`
	PatchURL           string     `json:"patch_url,omitempty"`
	IssueURL           string     `json:"issue_url,omitempty"`
	CommitsURL         string     `json:"commits_url,omitempty"`
	ReviewCommentsURL  string     `json:"review_comments_url,omitempty"`
	ReviewCommentURL   string     `json:"review_comment_url,omitempty"`
	CommentsURL        string     `json:"comments_url,omitempty"`
	StatusesURL        string     `json:"statuses_url,omitempty"`
	Number             int        `json:"number,omitempty"`
	State              string     `json:"state,omitempty"`
	Locked             bool       `json:"locked,omitempty"`
	Title              string     `json:"title,omitempty"`
	User               User       `json:"user,omitempty"`
	Body               string     `json:"body,omitempty"`
	Labels             []Label    `json:"labels,omitempty"`
	Milestone          Milestone  `json:"milestone,omitempty"`
	ActiveLockReason   string     `json:"active_lock_reason,omitempty"`
	CreatedAt          time.Time  `json:"created_at,omitempty"`
	UpdatedAt          time.Time  `json:"updated_at,omitempty"`
	ClosedAt           *time.Time `json:"closed_at,omitempty"`
	MergedAt           *time.Time `json:"merged_at,omitempty"`
	MergeCommitSHA     string     `json:"merge_commit_sha,omitempty"`
	Assignee           User       `json:"assignee,omitempty"`
	Assignees          []User     `json:"assignees,omitempty"`
	RequestedReviewers []User     `json:"requested_reviewers,omitempty"`
	RequestedTeams     []Team     `json:"requested_teams,omitempty"`
	Head               struct {
		Label string     `json:"label,omitempty"`
		Ref   string     `json:"ref,omitempty"`
		SHA   string     `json:"sha,omitempty"`
		User  User       `json:"user,omitempty"`
		Repo  Repository `json:"repo,omitempty"`
	} `json:"head,omitempty"`
	Base struct {
		Label string     `json:"label,omitempty"`
		Ref   string     `json:"ref,omitempty"`
		SHA   string     `json:"sha,omitempty"`
		User  User       `json:"user,omitempty"`
		Repo  Repository `json:"repo,omitempty"`
	} `json:"base,omitempty"`
	// _links is a map of link relations to their corresponding URLs
	Links map[string]struct {
		Href string `json:"href,omitempty"`
	} `json:"_links,omitempty"`
	AuthorAssociation string      `json:"author_association,omitempty"`
	AutoMerge         interface{} `json:"auto_merge,omitempty"`
	Draft             bool        `json:"draft,omitempty"`
}

type Label struct {
	ID          int64  `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
}

type Milestone struct {
	URL          string     `json:"url"`
	HTMLURL      string     `json:"html_url"`
	LabelsURL    string     `json:"labels_url"`
	ID           int        `json:"id"`
	NodeID       string     `json:"node_id"`
	Number       int        `json:"number"`
	State        string     `json:"state"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Creator      User       `json:"creator"`
	OpenIssues   int        `json:"open_issues"`
	ClosedIssues int        `json:"closed_issues"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	ClosedAt     *time.Time `json:"closed_at,omitempty"`
	DueOn        *time.Time `json:"due_on,omitempty"`
}

type Team struct {
	ID                  int64       `json:"id"`
	NodeID              string      `json:"node_id"`
	URL                 string      `json:"url"`
	HTMLURL             string      `json:"html_url"`
	Name                string      `json:"name"`
	Slug                string      `json:"slug"`
	Description         string      `json:"description"`
	Privacy             string      `json:"privacy"`
	Permission          string      `json:"permission"`
	NotificationSetting string      `json:"notification_setting"`
	MembersURL          string      `json:"members_url"`
	RepositoriesURL     string      `json:"repositories_url"`
	Parent              interface{} `json:"parent,omitempty"`
}
