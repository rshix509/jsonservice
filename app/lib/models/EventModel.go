package models

import "time"

type EachEvent struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Actor struct {
		ID         int    `json:"id"`
		Login      string `json:"login"`
		GravatarID string `json:"gravatar_id"`
		URL        string `json:"url"`
		AvatarURL  string `json:"avatar_url"`
	} `json:"actor"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Payload struct {
		Ref          string `json:"ref"`
		RefType      string `json:"ref_type"`
		MasterBranch string `json:"master_branch"`
		Description  string `json:"description"`
		PusherType   string `json:"pusher_type"`
	} `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}
