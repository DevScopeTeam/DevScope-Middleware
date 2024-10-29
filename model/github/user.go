package github

import (
	"time"
)

type User struct {
	AvatarURL               string    `json:"avatar_url"`
	EventsURL               string    `json:"events_url"`
	FollowersURL            string    `json:"followers_url"`
	FollowingURL            string    `json:"following_url"`
	GistsURL                string    `json:"gists_url"`
	GravatarID              string    `json:"gravatar_id"`
	HTMLURL                 string    `json:"html_url"`
	ID                      int       `json:"id"`
	NodeID                  string    `json:"node_id"`
	Login                   string    `json:"login"`
	OrganizationsURL        string    `json:"organizations_url"`
	ReceivedEventsURL       string    `json:"received_events_url"`
	ReposURL                string    `json:"repos_url"`
	SiteAdmin               bool      `json:"site_admin"`
	StarredURL              string    `json:"starred_url"`
	SubscriptionsURL        string    `json:"subscriptions_url"`
	Type                    string    `json:"type"`
	URL                     string    `json:"url"`
	Bio                     string    `json:"bio"`
	Blog                    string    `json:"blog"`
	Company                 string    `json:"company"`
	Email                   string    `json:"email"`
	Followers               int       `json:"followers"`
	Following               int       `json:"following"`
	Hireable                bool      `json:"hireable"`
	Location                string    `json:"location"`
	Name                    string    `json:"name"`
	PublicGists             int       `json:"public_gists"`
	PublicRepos             int       `json:"public_repos"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	Collaborators           int       `json:"collaborators"`
	DiskUsage               int       `json:"disk_usage"`
	OwnedPrivateRepos       int       `json:"owned_private_repos"`
	PrivateGists            int       `json:"private_gists"`
	TotalPrivateRepos       int       `json:"total_private_repos"`
	TwoFactorAuthentication bool      `json:"two_factor_authentication"`
}

type UserResp struct {
	Code int  `json:"code"`
	User User `json:"user"`
}
