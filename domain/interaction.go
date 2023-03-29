package domain

import "time"

type Interaction struct {
	UserID      int       `json:"user_id"`
	Service     string    `json:"service"`
	ContentType string    `json:"content_type"`
	ContentID   int       `json:"content_id"`
	Action      string    `json:"action"`
	CreatedAt   time.Time `json:"created_at"`
}

type RequestGet struct {
	Service     string
	ContentType string
	ContentId   int
}

type Total struct {
	Total       int
	TotalString string
}

type DeleteMongo struct {
	UserId      int    `json:"user_id"`
	Service     string `json:"service"`
	ContentType string `json:"content_type"`
	ContentID   int    `json:"content_id"`
}

type InteractionView struct {
	UserID      int       `json:"user_id"`
	Service     string    `json:"service"`
	ContentType string    `json:"content_type"`
	ContentID   int       `json:"content_id"`
	Action      string    `json:"action"`
	CreatedAt   time.Time `json:"created_at"`
	DeviceID    string    `json:"device_id"`
}

type InteractionFollow struct {
	UserID    int       `json:"user_id"`
	FollowTo  int       `json:"follow_to"`
	CreatedAt time.Time `json:"created_at"`
}

type DeleteFollowMongo struct {
	UserID   int `json:"user_id"`
	FollowTo int `json:"follow_to"`
}

type RequestGetFollow struct {
	UserID int `json:"user_id"`
}

type RequestGetFollower struct {
	FollowTo int `json:"follow_to"`
}

type TotalFollow struct {
	TotalFollowing       int
	TotalStringFollowing string
	TotalFollower        int
	TotalStringFollower  string
}
