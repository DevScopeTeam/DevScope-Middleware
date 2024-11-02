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

type EventResp struct {
	Code  int   `json:"code"`
	Event Event `json:"event"`
}

type EventListResp struct {
	Code int     `json:"code"`
	List []Event `json:"list"`
}
