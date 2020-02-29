package graph

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ngavinsir/mari/graph/generated"
	"github.com/ngavinsir/mari/graph/model"
	"github.com/ngavinsir/mari/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	u := &models.User{
		Email: input.Email,
		Name: null.StringFromPtr(input.Name),
		Password: input.Password,
	}

	err := u.Insert(ctx, r.DB, boil.Infer())
	if err != nil {
		graphql.AddErrorf(ctx, "Can't create new user")
	}

	gqlUser := &model.User{
		ID: u.ID,
		Name: &u.Name.String,
		Email: u.Email,
	}

	return gqlUser, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation resolver
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
// Query resolver
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
