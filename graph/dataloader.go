package graph

import (
	"context"
	"example/graph/models"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
)

const userloaderKey = "userloader"

func DataloaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			userloader := UserSliceLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []string) ([]*models.User, []error) {
					var users []*models.User

					err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
					if err != nil {
						return nil, []error{err}
					}

					u := make(map[string]*models.User, len(users))
					for _, user := range users {
						u[user.ID] = user
					}

					result := make([]*models.User, len(ids))

					for i, id := range ids {
						result[i] = u[id]
					}

					return result, nil
				},
			}

			ctx := context.WithValue(r.Context(), userloaderKey, &userloader)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func getUserloader(ctx context.Context) *UserSliceLoader {
	return ctx.Value(userloaderKey).(*UserSliceLoader)
}
