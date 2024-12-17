package graph

import (
	"context"
	"example/graph/models"
	"fmt"
)

// Meetups is the resolver for the meetups field.
func (r *meetupsResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	panic(fmt.Errorf("not implemented: CreateMeetup - createMeetup"))

}

func (r *Resolver) Meetup() MeetupResolver { return &meetupsResolver{r} }

type meetupsResolver struct{ *Resolver }