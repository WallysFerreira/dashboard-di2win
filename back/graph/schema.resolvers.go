package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"api/graph/model"
	"api/service"
	"context"
	"fmt"
	"log"
)

// Count is the resolver for the count field.
func (r *queryResolver) Count(ctx context.Context, groupBy string) (*model.Count, error) {
	panic(fmt.Errorf("not implemented: Count - count"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *int, segment *string) ([]*model.User, error) {
	rep := service.RepositorioPostgre{
		ConnStr: "user=postgres dbname=database sslmode=disable",
	}

	if id != nil {
		user_found, err := rep.FindUser(*id)
		if err != nil {
			log.Fatal(err)
		}

		user_return := model.User{
			ID:      fmt.Sprintf("%d", user_found.Id),
			Name:    user_found.Name,
			Segment: user_found.Segment,
		}

		return []*model.User{&user_return}, nil
	}

	if segment != nil {
		user_found := rep.FindUsers(*segment)
		users_return := []*model.User{}

		for _, value := range user_found {
			user_return := model.User{
				ID:      fmt.Sprintf("%d", value.Id),
				Name:    value.Name,
				Segment: value.Segment,
			}

			users_return = append(users_return, &user_return)
		}

		return users_return, nil
	}

	users_found := rep.FindUsers("")
	users_return := []*model.User{}

	for _, value := range users_found {
		user_return := model.User{
			ID:      fmt.Sprintf("%d", value.Id),
			Name:    value.Name,
			Segment: value.Segment,
		}

		users_return = append(users_return, &user_return)
	}

	return users_return, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

/*
func (r *queryResolver) Extract(ctx context.Context, userID *int, tipoDocumento *string, dataComeco *string, dataFinal *string) (*model.Contagem, error) {
	rep := service.RepositorioPostgre{
		ConnStr: "user=postgres dbname=database sslmode=disable",
	}

	filter := service.FiltroExtract{}

	if dataComeco != nil {
		parsed_data_start, err := time.Parse(time.RFC3339, *dataComeco)
		if err != nil {
			log.Fatal(err)
		}

		filter.DataStart = parsed_data_start
	}

	if dataFinal != nil {
		parsed_data_final, err := time.Parse(time.RFC3339, *dataFinal)
		if err != nil {
			log.Fatal(err)
		}

		filter.DataEnd = parsed_data_final
	}

	if tipoDocumento != nil {
		filter.DocType = *tipoDocumento
	}

	if userID != nil {
		filter.UserId = *userID
	}

	_, count := rep.FindExtracts(filter)

	return &model.Contagem{Count: count}, nil
}
*/
