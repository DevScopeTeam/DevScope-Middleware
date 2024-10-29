package github

import "time"

// Event represents a generic GitHub event.
type Event struct {
	ID        string     `json:"id"`
	Type      string     `json:"type"`
	Actor     Actor      `json:"actor"`
	Repo      Repository `json:"repo"`
	Payload   Payload    `json:"payload"`
	Public    bool       `json:"public"`
	CreatedAt time.Time  `json:"created_at"`
}

// Actor represents the user who triggered the event.
type Actor struct {
	ID           int64  `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

// Repository represents the repository associated with the event.
type Repository struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Payload is the payload of the event, which varies depending on the event type.
type Payload struct {
	Action       string   `json:"action,omitempty"`
	PushID       int64    `json:"push_id,omitempty"`
	Size         int      `json:"size,omitempty"`
	DistinctSize int      `json:"distinct_size,omitempty"`
	Ref          string   `json:"ref,omitempty"`
	Head         string   `json:"head,omitempty"`
	Before       string   `json:"before,omitempty"`
	Commits      []Commit `json:"commits,omitempty"`
}

// Commit represents a commit object within the payload of a PushEvent.
type Commit struct {
	SHA      string `json:"sha"`
	Author   Author `json:"author"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
	URL      string `json:"url"`
}

// Author represents the author of a commit.
type Author struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type EventResp struct {
	Code  int   `json:"code"`
	Event Event `json:"event"`
}

type EventListResp struct {
	Code int     `json:"code"`
	List []Event `json:"list"`
}
