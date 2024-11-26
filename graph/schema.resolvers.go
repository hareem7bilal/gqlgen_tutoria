package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.57

import (
	"context"
	"errors"
	"example/graph/model"
	"fmt"
)

var meetups = []*model.Meetup{
	&model.Meetup{
		ID:          "1",
		Name:        "Tech Innovators Meetup",
		Description: "Explore the latest in technology and innovation.",
		UserID:      "2",
	},
	{
		ID:          "2",
		Name:        "Startup Networking Event",
		Description: "A platform for entrepreneurs to connect and share ideas.",
		UserID:      "5",
	},
	{
		ID:          "3",
		Name:        "AI & Machine Learning Forum",
		Description: "Discuss the future of artificial intelligence and ML.",
		UserID:      "4",
	},
	{
		ID:          "4",
		Name:        "Web Developers Meetup",
		Description: "Learn about new tools and frameworks for web development.",
		UserID:      "1",
	},
	{
		ID:          "5",
		Name:        "Cloud Computing Conference",
		Description: "Dive into cloud architectures and best practices.",
		UserID:      "3",
	},
}
var users = []*model.User{
	&model.User{
		ID:       "1",
		Username: "JohnDoe",
		Email:    "john.doe@example.com",
	},
	{
		ID:       "2",
		Username: "JaneSmith",
		Email:    "jane.smith@example.com",
	},
	{
		ID:       "3",
		Username: "BobBrown",
		Email:    "bob.brown@example.com",
	},
	{
		ID:       "4",
		Username: "AliceWhite",
		Email:    "alice.white@example.com",
	},
	{
		ID:       "5",
		Username: "CharlieBlack",
		Email:    "charlie.black@example.com",
	},
}

// User is the resolver for the user field.
func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	user := new(model.User)
	for _, u := range users {
		if u.ID == obj.UserID {
			user = u
			break

		}
	}

	if user == nil {
		return nil, errors.New("User with id does not exist")

	}
	return user, nil
}

// CreateMeetup is the resolver for the createMeetup field.
func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	panic(fmt.Errorf("not implemented: CreateMeetup - createMeetup"))
}

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	r.MeetupList = append(r.MeetupList, meetups...)
	return r.MeetupList, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	r.UserList = append(r.UserList, users...)
	return r.UserList, nil
}

// Meetups is the resolver for the meetups field.
func (r *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	var meetupsSlice []*model.Meetup

	for _, m := range meetups {
		if m.UserID == obj.ID {
			meetupsSlice = append(meetupsSlice, m)

		}
	}

	return meetupsSlice, nil

}

// Meetup returns MeetupResolver implementation.
func (r *Resolver) Meetup() MeetupResolver { return &meetupResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
