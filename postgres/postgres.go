package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type DBLogger struct{}

func (d DBLogger) BeforeQuery(ctx context.Context, query *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(ctx context.Context, query *pg.QueryEvent) error {
	// Convert the query bytes to a string
	queryStr, err := query.FormattedQuery()
	if err == nil {
		fmt.Println(string(queryStr)) // Convert the bytes to a string for printing
	} else {
		fmt.Println("Error formatting query:", err)
	}
	return nil
}

func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)
	return db
}
