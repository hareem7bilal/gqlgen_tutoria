gqlgen-generate:
	go get github.com/99designs/gqlgen/codegen/config@v0.17.57
	go get github.com/99designs/gqlgen/internal/imports@v0.17.57
	go get github.com/99designs/gqlgen@v0.17.57
	go run github.com/99designs/gqlgen generate

generate-user-slice-loader:
	go run github.com/vektah/dataloaden UserSliceLoader string *example/graph/models.User

