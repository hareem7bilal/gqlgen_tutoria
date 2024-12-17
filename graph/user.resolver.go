package graph

import (
	"context"
	"example/graph/models"
)

// User is the resolver for the user field.
func (r *userResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return getUserloader(ctx).Load(obj.UserID)
}

func (r *Resolver) User() UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
