package graph

import (
	"context"
	"errors"
	"example/graph/models"
)

// CreateMeetup is the resolver for the createMeetup field.
func (r *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 5 {
		return nil, errors.New("Name should be atleast 5 characters")
	}
	if len(input.Description) < 5 {
		return nil, errors.New("Description should be atleast 5 characters")
	}
	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      "1",
	}
	return r.MeetupsRepo.CreateMeetup(meetup)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
