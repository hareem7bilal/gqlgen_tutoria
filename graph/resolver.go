package graph

import "example/graph/model"

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver struct containing the meetups and users data
type Resolver struct {
	MeetupList []*model.Meetup
	UserList   []*model.User
}


