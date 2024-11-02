package github

// Commit represents a GitHub commit object.
type Commit struct {
	URL         string        `json:"url"`
	SHA         string        `json:"sha"`
	NodeID      string        `json:"node_id,omitempty"`
	HTMLURL     string        `json:"html_url,omitempty"`
	CommentsURL string        `json:"comments_url,omitempty"`
	Commit      CommitDetails `json:"commit,omitempty"`
	Author      User          `json:"author"`
	Committer   User          `json:"committer,omitempty"`
	Parents     []Parent      `json:"parents,omitempty"`
	Message     string        `json:"message"`
	Distinct    bool          `json:"distinct"`
}

// CommitDetails represents the details of a commit.
type CommitDetails struct {
	URL          string       `json:"url"`
	Author       User         `json:"author"`
	Committer    User         `json:"committer"`
	Message      string       `json:"message"`
	Tree         Tree         `json:"tree"`
	CommentCount int          `json:"comment_count"`
	Verification Verification `json:"verification"`
}

// Tree represents the tree of a commit.
type Tree struct {
	URL string `json:"url"`
	SHA string `json:"sha"`
}

// Verification represents the verification status of a commit.
type Verification struct {
	Verified  bool   `json:"verified"`
	Reason    string `json:"reason"`
	Signature string `json:"signature"`
	Payload   string `json:"payload"`
}

// Parent represents a parent commit.
type Parent struct {
	URL string `json:"url"`
	SHA string `json:"sha"`
}
