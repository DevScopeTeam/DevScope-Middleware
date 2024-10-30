package github

import (
	"time"
)

// Issue represents a GitHub issue.
type Issue struct {
	ID                int          `json:"id"`
	NodeID            string       `json:"node_id"`
	URL               string       `json:"url"`
	RepositoryURL     string       `json:"repository_url"`
	LabelsURL         string       `json:"labels_url"`
	CommentsURL       string       `json:"comments_url"`
	EventsURL         string       `json:"events_url"`
	HTMLURL           string       `json:"html_url"`
	Number            int          `json:"number"`
	State             string       `json:"state"`
	Title             string       `json:"title"`
	Body              string       `json:"body"`
	User              *User        `json:"user"`
	Labels            []Label      `json:"labels"`
	Assignee          *User        `json:"assignee"`
	Assignees         []*User      `json:"assignees"`
	Milestone         *Milestone   `json:"milestone"`
	Locked            bool         `json:"locked"`
	ActiveLockReason  string       `json:"active_lock_reason"`
	Comments          int          `json:"comments"`
	PullRequest       *PullRequest `json:"pull_request"`
	ClosedAt          *time.Time   `json:"closed_at"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
	ClosedBy          *User        `json:"closed_by"`
	AuthorAssociation string       `json:"author_association"`
	StateReason       string       `json:"state_reason"`
}
