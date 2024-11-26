package model

type Meetup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID        string  `json:"userId"`
}

type NewMeetup struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}