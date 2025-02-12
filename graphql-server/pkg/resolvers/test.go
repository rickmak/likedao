package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	pkgContext "github.com/oursky/likedao/pkg/context"
	graphql1 "github.com/oursky/likedao/pkg/generated/graphql"
	"github.com/oursky/likedao/pkg/models"
)

func (r *mutationResolver) CreateTest(ctx context.Context, input models.CreateTest) (*models.Test, error) {
	res, err := pkgContext.GetMutatorsFromCtx(ctx).Test.CreateTest(input.String, input.Int)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *queryResolver) QueryTestByID(ctx context.Context, id models.NodeID) (*models.Test, error) {
	res, err := pkgContext.GetDataLoadersFromCtx(ctx).Test.Load(id.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *queryResolver) QueryTestsByIDs(ctx context.Context, ids []models.NodeID) ([]*models.Test, error) {
	objIds := models.ExtractObjectIDs(ids)
	res, errs := pkgContext.GetDataLoadersFromCtx(ctx).Test.LoadAll(objIds)

	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
